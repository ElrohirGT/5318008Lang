package tac_generator

import (
	"log"
	"strings"

	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/ElrohirGT/5318008Lang/type_checker"
)

// Helper function to get the correct scope name for TAC generation
func getTACScope(scope *type_checker.Scope) ScopeName {
	scopeName := ScopeName(scope.Name)

	// If we're in a class scope, use the constructor scope for TAC generation
	if scope.Type == type_checker.SCOPE_TYPES.CLASS {
		scopeName = ScopeName(scope.Name + "_" + type_checker.CONSTRUCTOR_NAME)
	}

	return scopeName
}

func (l Listener) ExitAdditiveExpr(ctx *p.AdditiveExprContext) {
	scopeName := getTACScope(l.GetCurrentScope())

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

	currentResult := leftVar
	currentPos := len(leftText)

	for i := 1; i < len(ctx.AllMultiplicativeExpr()); i++ {
		rightExpr := ctx.MultiplicativeExpr(i)
		rightText := rightExpr.GetText()
		rightVar := l.getOrCreateExpressionVariable(rightText, scopeName)

		var opType ArithmethicOpType
		if currentPos < len(fullText) && fullText[currentPos] == '+' {
			opType = ARITHMETHIC_OPERATION_TYPES.Add
		} else {
			opType = ARITHMETHIC_OPERATION_TYPES.Subtract
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

func (l Listener) getOrCreateExpressionVariable(exprText string, scopeName ScopeName) VariableName {
	// Check if already computed
	if result, found := l.Program.GetVariableFor(exprText, scopeName); found {
		return result
	}

	// Check for literal
	if _, isLiteral := l.TypeChecker.GetLiteralType(exprText); isLiteral {
		return VariableName(exprText)
	}

	return l.Program.GetOrGenerateVariable(exprText, scopeName)
}
