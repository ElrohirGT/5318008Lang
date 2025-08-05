package main

import (
	"fmt"
	"os"

	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/antlr4-go/antlr/v4"
)

type Listener struct{}

func (Listener) EnterProgram(c *p.ProgramContext) {
	fmt.Println("Got Text!", c.GetText())
}

func main() {
	inputStream, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		panic(err)
	}

	lexer := p.NewCompiscriptLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := p.NewCompiscriptParser(stream)

	parser.Program()
}
