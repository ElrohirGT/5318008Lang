package main

import (
	"fmt"

	p "github.com/ElrohirGT/5318008Lang/parser"
	s "github.com/ElrohirGT/5318008Lang/semantic"
	"github.com/antlr4-go/antlr/v4"
)

type Listener struct{}

func (Listener) EnterProgram(c *p.ProgramContext) {
}

func main() {
  	symbols := s.NewSymbolTable()
  	symbols.Variables["x"] = s.Content{DataType: s.INTEGER}
	symbols.Variables["y"] = s.Content{DataType: s.INTEGER}

  	input := "do {} while (x > 5 || (y == y));"

  	is := antlr.NewInputStream(input)
  	lexer := p.NewCompiscriptLexer(is)
  	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
  	p := p.NewCompiscriptParser(tokens)

  	tree := p.DoWhileStatement()

  	listener := s.NewSemanticListener(symbols)
  	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

  	if len(listener.Errors) > 0 {
		fmt.Println("Errores semánticos:")
		for _, err := range listener.Errors {
			fmt.Println(" -", err)
		}
	} else {
		fmt.Println("Análisis semántico sin errores")
	}
}
