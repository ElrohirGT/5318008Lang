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
}

func (l Listener) ExitFunctionDeclaration(ctx *p.FunctionDeclarationContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

func (s Listener) ExitLeftHandSide(ctx *p.LeftHandSideContext) {
	log.Println("Left hand side:", ctx.GetText())
	suffixes := ctx.AllSuffixOp()
	if len(suffixes) == 0 {
		return
	}

}
