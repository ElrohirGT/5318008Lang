package tac_generator

import (
	"fmt"
	"log"
	"strings"

	"github.com/ElrohirGT/5318008Lang/lib"
	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/ElrohirGT/5318008Lang/type_checker"
)

func (l Listener) ExitAdditiveExpr(ctx *p.AdditiveExprContext) {

	// runtime.Breakpoint()
	scopeName := getTACScope(l.GetCurrentScope())
	scope := l.GetCurrentScope()

	if len(ctx.AllMultiplicativeExpr()) == 1 {
		childExpr := ctx.MultiplicativeExpr(0).GetText()
		if result, found := l.Program.GetVariableFor(childExpr, scopeName); found {
			l.Program.UpsertTranslation(scopeName, ctx.GetText(), result)
		}
		return
	}

	fullText := ctx.GetText()
	leftExpr := ctx.MultiplicativeExpr(0)
	leftText := leftExpr.GetText()
	leftVar := l.getOrCreateExpressionVariable(leftText, scopeName)

	// Get the type of the left operand
	leftType, foundLeftType := scope.GetExpressionType(leftText)
	if !foundLeftType {
		log.Panicf("Failed to find type for expression: %s", leftText)
	}

	currentResult := leftVar
	currentResultType := leftType
	currentPos := len(leftText)

	for i := 1; i < len(ctx.AllMultiplicativeExpr()); i++ {
		rightExpr := ctx.MultiplicativeExpr(i)
		rightText := rightExpr.GetText()
		rightVar := l.getOrCreateExpressionVariable(rightText, scopeName)

		// Get the type of the right operand
		rightType, foundRightType := scope.GetExpressionType(rightText)
		if !foundRightType {
			log.Panicf("Failed to find type for expression: %s", rightText)
		}

		if currentPos < len(fullText) && fullText[currentPos] == '+' {
			resultVar := l.Program.GetNextVariable()

			if currentResultType == type_checker.BASE_TYPES.STRING &&
				rightType == type_checker.BASE_TYPES.STRING {

				s1, found := l.Program.GetArraySize(currentResult, scopeName)
				if !found {
					log.Panicf("Could not find array size for %s \n", currentResult)
				}
				s2, found := l.Program.GetArraySize(currentResult, scopeName)
				if !found {
					log.Panicf("Could not find array size for %s \n", rightVar)
				}

				newStrSize := s1 + s2 - 1
				l.AppendInstruction(scopeName, NewAllocInstruction(AllocInstruction{
					Target: resultVar,
					Size:   lib.AlignSize(newStrSize, lib.MIPS32_WORD_BYTE_SIZE), // minus one null terminator
				}))

				l.AppendInstruction(scopeName, NewParamInstruction(ParamInstruction{
					Parameter: LiteralOrVariable(currentResult),
				}))
				l.AppendInstruction(scopeName, NewParamInstruction(ParamInstruction{
					Parameter: LiteralOrVariable(rightVar),
				}))
				l.AppendInstruction(scopeName, NewParamInstruction(ParamInstruction{
					Parameter: LiteralOrVariable(resultVar),
				}))
				l.AppendInstruction(scopeName, NewCallInstruction(CallInstruction{
					SaveReturnOn:   lib.NewOpValue(resultVar),
					ProcedureName:  "concat_str",
					NumberOfParams: 3,
				}))

				// CONCAT instruction for strings
				// l.AppendInstruction(scopeName, NewConcatInstruction(ConcatInstruction{
				// 	Target:  resultVar,
				// 	String1: LiteralOrVariable(currentResult),
				// 	String2: LiteralOrVariable(rightVar),
				// }))

				currentResultType = type_checker.BASE_TYPES.STRING
			} else if currentResultType == type_checker.BASE_TYPES.INTEGER &&
				rightType == type_checker.BASE_TYPES.INTEGER {
				// ADD instruction for integers
				l.AppendInstruction(scopeName, NewArithmethicInstruction(ArithmethicInstruction{
					Signed: true,
					Type:   ARITHMETHIC_OPERATION_TYPES.Add,
					Target: resultVar,
					P1:     LiteralOrVariable(currentResult),
					P2:     LiteralOrVariable(rightVar),
				}))
				currentResultType = type_checker.BASE_TYPES.INTEGER
			} else {
				l.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					ctx.GetStop().GetColumn(),
					fmt.Sprintf("Cannot add %s and %s", currentResultType, rightType),
				)
				return
			}

			currentResult = resultVar
		} else {
			// Subtraction - only for numbers
			if currentResultType != type_checker.BASE_TYPES.INTEGER ||
				rightType != type_checker.BASE_TYPES.INTEGER {
				l.AddError(
					ctx.GetStart().GetLine(),
					ctx.GetStart().GetColumn(),
					ctx.GetStop().GetColumn(),
					"Subtraction only works with integers",
				)
				return
			}

			resultVar := l.Program.GetNextVariable()
			l.AppendInstruction(scopeName, NewArithmethicInstruction(ArithmethicInstruction{
				Signed: true,
				Type:   ARITHMETHIC_OPERATION_TYPES.Subtract,
				Target: resultVar,
				P1:     LiteralOrVariable(currentResult),
				P2:     LiteralOrVariable(rightVar),
			}))
			currentResult = resultVar
		}

		currentPos += 1 + len(rightText)
	}

	l.Program.UpsertTranslation(scopeName, ctx.GetText(), currentResult)
}

