package tac_generator

import (
	"bytes"

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
}

func NewListener() Listener {
	return Listener{}
}

// Generates the final TAC contents.
func (l *Listener) Generate(buff *bytes.Buffer) {
	// FIXME: Prince needs to fill this!
}
