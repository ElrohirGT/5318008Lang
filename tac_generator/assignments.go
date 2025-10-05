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

	if scope.Type == type_checker.SCOPE_TYPES.CLASS {
		scopeName = ScopeName(scope.Name + "_" + type_checker.CONSTRUCTOR_NAME)
		variableName = "this." + variableName
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

	createAssignment(l, scope, scopeName, isLiteral, variableName, exprType, variableValue)
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
	exprType, found := l.GetCurrentScope().GetExpressionType(exprText)
	if !found {
		log.Panicf(
			"Failed to get type for expression: `%s`",
			exprText,
		)
	}

	_, isLiteral := l.TypeChecker.GetLiteralType(exprText)

	log.Printf(
		"Assignment to variable `%s`, with value: `%s`, which is a literal? %t",
		originalName,
		exprText,
		isLiteral,
	)

	createAssignment(l, scope, scopeName, isLiteral, originalName, exprType, exprText)
}

func createAssignment(
	l Listener,
	scope *type_checker.Scope,
	scopeName ScopeName,
	isLiteral bool,
	variableName string,
	exprType type_checker.TypeIdentifier,
	exprText string,
) {
	if isLiteral {
		literalType, literalValue := literalToTAC(exprText, exprType)
		l.CreateAssignment(scopeName, variableName, literalType, literalValue)
	} else {
		exprVar, found := l.Program.GetVariableFor(exprText, scopeName)
		if !found {
			log.Panicf("Failed to find a variable for the expression:\n%s", exprText)
		}

		if length, found := scope.GetArrayLength(exprText); found {
			scope.UpsertArrayLength(variableName, length)
			l.Program.UpsertTranslation(scopeName, variableName, exprVar)
		} else if tacType, found := l.MapBaseTypeToTacType(exprType); found {
			l.CreateAssignment(scopeName, variableName, tacType, string(exprVar))
		} else {
			l.Program.UpsertTranslation(scopeName, variableName, exprVar)
		}
	}
}
