package tac_generator

import (
	"log"
	"strconv"

	"github.com/ElrohirGT/5318008Lang/type_checker"
)

func literalToTAC(literal string, literalType type_checker.TypeIdentifier) (VariableType, string) {
	literalValue := "**SKILL ISSUE VALUE**"
	varType := VARIABLE_TYPES.I32

	switch literalType {
	case type_checker.BASE_TYPES.BOOLEAN:
		varType = VARIABLE_TYPES.I8
		switch literal {
		case type_checker.LITERAL_VALUES.False:
			literalValue = "0"
		case type_checker.LITERAL_VALUES.True:
			literalValue = strconv.FormatInt(^0, 10)
		default:
			log.Panicf(
				"Expression: `%s`\nis of type `%s`\nbut it isn't `%s` nor `%s`",
				literal,
				literalType,
				type_checker.LITERAL_VALUES.False,
				type_checker.LITERAL_VALUES.True,
			)
		}
	case type_checker.BASE_TYPES.STRING:
		return VARIABLE_TYPES.I32, literal
	case type_checker.BASE_TYPES.INTEGER:
		literalValue = literal
	case type_checker.BASE_TYPES.NULL, type_checker.BASE_TYPES.INVALID, type_checker.BASE_TYPES.UNKNOWN:
		log.Panicf(
			"Literal expression: `%s` is of invalid type! `%s`",
			literal,
			literalType,
		)
	default:
		log.Panicf(
			"You shouldn't create an assignment for the type: `%s`\nCheckout the array example to see what should be done!",
			literalType,
		)
	}

	return varType, literalValue
}
