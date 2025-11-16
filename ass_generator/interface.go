package assgenerator

import (
	"bytes"
)

type AssGenerator interface {
	GenerateTo(*bytes.Buffer) error
}
