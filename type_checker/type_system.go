package type_checker

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

// METHODS FOR HANDLING TYPE CHECKING

func (l Listener) ExitAdditiveExpr(ctx *p.AdditiveExprContext) {
	exprs := ctx.AllMultiplicativeExpr()
	if len(exprs) <= 1 {
		return
	}

	firstExpr := exprs[0]
	line := ctx.GetStart().GetLine()

	colStart := firstExpr.GetStart().GetColumn()
	colEnd := colStart + len(firstExpr.GetText())

	referenceType, available := l.ScopeManager.CurrentScope.GetExpressionType(firstExpr.GetText())
	if !available {
		l.AddError(line, colStart, colEnd, fmt.Sprintf("Exit Add (firstExpr): `%s` doesn't have a type!", firstExpr.GetText()))
		return
	}

	previousExprStart := firstExpr.GetStart().GetStart()
	previousExprEnd := firstExpr.GetStop().GetStop()

	for i, expr := range exprs[1:] {
		exprType, available := l.ScopeManager.CurrentScope.GetExpressionType(expr.GetText())
		exprColStart := expr.GetStart().GetColumn()
		exprColEnd := exprColStart + len(expr.GetText())

		if !available {
			l.AddError(line, exprColStart, exprColEnd, fmt.Sprintf("Exit Add (expr - second): `%s` doesn't have a type!", exprType))
		}

		if exprType != referenceType {
			stream := ctx.GetStart().GetInputStream()
			leftStart := expr.GetStart().GetColumn()
			leftEnd := ctx.MultiplicativeExpr(i + 1).GetStop().GetColumn()
			l.AddError(line,
				leftStart,
				leftEnd,
				"Can't add:",
				fmt.Sprintf("leftSide: `%s` of type `%s`",
					stream.GetText(previousExprStart, previousExprEnd),
					referenceType,
				),
				fmt.Sprintf("rightSide: `%s` of type `%s`",
					ctx.MultiplicativeExpr(i+1).GetText(),
					exprType,
				),
			)
		}

		previousExprStart = expr.GetStart().GetStart()
		previousExprEnd = expr.GetStop().GetStop()
	}

	log.Printf("Adding expression `%s` of type `%s`", ctx.GetText(), referenceType)
	l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), referenceType)
}

func (l Listener) ExitMultiplicativeExpr(ctx *p.MultiplicativeExprContext) {
	exprs := ctx.AllUnaryExpr()
	if len(exprs) <= 1 {
		// If there's only one expression, just inherit its type
		if len(exprs) == 1 {
			firstExpr := exprs[0]
			referenceType, available := l.ScopeManager.CurrentScope.GetExpressionType(firstExpr.GetText())
			if available {
				l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), referenceType)
			}
		}
		return
	}

	firstExpr := exprs[0]
	line := ctx.GetStart().GetLine()

	colStart := firstExpr.GetStart().GetColumn()
	colEnd := colStart + len(firstExpr.GetText())

	referenceType, available := l.ScopeManager.CurrentScope.GetExpressionType(firstExpr.GetText())
	if !available {
		l.AddError(line, colStart, colEnd, fmt.Sprintf("Exit Multiply (firstExpr): `%s` doesn't have a type!", firstExpr.GetText()))
		return
	}

	// Check that it's a numeric type (integer) for arithmetic operations
	if referenceType != BASE_TYPES.INTEGER {
		l.AddError(line, colStart, colEnd, fmt.Sprintf("Multiplicative operations require integer operands, but got `%s`", referenceType))
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
		return
	}

	previousExprStart := firstExpr.GetStart().GetStart()
	previousExprEnd := firstExpr.GetStop().GetStop()

	for i, expr := range exprs[1:] {
		exprType, available := l.ScopeManager.CurrentScope.GetExpressionType(expr.GetText())
		exprColStart := expr.GetStart().GetColumn()
		exprColEnd := exprColStart + len(expr.GetText())

		if !available {
			l.AddError(line, exprColStart, exprColEnd, fmt.Sprintf("Exit Multiply (expr - operand): `%s` doesn't have a type!", expr.GetText()))
			l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
			return
		}

		if exprType != referenceType {
			stream := ctx.GetStart().GetInputStream()
			leftStart := expr.GetStart().GetColumn()
			leftEnd := ctx.UnaryExpr(i + 1).GetStop().GetColumn()
			l.AddError(line,
				leftStart,
				leftEnd,
				"Can't perform multiplicative operation:",
				fmt.Sprintf("leftSide: `%s` of type `%s`",
					stream.GetText(previousExprStart, previousExprEnd),
					referenceType,
				),
				fmt.Sprintf("rightSide: `%s` of type `%s`",
					ctx.UnaryExpr(i+1).GetText(),
					exprType,
				),
			)
			l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
			return
		}

		if i < len(ctx.AllUnaryExpr()) && ctx.AllUnaryExpr() != nil {
			if expr.GetText() == "0" {
				l.AddError(line, exprColStart, exprColEnd, "Division by zero is not allowed")
				l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
				return
			}
		}

		previousExprStart = expr.GetStart().GetStart()
		previousExprEnd = expr.GetStop().GetStop()
	}

	log.Printf("Setting multiplicative expression `%s` to type `%s`", ctx.GetText(), referenceType)
	l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), referenceType)
}

