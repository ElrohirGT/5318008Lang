package tac_generator

import (
	"fmt"
	"log"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

func (l Listener) EnterFunctionDeclaration(ctx *p.FunctionDeclarationContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}

	scope := l.GetCurrentScope()
	baseScopeName := ScopeName(scope.Name)
	scopeName := baseScopeName

	counter := 0
	for l.Program.FunctionScopes.Exists(scopeName) {
		counter++
		scopeName = ScopeName(fmt.Sprintf("%s_%d", baseScopeName, counter))
	}

	l.TACScope[scope] = scopeName

	parentName := l.GetParentScopeName()
	l.Program.InsertIfNotExists(scopeName, parentName)

	l.Program.FunctionScopes.Add(scopeName)

	log.Println("Adding parameters for function:", scopeName)
	_, isMethod := l.TypeChecker.ScopeManager.SearchClassScope()
	if isMethod {
		thisTacName := l.Program.GetOrGenerateVariable("this", scopeName)
		l.AppendInstruction(
			scopeName,
			NewLoadInstruction(LoadInstruction{thisTacName}).
				AddComment("(this)"),
		)
		l.Program.UpsertTranslation(scopeName, "this", thisTacName)
	}

	if params := ctx.Parameters(); params != nil {
		args := params.AllParameter()
		for idx := range args {
			paramExpr := args[idx].Identifier().GetText()
			log.Println("Appending parameter", paramExpr, "for scope", scopeName)
			l.AppendInstruction(
				scopeName,
				NewLoadInstruction(LoadInstruction{l.Program.GetOrGenerateVariable(paramExpr, scopeName)}).
					AddComment("("+paramExpr+")"),
			)
		}
	}
}

func (l Listener) ExitProgram(ctx *p.ProgramContext) {
	scope := l.GetCurrentScope()
	scopeName := ScopeName(scope.Name)

	l.AppendInstruction(scopeName, NewEndInstruction())
}

func (l Listener) ExitFunctionDeclaration(ctx *p.FunctionDeclarationContext) {
	scope := l.GetCurrentScope()
	var scopeName ScopeName
	scopeName, found := l.TACScope[scope]

	if !found {
		scopeName = ScopeName(scope.Name)
	}

	l.TypeChecker.ScopeManager.ReplaceWithParent()
	l.AppendInstruction(scopeName, NewEndInstruction())
}

func (l Listener) ExitReturnStatement(ctx *p.ReturnStatementContext) {
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
