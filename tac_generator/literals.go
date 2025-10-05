package tac_generator

import (
	"log"
	"strconv"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

// let a = [1,2,3];

func (l Listener) ExitArrayLiteral(ctx *p.ArrayLiteralContext) {
	arrayExpr := ctx.GetText()
	scope := l.GetCurrentScope()
	scopeName := ScopeName(scope.Name)
	expressions := ctx.AllConditionalExpr()

	arrayLength := len(expressions)
	scope.UpsertArrayLength(ctx.GetText(), uint(arrayLength))

	arrayType, found := scope.GetExpressionType(arrayExpr)
	if !found {
		log.Panicf(
			"Failed to get type for array: `%s`",
			arrayExpr,
		)
	}
	arrayTypeInfo, found := l.TypeChecker.GetTypeInfo(arrayType)
	if !found {
		log.Panicf(
			"Failed to get type information for array: `%s`",
			arrayType,
		)
	}

	elemType := arrayTypeInfo.ArrayType.GetValue().Type
	elemTypeInfo, found := l.TypeChecker.GetTypeInfo(elemType)
	if !found {
		log.Panicf(
			"Failed to get type information for array element: `%s`",
			elemType,
		)
	}

	log.Printf(
		"Array: `%s` has items of type: `%s` whose size is: `%d`",
		arrayExpr,
		elemType,
		elemTypeInfo.Size,
	)

	varName := l.Program.GetOrGenerateVariable(arrayExpr, ScopeName(scope.Name))
	allocSize := elemTypeInfo.Size * uint(arrayLength)

	l.AppendInstruction(
		scopeName,
		NewAllocInstruction(AllocInstruction{
			Target: varName,
			Size:   uint(allocSize),
		}),
	)

	for i, expr := range expressions {
		instValue := expr.GetText()
		if literalType, isLiteral := l.TypeChecker.GetLiteralType(instValue); isLiteral {
			_, instValue = literalToTAC(instValue, literalType)
		} else {
			varName, found = l.Program.GetVariableFor(instValue, ScopeName(scope.Name))
			if !found {
				log.Panicf(
					"Can't find variable for non literal value: `%s`",
					instValue,
				)
			}
			instValue = string(varName)
		}

		l.AppendInstruction(
			scopeName,
			NewSetWithOffsetInstruction(SetWithOffsetInstruction{
				Target: varName,
				Offset: LiteralOrVariable(strconv.FormatInt(int64(i*int(elemTypeInfo.Size)), 10)),
				Value:  LiteralOrVariable(instValue),
			}),
		)
	}
}
