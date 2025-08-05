package main

import (
	"fmt"
	"os"

	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/antlr4-go/antlr/v4"
)

type Listener struct {
	*p.BaseCompiscriptListener
	Errors []string
}

func (Listener) EnterProgram(c *p.ProgramContext) {
	fmt.Println("Got Text!", c.GetText())
}

func (l Listener) HasErrors() bool {
	return len(l.Errors) > 0
}

func main() {
	inputStream, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		panic(err)
	}

	lexer := p.NewCompiscriptLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := p.NewCompiscriptParser(stream)

	tree := parser.Program()

	walker := antlr.NewParseTreeWalker()
	listener := Listener{}
	walker.Walk(listener, tree)

	if listener.HasErrors() {
		for _, error := range listener.Errors {
			fmt.Fprintf(os.Stderr, "Type checking error: %s", error)
		}
	} else {
		fmt.Println("No type errors found!")
	}
}