func (l Listener) ExitLiteralExpr(ctx *p.LiteralExprContext) {
	strRepresentation := ctx.GetText()
	switch strRepresentation {
	case "null":
		l.ScopeManager.CurrentScope.UpsertExpressionType(strRepresentation, BASE_TYPES.NULL)
		l.UpsertLiteral(ctx.GetText(), BASE_TYPES.NULL)
	case "true", "false":
		l.ScopeManager.CurrentScope.UpsertExpressionType(strRepresentation, BASE_TYPES.BOOLEAN)
		l.UpsertLiteral(ctx.GetText(), BASE_TYPES.BOOLEAN)
	default:
		literal := ctx.Literal()
		if literal != nil {
			literalExpr := literal.GetText()
			_, err := strconv.ParseInt(literalExpr, 10, 64)
			if err != nil {
				log.Println("Adding", literalExpr, "as an expresion of type", BASE_TYPES.STRING)
				l.ScopeManager.CurrentScope.UpsertExpressionType(literalExpr, BASE_TYPES.STRING)
				// NOTE: Since strings are actually u8 arrays. They're not really a literal.
				// l.UpsertLiteral(literalExpr, BASE_TYPES.STRING)
			} else {
				log.Println("Adding", literalExpr, "as an expresion of type", BASE_TYPES.INTEGER)
				l.ScopeManager.CurrentScope.UpsertExpressionType(literalExpr, BASE_TYPES.INTEGER)
				l.UpsertLiteral(literalExpr, BASE_TYPES.INTEGER)
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
		if !hasInitialExpr {
			l.AddError(line, colStartT, colEndI,
				fmt.Sprintf("Variable `%s` must have an initial expression assigned to it! Maybe try assigning a value or a type?", name.GetText()),
			)

			if isInsideClassDeclaration {
				classExprName := "this." + name.GetText()
				l.ScopeManager.CurrentScope.UpsertExpressionType(classExprName, BASE_TYPES.INVALID)
				l.ModifyClassTypeInfo(TypeIdentifier(l.ScopeManager.CurrentScope.Name), func(cti *ClassTypeInfo) {
					cti.UpsertField(name.GetText(), BASE_TYPES.INVALID)
				})
			} else {
				l.ScopeManager.CurrentScope.UpsertExpressionType(name.GetText(), BASE_TYPES.INVALID)
			}
		} else {
			declarationText := declarationExpr.ConditionalExpr().GetText()

			inferedType := BASE_TYPES.INVALID
			if declarationText == "[]" {
				l.AddError(
					line,
					colStartI,
					colEndI,
					"Can't initialize empty array variable without a type!",
				)
			} else {
				t, found := l.ScopeManager.CurrentScope.GetExpressionType(declarationText)
				if !found {
					l.AddError(line, colStartI, colEndI, fmt.Sprintf(
						"Couldn't infer the type of variable `%s`, initialized with: `%s`",
						name.GetText(),
						declarationText,
					))
				} else {
					inferedType = t
				}
			}

			if isInsideClassDeclaration {
				newExprName := "this." + name.GetText()
				log.Printf("Inferring type of `%s` as `%s`\n", newExprName, inferedType)
				l.ModifyClassTypeInfo(TypeIdentifier(l.ScopeManager.CurrentScope.Name), func(cti *ClassTypeInfo) {
					cti.UpsertField(name.GetText(), inferedType)
				})
				l.ScopeManager.CurrentScope.UpsertExpressionType(newExprName, inferedType)
			} else {
				_, found := l.ScopeManager.CurrentScope.GetOnlyInScope(name.GetText())
				if found {
					l.AddError(
						line,
						colStartI,
						colEndI,
						fmt.Sprintf("Can't redeclare variable `%s` in the same scope!", name.GetText()),
					)
				} else {
					log.Printf("Inferring type of `%s` as `%s`\n", name.GetText(), inferedType)
					l.ScopeManager.CurrentScope.UpsertExpressionType(name.GetText(), inferedType)
				}
			}
		}
	} else {
		annotText := typeAnnot.Type_().GetText()
		isArray := strings.HasSuffix(annotText, "[]")

		// FIXME: We shouldn't be making string manipulation here!
		declarationType := TypeIdentifier(annotText)
		if !isArray {
			log.Println("Variable", name.GetText(), "has type", declarationType)

			if !l.TypeExists(declarationType) {
				// FIXME: Make error more specific
				l.AddError(line, colStartT, colEndT, fmt.Sprintf(
					"%s doesn't exist!",
					declarationType,
				))
				declarationType = BASE_TYPES.INVALID
			}
		} else {
			baseType, _, _ := strings.Cut(annotText, "[]")
			log.Println("Array variable", name.GetText(), "has base type", baseType)

			if !l.TypeExists(TypeIdentifier(baseType)) {
				// FIXME: Make error more specific
				l.AddError(line, colStartT, colEndT,
					baseType+" doesn't exist!",
				)
				declarationType = BASE_TYPES.INVALID
			} else {
				declarationType = NewArrayTypeIdentifier(TypeIdentifier(annotText[:len(annotText)-2]))
			}
		}

		if hasInitialExpr {
			declExpr := declarationExpr.ConditionalExpr() // solo la expresión del valor
			colStartD := declExpr.GetStart().GetColumn()
			colEndD := colStartD + len(declExpr.GetText())

			exprText := declarationExpr.ConditionalExpr().GetText()
			if exprText != "[]" {
				log.Println("Known expressions", l.ScopeManager.CurrentScope.typesByExpression)
				initialExprType, exists := l.ScopeManager.CurrentScope.GetExpressionType(exprText)
				if !exists {
					l.AddError(line, colStartI, colStartT, fmt.Sprintf(
						"Variable Declaration: `%s` doesn't have a type!",
						exprText,
					))
					initialExprType = BASE_TYPES.INVALID
				}

				if initialExprType != declarationType && declarationType != BASE_TYPES.INVALID {
					l.AddError(line, colStartD, colEndD, fmt.Sprintf(
						"The declaration of `%s` specifies a type of `%s` but `%s` was given",
						name,
						declarationType,
						initialExprType,
					))
				}
			}
		}

		if isInsideClassDeclaration {
			l.ModifyClassTypeInfo(TypeIdentifier(l.ScopeManager.CurrentScope.Name), func(cti *ClassTypeInfo) {
				cti.UpsertField(name.GetText(), declarationType)
			})
			l.ScopeManager.CurrentScope.UpsertExpressionType("this."+name.GetText(), declarationType)
		} else {
			_, found := l.ScopeManager.CurrentScope.GetOnlyInScope(name.GetText())
			if found {
				l.AddError(
					line,
					colStartI,
					colEndI,
					fmt.Sprintf("Can't redeclare variable `%s` in the same scope!", name.GetText()),
				)
			} else {
				l.ScopeManager.CurrentScope.UpsertExpressionType(name.GetText(), declarationType)
			}
		}
	}
}

func (l Listener) ExitConstantDeclaration(ctx *p.ConstantDeclarationContext) {
	line := ctx.GetStart().GetLine()
	name := ctx.Identifier()
	typeAnnot := ctx.TypeAnnotation()
	declarationExpr := ctx.ConditionalExpr()

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
	isInsideClassDeclaration := l.ScopeManager.CurrentScope.Type == SCOPE_TYPES.CLASS

	if !hasAnnotation {
		log.Println("Constant", name.GetText(), "does NOT have a type! We need to infer it...")
		declarationText := declarationExpr.GetText()

		inferedType := BASE_TYPES.INVALID
		if declarationText == "[]" {
			l.AddError(
				line,
				colStartI,
				colEndI,
				"Can't initialize empty array constant without a type!",
			)
		} else {
			t, found := l.ScopeManager.CurrentScope.GetExpressionType(declarationText)
			if !found {
				l.AddError(line, colStartI, colEndI, fmt.Sprintf(
					"Couldn't infer the type of variable `%s`, initialized with: `%s`",
					name.GetText(),
					declarationText,
				))
				inferedType = BASE_TYPES.INVALID
			} else {
				inferedType = t
			}
		}

		if isInsideClassDeclaration {
			newExprName := "this." + name.GetText()
			log.Printf("Inferring type of `%s` as `%s`\n", newExprName, inferedType)
			l.ModifyClassTypeInfo(TypeIdentifier(l.ScopeManager.CurrentScope.Name), func(cti *ClassTypeInfo) {
				cti.UpsertField(name.GetText(), inferedType)
				cti.ConstantFields.Add(name.GetText())
			})
			l.ScopeManager.CurrentScope.UpsertExpressionType(newExprName, inferedType)
			l.ScopeManager.CurrentScope.AddConstant(newExprName)
		} else {
			_, found := l.ScopeManager.CurrentScope.GetOnlyInScope(name.GetText())
			if found {
				l.AddError(
					line,
					colStartI,
					colEndI,
					fmt.Sprintf("Can't redeclare constant `%s` in the same scope!", name.GetText()),
				)
			} else {
				log.Printf("Inferring type of `%s` as `%s`\n", name.GetText(), inferedType)
				l.ScopeManager.CurrentScope.UpsertExpressionType(name.GetText(), inferedType)
				l.ScopeManager.CurrentScope.AddConstant(name.GetText())
			}
		}
	} else {
		declarationType := TypeIdentifier(typeAnnot.Type_().GetText())
		log.Println("Constant", name.GetText(), "has type", declarationType)

		if !l.TypeExists(declarationType) {
			l.AddError(line, colStartT, colEndT, fmt.Sprintf(
				"%s doesn't exist!",
				declarationType,
			))
			declarationType = BASE_TYPES.INVALID
		}

		declExpr := declarationExpr // solo la expresión del valor
		colStartD := declExpr.GetStart().GetColumn()
		colEndD := colStartD + len(declExpr.GetText())

		exprText := declarationExpr.GetText()
		if exprText != "[]" {
			log.Println("Known expressions", l.ScopeManager.CurrentScope.typesByExpression)
			initialExprType, exists := l.ScopeManager.CurrentScope.GetExpressionType(exprText)
			if !exists {
				l.AddError(line, colStartI, colStartT, fmt.Sprintf(
					"Constant Declaration: `%s` doesn't have a type!",
					exprText,
				))
				initialExprType = BASE_TYPES.INVALID
			}

			if initialExprType != declarationType && declarationType != BASE_TYPES.INVALID {
				l.AddError(line, colStartD, colEndD, fmt.Sprintf(
					"The declaration of `%s` specifies a type of `%s` but `%s` was given",
					name,
					declarationType,
					initialExprType,
				))
			}
		}

		if isInsideClassDeclaration {
			fieldName := "this." + name.GetText()
			l.ModifyClassTypeInfo(TypeIdentifier(l.ScopeManager.CurrentScope.Name), func(cti *ClassTypeInfo) {
				cti.UpsertField(name.GetText(), declarationType)
				cti.ConstantFields.Add(name.GetText())
			})
			l.ScopeManager.CurrentScope.UpsertExpressionType(fieldName, declarationType)
			l.ScopeManager.CurrentScope.AddConstant(fieldName)
		} else {
			_, found := l.ScopeManager.CurrentScope.GetOnlyInScope(name.GetText())
			if found {
				l.AddError(
					line,
					colStartI,
					colEndI,
					fmt.Sprintf("Can't redeclare constant `%s` in the same scope!", name.GetText()),
				)
			} else {
				l.ScopeManager.CurrentScope.UpsertExpressionType(name.GetText(), declarationType)
				l.ScopeManager.CurrentScope.AddConstant(name.GetText())
			}
		}
	}
}

func (l Listener) ExitAssignment(ctx *p.AssignmentContext) {
	log.Println("Assignment found!", ctx.GetText())
	line := ctx.GetStart().GetLine()

	parts := []p.IAssignmentPartContext{}
	previousTypeId := BASE_TYPES.UNKNOWN
	var conditionalExpr p.IConditionalExprContext
	if thisCtx := ctx.ThisAssignment(); thisCtx != nil {
		log.Println("This assignment!")
		classScope, isInsideClassDeclaration := l.ScopeManager.SearchClassScope()
		if !isInsideClassDeclaration {
			l.AddError(line,
				thisCtx.GetStart().GetColumn(),
				thisCtx.GetStop().GetColumn(),
				"Can't use `this` outside of a class declaration scope!",
			)
			return
		}

		property := thisCtx.Identifier()
		typeInfo, found := l.GetTypeInfo(TypeIdentifier(classScope.Name))
		if !found {
			log.Panicf(
				"Trying to access property `%s` for undefined class: `%s`.\nIt should be defined by this point!",
				property.GetText(),
				classScope.Name,
			)
		}

		classInfo := typeInfo.ClassType.GetValue()
		propertyType, found := classInfo.GetFieldType(property.GetText(), &l)
		if !found {
			l.AddError(
				line,
				property.GetSourceInterval().Start,
				property.GetSourceInterval().Stop,
				fmt.Sprintf(
					"Trying to access field `%s` not defined in class `%s`.",
					property.GetText(),
					classInfo.Name,
				),
			)
			return
		}

		previousTypeId = propertyType
		parts = thisCtx.AllAssignmentPart()
		conditionalExpr = thisCtx.ConditionalExpr()
	} else if varCtx := ctx.VariableAssignment(); varCtx != nil {
		log.Println("Variable assignment!")

		varName := varCtx.Identifier()
		typeId, found := l.ScopeManager.CurrentScope.GetExpressionType(varName.GetText())
		if !found {
			l.AddError(
				line,
				varName.GetSymbol().GetColumn(),
				varName.GetSymbol().GetColumn()+len(varName.GetText()),
				fmt.Sprintf(
					"Undeclared variable `%s`",
					varName.GetText(),
				),
			)
		}

		previousTypeId = typeId
		parts = varCtx.AllAssignmentPart()
		conditionalExpr = varCtx.ConditionalExpr()
	}

	if previousTypeId == BASE_TYPES.UNKNOWN {
		l.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			ctx.GetStop().GetColumn(),
			"Can't assign to undeclared variable!",
		)
		return
	}

	for _, part := range parts {
		previousTypeInfo, found := l.GetTypeInfo(previousTypeId)
		if !found {
			l.AddError(
				ctx.GetStart().GetLine(),
				ctx.GetStart().GetColumn(),
				ctx.GetStop().GetColumn(),
				"Can't find type information for type `%s` in assignment\n%s",
				ctx.GetText(),
			)
			return
		}

		isArrayAccess := part.ConditionalExpr() != nil
		if isArrayAccess {
			log.Printf("Trying to access an array element on type `%s`!\n", previousTypeId)

			if !previousTypeInfo.ArrayType.HasValue() {
				l.AddError(
					line,
					part.GetStart().GetColumn(),
					part.GetStart().GetColumn()+len(part.GetText()),
					fmt.Sprintf("Can't index elements on non-array object with type `%s`!", previousTypeId),
				)
				return
			}
			arrayInfo := previousTypeInfo.ArrayType.GetValue()

			idxExpr := part.ConditionalExpr()
			idxId, found := l.ScopeManager.CurrentScope.GetExpressionType(idxExpr.GetText())
			if !found {
				l.AddError(
					part.GetStart().GetLine(),
					part.GetStart().GetColumn(),
					part.GetStop().GetColumn(),
					fmt.Sprintf("Can't index as `%s` because it doesn't have a type assigned to it.", idxExpr.GetText()),
				)
				return
			}

			if idxId != BASE_TYPES.INTEGER {
				l.AddError(
					part.GetStart().GetLine(),
					part.GetStart().GetColumn(),
					part.GetStop().GetColumn(),
					fmt.Sprintf("Can't index as `%s` because it isn't an integer!", idxExpr.GetText()),
				)
				return
			}
			previousTypeId = arrayInfo.Type
		} else {
			fieldName := part.Identifier()
			log.Printf("Trying to access an object field `%s` on type `%s`!\n", fieldName, previousTypeId)

			if !previousTypeInfo.ClassType.HasValue() {
				l.AddError(
					line,
					part.GetStart().GetColumn(),
					part.GetStart().GetColumn()+len(part.GetText()),
					"Can't access elements on an object that is not a class!",
				)
				return
			}

			classInfo := previousTypeInfo.ClassType.GetValue()
			fieldTypeId, found := classInfo.GetFieldType(fieldName.GetText(), &l)
			if !found {
				l.AddError(
					line,
					part.GetStart().GetColumn(),
					part.GetStart().GetColumn()+len(part.GetText()),
					fmt.Sprintf("Can't access unkonwn field `%s` for class `%s`", fieldName.GetText(), classInfo.Name),
				)
				return
			}

			previousTypeId = fieldTypeId
		}
	}

	assignTypeId, found := l.ScopeManager.CurrentScope.GetExpressionType(conditionalExpr.GetText())
	if !found {
		l.AddError(
			conditionalExpr.GetStart().GetLine(),
			conditionalExpr.GetStart().GetColumn(),
			conditionalExpr.GetStop().GetColumn(),
			fmt.Sprintf("Can't assign a value (`%s`) of an unknown type!", conditionalExpr.GetText()),
		)
		return
	}

	if previousTypeId != assignTypeId {
		l.AddError(
			line,
			ctx.GetStart().GetColumn(),
			ctx.GetStop().GetColumn(),
			fmt.Sprintf("Type mismatch in assignment! (`%s` != `%s`)", previousTypeId, assignTypeId),
		)
		return
	}

	log.Printf("Valid assignment! Types match: `%s`\n", previousTypeId)
}

func (l Listener) ExitPrimaryExpr(ctx *p.PrimaryExprContext) {
	var referenceType TypeIdentifier
	var found bool

	// Primary expresion is of type conditional just pass inherit its type
	if expr := ctx.ConditionalExpr(); expr != nil {
		referenceType, found = l.ScopeManager.CurrentScope.GetExpressionType(expr.GetText())
		if found {
			l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), referenceType)
		}
		return
	}

	if lhs := ctx.LeftHandSide(); lhs != nil {
		referenceType, found = l.ScopeManager.CurrentScope.GetExpressionType(lhs.GetText())
		if found {
			l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), referenceType)
		}
		return
	}

	if literal := ctx.LiteralExpr(); literal != nil {
		referenceType, found = l.ScopeManager.CurrentScope.GetExpressionType(literal.GetText())
		if found {
			l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), referenceType)
		}
		return
	}
}

