package applib

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/ElrohirGT/5318008Lang/lib"
	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/ElrohirGT/5318008Lang/tac_generator"
	"github.com/ElrohirGT/5318008Lang/type_checker"
	"github.com/antlr4-go/antlr/v4"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Grey   = "\033[90m"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
	hadError bool
	errors   []string
}

func generateErrorOutput(errors []string) error {
	b := strings.Builder{}
	for _, error := range errors {
		// b.WriteString("* ") <- Now added during analysis
		b.WriteString(error)
		b.WriteRune('\n')
	}
	return fmt.Errorf(Red+"=== ERRORS ===\n%s"+Reset, b.String())
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer,
	offendingSymbol any, line, column int,
	msg string, e antlr.RecognitionException) {

	l.hadError = true

	errMsg := fmt.Sprintf("line %d:%d %s", line, column, msg)

	fmt.Println("Sintax error:", errMsg)

	// also collect for later
	l.errors = append(l.errors, errMsg)
}

type CompilerConfig struct {
	TACBuffer lib.Optional[*bytes.Buffer]
	ASMBuffer lib.Optional[*bytes.Buffer]
}

func TestableMain(reader io.Reader, config CompilerConfig) error {
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
	typeListener := type_checker.NewListener()
	walker.Walk(typeListener, tree)

	if typeListener.HasErrors() {
		return generateErrorOutput(*typeListener.Errors)
	}

	tacListener := tac_generator.NewListener()
	walker.Walk(tacListener, tree)

	// FIXME: Does TAC generation should have errors?
	// if tacListener.HasErrors() {
	// 	return generateErrorOutput(*tacListener.Errors)
	// }

	if config.TACBuffer.HasValue() {
		tacListener.Generate(config.TACBuffer.GetValue())
	}

	return nil
}
