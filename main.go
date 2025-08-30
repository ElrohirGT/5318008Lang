package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ElrohirGT/5318008Lang/listener"
	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/antlr4-go/antlr/v4"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

func generateErrorOutput(errors []string) error {
	b := strings.Builder{}
	for _, error := range errors {
		// b.WriteString("* ") <- Now added during analysis
		b.WriteString(error)
		b.WriteRune('\n')
	}
	return fmt.Errorf(Red+"=== ERRORS ===\n%s"+Reset, b.String())
}

func testableMain(reader io.Reader) error {
	inputStream := antlr.NewIoStream(reader)

	lexer := p.NewCompiscriptLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := p.NewCompiscriptParser(stream)

	tree := parser.Program()

	walker := antlr.NewParseTreeWalker()
	lis := listener.NewListener()
	walker.Walk(lis, tree)

	if lis.HasErrors() {
		return generateErrorOutput(*lis.Errors)
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
