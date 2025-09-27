package tac_generator

import (
	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/ElrohirGT/5318008Lang/type_checker"
)

// let a = 5;
// = t1 i32 5

// a = sumar(1,2);

func (l Listener) ExitAssignment(ctx *p.AssignmentContext) {
	// currentScope := l.TypeChecker.ScopeManager.CurrentScope
	var assignExpr p.IConditionalExprContext
	if ctx.ThisAssignment() != nil {
		assignExpr = ctx.ThisAssignment().ConditionalExpr()
	} else if ctx.VariableAssignment() != nil {
		assignExpr = ctx.VariableAssignment().ConditionalExpr()
	}

	literalType, found := l.TypeChecker.GetLiteralType(assignExpr.GetText())
	if !found {
		// FIXME: What do we do when the assignment is not a literal!?
	} else {
		literalValue := ""
		switch literalType {
		case type_checker.BASE_TYPES.BOOLEAN:
		case type_checker.BASE_TYPES.INTEGER:
			literalValue = assignExpr.GetText()
		case type_checker.BASE_TYPES.STRING:
		case type_checker.BASE_TYPES.NULL:
		default:
			// FIXME: It's an array! Handle array cases.
		}

		l.AppendInstruction(NewAssignmentInstruction(AssignmentInstruction{
			Target: l.Program.GetNextVariableName(),
			Type:   VARIABLE_TYPES.I32,
			Value:  LiteralOrVariable(literalValue),
		}))
	}
}
