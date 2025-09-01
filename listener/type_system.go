package listener

import (
	"fmt"
	"log"
	"strconv"

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

	for i, expr := range exprs[1:] {
		exprType, available := l.ScopeManager.CurrentScope.GetExpressionType(expr.GetText())
		exprColStart := expr.GetStart().GetColumn()
		exprColEnd := exprColStart + len(expr.GetText())

		if !available {
			l.AddError(line, exprColStart, exprColEnd, fmt.Sprintf("Exit Add (expr - second): `%s` doesn't have a type!", exprType))
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
			inferedType, found := l.ScopeManager.CurrentScope.GetExpressionType(declarationText)
			if !found {
				l.AddError(line, colStartI, colEndI, fmt.Sprintf(
					"Couldn't infer the type of variable `%s`, initialized with: `%s`",
					name.GetText(),
					declarationText,
				))
				inferedType = BASE_TYPES.INVALID
			}

			if isInsideClassDeclaration {
				newExprName := "this." + name.GetText()
				log.Printf("Inferring type of `%s` as `%s`\n", newExprName, inferedType)
				l.ModifyClassTypeInfo(TypeIdentifier(l.ScopeManager.CurrentScope.Name), func(cti *ClassTypeInfo) {
					cti.UpsertField(name.GetText(), inferedType)
				})
				// FIXME: Check if variable is already declared!
				l.ScopeManager.CurrentScope.UpsertExpressionType(newExprName, inferedType)
			} else {
				log.Printf("Inferring type of `%s` as `%s`\n", name.GetText(), inferedType)
				// FIXME: Check if variable is already declared!
				l.ScopeManager.CurrentScope.UpsertExpressionType(name.GetText(), inferedType)
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
			declarationType = BASE_TYPES.INVALID
		}

		if hasInitialExpr {
			declExpr := declarationExpr.ConditionalExpr() // solo la expresión del valor
			colStartD := declExpr.GetStart().GetColumn()
			colEndD := colStartD + len(declExpr.GetText())

			exprText := declarationExpr.ConditionalExpr().GetText()
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

		if isInsideClassDeclaration {
			l.ModifyClassTypeInfo(TypeIdentifier(l.ScopeManager.CurrentScope.Name), func(cti *ClassTypeInfo) {
				cti.UpsertField(name.GetText(), declarationType)
			})
			// FIXME: Check if variable is already declared!
			l.ScopeManager.CurrentScope.UpsertExpressionType("this."+name.GetText(), declarationType)
		} else {
			// FIXME: Check if variable is already declared!
			l.ScopeManager.CurrentScope.UpsertExpressionType(name.GetText(), declarationType)
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
		inferedType, found := l.ScopeManager.CurrentScope.GetExpressionType(declarationText)
		if !found {
			l.AddError(line, colStartI, colEndI, fmt.Sprintf(
				"Couldn't infer the type of variable `%s`, initialized with: `%s`",
				name.GetText(),
				declarationText,
			))
			inferedType = BASE_TYPES.INVALID
		}

		if isInsideClassDeclaration {
			newExprName := "this." + name.GetText()
			log.Printf("Inferring type of `%s` as `%s`\n", newExprName, inferedType)
			l.ModifyClassTypeInfo(TypeIdentifier(l.ScopeManager.CurrentScope.Name), func(cti *ClassTypeInfo) {
				cti.UpsertField(name.GetText(), inferedType)
				cti.ConstantFields.Add(name.GetText())
			})
			// FIXME: Check if variable is already declared!
			l.ScopeManager.CurrentScope.UpsertExpressionType(newExprName, inferedType)
			l.ScopeManager.CurrentScope.AddConstant(newExprName)
		} else {
			log.Printf("Inferring type of `%s` as `%s`\n", name.GetText(), inferedType)
			// FIXME: Check if variable is already declared!
			l.ScopeManager.CurrentScope.UpsertExpressionType(name.GetText(), inferedType)
			l.ScopeManager.CurrentScope.AddConstant(name.GetText())
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

		if isInsideClassDeclaration {
			fieldName := "this." + name.GetText()
			l.ModifyClassTypeInfo(TypeIdentifier(l.ScopeManager.CurrentScope.Name), func(cti *ClassTypeInfo) {
				cti.UpsertField(name.GetText(), declarationType)
				cti.ConstantFields.Add(name.GetText())
			})
			// FIXME: Check if variable is already declared!
			l.ScopeManager.CurrentScope.UpsertExpressionType(fieldName, declarationType)
			l.ScopeManager.CurrentScope.AddConstant(fieldName)
		} else {
			// FIXME: Check if variable is already declared!
			l.ScopeManager.CurrentScope.UpsertExpressionType(name.GetText(), declarationType)
			l.ScopeManager.CurrentScope.AddConstant(name.GetText())
		}
	}
}

func (l Listener) ExitThisAssignment(ctx *p.ThisAssignmentContext) {
	log.Println("`This` Variable Assignment!", ctx.GetText())
	line := ctx.GetStart().GetLine()

	identifiers := ctx.AllIdentifier()
	firstExpr := ctx.Identifier(0)
	colStartF := firstExpr.GetSymbol().GetColumn()
	colEndF := colStartF + len(firstExpr.GetText())

	classScope, isInsideClassDeclaration := l.ScopeManager.SearchClassScope()
	if !isInsideClassDeclaration {
		l.AddError(line, colStartF, colEndF, "Can't use `this` outside of a class declaration scope!")
		return
	}

	previousIdentifier := "this"
	previousType := TypeIdentifier(classScope.Name)
	for i, identifier := range identifiers {
		info, found := l.GetTypeInfo(previousType)
		if !found {
			l.AddError(line, colStartF, colEndF, fmt.Sprintf(
				"Undeclared type `%s` for variable `%s`",
				previousType,
				previousIdentifier,
			))
			return
		}

		if !info.ClassType.HasValue() {
			l.AddError(line, colStartF, colEndF, fmt.Sprintf(
				"Trying to access a field `%s` from type `%s` but `%s` is not a class!",
				identifier.GetText(),
				previousType,
				previousIdentifier,
			))
			return
		}

		colStartI := identifier.GetSymbol().GetColumn()
		colEndI := colStartI + len(identifier.GetText())
		classInfo := info.ClassType.GetValue()
		fieldType, hasField := classInfo.Fields[identifier.GetText()]
		if !hasField {
			log.Printf("Field `%s` not found in class:\n%#v\n", identifier.GetText(), classInfo)
			l.AddError(line, colStartI, colEndI, fmt.Sprintf(
				"Trying to access field `%s` not defined in class `%s`!",
				identifier.GetText(),
				classInfo.Name,
			))
			return
		}
		previousType = fieldType

		isLastIdentifier := i == len(identifiers)-1-1
		if isLastIdentifier {
			assignExpr := ctx.ConditionalExpr()
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

			if fieldType != assignType {
				l.AddError(line, colStartI, colEndA, fmt.Sprintf(
					"Trying to assign `%s` to variable `%s` but types don't match! (`%s` != `%s`)",
					assignExpr.GetText(),
					identifier.GetText(),
					fieldType,
					assignType,
				))
				return
			}
		}
	}
}

func (l Listener) ExitVariableAssignment(ctx *p.VariableAssignmentContext) {
	log.Println("General Variable Assignment!", ctx.GetText())
	line := ctx.GetStart().GetLine()

	identifiers := ctx.AllIdentifier()
	isPropertyAssignment := len(identifiers) > 1
	if isPropertyAssignment {
		firstExpr := ctx.Identifier(0)
		colStartF := firstExpr.GetSymbol().GetColumn()
		colEndF := colStartF + len(firstExpr.GetText())

		initialType, found := l.ScopeManager.CurrentScope.GetExpressionType(firstExpr.GetText())
		if !found {
			l.AddError(line, colStartF, colEndF, fmt.Sprintf(
				"Undeclared variable `%s`",
				firstExpr.GetText(),
			))
			return
		}

		if initialType == BASE_TYPES.INVALID {
			l.AddError(line, colStartF, colEndF, fmt.Sprintf(
				"Can't access properties of invalid variable `%s`",
				firstExpr.GetText(),
			))
			return
		}

		previousType := initialType
		for i, identifier := range identifiers[1:] {
			info, found := l.GetTypeInfo(previousType)
			if !found {
				l.AddError(line, colStartF, colEndF, fmt.Sprintf(
					"Undeclared type `%s` for variable `%s`",
					initialType,
					firstExpr.GetText(),
				))
				return
			}

			if !info.ClassType.HasValue() {
				l.AddError(line, colStartF, colEndF, fmt.Sprintf(
					"Trying to access a field `%s` from type `%s` but `%s` is not a class!",
					firstExpr.GetText(),
					initialType,
					firstExpr.GetText(),
				))
				return
			}

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
			previousType = fieldType

			isLastIdentifier := i == len(identifiers)-1-1
			if isLastIdentifier {
				assignExpr := ctx.ConditionalExpr()
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

				if fieldType != assignType {
					l.AddError(line, colStartI, colEndA, fmt.Sprintf(
						"Trying to assign `%s` to variable `%s` but types don't match! (`%s` != `%s`)",
						assignExpr.GetText(),
						identifier.GetText(),
						fieldType,
						assignType,
					))
					return
				}
			}

		}
	} else {
		varExpr := ctx.Identifier(0)
		colStartF := varExpr.GetSymbol().GetColumn()
		colEndF := colStartF + len(varExpr.GetText())

		varType, found := l.ScopeManager.CurrentScope.GetExpressionType(varExpr.GetText())
		if !found {
			l.AddError(line, colStartF, colEndF, fmt.Sprintf(
				"Undeclared variable `%s`",
				varExpr.GetText(),
			))
			return
		}

		if varType == BASE_TYPES.INVALID {
			l.AddError(line, colStartF, colEndF, fmt.Sprintf(
				"Can't assign to invalid variable: `%s`",
				varExpr.GetText(),
			))
			return
		}

		assignExpr := ctx.ConditionalExpr()
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

		if varType != assignType {
			l.AddError(line, colStartF, colEndA, fmt.Sprintf(
				"Trying to assign `%s` to variable `%s` but types don't match! (`%s` != `%s`)",
				assignExpr.GetText(),
				varExpr.GetText(),
				varType,
				assignType,
			))
			return
		}

	}
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

	identifierType, found := l.ScopeManager.CurrentScope.GetExpressionType(identifier)
	if found {
		parent := ctx.GetParent()
		if leftHandSide, ok := parent.(*p.LeftHandSideContext); ok {
			if len(leftHandSide.AllSuffixOp()) == 0 {
				l.ScopeManager.CurrentScope.UpsertExpressionType(leftHandSide.GetText(), identifierType)
			}
		}
	} else {
		line := ctx.GetStart().GetLine()
		colStart := ctx.Identifier().GetSymbol().GetColumn()
		colEnd := colStart + len(identifier)
		l.AddError(line, colStart, colEnd, fmt.Sprintf(
			"Undeclared identifier `%s`",
			identifier,
		))
	}
}

func (l Listener) EnterLiteralExpr(ctx *p.LiteralExprContext) {
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

	// FIXME: This fails with test.cps but it shouldn't!
	if !l.isValidCombination(primaryAtom, suffixOps) {
		l.AddError(line, colStart, colEnd, fmt.Sprintf(
			"Invalid expression structure: %s",
			ctx.GetText(),
		))
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
		// ✗ '(' arguments? ')'                    (CallExpr) - NOT allowed
		// ✗ '[' conditionalExpr ']'              (IndexExpr) - NOT allowed
		// ✗ '.' Identifier                       (PropertyAccessExpr) - NOT allowed

		if len(suffixOps) > 0 {
			log.Printf("✗ ThisExpr 'this' cannot have suffix operations")
			return false
		}

		log.Printf("✓ ThisExpr 'this' with no suffix operations")
		return true

	default:
		log.Printf("✗ Unknown primaryAtom type")
		return false
	}
}

func (l Listener) processValidLeftHandSide(ctx *p.LeftHandSideContext, primaryAtom p.IPrimaryAtomContext, suffixOps []p.ISuffixOpContext) {
	// FIXME: Check if INVALID is correct here!
	var currentType TypeIdentifier = BASE_TYPES.INVALID
	var currentExpr string = ctx.GetText()

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
			currentType = foundType
			log.Printf("'this' has type '%s'", currentType)
		} else {
			line := ctx.GetStart().GetLine()
			colStart := ctx.GetStart().GetColumn()
			colEnd := colStart + 4
			l.AddError(line, colStart, colEnd, "'this' is not available in current scope")
			return
		}
	}

	if _, ok := primaryAtom.(*p.IdentifierExprContext); ok {
		for i, suffixOp := range suffixOps {
			switch suffix := suffixOp.(type) {
			case *p.CallExprContext:
				log.Printf("Processing CallExpr [%d] on type '%s'", i, currentType)

			case *p.IndexExprContext:
				log.Printf("Processing IndexExpr [%d] on type '%s'", i, currentType)

			case *p.PropertyAccessExprContext:
				propertyName := suffix.Identifier().GetText()
				log.Printf("Processing PropertyAccessExpr [%d]: '.%s' on type '%s'", i, propertyName, currentType)
			}
		}
	}

	l.ScopeManager.CurrentScope.UpsertExpressionType(currentExpr, currentType)
	log.Printf("Final type for '%s': '%s'", currentExpr, currentType)
}
