package tac_generator

import (
	"strconv"
	"testing"
)

func Test_ProgramVariableCounter(t *testing.T) {
	p := Program{}
	vars := [][]VariableName{}
	for i := range 50 {
		vars = append(vars, []VariableName{
			p.GetNextVariableName(),
			VariableName("T" + strconv.FormatUint(uint64(i+1), 10)),
		})
	}

	for _, pair := range vars {
		actual := pair[0]
		expected := pair[1]

		if actual != expected {
			t.Errorf("(`%s` != `%s`)", actual, expected)
		}
	}
}
