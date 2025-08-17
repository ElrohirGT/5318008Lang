package listener

import (
	"fmt"
	"log"
	"strconv"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

var BASE_TYPES = struct {
	INTEGER TypeIdentifier
	BOOLEAN TypeIdentifier
	STRING  TypeIdentifier
	NULL    TypeIdentifier
	// Type used when the type system can't define the type yet
	// After evaluating the whole program NOTHING can have an unknown type
	UNKNOWN TypeIdentifier
}{
	INTEGER: "integer",
	BOOLEAN: "boolean",
	STRING:  "string",
	NULL:    "null",
	UNKNOWN: "unknown",
}

func (l Listener) ExitAdditiveExpr(ctx *p.AdditiveExprContext) {
	exprs := ctx.AllMultiplicativeExpr()
	firstExpr := exprs[0]
	firstType, available := l.CurrentScope.GetExpressionType(firstExpr.GetText())
	if !available {
		l.AddError(fmt.Sprintf("`%s` doesn't have a type!", firstExpr.GetText()))
		return
	}

	for _, expr := range exprs[1:] {
		exprType, available := l.CurrentScope.GetExpressionType(expr.GetText())
		if !available {
			l.AddError(fmt.Sprintf("`%s` doesn't have a type!", exprType))
		}

		if exprType != firstType {
			l.AddError(fmt.Sprintf(
				"Can't add:\n * leftSide: `%s` of type `%s`\n * rightSide: `%s` of type `%s`",
				firstExpr.GetText(),
				firstType,
				expr.GetText(),
				exprType,
			))
		}
	}

	log.Printf("Adding expression `%s` of type `%s`", ctx.GetText(), firstType)
	l.CurrentScope.AddExpressionType(ctx.GetText(), firstType)
}

func (l Listener) ExitLiteralExpr(ctx *p.LiteralExprContext) {
	strRepresentation := ctx.GetText()
	switch strRepresentation {
	case "null":
		l.CurrentScope.AddExpressionType(strRepresentation, BASE_TYPES.NULL)
	case "true", "false":
		l.CurrentScope.AddExpressionType(strRepresentation, BASE_TYPES.BOOLEAN)
	default:
		literal := ctx.Literal()
		if literal != nil {
			literalExpr := literal.GetText()
			_, err := strconv.ParseInt(literalExpr, 10, 64)
			if err != nil {
				log.Println("Adding", literalExpr, "as an expresion of type", BASE_TYPES.STRING)
				l.CurrentScope.AddExpressionType(literalExpr, BASE_TYPES.STRING)
			} else {
				log.Println("Adding", literalExpr, "as an expresion of type", BASE_TYPES.INTEGER)
				l.CurrentScope.AddExpressionType(literalExpr, BASE_TYPES.INTEGER)
			}
		}
	}
}

func (l Listener) ExitVariableDeclaration(ctx *p.VariableDeclarationContext) {
	line := ctx.GetStart().GetLine()
	name := ctx.Identifier()

	typeAnnot := ctx.TypeAnnotation()
	hasAnnotation := typeAnnot != nil

	declarationExpr := ctx.Initializer()
	hasInitialExpr := declarationExpr != nil

	if !hasAnnotation {
		log.Println("Variable", name.GetText(), "does NOT have a type! We need to infer it...")
		if hasInitialExpr {
			declarationText := declarationExpr.Expression().GetText()
			inferedType, found := l.CurrentScope.GetExpressionType(declarationText)
			if !found {
				l.AddError(fmt.Sprintf(
					"(line: %d) Couldn't infer the type of variable `%s`, initialized with: `%s`",
					line,
					name.GetText(),
					declarationText,
				))
			} else {
				l.CurrentScope.AddExpressionType(name.GetText(), inferedType)
			}
		} else {
			l.CurrentScope.AddExpressionType(name.GetText(), BASE_TYPES.UNKNOWN)
		}
	} else {
		declarationType := TypeIdentifier(typeAnnot.Type_().GetText())
		log.Println("Variable", name.GetText(), "has type", declarationType)

		if !l.TypeExists(declarationType) {
			l.AddError(fmt.Sprintf(
				"(line: %d) %s doesn't exist!",
				line,
				declarationType,
			))
		}

		if hasInitialExpr {
			exprText := declarationExpr.Expression().GetText()
			log.Println("Known expressions", l.CurrentScope.typesByExpression)
			initialExprType, exists := l.CurrentScope.GetExpressionType(exprText)
			if !exists {
				l.AddError(fmt.Sprintf(
					"(line: %d) `%s` doesn't have a type!",
					line,
					exprText,
				))
			}

			if initialExprType != declarationType {
				l.AddError(fmt.Sprintf(
					"(line: %d) The declaration of `%s` specifies a type of `%s` but `%s` was given",
					line,
					name,
					declarationType,
					initialExprType,
				))
			}
		}

		l.CurrentScope.AddExpressionType(name.GetText(), declarationType)
	}
}
