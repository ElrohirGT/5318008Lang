package tac_generator

import (
	"log"
	"strconv"

	"github.com/ElrohirGT/5318008Lang/lib"
	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/ElrohirGT/5318008Lang/type_checker"
	"github.com/antlr4-go/antlr/v4"
)

func (l Listener) ExitLeftHandSide(ctx *p.LeftHandSideContext) {
	log.Printf("Translating leftHandSide: %s", ctx.GetText())
	suffixes := ctx.AllSuffixOp()
	handleAtomAndSuffixes(l, ctx.PrimaryAtom(), &suffixes)
}

func (l Listener) ExitStandaloneExpresion(ctx *p.StandaloneExpresionContext) {
	log.Printf("Translating standalone expression: %s", ctx.GetText())
	suffixes := ctx.AllSuffixOp()
	handleAtomAndSuffixes(l, ctx.StandaloneAtom(), &suffixes)
}

func handleAtomAndSuffixes(l Listener, primaryCtx any, suffixes *[]p.ISuffixOpContext) {
	scope := l.GetCurrentScope()
	scopeName := ScopeName(scope.Name)

	primaryExpr := "**INVALID PRIMARY EXPR**"
	switch innerCtx := primaryCtx.(type) {
	case *p.IdentifierExprContext:
		primaryExpr = innerCtx.GetText()
		if len(*suffixes) == 0 {
			return
		}
	case *p.StandaloneIdentifierExprContext:
		primaryExpr = innerCtx.GetText()
		if len(*suffixes) == 0 {
			return
		}
	case *p.ThisExprContext:
		primaryExpr = innerCtx.GetText()
	case *p.StandaloneThisExprContext:
		primaryExpr = innerCtx.GetText()

	case *p.NewExprContext:
		primaryExpr = innerCtx.GetText()
		handleConstructorCall(l, primaryExpr, scopeName, innerCtx.Identifier(), innerCtx.Arguments())
	case *p.StandaloneNewExprContext:
		primaryExpr = innerCtx.GetText()
		handleConstructorCall(l, primaryExpr, scopeName, innerCtx.Identifier(), innerCtx.Arguments())
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
			paramCount := 0
			if args := suffixCtx.Arguments(); args != nil {
				paramCount = len(args.AllConditionalExpr())

				for _, arg := range args.AllConditionalExpr() {
					literalType, isLiteral := l.TypeChecker.GetLiteralType(arg.GetText())
					if isLiteral {
						_, literalValue := literalToTAC(arg.GetText(), literalType)
						l.AppendInstruction(scopeName, NewParamInstruction(ParamInstruction{LiteralOrVariable(literalValue)}))
					} else {
						varName, found := l.Program.GetVariableFor(arg.GetText(), scopeName)
						if !found {
							log.Panicf(
								"Variable for expression: `%s` not found!",
								arg.GetText(),
							)
						}

						l.AppendInstruction(scopeName, NewParamInstruction(ParamInstruction{LiteralOrVariable(varName)}).AddComment("("+arg.GetText()+")"))
					}
				}
			}

			classInfo := previousTypeInfo.ClassType.GetValue()
			methodName := suffixCtx.Identifier().GetText()
			funcInfo, found := classInfo.Methods[methodName]
			if !found {
				log.Panicf(
					"Failed to find method info of: `%s`",
					methodName,
				)
			}

			returnNonNull := funcInfo.ReturnType != type_checker.BASE_TYPES.NULL && funcInfo.ReturnType != type_checker.BASE_TYPES.UNKNOWN
			switch funcInfo.ReturnType {
			// FIXME: Unknown should be a type to throw but we're going to accept Unknown too
			// case type_checker.BASE_TYPES.UNKNOWN, type_checker.BASE_TYPES.INVALID:
			case type_checker.BASE_TYPES.INVALID:
				log.Panicf(
					"Function `%s` should not be able to return `%s`!",
					methodName,
					funcInfo.ReturnType,
				)
			}

			saveOnReturn := lib.NewOpEmpty[VariableName]()
			if returnNonNull {
				saveOnReturn = lib.NewOpValue(tempName)
			}

			l.AppendInstruction(scopeName, NewParamInstruction(ParamInstruction{
				Parameter: LiteralOrVariable(previousInChain),
			}).AddComment("(previousInChain := "+primaryExpr+getUntil(suffixes, i-1)+")"))

			l.AppendInstruction(scopeName, NewCallInstruction(CallInstruction{
				SaveReturnOn:   saveOnReturn,
				ProcedureName:  ScopeName(string(previousType) + "_" + methodName),
				NumberOfParams: uint(paramCount),
			}).AddComment("("+varName+")"))

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

		case *p.CallExprContext:
			paramCount := 0
			if args := suffixCtx.Arguments(); args != nil {
				paramCount = len(args.AllConditionalExpr())

				for _, arg := range args.AllConditionalExpr() {
					literalType, isLiteral := l.TypeChecker.GetLiteralType(arg.GetText())
					if isLiteral {
						_, literalValue := literalToTAC(arg.GetText(), literalType)
						l.AppendInstruction(scopeName, NewParamInstruction(ParamInstruction{LiteralOrVariable(literalValue)}))
					} else {
						varName, found := l.Program.GetVariableFor(arg.GetText(), scopeName)
						if !found {
							log.Panicf(
								"Variable for expression: `%s` not found!",
								arg.GetText(),
							)
						}

						l.AppendInstruction(scopeName, NewParamInstruction(ParamInstruction{LiteralOrVariable(varName)}).AddComment("("+arg.GetText()+")"))
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

			l.AppendInstruction(scopeName, NewCallInstruction(CallInstruction{
				SaveReturnOn:   saveOnReturn,
				ProcedureName:  ScopeName(previousExpr),
				NumberOfParams: uint(paramCount),
			}).AddComment("("+varName+")"))

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
						log.Panicf(
							"Failed to parse integer literal: `%s`",
							condExpr,
						)

					}

					if idx < 0 {
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

			l.AppendInstruction(scopeName, NewLoadWithOffsetInstruction(LoadWithOffsetInstruction{
				Target: tempName,
				Source: previousInChain,
				Offset: LiteralOrVariable(offset),
			}))

			previousInChain = tempName
			previousType = elemType
			previousTypeInfo = elemTypeInfo
		case *p.PropertyAccessExprContext:
			fieldName := suffixCtx.Identifier().GetText()
			classInfo := previousTypeInfo.ClassType.GetValue()

			computedOffset := classInfo.GetFieldOffset(l.TypeChecker, fieldName)
			offset := strconv.FormatUint(uint64(computedOffset), 10)

			l.AppendInstruction(scopeName, NewLoadWithOffsetInstruction(LoadWithOffsetInstruction{
				Target: tempName,
				Source: previousInChain,
				Offset: LiteralOrVariable(offset),
			}))

			previousInChain = tempName

			fieldType, found := classInfo.GetFieldType(fieldName, l.TypeChecker)
			if !found {
				log.Panicf(
					"Failed to find type for field `%s` of class `%s`",
					fieldName,
					classInfo.Name,
				)
			}
			previousType = fieldType

			fieldTypeInfo, found := l.TypeChecker.GetTypeInfo(fieldType)
			if !found {
				log.Panicf(
					"Failed to find the type information for field `%s` of class `%s`",
					fieldName,
					classInfo.Name,
				)
			}
			previousTypeInfo = fieldTypeInfo
		}
	}
}

func handleConstructorCall(
	l Listener,
	completeExpr string,
	scopeName ScopeName,
	identifier antlr.TerminalNode,
	args p.IArgumentsContext,
) {
	className := identifier.GetText()
	typeInfo, found := l.TypeChecker.GetTypeInfo(type_checker.TypeIdentifier(className))
	if !found {
		log.Panicf(
			"Failed to find class (%s) type information! When translating constructor.",
			className,
		)
	}
	log.Printf("Constructing instance of: `%s`", className)

	classRefTac := l.Program.GetOrGenerateVariable(completeExpr, scopeName)
	log.Printf("Variable for: `%s` is `%s`", completeExpr, classRefTac)
	l.AppendInstruction(scopeName, NewAllocInstruction(AllocInstruction{
		Target: classRefTac,
		Size:   typeInfo.Size,
	}))

	var argCount uint = 1 // The default argument is the class ref!
	if args != nil {
		argExprs := args.AllConditionalExpr()
		argCount += uint(len(argExprs))
		maxIdx := len(argExprs) - 1
		for idx := maxIdx; idx >= 0; idx -= 1 {
			currentArg := argExprs[idx].GetText()

			var comment string
			var argValue string
			if literalType, isLiteral := l.TypeChecker.GetLiteralType(currentArg); isLiteral {
				_, argValue = literalToTAC(currentArg, literalType)
			} else {
				comment = "(" + currentArg + ")"
				tacName, found := l.Program.GetVariableFor(currentArg, scopeName)
				if !found {
					log.Panicf(
						"Failed to find tac variable for expression: `%s`",
						currentArg,
					)
				}

				argValue = string(tacName)
			}

			l.AppendInstruction(scopeName, NewParamInstruction(ParamInstruction{LiteralOrVariable(argValue)}).AddComment(comment))
		}
	}

	l.AppendInstruction(scopeName, NewParamInstruction(ParamInstruction{LiteralOrVariable(classRefTac)}).AddComment("(this)"))
	l.AppendInstruction(scopeName, NewCallInstruction(CallInstruction{
		SaveReturnOn:   lib.NewOpEmpty[VariableName](),
		ProcedureName:  ScopeName(className + "_" + type_checker.CONSTRUCTOR_NAME),
		NumberOfParams: argCount,
	}).AddComment("("+completeExpr+")"))
}

func getUntil(suffixes *[]p.ISuffixOpContext, maxIdx int) string {
	buff := ""
	for i := range maxIdx + 1 {
		buff += (*suffixes)[i].GetText()
	}
	return buff
}
