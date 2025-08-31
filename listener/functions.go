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

func (l Listener) ExitReturnStatement(ctx *p.ReturnStatementContext) {
	line := ctx.GetStart().GetLine()
	colStart := ctx.GetStart().GetColumn()
	colEnd := colStart + len(ctx.GetText())

	funcScope, inFunctionScope := l.ScopeManager.SearchScopeByType(SCOPE_TYPES.FUNCTION)
	if !inFunctionScope {
		l.AddError(line, colStart, colEnd, "'return' statement outside functionscope.")
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
		if expectedReturnType != BASE_TYPES.UNKNOWN && expectedReturnType != returnValueType {
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
