package main

import (
	"fmt"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

type Listener struct{}

func (Listener) EnterProgram(c *p.ProgramContext) {
}

func main() {
	fmt.Println("Hello World!")
}
