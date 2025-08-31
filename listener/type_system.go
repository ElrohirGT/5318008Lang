package listener

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

// METHODS FOR HANDLING TYPE CHECKING

func (l Listener) ExitAdditiveExpr(ctx *p.AdditiveExprContext) {
	line := ctx.GetStart().GetLine()
	exprs := ctx.AllMultiplicativeExpr()
	firstExpr := exprs[0]

	colStart := firstExpr.GetStart().GetColumn()
	colEnd := colStart + len(firstExpr.GetText())

	referenceType, available := l.ScopeManager.CurrentScope.GetExpressionType(firstExpr.GetText())
	if !available {
		l.AddError(line, colStart, colEnd, fmt.Sprintf("Exit Add (firstExpr): `%s` doesn't have a type!", firstExpr.GetText()))
		return
	}

	defineLater := make([]p.IMultiplicativeExprContext, 0, len(exprs))

	if referenceType == BASE_TYPES.UNKNOWN {
		defineLater = append(defineLater, firstExpr)

		for _, expr := range exprs[1:] {
			exprType, available := l.ScopeManager.CurrentScope.GetExpressionType(expr.GetText())
			exprColStart := expr.GetStart().GetColumn()
			exprColEnd := exprColStart + len(expr.GetText())

			if !available {
				l.AddError(line, exprColStart, exprColEnd, fmt.Sprintf("Exit Add (expr): `%s` doesn't have a type!", exprType))
			}

			if exprType != BASE_TYPES.UNKNOWN {
				referenceType = exprType
				break
			}
		}

		if referenceType == BASE_TYPES.UNKNOWN {
			l.AddError(line, colStart, colEnd, "Can't infer the value of multiplication/addition! All types were unknown...")
			return
		}
	}

	for i, expr := range exprs[1:] {
		exprType, available := l.ScopeManager.CurrentScope.GetExpressionType(expr.GetText())
		exprColStart := expr.GetStart().GetColumn()
		exprColEnd := exprColStart + len(expr.GetText())

		if !available {
			l.AddError(line, exprColStart, exprColEnd, fmt.Sprintf("Exit Add (expr - second): `%s` doesn't have a type!", exprType))
		}

		if exprType == BASE_TYPES.UNKNOWN {
			defineLater = append(defineLater, expr)
			continue
		}

		if exprType != referenceType {
			stream := ctx.GetStart().GetInputStream()
			leftStart := exprs[0].GetStart().GetColumn()
			leftEnd := exprs[i].GetStop().GetColumn() + 1
			l.AddError(line,
				leftStart,
				leftEnd,
				"Can't add:",
				fmt.Sprintf("leftSide: `%s` of type `%s`",
					stream.GetText(exprs[0].GetStart().GetStart(), exprs[i].GetStart().GetStop()),
					referenceType),
				fmt.Sprintf("rightSide: `%s` of type `%s`",
					ctx.MultiplicativeExpr(i+1).GetText(),
					exprType,
				),
			)
		}
	}

	classScope, isInsideClassDeclaration := l.ScopeManager.SearchClassScope()
	for _, expr := range defineLater {
		log.Printf("Inferring `%s` as type of `%s`\n", expr.GetText(), referenceType)
		// FIXME: There should be a better way to manage `this.`
		exprStartsWithThis := strings.HasPrefix(expr.GetText(), "this.")
		if isInsideClassDeclaration && exprStartsWithThis {
			fieldName := expr.GetText()[len("this."):]
			l.ModifyClassTypeInfo(TypeIdentifier(classScope.Name), func(cti *ClassTypeInfo) {
				cti.UpsertField(fieldName, referenceType)
			})
		}
		l.ScopeManager.CurrentScope.UpsertExpressionType(expr.GetText(), referenceType)
	}

	log.Printf("Adding expression `%s` of type `%s`", ctx.GetText(), referenceType)
	l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), referenceType)
}

func (l Listener) ExitLiteralExpr(ctx *p.LiteralExprContext) {
	strRepresentation := ctx.GetText()
	switch strRepresentation {
	case "null":
		l.ScopeManager.CurrentScope.UpsertExpressionType(strRepresentation, BASE_TYPES.NULL)
	case "true", "false":
		l.ScopeManager.CurrentScope.UpsertExpressionType(strRepresentation, BASE_TYPES.BOOLEAN)
	default:
		literal := ctx.Literal()
		if literal != nil {
			literalExpr := literal.GetText()
			_, err := strconv.ParseInt(literalExpr, 10, 64)
			if err != nil {
				log.Println("Adding", literalExpr, "as an expresion of type", BASE_TYPES.STRING)
				l.ScopeManager.CurrentScope.UpsertExpressionType(literalExpr, BASE_TYPES.STRING)
			} else {
				log.Println("Adding", literalExpr, "as an expresion of type", BASE_TYPES.INTEGER)
				l.ScopeManager.CurrentScope.UpsertExpressionType(literalExpr, BASE_TYPES.INTEGER)
			}
		}
	}
}

