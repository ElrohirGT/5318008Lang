package type_checker

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

	isInsideClassDeclaration := l.ScopeManager.CurrentScope.Type == SCOPE_TYPES.CLASS

	// Check duplicates in different contexts
	if isInsideClassDeclaration {
		// Check for duplicate methods in class
		className := l.ScopeManager.CurrentScope.Name
		typeInfo, found := l.GetTypeInfo(TypeIdentifier(className))
		if found && typeInfo.ClassType.HasValue() {
			classInfo := typeInfo.ClassType.GetValue()

			// Check if method already exists
			if funcName.GetText() == CONSTRUCTOR_NAME {

				log.Printf("DEBUG: Checking constructor for class %s", className)
				// log.Printf("DEBUG: Current constructor: ReturnType=%s, ParamCount=%d",
				// 	classInfo.Constructor.ReturnType, len(classInfo.Constructor.ParameterList))

				hasExistingConstructor := classInfo.constructor != nil

				if hasExistingConstructor {
					l.AddError(line, nameColStart, nameColEnd, fmt.Sprintf(
						"Constructor for class `%s` is already declared!",
						className,
					))
					return
				}

			} else {
				// Check if method already exists
				if _, exists := classInfo.Methods[funcName.GetText()]; exists {
					l.AddError(line, nameColStart, nameColEnd, fmt.Sprintf(
						"Method `%s` is already declared in class `%s`!",
						funcName.GetText(),
						className,
					))
					return
				}
			}
		}
	} else {
		// Check for duplicate functions in current scope (non-class context)
		if _, exists := l.ScopeManager.CurrentScope.functions[funcName.GetText()]; exists {
			l.AddError(line, nameColStart, nameColEnd, fmt.Sprintf(
				"Function `%s` is already declared in this scope!",
				funcName.GetText(),
			))
			return
		}

		scope := l.ScopeManager.CurrentScope.Father
		for scope != nil {
			if _, exists := scope.functions[funcName.GetText()]; exists {
				if scope.Type == SCOPE_TYPES.GLOBAL {
					l.AddWarning(fmt.Sprintf(
						"Function `%s` shadows a builtin function",
						funcName.GetText(),
					), fmt.Sprintf("line %d", line))
				} else {
					l.AddError(line, nameColStart, nameColEnd, fmt.Sprintf(
						"Function `%s` conflicts with function in parent scope",
						funcName.GetText(),
					))
					return
				}
				break
			}
			scope = scope.Father
		}
	}

	// Parse function parameters
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

	// Parse return type
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

	// Register the function/method in the appropriate place
	if isInsideClassDeclaration {
		className := l.ScopeManager.CurrentScope.Name
		l.ModifyClassTypeInfo(TypeIdentifier(className), func(cti *ClassTypeInfo) {
			if funcName.GetText() == CONSTRUCTOR_NAME {
				if info.ReturnType != BASE_TYPES.UNKNOWN && info.ReturnType != TypeIdentifier(className) {
					l.AddError(
						line,
						ctx.GetStart().GetColumn(),
						ctx.GetStop().GetColumn(),
						fmt.Sprintf("Can't return `%s` for a constructor of type `%s`", info.ReturnType, className),
					)
				} else {
					cti.constructor = &info
					log.Printf("DEBUG: Set constructor for class %s: %+v", className, info)
				}
			} else {
				cti.UpsertMethod(funcName.GetText(), info)
			}
		})
	} else {
		l.ScopeManager.CurrentScope.UpsertFunctionDef(funcName.GetText(), info)
	}

	// Create and setup function scope
	var funcScope *Scope
	if isInsideClassDeclaration {
		className := l.ScopeManager.CurrentScope.Name
		funcScope = NewScope(className+"_"+funcName.GetText(), SCOPE_TYPES.FUNCTION)
	} else {
		funcScope = NewScope(funcName.GetText(), SCOPE_TYPES.FUNCTION)
	}

	// Set up function scope with parameters and return type info
	for _, param := range info.ParameterList {
		funcScope.UpsertExpressionType(param.Name, param.Type)
	}

	funcScope.expectedReturnType = returnType
	funcScope.hasReturnStatement = false

	// CRITICAL: Add the function scope as a child and then enter it
	l.ScopeManager.CurrentScope.AddChildScope(funcScope)
	l.ScopeManager.ReplaceCurrent(funcScope)

	log.Printf("Entered function scope: %s (type: %s)", funcScope.Name, funcScope.Type)
}

