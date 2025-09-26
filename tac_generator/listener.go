package tac_generator

import (
	"bytes"

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
}

func NewListener(typeChecker *type_checker.Listener) Listener {
	return Listener{
		Program:     NewProgram(),
		TypeChecker: typeChecker,
	}
}

// Generates the final TAC contents.
func (l *Listener) Generate(buff *bytes.Buffer) {
	// FIXME: Prince needs to fill this!
}

func (l *Listener) AppendInstruction(inst Instruction) {
	currentScope := l.TypeChecker.ScopeManager.CurrentScope
	scopeInstructions := l.Program.Scopes[ScopeName(currentScope.Name)]
	scopeInstructions = append(scopeInstructions, inst)
	l.Program.Scopes[ScopeName(currentScope.Name)] = scopeInstructions
}