func (l Listener) ExitIdentifierExpr(ctx *p.IdentifierExprContext) {
	identifier := ctx.Identifier().GetText()
	log.Printf("Processing identifier expression: %s", identifier)

	// First check if it's a regular expression/variable
	identifierType, found := l.ScopeManager.CurrentScope.GetExpressionType(identifier)
	if found {
		parent := ctx.GetParent()
		if leftHandSide, ok := parent.(*p.LeftHandSideContext); ok {
			if len(leftHandSide.AllSuffixOp()) == 0 {
				l.ScopeManager.CurrentScope.UpsertExpressionType(leftHandSide.GetText(), identifierType)
			}
		}
		return
	}

	// If not found as expression, check if it's a function
	if funcInfo, functionFound := l.findFunctionInfo(identifier); functionFound {
		log.Printf("Identifier %s is a function, registering its return type", identifier)
		l.ScopeManager.CurrentScope.UpsertExpressionType(identifier, funcInfo.ReturnType)

		parent := ctx.GetParent()
		if leftHandSide, ok := parent.(*p.LeftHandSideContext); ok {
			if len(leftHandSide.AllSuffixOp()) == 0 {
				l.ScopeManager.CurrentScope.UpsertExpressionType(leftHandSide.GetText(), funcInfo.ReturnType)
			}
		}
		return
	}

	// If neither expression nor function, it's undeclared
	line := ctx.GetStart().GetLine()
	colStart := ctx.Identifier().GetSymbol().GetColumn()
	colEnd := colStart + len(identifier)
	l.AddError(line, colStart, colEnd, fmt.Sprintf(
		"Undeclared identifier `%s`",
		identifier,
	))
}

