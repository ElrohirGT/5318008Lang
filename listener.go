package main

import (
	"fmt"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

type Listener struct {
	*p.BaseCompiscriptListener
	Errors []string
}

func (Listener) EnterProgram(ctx *p.ProgramContext) {
	fmt.Println("Enter program!", ctx.GetText())
}

func (Listener) ExitProgram(ctx *p.ProgramContext) {
	fmt.Println("Exit program!", ctx.GetText())
}

func (l Listener) HasErrors() bool {
	return len(l.Errors) > 0
}
