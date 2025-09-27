package tac_generator

import (
	"strconv"

	p "github.com/ElrohirGT/5318008Lang/parser"
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

	literalType, literalValue := l.GetLiteralType(assignExpr.GetText())
	if literalType != LITERAL_TYPES.NotALiteral {
		// FIXME: What do we do when the assignment is not a literal!?
	} else {
		value := "**INVALID**"
		switch innerValue := literalValue.(type) {
		case string:
			value = innerValue
		case int64:
			value = strconv.FormatInt(innerValue, 10)
		case bool:
			value = strconv.FormatBool(innerValue)
		}

		l.AppendInstruction(NewAssignmentInstruction(AssignmentInstruction{
			Target: l.Program.GetNextVariableName(),
			Type:   VARIABLE_TYPES.I32,
			Value:  LiteralOrVariable(value),
		}))
	}
}