var LITERAL_VALUES = struct {
	Null  string
	True  string
	False string
}{
	Null:  "null",
	True:  "true",
	False: "false",
}

func (l Listener) EnterLiteralExpr(ctx *p.LiteralExprContext) {
	// fmt.Println("ENTERING LITERAL EXPRESION: " + ctx.GetText())
	strRepresentation := ctx.GetText()
	switch strRepresentation {
	case LITERAL_VALUES.Null:
		l.ScopeManager.CurrentScope.UpsertExpressionType(strRepresentation, BASE_TYPES.NULL)
	case LITERAL_VALUES.True, LITERAL_VALUES.False:
		l.ScopeManager.CurrentScope.UpsertExpressionType(strRepresentation, BASE_TYPES.BOOLEAN)
	default:
		literal := ctx.Literal()
		if literal != nil {
			literalExpr := literal.GetText()
			_, err := strconv.ParseInt(literalExpr, 10, 64)
			if err != nil {
				log.Println("Adding", literalExpr, "as an expression of type", BASE_TYPES.STRING)
				l.ScopeManager.CurrentScope.UpsertExpressionType(literalExpr, BASE_TYPES.STRING)
				l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.STRING)
			} else {
				log.Println("Adding", literalExpr, "as an expression of type", BASE_TYPES.INTEGER)
				l.ScopeManager.CurrentScope.UpsertExpressionType(literalExpr, BASE_TYPES.INTEGER)
				l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INTEGER)
			}
		}
	}
}

