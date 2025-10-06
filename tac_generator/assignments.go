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

func (l Listener) ExitVariableDeclaration(ctx *p.VariableDeclarationContext) {
	scope := l.GetCurrentScope()
	scopeName := ScopeName(scope.Name)
	variableName := ctx.Identifier().GetText()

	exprText := ""
	if init := ctx.Initializer(); init != nil {
		exprText = init.ConditionalExpr().GetText()
	}

	if scope.Type == type_checker.SCOPE_TYPES.CLASS {
		scopeName = ScopeName(scope.Name + "_" + type_checker.CONSTRUCTOR_NAME)
		variableName = "this." + variableName
	}

	exprType, foundType := scope.GetExpressionType(variableName)
	if !foundType {
		log.Panicf(
			"Variable with name: `%s`\nnot found in current scope!%#v",
			variableName,
			*l.GetCurrentScope(),
		)
	}

	isLiteral := false
	variableValue := exprText
	if exprText == "" && exprType != type_checker.BASE_TYPES.STRING {
		isLiteral = true
		switch exprType {
		case type_checker.BASE_TYPES.BOOLEAN:
			variableValue = "0"
		case type_checker.BASE_TYPES.INTEGER:
			variableValue = "0"
		case type_checker.BASE_TYPES.INVALID, type_checker.BASE_TYPES.NULL, type_checker.BASE_TYPES.UNKNOWN:
			log.Panicf("Variable with name: `%s` has an invalid type! `%s`", variableName, exprType)
		}
	} else {
		_, isLiteral = l.TypeChecker.GetLiteralType(exprText)
	}

	log.Printf(
		"Declaring variable `%s`, with value: `%s`, which is a literal? %t",
		variableName,
		exprText,
		isLiteral,
	)

	createAssignment(l, scope, scopeName, isLiteral, variableName, exprType, variableValue, false)
}

func (l Listener) ExitThisAssignment(ctx *p.ThisAssignmentContext) {
	scope := l.GetCurrentScope()
	scopeName := ScopeName(scope.Name)

	parts := ctx.AllAssignmentPart()
	expr := ctx.ConditionalExpr()
	fieldName := ctx.Identifier().GetText()

	classScope, found := l.GetCurrentScope().SearchClassScope()
	if !found {
		log.Panicf(
			"Can't access a field `%s` outside of class scope!",
			fieldName,
		)
	}

	typeInfo, found := l.TypeChecker.GetTypeInfo(type_checker.TypeIdentifier(classScope.Name))
	if !found {
		log.Panicf(
			"Can't get type information of class: `%s`",
			classScope.Name,
		)
	}
	classInfo := typeInfo.ClassType.GetValue()

	previousType, found := classInfo.GetFieldType(fieldName, l.TypeChecker)
	if !found {
		log.Panicf(
			"Failed to get type for field: `%s`",
			fieldName,
		)
	}

	previousTypeInfo, found := l.TypeChecker.GetTypeInfo(previousType)
	if !found {
		log.Panicf(
			"Failed to get type information for type: `%s`",
			previousType,
		)
	}

	thisTacName, found := l.Program.GetVariableFor("this", scopeName)
	if !found {
		log.Panicf(
			"Failed to get tac variable for `this`, even though we're on a class scope!\nScope name: `%s`",
			scopeName,
		)
	}

	offsetTacName := l.Program.GetNextVariable()
	computedOffset := classInfo.GetFieldOffset(l.TypeChecker, fieldName)
	l.AppendInstruction(scopeName, NewAssignmentInstruction(AssignmentInstruction{
		Target: offsetTacName,
		Type:   VARIABLE_TYPES.I32,
		Value:  LiteralOrVariable(strconv.FormatUint(uint64(computedOffset), 10)),
	}))

	varNameUntilNow := "this." + fieldName
	for _, part := range parts {
		switch partCtx := part.(type) {
		case *p.FieldAssignmentPartExprContext:
			newFieldName := partCtx.Identifier().GetText()
			classInfo := previousTypeInfo.ClassType.GetValue()

			varNameUntilNow += partCtx.GetText()
			innerOffset := classInfo.GetFieldOffset(l.TypeChecker, newFieldName)
			l.AppendInstruction(scopeName, NewArithmethicInstruction(ArithmethicInstruction{
				Signed: false,
				Type:   ARITHMETHIC_OPERATION_TYPES.Add,
				Target: offsetTacName,
				P1:     LiteralOrVariable(offsetTacName),
				P2:     LiteralOrVariable(strconv.FormatUint(uint64(innerOffset), 10)),
			}))

			previousType, found = classInfo.GetFieldType(newFieldName, l.TypeChecker)
			if !found {
				log.Panicf(
					"Can't find field `%s` for class `%s`",
					newFieldName,
					classInfo.Name,
				)
			}

			previousTypeInfo, found = l.TypeChecker.GetTypeInfo(previousType)
			if !found {
				log.Panicf(
					"Can't find type information for: `%s`",
					previousType,
				)
			}

		case *p.IndexAssignmentPartExprContext:
			offset := "**INVALID OFFSET**"

			elemType := previousTypeInfo.ArrayType.GetValue().Type
			elemTypeInfo, found := l.TypeChecker.GetTypeInfo(elemType)
			if !found {
				log.Panicf(
					"Failed to find the type information for the element in:\n%s",
					varNameUntilNow,
				)
			}

			condExpr := partCtx.ConditionalExpr().GetText()
			litType, found := l.TypeChecker.GetLiteralType(condExpr)
			if !found {
				varName, found := l.Program.GetVariableFor(condExpr, scopeName)
				if !found {
					log.Panicf("Can't find the result variable for expr: `%s`", condExpr)
				}
				offset = string(varName)
			} else {
				if litType != type_checker.BASE_TYPES.INTEGER {
					log.Panicf("Can't index with type that is not integer: %s", litType)
				} else {
					idx, err := strconv.ParseInt(condExpr, 10, 64)
					if err != nil {
						log.Panicf(
							"Failed to parse integer literal: `%s`",
							condExpr,
						)

					}

					if idx < 0 {
						l.AddError(
							partCtx.GetStart().GetLine(),
							partCtx.GetStart().GetColumn(),
							partCtx.GetStop().GetColumn(),
							"Can't index an array with a negative value!",
						)
						return
					}

					elemSize := int64(elemTypeInfo.Size)
					if elemSize <= 0 {
						log.Panicf(
							"The element size is 0 or negative! But it should be at least 1.\nElem type: %s\n%s",
							elemType,
							varNameUntilNow,
						)
					}
					offset = strconv.FormatInt(elemSize*idx, 10)
				}
			}

			varNameUntilNow += partCtx.GetText()
			l.AppendInstruction(scopeName, NewArithmethicInstruction(ArithmethicInstruction{
				Signed: false,
				Type:   ARITHMETHIC_OPERATION_TYPES.Add,
				Target: offsetTacName,
				P1:     LiteralOrVariable(offsetTacName),
				P2:     LiteralOrVariable(offset),
			}))

			previousType = elemType
			previousTypeInfo = elemTypeInfo
		}
	}

	exprTac := "**INVALID EXPRESSION TAC**"
	exprText := expr.GetText()
	if literalType, isLiteral := l.TypeChecker.GetLiteralType(exprText); isLiteral {
		_, exprTac = literalToTAC(exprText, literalType)
	} else {
		exprTacName, found := l.Program.GetVariableFor(exprText, scopeName)
		if !found {
			log.Panicf(
				"Failed to get TAC variable for expression: `%s`",
				exprText,
			)
		}

		exprTac = string(exprTacName)
	}

	l.AppendInstruction(scopeName, NewSetWithOffsetInstruction(SetWithOffsetInstruction{
		Target: thisTacName,
		Offset: LiteralOrVariable(offsetTacName),
		Value:  LiteralOrVariable(exprTac),
	}))
}

