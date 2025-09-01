package listener

import (
	"fmt"
	"slices"

	"github.com/antlr4-go/antlr/v4"
)

func expresionOfTheSameType(
	requiredTypes []TypeIdentifier,
	scopeManager *ScopeManager,
	expresions ...antlr.ParserRuleContext) []string {

	errors := make([]string, 0, len(expresions))

	if len(expresions) == 0 {
		return errors
	}

	generalType, available := scopeManager.CurrentScope.GetExpressionType(expresions[0].GetText())
	if !available {
		msg := fmt.Sprintf("`%s` symbol is not registered in scope!", expresions[0].GetText())
		errors = append(errors, msg)
		return errors
	}

	if !slices.Contains(requiredTypes, generalType) {
		msg := fmt.Sprintf("expresion `%s` should be a ", expresions[0].GetText())
		for _, s := range requiredTypes {
			msg += string(s) + " "
		}
		errors = append(errors, msg)
		return errors
	}

	for _, expr := range expresions[1:] {
		referenceType, available := scopeManager.CurrentScope.GetExpressionType(expr.GetText())
		if !available {
			msg := fmt.Sprintf("`%s` symbol is not registered in scope!", expr.GetText())
			errors = append(errors, msg)
		} else if referenceType == BASE_TYPES.UNKNOWN {
			msg := fmt.Sprintf("`%s` doesn't have a type!", expr.GetText())
			errors = append(errors, msg)
		} else if referenceType == BASE_TYPES.INVALID {
			msg := fmt.Sprintf("`%s` has an invalid type!", expr.GetText())
			errors = append(errors, msg)
		} else if referenceType != generalType {
			msg := fmt.Sprintf("`%s` should be type `%s`", expr.GetText(), string(generalType))
			errors = append(errors, msg)
		}
	}
	return errors
}