func (l Listener) ExitFunctionDeclaration(ctx *p.FunctionDeclarationContext) {
	log.Printf("Attempting to exit function scope. Current scope: %s (type: %s)",
		l.ScopeManager.CurrentScope.Name, l.ScopeManager.CurrentScope.Type)

	if l.ScopeManager.CurrentScope.Type != SCOPE_TYPES.FUNCTION {
		log.Printf("Warning: Attempting to exit function scope but current scope is %s. This likely means the function declaration had errors during entry.",
			l.ScopeManager.CurrentScope.Type)
		return
	}

	line := ctx.GetStart().GetLine()
	funcName := ctx.Identifier().GetText()
	funcScope := l.ScopeManager.CurrentScope
	expectedReturnType := funcScope.expectedReturnType

	l.ScopeManager.ReplaceWithParent()
	log.Printf("Exited function scope, now in: %s (type: %s)",
		l.ScopeManager.CurrentScope.Name, l.ScopeManager.CurrentScope.Type)

	isInsideClassDeclaration := l.ScopeManager.CurrentScope.Type == SCOPE_TYPES.CLASS
	if isInsideClassDeclaration {
		className := l.ScopeManager.CurrentScope.Name
		if funcName == CONSTRUCTOR_NAME && className != string(funcScope.inferredReturnType) && funcScope.inferredReturnType != BASE_TYPES.UNKNOWN {
			l.AddError(
				line,
				ctx.GetStart().GetColumn(),
				ctx.GetStop().GetColumn(),
				fmt.Sprintf("Constructor of type `%s` cannot return type `%s`", className, funcScope.inferredReturnType),
			)
			return
		}
	}

	if expectedReturnType == BASE_TYPES.UNKNOWN {
		if funcScope.inferredReturnType != BASE_TYPES.UNKNOWN {
			expectedReturnType = funcScope.inferredReturnType
			l.updateFunctionReturnType(funcName, expectedReturnType, ctx)
		}
	}

	if expectedReturnType != BASE_TYPES.UNKNOWN && expectedReturnType != BASE_TYPES.INVALID {
		allPathsReturn := l.checkAllPathsReturn(ctx)

		if !allPathsReturn {
			nameColStart := ctx.Identifier().GetSymbol().GetColumn()
			nameColEnd := nameColStart + len(funcName)
			l.AddError(line, nameColStart, nameColEnd, fmt.Sprintf(
				"Function `%s` with return type `%s` must have at least one return statement",
				funcName,
				expectedReturnType,
			))
		}
	}

}

func (l Listener) updateFunctionReturnType(funcName string, returnType TypeIdentifier, ctx *p.FunctionDeclarationContext) {
	isInsideClassDeclaration := l.ScopeManager.CurrentScope.Type == SCOPE_TYPES.CLASS

	if isInsideClassDeclaration {
		className := l.ScopeManager.CurrentScope.Name
		l.ModifyClassTypeInfo(TypeIdentifier(className), func(cti *ClassTypeInfo) {
			// FIXME: The constructor should only return the type of the class!
			// It doesn't make sense for it to return other thing!
			if funcName == CONSTRUCTOR_NAME && returnType != TypeIdentifier(className) {
				l.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					ctx.GetStop().GetColumn(),
					fmt.Sprintf("Can't return a type of `%s` for a constructor of `%s`", returnType, className),
				)
			} else if methodInfo, exists := cti.Methods[funcName]; exists {
				methodInfo.ReturnType = returnType
				cti.Methods[funcName] = methodInfo
			}
		})
	} else {
		if funcInfo, exists := l.ScopeManager.CurrentScope.functions[funcName]; exists {
			funcInfo.ReturnType = returnType
			l.ScopeManager.CurrentScope.functions[funcName] = funcInfo
			l.ScopeManager.CurrentScope.UpsertFunctionDef(funcName, funcInfo)
		}
	}
}

