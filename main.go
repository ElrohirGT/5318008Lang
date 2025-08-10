package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/antlr4-go/antlr/v4"
)

func generateErrorOutput(errors []string) error {
	b := strings.Builder{}
	for _, error := range errors {
		b.WriteString("* ")
		b.WriteString(error)
		b.WriteRune('\n')
	}
	return fmt.Errorf("=== ERRORS ===\n%s", b.String())
}

func testableMain(reader io.Reader) error {
	inputStream := antlr.NewIoStream(reader)

	lexer := p.NewCompiscriptLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := p.NewCompiscriptParser(stream)

	tree := parser.Program()

	walker := antlr.NewParseTreeWalker()
	listener := NewListener()
	walker.Walk(listener, tree)

	if listener.HasErrors() {
		return generateErrorOutput(*listener.Errors)
	}

	return nil
}

func main() {
	filePath := os.Args[1]
	reader, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	err = testableMain(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println("No type errors found!")
}
