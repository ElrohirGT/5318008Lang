package tac_generator

import (
	"testing"
)

func Test_ProgramVariableCounter(t *testing.T) {
	p := Program{}
	vars := [][]VariableName{}
	for range 50 {
		vars = append(vars, []VariableName{
			p.GetOrGenerateVariable("a", "GLOBAL"),
			VariableName("t1"),
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