func (l Listener) ExitVariableDeclaration(ctx *p.VariableDeclarationContext) {
	line := ctx.GetStart().GetLine()
	name := ctx.Identifier()
	typeAnnot := ctx.TypeAnnotation()
	declarationExpr := ctx.Initializer()

	// Possible errors
	// Error in identifier
	colStartI := name.GetSymbol().GetColumn()
	colEndI := colStartI + len(name.GetText())

	// Error in Type
	var colStartT, colEndT int
	if typeAnnot != nil {
		colStartT = typeAnnot.GetStart().GetColumn()
		colEndT = colStartT + len(typeAnnot.GetText())
	}

	hasAnnotation := typeAnnot != nil

	hasInitialExpr := declarationExpr != nil

	isInsideClassDeclaration := l.ScopeManager.CurrentScope.Type == SCOPE_TYPES.CLASS

	if !hasAnnotation {
		log.Println("Variable", name.GetText(), "does NOT have a type! We need to infer it...")
		if hasInitialExpr {
			declarationText := declarationExpr.Expression().GetText()
			inferedType, found := l.ScopeManager.CurrentScope.GetExpressionType(declarationText)
			if !found {
				l.AddError(line, colStartI, colEndI, fmt.Sprintf(
					"Couldn't infer the type of variable `%s`, initialized with: `%s`",
					name.GetText(),
					declarationText,
				))
			} else {
				if isInsideClassDeclaration {
					newExprName := "this." + name.GetText()
					log.Printf("Inferring type of `%s` as `%s`\n", newExprName, inferedType)
					l.ModifyClassTypeInfo(TypeIdentifier(l.ScopeManager.CurrentScope.Name), func(cti *ClassTypeInfo) {
						cti.UpsertField(name.GetText(), inferedType)
					})
					l.ScopeManager.CurrentScope.UpsertExpressionType(newExprName, inferedType)
				} else {
					log.Printf("Inferring type of `%s` as `%s`\n", name.GetText(), inferedType)
					l.ScopeManager.CurrentScope.UpsertExpressionType(name.GetText(), inferedType)
				}
			}
		} else {
			if isInsideClassDeclaration {
				l.ModifyClassTypeInfo(TypeIdentifier(l.ScopeManager.CurrentScope.Name), func(cti *ClassTypeInfo) {
					cti.UpsertField(name.GetText(), BASE_TYPES.UNKNOWN)
				})
				l.ScopeManager.CurrentScope.UpsertExpressionType("this."+name.GetText(), BASE_TYPES.UNKNOWN)
			} else {
				l.ScopeManager.CurrentScope.UpsertExpressionType(name.GetText(), BASE_TYPES.UNKNOWN)
			}
		}
	} else {
		declarationType := TypeIdentifier(typeAnnot.Type_().GetText())

		log.Println("Variable", name.GetText(), "has type", declarationType)

		if !l.TypeExists(declarationType) {
			l.AddError(line, colStartT, colEndT, fmt.Sprintf(
				"%s doesn't exist!",
				declarationType,
			))
		}

		if hasInitialExpr {
			declExpr := declarationExpr.Expression() // solo la expresión del valor
			colStartD := declExpr.GetStart().GetColumn()
			colEndD := colStartD + len(declExpr.GetText())

			exprText := declarationExpr.Expression().GetText()
			log.Println("Known expressions", l.ScopeManager.CurrentScope.typesByExpression)
			initialExprType, exists := l.ScopeManager.CurrentScope.GetExpressionType(exprText)
			if !exists {
				l.AddError(line, colStartI, colStartT, fmt.Sprintf(
					"Variable Declaration: `%s` doesn't have a type!",
					exprText,
				))
			}

			if initialExprType != declarationType {
				l.AddError(line, colStartD, colEndD, fmt.Sprintf(
					"The declaration of `%s` specifies a type of `%s` but `%s` was given",
					name,
					declarationType,
					initialExprType,
				))
			}
		}

		if isInsideClassDeclaration {
			l.ModifyClassTypeInfo(TypeIdentifier(l.ScopeManager.CurrentScope.Name), func(cti *ClassTypeInfo) {
				cti.UpsertField(name.GetText(), declarationType)
			})
			l.ScopeManager.CurrentScope.UpsertExpressionType("this."+name.GetText(), declarationType)
		} else {
			l.ScopeManager.CurrentScope.UpsertExpressionType(name.GetText(), declarationType)
		}
	}
}