func (l Listener) EnterIdentifierExpr(ctx *p.IdentifierExprContext) {
	identifier := ctx.Identifier().GetText()

	identifierType, found := l.ScopeManager.CurrentScope.GetExpressionType(identifier)
	if found {
		parent := ctx.GetParent()
		if leftHandSide, ok := parent.(*p.LeftHandSideContext); ok {
			if len(leftHandSide.AllSuffixOp()) == 0 {
				l.ScopeManager.CurrentScope.UpsertExpressionType(leftHandSide.GetText(), identifierType)
			}
		}
	}
}

func (l Listener) ExitLeftHandSide(ctx *p.LeftHandSideContext) {
	primaryAtom := ctx.PrimaryAtom()
	suffixOps := ctx.AllSuffixOp()

	line := ctx.GetStart().GetLine()
	colStart := ctx.GetStart().GetColumn()
	colEnd := colStart + len(ctx.GetText())

	log.Printf("Validating leftHandSide: %s", ctx.GetText())

	if !l.isValidCombination(primaryAtom, suffixOps) {
		l.AddError(line, colStart, colEnd, "Invalid expression structure: "+ctx.GetText())
		return
	}

	l.processValidLeftHandSide(ctx, primaryAtom, suffixOps)
}

func (l Listener) isValidCombination(primaryAtom p.IPrimaryAtomContext, suffixOps []p.ISuffixOpContext) bool {
	switch atom := primaryAtom.(type) {
	case *p.IdentifierExprContext:
		// ✓ '(' arguments? ')'                    (CallExpr)
		// ✓ '[' conditionalExpr ']'              (IndexExpr)
		// ✓ '.' Identifier                       (PropertyAccessExpr)
		log.Printf("✓ Identifier '%s' - all suffix operations allowed", atom.Identifier().GetText())
		return true

	case *p.NewExprContext:
		// 'new' Identifier '(' arguments? ')' can be followed by:
		// ✗ '(' arguments? ')'                    (CallExpr) - NOT allowed
		// ✗ '[' conditionalExpr ']'              (IndexExpr) - NOT allowed
		// ✗ '.' Identifier                       (PropertyAccessExpr) - NOT allowed

		className := atom.Identifier().GetText()

		if len(suffixOps) > 0 {
			log.Printf("✗ NewExpr 'new %s(...)' cannot have suffix operations", className)
			return false
		}

		log.Printf("✓ NewExpr 'new %s(...)' with no suffix operations", className)
		return true

	case *p.ThisExprContext:
		// ✓ '.' Identifier '(' arguments? ')'         # MethodCallExpr
		// ✗ '(' arguments? ')'                    (CallExpr) - NOT allowed
		// ✗ '[' conditionalExpr ']'              (IndexExpr) - NOT allowed
		// ✓  '.' Identifier                       (PropertyAccessExpr) - NOT allowed

		if len(suffixOps) > 0 {
			switch suffixOps[0].(type) {
			case *p.MethodCallExprContext, *p.PropertyAccessExprContext:
				log.Printf("✓ ThisExpr 'this' with a property/method call operation")

			default:
				log.Printf("✗ ThisExpr 'this' cannot have suffix operations")
				return false
			}
		}

		log.Printf("✓ ThisExpr 'this' with no suffix operations")
		return true

	default:
		log.Printf("✗ Unknown primaryAtom type: %#v", primaryAtom)
		return false
	}
}

