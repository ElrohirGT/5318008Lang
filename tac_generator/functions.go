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
	l.Program.UpsertScope(ScopeName(scope.Name))
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

		l.AppendInstruction(NewReturnInstruction(ReturnInstruction{LiteralOrVariable(returnValue)}))
	}
}
