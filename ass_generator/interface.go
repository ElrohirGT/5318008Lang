package assgenerator

import (
	"bytes"

	"github.com/ElrohirGT/5318008Lang/tac_generator"
)

type AssGenerator interface {
	GenerateTo(*tac_generator.Listener, *bytes.Buffer) error
}