func (l Listener) processValidLeftHandSide(ctx *p.LeftHandSideContext, primaryAtom p.IPrimaryAtomContext, suffixOps []p.ISuffixOpContext) {
	// FIXME: Check if INVALID is correct here!
	currentType := BASE_TYPES.INVALID
	currentExpr := ctx.GetText()

	switch atom := primaryAtom.(type) {
	case *p.IdentifierExprContext:
		identifier := atom.Identifier().GetText()
		if foundType, exists := l.ScopeManager.CurrentScope.GetExpressionType(identifier); exists {
			currentType = foundType
			log.Printf("Identifier '%s' has type '%s'", identifier, currentType)
		} else {
			line := ctx.GetStart().GetLine()
			colStart := atom.Identifier().GetSymbol().GetColumn()
			colEnd := colStart + len(identifier)
			l.AddError(line, colStart, colEnd, fmt.Sprintf(
				"Undeclared identifier '%s'",
				identifier,
			))
			return
		}

	case *p.NewExprContext:
		className := atom.Identifier().GetText()
		currentType = TypeIdentifier(className)
		log.Printf("NewExpr creates instance of type '%s'", currentType)

	case *p.ThisExprContext:

		if foundType, exists := l.ScopeManager.CurrentScope.GetExpressionType("this"); exists {
			if len(suffixOps) == 0 {
				currentType = foundType
			}
			log.Printf("'this' has type '%s'", currentType)
		} else {
			line := ctx.GetStart().GetLine()
			colStart := ctx.GetStart().GetColumn()
			colEnd := colStart + 4
			l.AddError(line, colStart, colEnd, "'this' is not available in current scope")
			return
		}

	}

	for i, suffixOp := range suffixOps {
		switch suffix := suffixOp.(type) {
		case *p.CallExprContext:
			log.Printf("Processing CallExpr [%d] on type '%s'", i, currentType)

		case *p.MethodCallExprContext:
			methodName := suffix.Identifier().GetText()
			log.Printf("Processing MethodCallExpr [%d; %s] on type '%s'", i, methodName, currentType)

			tInfo, found := l.GetTypeInfo(currentType)
			if !found {
				l.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					ctx.GetStop().GetColumn(),
					fmt.Sprintf("Type `%s` is not defined!", currentType),
				)
			} else {

				if !tInfo.ClassType.HasValue() {
					l.AddError(
						suffix.GetStart().GetLine(),
						suffix.GetStart().GetColumn(),
						suffix.GetStop().GetColumn(),
						fmt.Sprintf("Can't call method `%s` on non-class type `%s`", methodName, currentType),
					)
				} else {
					classInfo := tInfo.ClassType.GetValue()
					methodInfo, found := classInfo.Methods[methodName]
					if !found {
						l.AddError(
							suffix.GetStart().GetLine(),
							suffix.GetStart().GetColumn(),
							suffix.GetStop().GetColumn(),
							fmt.Sprintf("Can't call undefined method `%s` on type `%s`", methodName, currentType),
						)
					} else {
						currentType = methodInfo.ReturnType
					}
				}
			}

		case *p.IndexExprContext:
			log.Printf("Processing IndexExpr [%d] on type '%s'", i, currentType)
			extractedBaseType, _, _ := strings.Cut(string(currentType), "[]")
			currentType = TypeIdentifier(extractedBaseType)

		case *p.PropertyAccessExprContext:
			propertyName := suffix.Identifier().GetText()
			log.Printf("Processing PropertyAccessExpr [%d]: '.%s'", i, propertyName)
			classScope, found := l.ScopeManager.SearchClassScope()
			if !found {
				log.Printf("Can't access property on `this` when not inside a class declaration!")
				l.AddError(
					suffix.GetStart().GetLine(),
					suffix.GetStart().GetColumn(),
					suffix.GetStop().GetColumn(),
					"Can't access property on `this` when not inside a class declaration!",
				)
			} else {
				tf, found := l.GetTypeInfo(TypeIdentifier(classScope.Name))
				if !found {
					l.AddError(
						suffix.GetStart().GetLine(),
						suffix.GetStart().GetColumn(),
						suffix.GetStop().GetColumn(),
						"Can't access property on `this` when not inside a class declaration!",
					)
				} else {
					tfInfo := tf.ClassType.GetValue()
					fieldType, found := tfInfo.GetFieldType(propertyName, &l)
					if !found {
						l.AddError(
							suffix.GetStart().GetLine(),
							suffix.GetStart().GetColumn(),
							suffix.GetStop().GetColumn(),
							fmt.Sprintf("Can't access property `%s`! Not defined in class `%s` or parent classes!", propertyName[1:], tfInfo.Name),
						)
					} else {
						currentType = fieldType
					}
				}
			}
		}
	}

	l.ScopeManager.CurrentScope.UpsertExpressionType(currentExpr, currentType)
	log.Printf("Final type for '%s': '%s'", currentExpr, currentType)
}