func (l Listener) ExitReturnStatement(ctx *p.ReturnStatementContext) {
	line := ctx.GetStart().GetLine()
	colStart := ctx.GetStart().GetColumn()
	colEnd := colStart + len(ctx.GetText())

	l.ScopeManager.CurrentScope.terminated = true

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
		// Analyze Assignments
		if primaryAtom := leftHandSide.PrimaryAtom(); primaryAtom != nil {
			if identifierExpr, ok := primaryAtom.(*p.IdentifierExprContext); ok {
				funcName = identifierExpr.Identifier().GetText()
				funcNameColStart = identifierExpr.Identifier().GetSymbol().GetColumn()
				funcNameColEnd = funcNameColStart + len(funcName)
			}
		}
	} else {
		// Analyze Standalone expression
		current := parent
		for current != nil {
			if standaloneExpr, ok := current.(*p.StandaloneExpresionContext); ok {
				if standaloneAtom := standaloneExpr.StandaloneAtom(); standaloneAtom != nil {
					if identifierExpr, ok := standaloneAtom.(*p.StandaloneIdentifierExprContext); ok {
						funcName = identifierExpr.Identifier().GetText()
						funcNameColStart = identifierExpr.Identifier().GetSymbol().GetColumn()
						funcNameColEnd = funcNameColStart + len(funcName)
						break
					}
				}
			}
			current = current.GetParent()
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
	var isStandaloneExpression bool

	if leftHandSide, ok := parent.(*p.LeftHandSideContext); ok {
		// Analyze assignments
		if primaryAtom := leftHandSide.PrimaryAtom(); primaryAtom != nil {
			if identifierExpr, ok := primaryAtom.(*p.IdentifierExprContext); ok {
				funcName = identifierExpr.Identifier().GetText()
			}
		}
	} else {
		// Analyze Stand Alone expression / Traverse the tree until it finds an atom
		current := parent
		for current != nil {
			if standaloneExpr, ok := current.(*p.StandaloneExpresionContext); ok {
				isStandaloneExpression = true
				if standaloneAtom := standaloneExpr.StandaloneAtom(); standaloneAtom != nil {
					if identifierExpr, ok := standaloneAtom.(*p.StandaloneIdentifierExprContext); ok {
						funcName = identifierExpr.Identifier().GetText()
						break
					}
				}
			}
			current = current.GetParent()
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

	// Validate argument types
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

	// Set return type for the expression
	if isStandaloneExpression {
		current := parent
		for current != nil {
			if standaloneExpr, ok := current.(*p.StandaloneExpresionContext); ok {
				exprText := standaloneExpr.GetText()
				actualExpr := exprText[:len(exprText)-1]
				log.Printf("Setting return type of standalone expression `%s` to `%s`", actualExpr, funcInfo.ReturnType)
				l.ScopeManager.CurrentScope.UpsertExpressionType(actualExpr, funcInfo.ReturnType)
				break
			}
			current = current.GetParent()
		}
	} else {
		if leftHandSide, ok := parent.(*p.LeftHandSideContext); ok {
			fullExpr := leftHandSide.GetText()
			log.Printf("Setting return type of `%s` to `%s`", fullExpr, funcInfo.ReturnType)
			l.ScopeManager.CurrentScope.UpsertExpressionType(fullExpr, funcInfo.ReturnType)
		}
	}

	l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), funcInfo.ReturnType)
}

func (l Listener) findFunctionInfo(funcName string) (MethodInfo, bool) {
	log.Printf("Searching for function: %s", funcName)

	// First check current scope and walk up
	scope := l.ScopeManager.CurrentScope
	scopeDepth := 0
	for scope != nil {
		log.Printf("Checking scope: %s (depth %d), functions: %v", scope.Name, scopeDepth, scope.functions)
		if funcInfo, found := scope.functions[funcName]; found {
			log.Printf("Found function %s in scope %s", funcName, scope.Name)
			return funcInfo, true
		}
		scope = scope.Father
		scopeDepth++
	}

	// Check if it's a method in the current class (if we're in a class)
	if classScope, inClass := l.ScopeManager.SearchClassScope(); inClass {
		log.Printf("Checking class scope: %s", classScope.Name)
		if typeInfo, found := l.GetTypeInfo(TypeIdentifier(classScope.Name)); found {
			if typeInfo.ClassType.HasValue() {
				classInfo := typeInfo.ClassType.GetValue()
				if methodInfo, found := classInfo.Methods[funcName]; found {
					log.Printf("Found method %s in class %s", funcName, classScope.Name)
					return methodInfo, true
				}
			}
		}
	}

	log.Printf("Function %s not found anywhere", funcName)
	return MethodInfo{}, false
}

func (l Listener) ExitStandaloneExpresion(ctx *p.StandaloneExpresionContext) {
	line := ctx.GetStart().GetLine()
	colStart := ctx.GetStart().GetColumn()
	colEnd := ctx.GetStop().GetColumn()

	log.Printf("Processing standalone expression: %s", ctx.GetText())

	standaloneAtom := ctx.StandaloneAtom()
	suffixOps := ctx.AllSuffixOp()

	hasCallExpression := false
	for _, suffixOp := range suffixOps {
		if _, ok := suffixOp.(*p.CallExprContext); ok {
			hasCallExpression = true
			break
		}
		if _, ok := suffixOp.(*p.MethodCallExprContext); ok {
			hasCallExpression = true
			break
		}
	}

	if !hasCallExpression {
		switch atom := standaloneAtom.(type) {
		case *p.StandaloneIdentifierExprContext:
			l.AddError(line, colStart, colEnd, fmt.Sprintf(
				"Identifier `%s` cannot be used as a standalone statement",
				atom.Identifier().GetText(),
			))
			return
		case *p.StandaloneNewExprContext:
			log.Printf("Standalone new expression is valid: %s", ctx.GetText())
		case *p.StandaloneThisExprContext:
			l.AddError(line, colStart, colEnd, "Cannot use 'this' as a standalone statement")
			return
		}
	}

	// Check type for stand alone expression
	exprText := ctx.GetText()
	actualExpr := exprText[:len(exprText)-1]

	if hasCallExpression {
		_, found := l.ScopeManager.CurrentScope.GetExpressionType(actualExpr)
		if !found {
			log.Printf("Warning: Standalone expression `%s` has no type information, but this might be normal for void functions", actualExpr)
		}
	}
}

func (l Listener) ExitStandaloneIdentifierExpr(ctx *p.StandaloneIdentifierExprContext) {
	identifier := ctx.Identifier().GetText()
	log.Printf("Processing standalone identifier: %s", identifier)

	identifierType, found := l.ScopeManager.CurrentScope.GetExpressionType(identifier)
	if found {
		parent := ctx.GetParent()
		if standaloneAtom, ok := parent.(*p.StandaloneAtomContext); ok {
			l.ScopeManager.CurrentScope.UpsertExpressionType(standaloneAtom.GetText(), identifierType)
		}
	} else {
		if _, functionFound := l.findFunctionInfo(identifier); functionFound {
			log.Printf("Identifier %s is a function, not throwing error", identifier)
			return
		}

		line := ctx.GetStart().GetLine()
		colStart := ctx.Identifier().GetSymbol().GetColumn()
		colEnd := colStart + len(identifier)
		l.AddError(line, colStart, colEnd, fmt.Sprintf(
			"Undeclared identifier `%s`",
			identifier,
		))
	}
}

func (l Listener) ExitStandaloneNewExpr(ctx *p.StandaloneNewExprContext) {
	line := ctx.GetStart().GetLine()
	className := ctx.Identifier().GetText()
	classColStart := ctx.Identifier().GetSymbol().GetColumn()
	classColEnd := classColStart + len(className)

	log.Printf("Processing standalone new expression: new %s(...)", className)

	// Check if the class exists
	if !l.TypeExists(TypeIdentifier(className)) {
		l.AddError(line, classColStart, classColEnd, fmt.Sprintf(
			"Class `%s` doesn't exist!",
			className,
		))
		return
	}

	// Get class info and validate constructor call
	typeInfo, found := l.GetTypeInfo(TypeIdentifier(className))
	if !found {
		l.AddError(line, classColStart, classColEnd, fmt.Sprintf(
			"Type information for class `%s` not found!",
			className,
		))
		return
	}

	if !typeInfo.ClassType.HasValue() {
		l.AddError(line, classColStart, classColEnd, fmt.Sprintf(
			"`%s` is not a class type!",
			className,
		))
		return
	}

	classInfo := typeInfo.ClassType.GetValue()

	// Validate constructor arguments
	exprArguments := []p.IConditionalExprContext{}
	if ctx.Arguments() != nil {
		exprArguments = ctx.Arguments().AllConditionalExpr()
	}

	constructor := classInfo.GetConstructor(&l)
	if len(constructor.ParameterList) != len(exprArguments) {
		l.AddError(line, classColStart, classColEnd, fmt.Sprintf(
			"Constructor for `%s` expects %d arguments but %d given",
			className,
			len(constructor.ParameterList),
			len(exprArguments),
		))
		return
	}

	// Validate argument types
	for i, param := range constructor.ParameterList {
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
				"Constructor parameter `%s` expects type `%s` but got `%s`",
				param.Name,
				param.Type,
				argType,
			))
		}
	}

	// Set the type for this expression
	l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), TypeIdentifier(className))

	// Set type for parent standalone atom
	parent := ctx.GetParent()
	if standaloneAtom, ok := parent.(*p.StandaloneAtomContext); ok {
		l.ScopeManager.CurrentScope.UpsertExpressionType(standaloneAtom.GetText(), TypeIdentifier(className))
	}
}

