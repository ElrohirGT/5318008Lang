package tac_generator

import (
	"log"
	"strconv"

	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/ElrohirGT/5318008Lang/type_checker"
)

func (l Listener) ExitStandaloneExpresion(ctx *p.StandaloneExpresionContext) {
	scopeName := ScopeName(l.GetCurrentScope().Name)
	suffixes := ctx.AllSuffixOp()
	if len(suffixes) == 0 {
		return
	}

	log.Printf("Translation standalone expression:\n%s", ctx.GetText())
	primaryExpr := ctx.StandaloneAtom().GetText()
	previousExpr := ctx.StandaloneAtom().GetText()

	previousType, found := l.TypeChecker.ScopeManager.CurrentScope.GetExpressionType(previousExpr)
	if !found {
		log.Panicf("Can't find type of expression: `%s`", previousExpr)
	}
	previousTypeInfo, found := l.TypeChecker.GetTypeInfo(previousType)
	if !found {
		log.Panicf("Can't find the type information of: `%s`", previousType)
	}

	previousInChain := l.Program.GetOrGenerateVariable(
		previousExpr,
		scopeName,
	)

	for i, suffix := range suffixes {
		switch suffixCtx := suffix.(type) {
		case *p.MethodCallExprContext:
		case *p.CallExprContext:
		case *p.IndexExprContext:
			varName := primaryExpr + getUntil(&suffixes, i)
			tempName := l.Program.GetOrGenerateVariable(
				varName,
				scopeName,
			)
			offset := "**INVALID OFFSET**"

			condExpr := suffixCtx.ConditionalExpr().GetText()
			litType, found := l.TypeChecker.GetLiteralType(condExpr)
			if !found {
				// TODO: Math with register offset is dynamic
			} else {
				if litType != type_checker.BASE_TYPES.INTEGER {
					log.Panicf("Can't index with type that is not integer: %s", litType)
				} else {
					idx, err := strconv.ParseInt(condExpr, 10, 64)
					if err != nil {
						l.AddError(
							suffixCtx.GetStart().GetLine(),
							suffixCtx.GetStart().GetColumn(),
							suffixCtx.GetStop().GetColumn(),
							"Can't index an array with a negative value!",
						)
						return
					}
					elemType := previousTypeInfo.ArrayType.GetValue().Type
					elemTypeInfo, found := l.TypeChecker.GetTypeInfo(elemType)
					if !found {
						log.Panicf(
							"Failed to find the type information for the element `%d` in:\n%s",
							idx,
							varName,
						)
					}

					elemSize := int64(elemTypeInfo.Size)
					if elemSize <= 0 {
						log.Panicf(
							"The element size is 0 or negative! But it should be at least 1.\nElem type: %s\n%s",
							elemType,
							varName,
						)
					}
					offset = strconv.FormatInt(elemSize*idx, 10)
				}
			}

			l.AppendInstruction(NewLoadWithOffsetInstruction(LoadWithOffsetInstruction{
				Target: tempName,
				Source: previousInChain,
				Offset: LiteralOrVariable(offset),
			}))
		case *p.PropertyAccessExprContext:
		}
	}
}

func getUntil(suffixes *[]p.ISuffixOpContext, maxIdx int) string {
	buff := ""
	for i := range maxIdx + 1 {
		buff += (*suffixes)[i].GetText()
	}
	return buff
}
