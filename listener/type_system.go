package listener

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
	line := ctx.GetStart().GetLine()
	exprs := ctx.AllMultiplicativeExpr()
	firstExpr := exprs[0]
	referenceType, available := l.ScopeManager.CurrentScope.GetExpressionType(firstExpr.GetText())
	if !available {
		l.AddError(fmt.Sprintf("(line: %d) `%s` doesn't have a type!", line, firstExpr.GetText()))
		return
	}

	defineLater := make([]p.IMultiplicativeExprContext, 0, len(exprs))

	if referenceType == BASE_TYPES.UNKNOWN {
		defineLater = append(defineLater, firstExpr)

		for _, expr := range exprs[1:] {
			exprType, available := l.ScopeManager.CurrentScope.GetExpressionType(expr.GetText())
			if !available {
				l.AddError(fmt.Sprintf("(line: %d) `%s` doesn't have a type!", line, exprType))
			}

			if exprType != BASE_TYPES.UNKNOWN {
				referenceType = exprType
				break
			}
		}

		if referenceType == BASE_TYPES.UNKNOWN {
			l.AddError(fmt.Sprintf("(line: %d) Can't infer the value of multiplication/addition! All types were unknown...", line))
			return
		}
	}

	for _, expr := range exprs[1:] {
		exprType, available := l.ScopeManager.CurrentScope.GetExpressionType(expr.GetText())
		if !available {
			l.AddError(fmt.Sprintf("(line: %d) `%s` doesn't have a type!", line, exprType))
		}

		if exprType == BASE_TYPES.UNKNOWN {
			defineLater = append(defineLater, expr)
			continue
		}

		if exprType != referenceType {
			l.AddError(fmt.Sprintf(
				"(line: %d) Can't add:\n * leftSide: `%s` of type `%s`\n * rightSide: `%s` of type `%s`",
				line,
				firstExpr.GetText(),
				referenceType,
				expr.GetText(),
				exprType,
			))
		}
	}

	classScope, isInsideClassDeclaration := l.ScopeManager.SearchClassScope()
	for _, expr := range defineLater {
		log.Printf("Infering `%s` as type of `%s`\n", expr.GetText(), referenceType)
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
	hasAnnotation := typeAnnot != nil

	declarationExpr := ctx.Initializer()
	hasInitialExpr := declarationExpr != nil

	isInsideClassDeclaration := l.ScopeManager.CurrentScope.Type == SCOPE_TYPES.CLASS

	if !hasAnnotation {
		log.Println("Variable", name.GetText(), "does NOT have a type! We need to infer it...")
		if hasInitialExpr {
			declarationText := declarationExpr.Expression().GetText()
			inferedType, found := l.ScopeManager.CurrentScope.GetExpressionType(declarationText)
			if !found {
				l.AddError(fmt.Sprintf(
					"(line: %d) Couldn't infer the type of variable `%s`, initialized with: `%s`",
					line,
					name.GetText(),
					declarationText,
				))
			} else {
				if isInsideClassDeclaration {
					l.ModifyClassTypeInfo(TypeIdentifier(l.ScopeManager.CurrentScope.Name), func(cti *ClassTypeInfo) {
						cti.UpsertField(name.GetText(), inferedType)
					})
					l.ScopeManager.CurrentScope.UpsertExpressionType("this."+name.GetText(), inferedType)
				} else {
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
			l.AddError(fmt.Sprintf(
				"(line: %d) %s doesn't exist!",
				line,
				declarationType,
			))
		}

		if hasInitialExpr {
			exprText := declarationExpr.Expression().GetText()
			log.Println("Known expressions", l.ScopeManager.CurrentScope.typesByExpression)
			initialExprType, exists := l.ScopeManager.CurrentScope.GetExpressionType(exprText)
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
