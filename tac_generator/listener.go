package tac_generator

import (
	"bytes"
	"fmt"
	"strings"

	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/ElrohirGT/5318008Lang/type_checker"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

// Typing system scanner, responsable of the semantic during compiscript code.
// Handles the notion of types, definitions and scope management.
type Listener struct {
	*p.BaseCompiscriptListener
	TypeChecker *type_checker.Listener
	Program     *Program
	Errors      *[]string
}

func NewListener(typeChecker *type_checker.Listener) Listener {
	return Listener{
		Program:     NewProgram(),
		TypeChecker: typeChecker,
		Errors:      &[]string{},
	}
}

// Generates the final TAC contents.
func (l *Listener) Generate(buff *bytes.Buffer) error {
	// FIXME: Prince needs to fill this!
	return nil
}

func (l *Listener) AppendInstruction(inst Instruction) {
	currentScope := l.TypeChecker.ScopeManager.CurrentScope
	scopeInstructions := l.Program.Scopes[ScopeName(currentScope.Name)]
	scopeInstructions = append(scopeInstructions, inst)
	l.Program.Scopes[ScopeName(currentScope.Name)] = scopeInstructions
}

// FIXME: Improve error handling:
// Centralize error handling in lib maybe?
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
