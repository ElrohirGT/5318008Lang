package tac_generator

import (
	"bytes"
	"strconv"

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

type LiteralType string

var LITERAL_TYPES = struct {
	Integer     LiteralType
	String      LiteralType
	Bool        LiteralType
	NotALiteral LiteralType
}{
	Integer:     "INT",
	String:      "STR",
	Bool:        "BOOL",
	NotALiteral: "NAL",
}

func (l *Listener) GetLiteralType(expression string) (LiteralType, any) {
	n, err := strconv.ParseInt(expression, 10, 64)
	if err == nil {
		return LITERAL_TYPES.Integer, n
	}

	t, err := strconv.ParseBool(expression)
	if err == nil {
		return LITERAL_TYPES.Bool, t
	}

	return LITERAL_TYPES.NotALiteral, nil
}