func (l Listener) ExitMultiplicativeExpr(ctx *p.MultiplicativeExprContext) {
	scopeName := getTACScope(l.GetCurrentScope())

	if len(ctx.AllUnaryExpr()) == 1 {
		childExpr := ctx.UnaryExpr(0).GetText()
		if result, found := l.Program.GetVariableFor(childExpr, scopeName); found {
			l.Program.UpsertTranslation(scopeName, ctx.GetText(), result)
		}
		return
	}

	fullText := ctx.GetText()
	leftExpr := ctx.UnaryExpr(0)
	leftText := leftExpr.GetText()
	leftVar := l.getOrCreateExpressionVariable(leftText, scopeName)

	currentResult := leftVar
	currentPos := len(leftText)

	for i := 1; i < len(ctx.AllUnaryExpr()); i++ {
		rightExpr := ctx.UnaryExpr(i)
		rightText := rightExpr.GetText()
		rightVar := l.getOrCreateExpressionVariable(rightText, scopeName)

		var opType ArithmethicOpType
		if currentPos < len(fullText) {
			op := fullText[currentPos]
			switch op {
			case '*':
				opType = ARITHMETHIC_OPERATION_TYPES.Multiplication
			case '/':
				opType = ARITHMETHIC_OPERATION_TYPES.Divide
			default:
				log.Panicf("Unknown operator '%c' in expression: %s", op, fullText)
			}
		}

		resultVar := l.Program.GetNextVariable()

		l.AppendInstruction(scopeName, NewArithmethicInstruction(ArithmethicInstruction{
			Signed: true,
			Type:   opType,
			Target: resultVar,
			P1:     LiteralOrVariable(currentResult),
			P2:     LiteralOrVariable(rightVar),
		}))

		currentResult = resultVar
		currentPos += 1 + len(rightText)
	}

	l.Program.UpsertTranslation(scopeName, ctx.GetText(), currentResult)
}

func (l Listener) ExitUnaryExpr(ctx *p.UnaryExprContext) {
	scopeName := getTACScope(l.GetCurrentScope())

	if ctx.PrimaryExpr() != nil {
		childText := ctx.PrimaryExpr().GetText()
		if result, found := l.Program.GetVariableFor(childText, scopeName); found {
			l.Program.UpsertTranslation(scopeName, ctx.GetText(), result)
		}
		return
	}

	childExpr := ctx.UnaryExpr()
	childText := childExpr.GetText()
	childVar := l.getOrCreateExpressionVariable(childText, scopeName)

	resultVar := l.Program.GetNextVariable()

	exprText := ctx.GetText()
	if strings.HasPrefix(exprText, "-") {
		l.AppendInstruction(scopeName, NewArithmethicInstruction(ArithmethicInstruction{
			Signed: true,
			Type:   ARITHMETHIC_OPERATION_TYPES.Subtract,
			Target: resultVar,
			P1:     "0",
			P2:     LiteralOrVariable(childVar),
		}))
	}

	l.Program.UpsertTranslation(scopeName, ctx.GetText(), resultVar)
}

func (l Listener) ExitPrimaryExpr(ctx *p.PrimaryExprContext) {
	scopeName := getTACScope(l.GetCurrentScope())
	exprText := ctx.GetText()

	if _, found := l.Program.GetVariableFor(exprText, scopeName); found {
		return
	}

	if ctx.LeftHandSide() != nil {
		lhs := ctx.LeftHandSide()

		if lhs.PrimaryAtom() != nil {
			atom := lhs.PrimaryAtom()

			switch atomCtx := atom.(type) {
			case *p.IdentifierExprContext:
				varName := atomCtx.Identifier().GetText()

				tacVar := l.Program.GetOrGenerateVariable(varName, scopeName)
				l.Program.UpsertTranslation(scopeName, exprText, tacVar)

				// TODO: cases for new and this
			}
		}
	} else if ctx.ConditionalExpr() != nil {
		childText := ctx.ConditionalExpr().GetText()
		if result, found := l.Program.GetVariableFor(childText, scopeName); found {
			l.Program.UpsertTranslation(scopeName, exprText, result)
		}
	}
}
