package tac_generator

import (
	"log"
	"strconv"

	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/ElrohirGT/5318008Lang/type_checker"
)

// a = 5;
// = t1 i32 5

// a = sumar(1,2);
// PARAM 1
// PARAM 2
// CALLRET a sumar 2

func (l Listener) ExitAssignment(ctx *p.AssignmentContext) {
	var assignExpr p.IConditionalExprContext
	if ctx.ThisAssignment() != nil {
		assignExpr = ctx.ThisAssignment().ConditionalExpr()
	} else if ctx.VariableAssignment() != nil {
		assignExpr = ctx.VariableAssignment().ConditionalExpr()
	}
	exprText := assignExpr.GetText()

	literalType, found := l.TypeChecker.GetLiteralType(exprText)
	if !found {
		// FIXME: What do we do when the assignment is not a literal!?
	} else {
		literalValue := "**SKILL ISSUE VALUE**"
		switch literalType {
		case type_checker.BASE_TYPES.BOOLEAN:
			switch exprText {
			case type_checker.LITERAL_VALUES.False:
				literalValue = "0"
			case type_checker.LITERAL_VALUES.True:
				literalValue = strconv.FormatInt(^0, 10)
			default:
				log.Panicf(
					"Expression: `%s`\nis of type `%s`\nbut it isn't `%s` nor `%s`",
					exprText,
					literalType,
					type_checker.LITERAL_VALUES.False,
					type_checker.LITERAL_VALUES.True,
				)
			}

			l.AppendInstruction(NewAssignmentInstruction(AssignmentInstruction{
				Target: l.Program.GetNextVariableName(),
				Type:   VARIABLE_TYPES.I8,
				Value:  LiteralOrVariable(literalValue),
			}))
		case type_checker.BASE_TYPES.INTEGER:
			literalValue = assignExpr.GetText()

			l.AppendInstruction(NewAssignmentInstruction(AssignmentInstruction{
				Target: l.Program.GetNextVariableName(),
				Type:   VARIABLE_TYPES.I32,
				Value:  LiteralOrVariable(literalValue),
			}))
		case type_checker.BASE_TYPES.STRING:
		case type_checker.BASE_TYPES.NULL:
		default:
			// FIXME: It's an array! Handle array cases.
		}

	}
}
