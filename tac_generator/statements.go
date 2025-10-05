package tac_generator

import (
	"log"
	"strconv"

	"github.com/ElrohirGT/5318008Lang/lib"
	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/ElrohirGT/5318008Lang/type_checker"
)

func (l Listener) ExitLeftHandSide(ctx *p.LeftHandSideContext) {
	log.Printf("Translating leftHandSide: %s", ctx.GetText())
	suffixes := ctx.AllSuffixOp()
	handleAtomAndSuffixes(l, ctx.PrimaryAtom().GetText(), &suffixes)
}

func (l Listener) ExitStandaloneExpresion(ctx *p.StandaloneExpresionContext) {
	log.Printf("Translating standalone expression: %s", ctx.GetText())
	suffixes := ctx.AllSuffixOp()
	handleAtomAndSuffixes(l, ctx.StandaloneAtom().GetText(), &suffixes)
}

func handleAtomAndSuffixes(l Listener, primaryExpr string, suffixes *[]p.ISuffixOpContext) {
	scope := l.GetCurrentScope()
	scopeName := ScopeName(scope.Name)
	if len(*suffixes) == 0 {
		return
	}

	previousExpr := primaryExpr
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

	for i, suffix := range *suffixes {
		varName := primaryExpr + getUntil(suffixes, i)
		tempName := l.Program.GetOrGenerateVariable(
			varName,
			scopeName,
		)

		switch suffixCtx := suffix.(type) {
		case *p.MethodCallExprContext:
		case *p.CallExprContext:
			paramCount := 0
			if args := suffixCtx.Arguments(); args != nil {
				paramCount = len(args.AllConditionalExpr())

				for _, arg := range args.AllConditionalExpr() {
					literalType, isLiteral := l.TypeChecker.GetLiteralType(arg.GetText())
					if isLiteral {
						_, literalValue := literalToTAC(arg.GetText(), literalType)
						l.AppendInstruction(NewParamInstruction(ParamInstruction{LiteralOrVariable(literalValue)}))
					} else {
						varName, found := l.Program.GetVariableFor(arg.GetText(), scopeName)
						if !found {
							log.Panicf(
								"Variable for expression: `%s` not found!",
								arg.GetText(),
							)
						}

						l.AppendInstruction(NewParamInstruction(ParamInstruction{LiteralOrVariable(varName)}))
					}
				}
			}

			funcInfo, found := scope.GetFunctionDef(previousExpr)
			if !found {
				log.Panicf(
					"Failed to find function info of: `%s`",
					previousExpr,
				)
			}
			returnNonNull := funcInfo.ReturnType != type_checker.BASE_TYPES.NULL
			switch funcInfo.ReturnType {
			case type_checker.BASE_TYPES.UNKNOWN, type_checker.BASE_TYPES.INVALID:
				log.Panicf(
					"Function `%s` should not be able to return `%s`!",
					previousExpr,
					funcInfo.ReturnType,
				)
			}

			saveOnReturn := lib.NewOpEmpty[VariableName]()
			if returnNonNull {
				saveOnReturn = lib.NewOpValue(tempName)
			}

			l.AppendInstruction(NewCallInstruction(CallInstruction{
				SaveReturnOn:   saveOnReturn,
				ProcedureName:  ScopeName(previousExpr),
				NumberOfParams: uint(paramCount),
			}))

			previousInChain = tempName
			if returnNonNull {
				returnTypeInfo, found := l.TypeChecker.GetTypeInfo(funcInfo.ReturnType)
				if !found {
					log.Panicf(
						"Failed to find return type info for function: `%s`",
						previousExpr,
					)
				}
				previousTypeInfo = returnTypeInfo
			}

		case *p.IndexExprContext:
			offset := "**INVALID OFFSET**"

			elemType := previousTypeInfo.ArrayType.GetValue().Type
			elemTypeInfo, found := l.TypeChecker.GetTypeInfo(elemType)
			if !found {
				log.Panicf(
					"Failed to find the type information for the element in:\n%s",
					varName,
				)
			}

			condExpr := suffixCtx.ConditionalExpr().GetText()
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
						l.AddError(
							suffixCtx.GetStart().GetLine(),
							suffixCtx.GetStart().GetColumn(),
							suffixCtx.GetStop().GetColumn(),
							"Can't index an array with a negative value!",
						)
						return
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

			previousInChain = tempName
			previousTypeInfo = elemTypeInfo
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
