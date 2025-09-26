package tac_generator

import (
	"bytes"
	"fmt"
	"strings"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

// Typing system scanner, responsable of the semantic during compiscript code.
// Handles the notion of types, definitions and scope management.
type Listener struct {
	*p.BaseCompiscriptListener
	Errors *[]string
}

func NewListener() Listener {
	return Listener{
		Errors: &[]string{},
	}
}

func (l *Listener) Generate(buff *bytes.Buffer) {

}

// ERROR LOGGING
// ====================

// FIXME: Merge error logging to be the same as the one used for the other listener...
func (l Listener) AddError(line int, columnStart int, columnEnd int, content string, details ...string) {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(Red+"* Error: (line: %d, column: %d-%d) %s"+Reset, line, columnStart, columnEnd, content))

	for _, v := range details {
		b.WriteString("\n * " + v)
	}

	*l.Errors = append(*l.Errors, b.String())
}

func (l Listener) AddWarning(content string, line string, details ...string) {
	*l.Errors = append(*l.Errors, "Warning: "+content)
}

func (l Listener) HasErrors() bool {
	return len(*l.Errors) > 0
}
