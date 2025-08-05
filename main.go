package main

import (
	"fmt"

	p "github.com/ElrohirGT/parser"
)

type Listener struct{}

func (Listener) EnterProgram(c *p.ProgramContext) {
}

func main() {
	fmt.Println("Hello World!")
}