func (l Listener) ExitAssignment(ctx *p.AssignmentContext) {
	log.Println("General Assignment!", ctx.GetText())
	line := ctx.GetStart().GetLine()

	isPropertyAssignment := len(ctx.AllExpression()) > 1
	classScope, _ := l.ScopeManager.SearchClassScope()
	if isPropertyAssignment {
		firstExpr := ctx.Expression(0)
		colStartF := firstExpr.GetStart().GetColumn()
		colEndF := colStartF + len(firstExpr.GetText())

		t, found := l.ScopeManager.CurrentScope.GetExpressionType(firstExpr.GetText())
		if !found {
			l.AddError(line, colStartF, colEndF, fmt.Sprintf(
				"Undeclared variable `%s`",
				firstExpr.GetText(),
			))
			return
		}

		info, found := l.GetTypeInfo(t)
		if !found {
			l.AddError(line, colStartF, colEndF, fmt.Sprintf(
				"Undeclared type `%s` for variable `%s`",
				t,
				firstExpr.GetText(),
			))
			return
		}

		if !info.ClassType.HasValue() {
			l.AddError(line, colStartF, colEndF, fmt.Sprintf(
				"Trying to access a field `%s` from type `%s` but `%s` is not a class!",
				firstExpr.GetText(),
				t,
				firstExpr.GetText(),
			))
			return
		}

		identifier := ctx.Identifier()
		colStartI := identifier.GetSymbol().GetColumn()
		colEndI := colStartI + len(identifier.GetText())
		classInfo := info.ClassType.GetValue()
		fieldType, hasField := classInfo.Fields[identifier.GetText()]
		if !hasField {
			l.AddError(line, colStartI, colEndI, fmt.Sprintf(
				"Trying to access field `%s` not defined in class `%s`!",
				identifier.GetText(),
				classInfo.Name,
			))
			return
		}

		assignExpr := ctx.Expression(1)
		colStartA := assignExpr.GetStart().GetColumn()
		colEndA := colStartA + len(assignExpr.GetText())
		assignType, found := l.ScopeManager.CurrentScope.GetExpressionType(assignExpr.GetText())
		if !found {
			l.AddError(line, colStartA, colEndA, fmt.Sprintf(
				"Type of expression `%s` not found!",
				assignExpr.GetText(),
			))
			return
		}

		if fieldType == BASE_TYPES.UNKNOWN && assignType == BASE_TYPES.UNKNOWN {
			l.AddError(line, colStartI, colEndI, fmt.Sprintf(
				"Trying to assign `%s` into `%s` but I don't know the types of both! Please give me hints!",
				assignExpr.GetText(),
				identifier.GetText(),
			))
			return
		}

		if fieldType == BASE_TYPES.UNKNOWN {
			log.Printf("Inferring `%s` as type `%s`", "this."+identifier.GetText(), assignType)
			classScope.UpsertExpressionType("this."+identifier.GetText(), assignType)
			l.ModifyClassTypeInfo(TypeIdentifier(classScope.Name), func(cti *ClassTypeInfo) {
				cti.UpsertField(identifier.GetText(), assignType)
			})
			fieldType = assignType
		}

		if fieldType != assignType {
			l.AddError(line, colStartI, colEndI, fmt.Sprintf(
				"Trying to assign `%s` to field `%s` but types don't match! (`%s` != `%s`)",
				assignExpr.GetText(),
				identifier.GetText(),
				fieldType,
				assignType,
			))
			return
		}
	} else {
		identifier := ctx.Identifier()
		colStartI := identifier.GetSymbol().GetColumn()
		colEndI := colStartI + len(identifier.GetText())

		identifierType, found := l.ScopeManager.CurrentScope.GetExpressionType(identifier.GetText())
		if !found {
			l.AddError(line, colStartI, colEndI, fmt.Sprintf(
				"Undeclared variable `%s`!",
				identifier.GetText(),
			))
			return
		}

		assignExpr := ctx.Expression(0)
		colStartA := assignExpr.GetStart().GetColumn()
		colEndA := colStartA + len(assignExpr.GetText())

		assignType, found := l.ScopeManager.CurrentScope.GetExpressionType(assignExpr.GetText())
		if !found {
			l.AddError(line, colStartA, colEndA, fmt.Sprintf(
				"Expression `%s` doesn't have a type!",
				assignExpr,
			))
		}

		if identifierType == BASE_TYPES.UNKNOWN && assignType == BASE_TYPES.UNKNOWN {
			l.AddError(line, colStartI, colEndA, fmt.Sprintf(
				"Can't assign `%s` = `%s` because both type are unknown! Use one of them first or write type hints!",
				identifier.GetText(),
				assignExpr.GetText(),
			))
			return
		}

		if identifierType == BASE_TYPES.UNKNOWN {
			log.Printf("Inferring type of `%s` as `%s`\n", identifier.GetText(), assignType)
			l.ScopeManager.CurrentScope.UpsertExpressionType(identifier.GetText(), assignType)
			identifierType = assignType
		}

		if identifierType != assignType {
			l.AddError(line, colStartI, colEndA, fmt.Sprintf(
				"Trying to assign `%s` to variable `%s` but types don't match! (`%s` != `%s`)",
				assignExpr.GetText(),
				identifier.GetText(),
				identifierType,
				assignType,
			))
			return
		}
	}
}

