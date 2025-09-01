package listener

import (
	"fmt"
	"slices"

	"github.com/antlr4-go/antlr/v4"
)

// Check if a given list of expresions are of the same type
func expresionsOfTheSameType(
	scopeManager *ScopeManager,
	expresions ...antlr.ParserRuleContext) (bool, TypeIdentifier) {

	if len(expresions) == 0 {
		return true, BASE_TYPES.UNKNOWN
	}

	generalType, available := scopeManager.CurrentScope.GetExpressionType(expresions[0].GetText())
	if len(expresions) == 1 {
		if !available {
			return false, BASE_TYPES.UNKNOWN
		}
		return true, generalType
	}

	for _, expr := range expresions[1:] {
		referenceType, available := scopeManager.CurrentScope.GetExpressionType(expr.GetText())
		if !available || referenceType != generalType {
			return false, BASE_TYPES.UNKNOWN
		}
	}
	return true, generalType
}

// Check if a given list of expresions are all of the same type,
// and this types can only be one of the specified on requiredTypes.
func expresionsOfTheRequiredType(
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
		for i, s := range requiredTypes {
			msg += string(s)
			if i != len(requiredTypes)-1 {
				msg += ", "
			}
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
