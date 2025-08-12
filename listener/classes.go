package listener

import (
	"fmt"
	"log"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

func (l Listener) EnterClassDeclaration(ctx *p.ClassDeclarationContext) {
	identifiers := ctx.AllIdentifier()
	className := identifiers[0]

	log.Println("Declaring", className)
	l.KnownTypes.Add(className.GetText())
	l.Scopes.Push(NewScope(fmt.Sprintf("Scope_%s", className), SCOPE_TYPES.CLASS))
}

func (l Listener) ExitClassDeclaration(ctx *p.ClassDeclarationContext) {
	identifiers := ctx.AllIdentifier()
	className := identifiers[0]

	log.Println("Leaving declaration for", className)
	if op := l.Scopes.Peek(); op.GetValue().Type != SCOPE_TYPES.CLASS {
		panic("Leaving class declaration but last scope is not of type class!")
	}

	l.Scopes.Pop()
}
