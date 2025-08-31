package listener

import (
	"fmt"
	"log"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

func (l Listener) EnterFunctionDeclaration(ctx *p.FunctionDeclarationContext) {
	line := ctx.GetStart().GetLine()
	funcName := ctx.Identifier()
	nameColStart := funcName.GetSymbol().GetColumn()
	nameColEnd := nameColStart + len(funcName.GetText())

	// Check duplicates in class scope
	isInsideClassDeclaration := l.ScopeManager.CurrentScope.Type == SCOPE_TYPES.CLASS
	if !isInsideClassDeclaration {
		if _, exists := l.ScopeManager.CurrentScope.functions[funcName.GetText()]; exists {
			l.AddError(line, nameColStart, nameColEnd, fmt.Sprintf(
				"Function `%s` is already declared in this scope!",
				funcName.GetText(),
			))
			return
		}
	}

	funcParams := []ParameterInfo{}
	if ctx.Parameters() != nil {
		paramNames := make(map[string]bool)
		for _, paramCtx := range ctx.Parameters().AllParameter() {
			name := paramCtx.Identifier()
			paramColStart := name.GetSymbol().GetColumn()
			paramColEnd := paramColStart + len(name.GetText())

			// Check for duplicate parameters
			if paramNames[name.GetText()] {
				l.AddError(line, paramColStart, paramColEnd, fmt.Sprintf(
					"Duplicate parameter name `%s` in function `%s`",
					name.GetText(),
					funcName.GetText(),
				))
				continue
			}
			paramNames[name.GetText()] = true

			paramType := BASE_TYPES.UNKNOWN
			if paramCtx.Type_() != nil {
				paramType = TypeIdentifier(paramCtx.Type_().GetText())
				typeColStart := paramCtx.Type_().GetStart().GetColumn()
				typeColEnd := typeColStart + len(paramCtx.Type_().GetText())

				if !l.TypeExists(paramType) {
					l.AddError(line, typeColStart, typeColEnd, fmt.Sprintf(
						"Parameter type `%s` doesn't exist!",
						paramType,
					))
				}
			}

			funcParams = append(funcParams, ParameterInfo{
				Name: name.GetText(),
				Type: paramType,
			})
		}
	}

	returnType := BASE_TYPES.UNKNOWN
	if ctx.Type_() != nil {
		returnType = TypeIdentifier(ctx.Type_().GetText())
		retColStart := ctx.Type_().GetStart().GetColumn()
		retColEnd := retColStart + len(ctx.Type_().GetText())

		if !l.TypeExists(returnType) {
			l.AddError(line, retColStart, retColEnd, fmt.Sprintf(
				"Return type `%s` doesn't exist!",
				returnType,
			))
			returnType = BASE_TYPES.INVALID
		}
	}

	info := MethodInfo{
		ParameterList: funcParams,
		ReturnType:    returnType,
	}

	var funcScope *Scope
	if isInsideClassDeclaration {
		className := l.ScopeManager.CurrentScope.Name
		l.ModifyClassTypeInfo(TypeIdentifier(className), func(cti *ClassTypeInfo) {
			if funcName.GetText() == CONSTRUCTOR_NAME {
				cti.Constructor = info
			} else {
				cti.UpsertMethod(funcName.GetText(), info)
			}
		})
		funcScope = NewScope(className+"_"+funcName.GetText(), SCOPE_TYPES.FUNCTION)
	} else {
		l.ScopeManager.CurrentScope.UpsertFunctionDef(funcName.GetText(), info)
		funcScope = NewScope(funcName.GetText(), SCOPE_TYPES.FUNCTION)

		l.ScopeManager.CurrentScope.UpsertExpressionType(funcName.GetText(), info.ReturnType)
	}

	for _, param := range info.ParameterList {
		funcScope.UpsertExpressionType(param.Name, param.Type)
	}

	funcScope.expectedReturnType = returnType
	funcScope.hasReturnStatement = false

	l.ScopeManager.AddToCurrent(funcScope)
	l.ScopeManager.ReplaceCurrent(funcScope)
}

func (l Listener) ExitFunctionDeclaration(ctx *p.FunctionDeclarationContext) {
	if l.ScopeManager.CurrentScope.Type != SCOPE_TYPES.FUNCTION {
		log.Panicf("Trying to exit function scope but scope is not of type function! %#v", l.ScopeManager.CurrentScope)
	}

	line := ctx.GetStart().GetLine()
	funcName := ctx.Identifier().GetText()

	expectedReturnType := l.ScopeManager.CurrentScope.expectedReturnType

	if expectedReturnType == BASE_TYPES.UNKNOWN {
		if l.ScopeManager.CurrentScope.inferredReturnType != BASE_TYPES.UNKNOWN {
			expectedReturnType = l.ScopeManager.CurrentScope.inferredReturnType
			l.updateFunctionReturnType(funcName, expectedReturnType)
		}
	}

	if expectedReturnType != BASE_TYPES.UNKNOWN && expectedReturnType != BASE_TYPES.INVALID {
		if !l.ScopeManager.CurrentScope.hasReturnStatement {
			nameColStart := ctx.Identifier().GetSymbol().GetColumn()
			nameColEnd := nameColStart + len(funcName)
			l.AddError(line, nameColStart, nameColEnd, fmt.Sprintf(
				"Function `%s` with return type `%s` must have at least one return statement",
				funcName,
				expectedReturnType,
			))
		}
	}

	l.ScopeManager.ReplaceWithParent()
}

func (l Listener) updateFunctionReturnType(funcName string, returnType TypeIdentifier) {
	isInsideClassDeclaration := l.ScopeManager.CurrentScope.Type == SCOPE_TYPES.CLASS

	if isInsideClassDeclaration {
		className := l.ScopeManager.CurrentScope.Name
		l.ModifyClassTypeInfo(TypeIdentifier(className), func(cti *ClassTypeInfo) {
			if funcName == CONSTRUCTOR_NAME {
				cti.Constructor.ReturnType = returnType
			} else if methodInfo, exists := cti.Methods[funcName]; exists {
				methodInfo.ReturnType = returnType
				cti.Methods[funcName] = methodInfo
			}
		})
	} else {
		if funcInfo, exists := l.ScopeManager.CurrentScope.functions[funcName]; exists {
			funcInfo.ReturnType = returnType
			l.ScopeManager.CurrentScope.functions[funcName] = funcInfo
		}
		l.ScopeManager.CurrentScope.UpsertExpressionType(funcName, returnType)
	}
}

func (l Listener) ExitReturnStatement(ctx *p.ReturnStatementContext) {
	line := ctx.GetStart().GetLine()
	colStart := ctx.GetStart().GetColumn()
	colEnd := colStart + len(ctx.GetText())

	funcScope, inFunctionScope := l.ScopeManager.SearchScopeByType(SCOPE_TYPES.FUNCTION)
	if !inFunctionScope {
		l.AddError(line, colStart, colEnd, "'return' statement outside function scope.")
		return
	}

	funcScope.hasReturnStatement = true

	if ctx.ConditionalExpr() != nil {
		returnExpr := ctx.ConditionalExpr()
		exprColStart := returnExpr.GetStart().GetColumn()
		exprColEnd := exprColStart + len(returnExpr.GetText())

		returnValueType, found := l.ScopeManager.CurrentScope.GetExpressionType(returnExpr.GetText())
		if !found {
			l.AddError(line, exprColStart, exprColEnd, fmt.Sprintf(
				"Return expression `%s` has unknown type",
				returnExpr.GetText(),
			))
			return
		}

		expectedReturnType := funcScope.expectedReturnType

		if expectedReturnType == BASE_TYPES.UNKNOWN {
			funcScope.inferredReturnType = returnValueType
			log.Printf("Inferred return type for function: %s", returnValueType)
		} else if expectedReturnType != returnValueType {
			l.AddError(line, exprColStart, exprColEnd, fmt.Sprintf(
				"Return type mismatch: expected `%s` but got `%s`",
				expectedReturnType,
				returnValueType,
			))
		}
	} else {
		expectedReturnType := funcScope.expectedReturnType
		if expectedReturnType != BASE_TYPES.UNKNOWN && expectedReturnType != BASE_TYPES.INVALID {
			l.AddError(line, colStart, colEnd, fmt.Sprintf(
				"Function expects return type `%s` but return statement has no value",
				expectedReturnType,
			))
		}
	}
}

func (l Listener) EnterCallExpr(ctx *p.CallExprContext) {
	line := ctx.GetStart().GetLine()
	parent := ctx.GetParent()
	var funcName string
	var funcNameColStart, funcNameColEnd int

	if leftHandSide, ok := parent.(*p.LeftHandSideContext); ok {
		if primaryAtom := leftHandSide.PrimaryAtom(); primaryAtom != nil {
			if identifierExpr, ok := primaryAtom.(*p.IdentifierExprContext); ok {
				funcName = identifierExpr.Identifier().GetText()
				funcNameColStart = identifierExpr.Identifier().GetSymbol().GetColumn()
				funcNameColEnd = funcNameColStart + len(funcName)
			}
		}
	}

	if funcName == "" {
		l.AddError(line, ctx.GetStart().GetColumn(), ctx.GetStop().GetColumn()+1,
			"Could not determine function name for call expression")
		return
	}

	log.Printf("Processing function call: %s", funcName)

	exprArguments := []p.IConditionalExprContext{}
	if ctx.Arguments() != nil {
		exprArguments = ctx.Arguments().AllConditionalExpr()
	}

	funcInfo, found := l.findFunctionInfo(funcName)
	if !found {
		l.AddError(line, funcNameColStart, funcNameColEnd, fmt.Sprintf(
			"Undefined function `%s`",
			funcName,
		))
		return
	}

	if len(funcInfo.ParameterList) != len(exprArguments) {
		l.AddError(line, funcNameColStart, funcNameColEnd, fmt.Sprintf(
			"Function `%s` expects %d arguments but %d given",
			funcName,
			len(funcInfo.ParameterList),
			len(exprArguments),
		))
		return
	}
}

func (l Listener) ExitCallExpr(ctx *p.CallExprContext) {
	line := ctx.GetStart().GetLine()
	parent := ctx.GetParent()
	var funcName string

	if leftHandSide, ok := parent.(*p.LeftHandSideContext); ok {
		if primaryAtom := leftHandSide.PrimaryAtom(); primaryAtom != nil {
			if identifierExpr, ok := primaryAtom.(*p.IdentifierExprContext); ok {
				funcName = identifierExpr.Identifier().GetText()
			}
		}
	}

	if funcName == "" {
		return
	}

	exprArguments := []p.IConditionalExprContext{}
	if ctx.Arguments() != nil {
		exprArguments = ctx.Arguments().AllConditionalExpr()
	}

	funcInfo, found := l.findFunctionInfo(funcName)
	if !found {
		return
	}

	if len(funcInfo.ParameterList) != len(exprArguments) {
		return
	}

	for i, param := range funcInfo.ParameterList {
		if i >= len(exprArguments) {
			break
		}

		argExpr := exprArguments[i]
		argColStart := argExpr.GetStart().GetColumn()
		argColEnd := argColStart + len(argExpr.GetText())

		argType, found := l.ScopeManager.CurrentScope.GetExpressionType(argExpr.GetText())
		if !found {
			l.AddError(line, argColStart, argColEnd, fmt.Sprintf(
				"Type of argument `%s` not found",
				argExpr.GetText(),
			))
			continue
		}

		if param.Type != BASE_TYPES.UNKNOWN && param.Type != argType {
			l.AddError(line, argColStart, argColEnd, fmt.Sprintf(
				"Function `%s` parameter `%s` expects type `%s` but got `%s`",
				funcName,
				param.Name,
				param.Type,
				argType,
			))
		}
	}

	if leftHandSide, ok := parent.(*p.LeftHandSideContext); ok {
		fullExpr := leftHandSide.GetText()
		log.Printf("Setting return type of `%s` to `%s`", fullExpr, funcInfo.ReturnType)
		l.ScopeManager.CurrentScope.UpsertExpressionType(fullExpr, funcInfo.ReturnType)
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), funcInfo.ReturnType)
	}
}

func (l Listener) findFunctionInfo(funcName string) (MethodInfo, bool) {
	// First check current scope and walk up
	scope := l.ScopeManager.CurrentScope
	for scope != nil {
		if funcInfo, found := scope.functions[funcName]; found {
			return funcInfo, true
		}
		scope = scope.Father
	}

	// Check if it's a method in the current class (if we're in a class)
	if classScope, inClass := l.ScopeManager.SearchClassScope(); inClass {
		if typeInfo, found := l.GetTypeInfo(TypeIdentifier(classScope.Name)); found {
			if typeInfo.ClassType.HasValue() {
				classInfo := typeInfo.ClassType.GetValue()
				if methodInfo, found := classInfo.Methods[funcName]; found {
					return methodInfo, true
				}
			}
		}
	}

	return MethodInfo{}, false
}
