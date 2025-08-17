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
	firstType, available := (*l.TypesByExpression)[firstExpr.GetText()]
	if !available {
		l.AddError(fmt.Sprintf("`%s` doesn't have a type!", firstExpr.GetText()))
		return
	}

	for _, expr := range exprs[1:] {
		exprType, available := (*l.TypesByExpression)[expr.GetText()]
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
	l.AddTypeByExpr(ctx.GetText(), firstType)
}

func (l Listener) ExitLiteralExpr(ctx *p.LiteralExprContext) {
	strRepresentation := ctx.GetText()
	switch strRepresentation {
	case "null":
		l.AddTypeByExpr(strRepresentation, BASE_TYPES.NULL)
	case "true", "false":
		l.AddTypeByExpr(strRepresentation, BASE_TYPES.BOOLEAN)
	default:
		literal := ctx.Literal()
		if literal != nil {
			literalExpr := literal.GetText()
			_, err := strconv.ParseInt(literalExpr, 10, 64)
			if err != nil {
				log.Println("Adding", literalExpr, "as an expresion of type", BASE_TYPES.STRING)
				l.AddTypeByExpr(literalExpr, BASE_TYPES.STRING)
			} else {
				log.Println("Adding", literalExpr, "as an expresion of type", BASE_TYPES.INTEGER)
				l.AddTypeByExpr(literalExpr, BASE_TYPES.INTEGER)
			}
		}
	}
}

func (l Listener) ExitVariableDeclaration(ctx *p.VariableDeclarationContext) {
	name := ctx.Identifier()

	typeAnnot := ctx.TypeAnnotation()
	hasAnnotation := typeAnnot != nil

	declarationExpr := ctx.Initializer()
	hasInitialExpr := declarationExpr != nil

	if !hasAnnotation {
		log.Println("Variable", name.GetText(), "does NOT have a type! We need to infer it...")
		if hasInitialExpr {
			declarationText := declarationExpr.Expression().GetText()
			inferedType, found := (*l.TypesByExpression)[declarationText]
			if !found {
				l.AddError(fmt.Sprintf("Couldn't infer the type of variable `%s`, initialized with: `%s`", name.GetText(), declarationText))
			} else {
				l.AddTypeByExpr(name.GetText(), inferedType)
			}
		}
		// FIXME: What type does a variable hold if it doesn't defines type or initializer?
		// Maybe null?
	} else {
		declarationType := TypeIdentifier(typeAnnot.Type_().GetText())
		log.Println("Variable", name.GetText(), "has type", declarationType)

		if !l.TypeExists(declarationType) {
			l.AddError(fmt.Sprintf("%s doesn't exist!", declarationType))
		}

		if hasInitialExpr {
			exprText := declarationExpr.Expression().GetText()
			log.Println("Known expressions", *l.TypesByExpression)
			initialExprType, exists := (*l.TypesByExpression)[exprText]
			if !exists {
				l.AddError(fmt.Sprintf("`%s` doesn't have a type!", exprText))
			}

			if initialExprType != declarationType {
				l.AddError(fmt.Sprintf("The declaration of `%s` specifies a type of `%s` but `%s` was given", name, declarationType, initialExprType))
			}
		}

		// FIXME: This variable declaration should be added only to the current active scope!
		l.AddTypeByExpr(name.GetText(), declarationType)
	}
}
