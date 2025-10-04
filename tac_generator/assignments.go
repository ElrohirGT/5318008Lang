package tac_generator

import (
	"log"

	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/ElrohirGT/5318008Lang/type_checker"
)

// a = 5;
// = t1 i32 5

// a = sumar(1,2);
// PARAM 1
// PARAM 2
// CALLRET a sumar 2

func (l Listener) ExitVariableDeclaration(ctx *p.VariableDeclarationContext) {
	variableName := ctx.Identifier().GetText()

	exprText := ""
	if init := ctx.Initializer(); init != nil {
		exprText = init.ConditionalExpr().GetText()
	}

	exprType, foundType := l.TypeChecker.ScopeManager.CurrentScope.GetExpressionType(variableName)
	if !foundType {
		log.Panicf(
			"Variable with name: `%s`\nnot found in current scope!%#v",
			variableName,
			*l.GetCurrentScope(),
		)
	}

	isLiteral := false
	variableValue := exprText
	if exprText == "" {
		isLiteral = true
		switch exprType {
		case type_checker.BASE_TYPES.STRING:
			variableValue = ""
		case type_checker.BASE_TYPES.BOOLEAN:
			variableValue = "0"
		case type_checker.BASE_TYPES.INTEGER:
			variableValue = "0"
		case type_checker.BASE_TYPES.INVALID, type_checker.BASE_TYPES.NULL, type_checker.BASE_TYPES.UNKNOWN:
			log.Panicf("Variable with name: `%s` has an invalid type! `%s`", variableName, exprType)
		default:
			// FIXME: Handle array type!
		}
	} else {
		_, isLiteral = l.TypeChecker.GetLiteralType(exprText)
	}

	log.Printf(
		"Declaring variable `%s`, with value: `%s`, which is a literal? %t",
		variableName,
		exprText,
		isLiteral,
	)
	if isLiteral {
		l.CreateLiteralAssignment(variableName, exprType, variableValue)
	} else {
		// FIXME: Manage case where is not literal
	}

}

func (l Listener) ExitAssignment(ctx *p.AssignmentContext) {
	var originalName string
	var assignExpr p.IConditionalExprContext
	if ctx.ThisAssignment() != nil {
		assignExpr = ctx.ThisAssignment().ConditionalExpr()
		originalName = ctx.ThisAssignment().Identifier().GetText()
	} else if ctx.VariableAssignment() != nil {
		assignExpr = ctx.VariableAssignment().ConditionalExpr()
		originalName = ctx.VariableAssignment().Identifier().GetText()
	}
	exprText := assignExpr.GetText()

	literalType, isLiteral := l.TypeChecker.GetLiteralType(exprText)
	if !isLiteral {
		// FIXME: What do we do when the assignment is not a literal!?
	} else {
		l.CreateLiteralAssignment(originalName, literalType, exprText)
	}
}