func (l Listener) EnterFunctionDeclaration(ctx *p.FunctionDeclarationContext) {
	// FIXME: ONLY CLASS METHODS IMPLEMENTED! PLEASE IMPLEMENT GENERAL METHODS!
	// For example, we need to infer the return type and compare it to the defined return type (if there's any)
	// remember that unknown exists for type inference!
	line := ctx.GetStart().GetLine()
	funcName := ctx.Identifier()

	funcParams := []ParameterInfo{}
	if ctx.Parameters() != nil {
		for _, paramCtx := range ctx.Parameters().AllParameter() {
			name := paramCtx.Identifier()
			paramType := BASE_TYPES.UNKNOWN
			typeColStart := paramCtx.Type_().GetStart().GetColumn()
			typeColEnd := typeColStart + len(paramCtx.Type_().GetText())
			if paramCtx.Type_() != nil {
				paramType = TypeIdentifier(paramCtx.Type_().GetText())
			}

			if !l.TypeExists(paramType) {
				l.AddError(line, typeColStart, typeColEnd, fmt.Sprintf(
					"Parameter type `%s` doesn't exist!",
					paramType,
				))
			} else {
				funcParams = append(funcParams, ParameterInfo{
					Name: name.GetText(),
					Type: paramType,
				})
			}
		}
	}

	returnType := BASE_TYPES.UNKNOWN
	if ctx.Type_() != nil {
		returnType = TypeIdentifier(ctx.Type_().GetText())
	}
	retColStart := ctx.Type_().GetStart().GetColumn()
	retColEnd := retColStart + len(ctx.Type_().GetText())

	if !l.TypeExists(returnType) {
		l.AddError(line, retColStart, retColEnd, fmt.Sprintf(
			"Return type `%s` doesn't exist!",
			returnType,
		))
		returnType = BASE_TYPES.INVALID
	}

	info := MethodInfo{
		ParameterList: funcParams,
		ReturnType:    returnType,
	}
	funcScope := NewScope(funcName.GetText(), SCOPE_TYPES.FUNCTION)
	isInsideClassDeclaration := l.ScopeManager.CurrentScope.Type == SCOPE_TYPES.CLASS
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
	}

	for _, param := range info.ParameterList {
		funcScope.UpsertExpressionType(param.Name, param.Type)
	}
	l.ScopeManager.CurrentScope.AddChildScope(funcScope)
	l.ScopeManager.ReplaceCurrent(funcScope)
}

func (l Listener) ExitFunctionDeclaration(ctx *p.FunctionDeclarationContext) {
	// FIXME: Check if the last return found inside this function matches the function type?
	if l.ScopeManager.CurrentScope.Type != SCOPE_TYPES.FUNCTION {
		log.Panicf("Trying to exit function scope but scope is not of type function! %#v", l.ScopeManager.CurrentScope)
	}

	l.ScopeManager.ReplaceWithParent()
}

func (s Listener) ExitLeftHandSide(ctx *p.LeftHandSideContext) {
	// FIXME: Methods and functions are called here!
	// For example `cell.setRow(15)` will be evaluated here!
	log.Println("LEFT HAND SIDE:", ctx.GetText())
	log.Println("CHILD", ctx.PrimaryAtom().GetChild(0))
	suffixes := ctx.AllSuffixOp()

	if len(suffixes) == 0 {
		// simple expression, no further evaluating is required
		return
	}
	// for _, suffixCtx := range suffixes {
	// 	// FIXME: Implement me!
	// }
}

func (l Listener) ExitReturnStatement(ctx *p.ReturnStatementContext) {
	line := ctx.GetStart().GetLine()
	colStart := ctx.GetStart().GetColumn()
	colEnd := colStart + len(ctx.GetText())

	if _, inFunctionScope := l.ScopeManager.SearchScopeByType(SCOPE_TYPES.FUNCTION); !inFunctionScope {
		l.AddError(line, colStart, colEnd, "'return' statement out function scope.")
	}
}

func (l Listener) ExitPrimaryExpr(ctx *p.PrimaryExprContext) {
	// Primary expresion is of type conditional just pass inherit its type
	expr := ctx.ConditionalExpr()
	if expr == nil {
		return
	}
	referenceType, _ := l.ScopeManager.CurrentScope.GetExpressionType(expr.GetText())

	l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), referenceType)
}