func (l Listener) ExitStandaloneThisExpr(ctx *p.StandaloneThisExprContext) {
	log.Printf("Processing standalone this expression")

	classScope, isInsideClass := l.ScopeManager.SearchClassScope()
	if !isInsideClass {
		line := ctx.GetStart().GetLine()
		colStart := ctx.GetStart().GetColumn()
		colEnd := colStart + 4 // length of "this"
		l.AddError(line, colStart, colEnd, "Cannot use 'this' outside of class scope")
		return
	}

	thisType := TypeIdentifier(classScope.Name)
	l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), thisType)

	parent := ctx.GetParent()
	if standaloneAtom, ok := parent.(*p.StandaloneAtomContext); ok {
		l.ScopeManager.CurrentScope.UpsertExpressionType(standaloneAtom.GetText(), thisType)
	}
}

func (l *Listener) checkAllPathsReturn(ctx *p.FunctionDeclarationContext) bool {
	if ctx.Block() == nil {
		return false
	}

	return l.blockAllPathsReturn(ctx.Block())
}

func (l *Listener) blockAllPathsReturn(blockCtx p.IBlockContext) bool {
	if blockCtx == nil {
		return false
	}

	statements := blockCtx.AllStatement()
	if len(statements) == 0 {
		return false
	}

	for i, stmt := range statements {
		if l.statementReturns(stmt) {
			if i < len(statements)-1 {
				nextStmt := statements[i+1]
				line := nextStmt.GetStart().GetLine()
				colStart := nextStmt.GetStart().GetColumn()
				colEnd := nextStmt.GetStop().GetColumn()
				l.AddError(line, colStart, colEnd,
					"Unreachable code detected after return statement",
					fmt.Sprintf("line %d", line),
				)
			}
			return true
		}

		if l.statementAllPathsReturn(stmt) {
			if i < len(statements)-1 {
				nextStmt := statements[i+1]
				line := nextStmt.GetStart().GetLine()
				colStart := nextStmt.GetStart().GetColumn()
				colEnd := nextStmt.GetStop().GetColumn()
				l.AddError(line, colStart, colEnd,
					"Unreachable code detected",
					fmt.Sprintf("line %d", line),
				)
			}
			return true
		}
	}

	return false
}

