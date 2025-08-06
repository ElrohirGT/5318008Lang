package main

import (
	"fmt"
	"log"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

type Listener struct {
	*p.BaseCompiscriptListener
	KnownTypes map[string]struct{}
	Errors     *[]string
}

func NewListener() Listener {
	baseTypes := map[string]struct{}{
		"integer": {},
		"float":   {},
		"boolean": {},
		"string":  {},
		"null":    {},
	}

	errors := []string{}

	return Listener{
		KnownTypes: baseTypes,
		Errors:     &errors,
	}
}

func (l Listener) EnterClassDeclaration(ctx *p.ClassDeclarationContext) {
	identifiers := ctx.AllIdentifier()
	className := identifiers[0]

	log.Println("Declaring", className)
	l.KnownTypes[className.GetText()] = struct{}{}
	log.Println("Declared types", l.KnownTypes)
}

func (l Listener) EnterVariableDeclaration(ctx *p.VariableDeclarationContext) {
	name := ctx.Identifier()
	typeAnnot := ctx.TypeAnnotation()
	if typeAnnot == nil {
		log.Println("Variable", name.GetText(), "does NOT have a type!")
		// TODO: Infer variable type and check if it exists
	} else {
		typeName := typeAnnot.Type_().GetText()
		log.Println("Variable", name.GetText(), "has type", typeName)

		if _, exists := l.KnownTypes[typeName]; !exists {
			*l.Errors = append(*l.Errors, fmt.Sprintf("ERROR: %s doesn't exist!", typeName))
		}

		// TODO: Check if the expr type matches defined type
	}
}

func (l Listener) HasErrors() bool {
	return len(*l.Errors) > 0
}