func (l Listener) ExitArrayLiteral(ctx *p.ArrayLiteralContext) {
	line := ctx.GetStart().GetLine()
	startCol := ctx.GetStart().GetColumn()
	endCol := ctx.GetStop().GetColumn()

	expressions := ctx.AllConditionalExpr()
	if len(expressions) == 0 {
		log.Printf("Array literal is empty! Must infer type according to usage...")
		return
	}

	firstExpr := expressions[0]
	firstExprType, found := l.ScopeManager.CurrentScope.GetExpressionType(firstExpr.GetText())
	if !found {
		l.AddError(
			line,
			startCol,
			endCol,
			fmt.Sprintf("Expression `%s` doesn't have a type!", firstExpr.GetText()),
		)
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
		return
	}

	for _, expr := range expressions[1:] {
		exprLine := expr.GetStart().GetLine()
		exprStartCol := expr.GetStart().GetColumn()
		exprEndCol := expr.GetStop().GetColumn()

		exprType, found := l.ScopeManager.CurrentScope.GetExpressionType(expr.GetText())
		if !found {
			l.AddError(
				exprLine,
				exprStartCol,
				exprEndCol,
				fmt.Sprintf("Expression `%s` doesn't have a type!", expr.GetText()),
			)
			l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
			return
		}

		if exprType != firstExprType {
			l.AddError(
				line,
				startCol,
				exprEndCol,
				"Can't define an array with multiple types! The following types are not equal:",
				fmt.Sprintf("`%s` of type `%s`", firstExpr.GetText(), firstExprType),
				fmt.Sprintf("`%s` of type `%s`", expr.GetText(), exprType),
			)
			l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
			return
		}
	}

	arrayId := NewArrayTypeIdentifier(firstExprType)
	l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), arrayId)
	l.UpsertTypeInfo(arrayId, NewTypeInfo_Array(ArrayTypeInfo{firstExprType}))
}