func (l Listener) ExitVariableAssignment(ctx *p.VariableAssignmentContext) {
	scope := l.GetCurrentScope()
	scopeName := ScopeName(scope.Name)

	varName := ctx.Identifier().GetText()

	parts := ctx.AllAssignmentPart()
	expr := ctx.ConditionalExpr()

	varNameUntilNow := ""
	previousType, found := scope.GetExpressionType(varName)
	if !found {
		log.Panicf(
			"Failed to get type for expression: `%s`",
			varName,
		)
	}

	previousTypeInfo, found := l.TypeChecker.GetTypeInfo(previousType)
	if !found {
		log.Panicf(
			"Failed to get type information for type: `%s`",
			previousType,
		)
	}

	previousTacName, found := l.Program.GetVariableFor(varName, scopeName)
	if !found {
		log.Panicf(
			"Failed to get tac name for variable: %s",
			varName,
		)
	}
	varNameUntilNow += varName

	for _, part := range parts {
		switch partCtx := part.(type) {
		case *p.FieldAssignmentPartExprContext:
			newFieldName := partCtx.Identifier().GetText()
			classInfo := previousTypeInfo.ClassType.GetValue()

			computedOffset := classInfo.GetFieldOffset(l.TypeChecker, newFieldName)
			offset := strconv.FormatUint(uint64(computedOffset), 10)

			varNameUntilNow += partCtx.GetText()
			target := l.Program.GetOrGenerateVariable(varNameUntilNow, scopeName)
			l.AppendInstruction(scopeName, NewLoadWithOffsetInstruction(LoadWithOffsetInstruction{
				Target: target,
				Source: previousTacName,
				Offset: LiteralOrVariable(offset),
			}))

			previousTacName = target
			previousType, found = classInfo.GetFieldType(newFieldName, l.TypeChecker)
			if !found {
				log.Panicf(
					"Can't find field `%s` for class `%s`",
					newFieldName,
					classInfo.Name,
				)
			}

			previousTypeInfo, found = l.TypeChecker.GetTypeInfo(previousType)
			if !found {
				log.Panicf(
					"Can't find type information for: `%s`",
					previousType,
				)
			}

		case *p.IndexAssignmentPartExprContext:
			offset := "**INVALID OFFSET**"

			elemType := previousTypeInfo.ArrayType.GetValue().Type
			elemTypeInfo, found := l.TypeChecker.GetTypeInfo(elemType)
			if !found {
				log.Panicf(
					"Failed to find the type information for the element in:\n%s",
					varNameUntilNow,
				)
			}

			condExpr := partCtx.ConditionalExpr().GetText()
			litType, found := l.TypeChecker.GetLiteralType(condExpr)
			if !found {
				varName, found := l.Program.GetVariableFor(condExpr, scopeName)
				if !found {
					log.Panicf("Can't find the result variable for expr: `%s`", condExpr)
				}
				offset = string(varName)
			} else {
				if litType != type_checker.BASE_TYPES.INTEGER {
					log.Panicf("Can't index with type that is not integer: %s", litType)
				} else {
					idx, err := strconv.ParseInt(condExpr, 10, 64)
					if err != nil {
						log.Panicf(
							"Failed to parse integer literal: `%s`",
							condExpr,
						)

					}

					if idx < 0 {
						l.AddError(
							partCtx.GetStart().GetLine(),
							partCtx.GetStart().GetColumn(),
							partCtx.GetStop().GetColumn(),
							"Can't index an array with a negative value!",
						)
						return
					}

					elemSize := int64(elemTypeInfo.Size)
					if elemSize <= 0 {
						log.Panicf(
							"The element size is 0 or negative! But it should be at least 1.\nElem type: %s\n%s",
							elemType,
							varNameUntilNow,
						)
					}
					offset = strconv.FormatInt(elemSize*idx, 10)
				}
			}

			varNameUntilNow += partCtx.GetText()
			target := l.Program.GetOrGenerateVariable(varNameUntilNow, scopeName)
			l.AppendInstruction(scopeName, NewLoadWithOffsetInstruction(LoadWithOffsetInstruction{
				Target: target,
				Source: previousTacName,
				Offset: LiteralOrVariable(offset),
			}))

			previousTacName = target
			previousType = elemType
			previousTypeInfo = elemTypeInfo
		}
	}

	exprText := expr.GetText()
	exprType, found := scope.GetExpressionType(exprText)
	if !found {
		log.Panicf(
			"Failed to get type for expression: `%s`",
			exprText,
		)
	}
	_, isLiteral := l.TypeChecker.GetLiteralType(exprText)

	log.Printf(
		"Assignment to variable:\n`%s`\nwith value: `%s`\nwhich is a literal? %t",
		varNameUntilNow,
		expr.GetText(),
		isLiteral,
	)

	createAssignment(l, scope, scopeName, isLiteral, varNameUntilNow, exprType, exprText, true)
}