func (l *Listener) statementReturns(stmt p.IStatementContext) bool {
	return stmt.ReturnStatement() != nil
}

func (l *Listener) statementAllPathsReturn(stmt p.IStatementContext) bool {
	// Check for direct return statement
	if stmt.ReturnStatement() != nil {
		return true
	}

	// Check if statement
	if ifStmt := stmt.IfStatement(); ifStmt != nil {
		// An if statement guarantees return only if:
		// Both if and else bodies exist and both return

		ifReturns := l.blockAllPathsReturn(ifStmt.IfBody().Block())

		if ifStmt.ElseBody() != nil {
			elseReturns := l.blockAllPathsReturn(ifStmt.ElseBody().Block())
			return ifReturns && elseReturns
		}

		// If there's no else branch, not all paths are guaranteed to return
		return false
	}

	// Check while statement
	if stmt.WhileStatement() != nil {
		// While loops don't guarantee return (could never execute)
		return false
	}

	// Check do-while statement
	if stmt.DoWhileStatement() != nil {
		// Do-while executes at least once, but still doesn't guarantee return
		// (could break out without returning)
		return false
	}

	// Check for statement
	if stmt.ForStatement() != nil {
		// For loops don't guarantee return
		return false
	}

	// Check switch statement
	if switchStmt := stmt.SwitchStatement(); switchStmt != nil {
		// Switch guarantees return only if:
		// All cases return AND there's a default case that returns

		hasDefault := switchStmt.DefaultCase() != nil
		if !hasDefault {
			return false // Without default, not all paths are covered
		}

		// Check if all cases return
		for _, switchCase := range switchStmt.AllSwitchCase() {
			if !l.caseBodyAllPathsReturn(switchCase.CaseBody()) {
				return false
			}
		}

		// Check if default case returns
		if hasDefault && !l.caseBodyAllPathsReturn(switchStmt.DefaultCase().CaseBody()) {
			return false
		}

		return hasDefault // Only if we have default and all cases return
	}

	// Check block statement
	if blockStmt := stmt.BlockStatement(); blockStmt != nil {
		return l.blockAllPathsReturn(blockStmt.Block())
	}

	// Check try statement
	if tryStmt := stmt.TryStatement(); tryStmt != nil {
		// Try-catch guarantees return only if both try and catch blocks return
		tryReturns := l.blockAllPathsReturn(tryStmt.Block())
		catchReturns := l.blockAllPathsReturn(tryStmt.CatchStatement().Block())
		return tryReturns && catchReturns
	}

	return false
}

func (l *Listener) caseBodyAllPathsReturn(caseBody p.ICaseBodyContext) bool {
	statements := caseBody.AllStatement()
	if len(statements) == 0 {
		return false
	}

	for _, stmt := range statements {
		if l.statementReturns(stmt) || l.statementAllPathsReturn(stmt) {
			return true
		}
	}

	return false
}
