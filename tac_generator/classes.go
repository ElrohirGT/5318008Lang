package tac_generator

import (
	"log"

	p "github.com/ElrohirGT/5318008Lang/parser"
	tc "github.com/ElrohirGT/5318008Lang/type_checker"
)

func (l Listener) EnterClassDeclaration(ctx *p.ClassDeclarationContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}

	scope := l.GetCurrentScope()
	// l.Program.UpsertScope(ScopeName(scope.Name))
	l.Program.UpsertScope(ScopeName(scope.Name + "_" + tc.CONSTRUCTOR_NAME))
}

func (l Listener) ExitClassDeclaration(ctx *p.ClassDeclarationContext) {
	if l.TypeChecker.ScopeManager.CurrentScope.Type != tc.SCOPE_TYPES.CLASS {
		panic("Trying to exit a class declaration but the scope is not of type class!")
	}

	log.Printf("Escaping class declaration: %s", ctx.AllIdentifier()[0].GetText())
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}
