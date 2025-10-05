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
	scope := l.GetCurrentScope()
	scopeName := ScopeName(scope.Name)
	variableName := ctx.Identifier().GetText()

	exprText := ""
	if init := ctx.Initializer(); init != nil {
		exprText = init.ConditionalExpr().GetText()
	}

	exprType, foundType := scope.GetExpressionType(variableName)
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
		l.CreateAssignment(variableName, exprType, variableValue)
	} else {
		if length, found := scope.GetArrayLength(exprText); found {
			scope.UpsertArrayLength(variableName, length)
		} else {
			exprVar, found := l.Program.GetVariableFor(exprText, scopeName)
			if !found {
				log.Panicf("Failed to find a variable for the expression:\n%s", exprText)
			}

			l.CreateAssignment(variableName, exprType, string(exprVar))
		}
	}
}

func (l Listener) ExitAssignment(ctx *p.AssignmentContext) {
	scope := l.GetCurrentScope()
	scopeName := ScopeName(scope.Name)

	var originalName string
	var assignExpr p.IConditionalExprContext
	if ctx.ThisAssignment() != nil {
		assignExpr = ctx.ThisAssignment().ConditionalExpr()
		originalName = "this." + ctx.ThisAssignment().Identifier().GetText()
	} else if ctx.VariableAssignment() != nil {
		assignExpr = ctx.VariableAssignment().ConditionalExpr()
		originalName = ctx.VariableAssignment().Identifier().GetText()
	}
	exprText := assignExpr.GetText()

	literalType, isLiteral := l.TypeChecker.GetLiteralType(exprText)

	log.Printf(
		"Assignment to variable `%s`, with value: `%s`, which is a literal? %t",
		originalName,
		exprText,
		isLiteral,
	)
	if isLiteral {
		l.CreateAssignment(originalName, literalType, exprText)
	} else {
		if length, found := scope.GetArrayLength(exprText); found {
			scope.UpsertArrayLength(originalName, length)
		} else {
			exprVar, found := l.Program.GetVariableFor(exprText, scopeName)
			if !found {
				log.Panicf("Failed to find a variable for the expression:\n%s", exprText)
			}

			l.CreateAssignment(originalName, literalType, string(exprVar))
		}
	}
}
