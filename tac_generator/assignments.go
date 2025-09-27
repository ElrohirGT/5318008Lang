package tac_generator

import (
	p "github.com/ElrohirGT/5318008Lang/parser"
)

func (l Listener) ExitAssignment(ctx *p.AssignmentContext) {
	// currentScope := l.TypeChecker.ScopeManager.CurrentScope
	l.AppendInstruction(NewAssignmentInstruction(AssignmentInstruction{
		Target: l.Program.GetNextVariableName(),
		Type:   VARIABLE_TYPES.I32,
	}))
}
