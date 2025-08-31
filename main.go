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

type ErrorListener struct {
	*antlr.DefaultErrorListener
	hadError bool
	errors   []string
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer,
	offendingSymbol interface{}, line, column int,
	msg string, e antlr.RecognitionException) {

	l.hadError = true

	errMsg := fmt.Sprintf("line %d:%d %s", line, column, msg)

	fmt.Println("Sintax error:", errMsg)

	// also collect for later
	l.errors = append(l.errors, errMsg)
}

func testableMain(reader io.Reader) error {
	inputStream := antlr.NewIoStream(reader)

	lexer := p.NewCompiscriptLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := p.NewCompiscriptParser(stream)

	// replace default console error listener
	parser.RemoveErrorListeners()
	errListener := &ErrorListener{}
	parser.AddErrorListener(errListener)

	tree := parser.Program()

	// bail early if syntax errors
	if errListener.hadError {
		return fmt.Errorf("syntax errors:\n%s", strings.Join(errListener.errors, "\n"))
	}

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