func createAssignment(
	l Listener,
	scope *type_checker.Scope,
	scopeName ScopeName,
	isLiteral bool,
	variableName string,
	exprType type_checker.TypeIdentifier,
	exprText string,
	isAssigment bool, // Flag to decide  if its and assignmetn or not
) {
	if isLiteral {
		literalType, literalValue := literalToTAC(exprText, exprType)
		// l.CreateAssignment(scopeName, variableName, literalType, literalValue)
		if isAssigment {
			// fmt.Println("THIS IS AN ASIGMENT OF:" + variableName)
			l.CreateVariableAssignment(scopeName, variableName, literalType, literalValue)
		} else {
			// fmt.Println("THIS IS A DECLARATION OF:" + variableName + " " + string(scopeName))
			l.CreateVariableDeclaration(scopeName, variableName, literalType, literalValue)
		}
	} else if exprText == "" && exprType == type_checker.BASE_TYPES.STRING {
		strRef := l.Program.GetOrGenerateVariable("EMPTY STRING", scopeName)
		l.AppendInstruction(scopeName, NewAllocInstruction(AllocInstruction{
			Target: strRef,
			Size:   1,
		}))
		l.AppendInstruction(scopeName, NewSetWithOffsetInstruction(SetWithOffsetInstruction{
			Target: strRef,
			Offset: "0",
			Value:  "0",
		}))
		l.Program.UpsertTranslation(scopeName, variableName, strRef)
	} else {
		exprVar, found := l.Program.GetVariableFor(exprText, scopeName)
		if !found {
			log.Panicf("Failed to find a variable for the expression:\n`%s`", exprText)
		}

		if length, found := scope.GetArrayLength(exprText); found {
			scope.UpsertArrayLength(variableName, length)
			l.Program.UpsertTranslation(scopeName, variableName, exprVar)
		} else if tacType, found := l.MapBaseTypeToTacType(exprType); found {
			if isAssigment {
				l.CreateVariableAssignment(scopeName, variableName, tacType, string(exprVar))
			} else {
				l.CreateVariableDeclaration(scopeName, variableName, tacType, string(exprVar))
			}
		} else {
			l.Program.UpsertTranslation(scopeName, variableName, exprVar)
		}
	}
}
