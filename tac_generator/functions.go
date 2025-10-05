package tac_generator

import (
	"log"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

func (l Listener) EnterFunctionDeclaration(ctx *p.FunctionDeclarationContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}

	scope := l.GetCurrentScope()
	scopeName := ScopeName(scope.Name)
	l.Program.UpsertScope(scopeName)

	log.Println("Adding parameters for function:", scopeName)
	if params := ctx.Parameters(); params != nil {
		args := params.AllParameter()
		maxIdx := len(args) - 1
		for idx := maxIdx; idx >= 0; idx -= 1 {
			paramExpr := args[idx].Identifier().GetText()
			log.Println("Appending parameter", paramExpr, "for scope", scopeName)
			l.AppendInstruction(
				scopeName,
				NewLoadInstruction(LoadInstruction{l.Program.GetOrGenerateVariable(paramExpr, scopeName)}),
			)
		}
	}

	_, isMethod := l.TypeChecker.ScopeManager.SearchClassScope()
	if isMethod {
		l.AppendInstruction(
			scopeName,
			NewLoadInstruction(LoadInstruction{l.Program.GetOrGenerateVariable("this", scopeName)}),
		)
	}
}

func (l Listener) ExitFunctionDeclaration(ctx *p.FunctionDeclarationContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterReturnStatement(ctx *p.ReturnStatementContext) {
	scope := l.GetCurrentScope()
	scopeName := ScopeName(scope.Name)
	if condExpr := ctx.ConditionalExpr(); condExpr != nil {
		condText := condExpr.GetText()

		var returnValue string
		if literalType, isLiteral := l.TypeChecker.GetLiteralType(condText); isLiteral {
			_, returnValue = literalToTAC(condText, literalType)
		} else {
			tacVar, found := l.Program.GetVariableFor(condExpr.GetText(), scopeName)
			if !found {
				log.Panicf(
					"Failed to get TAC variable for expression: `%s`",
					condExpr.GetText(),
				)
			}
			returnValue = string(tacVar)
		}

		l.AppendInstruction(
			scopeName,
			NewReturnInstruction(ReturnInstruction{LiteralOrVariable(returnValue)}),
		)
	}
}
