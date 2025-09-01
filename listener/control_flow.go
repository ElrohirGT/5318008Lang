package listener

import (
	"fmt"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

// ================
// 	BLOCK
// ================

func (l Listener) EnterBlockStatement(ctx *p.BlockStatementContext) {
	blockScope := NewScope("BLOCK", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(blockScope)
	l.ScopeManager.ReplaceCurrent(blockScope)
}

func (l Listener) ExitBlockStatement(ctx *p.BlockStatementContext) {
	l.ScopeManager.ReplaceWithParent()
}

// ================
// 	IF - ELSE
// ================

func (l Listener) ExitMustBoolExpr(ctx *p.MustBoolExprContext) {
	line := ctx.GetStart().GetLine()
	colStart := ctx.GetStart().GetColumn()
	colEnd := colStart + len(ctx.GetText())
	conditional := ctx.ConditionalExpr()

	generalType, available := l.ScopeManager.CurrentScope.GetExpressionType(conditional.GetText())
	// l.ScopeManager.CurrentScope.Print(3)

	if !available {
		msg := fmt.Sprintf("`%s` symbol is not registered in scope!", conditional.GetText())

		l.AddError(line, colStart, colEnd, msg)
		return
	}

	if generalType != BASE_TYPES.BOOLEAN {
		l.AddError(line, colStart, colEnd,
			fmt.Sprintf("`if` condition should be a `boolean` but `%s` was given.", generalType))
		return

	}
}

func (l Listener) EnterIfBody(ctx *p.IfBodyContext) {
	ifScope := NewScope("IF", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(ifScope)
	l.ScopeManager.ReplaceCurrent(ifScope)
}

func (l Listener) ExitIfBody(ctx *p.IfBodyContext) {
	l.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterElseBody(ctx *p.ElseBodyContext) {
	ifScope := NewScope("ELSE", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(ifScope)
	l.ScopeManager.ReplaceCurrent(ifScope)
}

func (l Listener) ExitElseBody(ctx *p.ElseBodyContext) {
	l.ScopeManager.ReplaceWithParent()
}

// ====================
// CONTROL KEYWORDS
// ====================

func (l Listener) ExitContinueStatement(ctx *p.ContinueStatementContext) {
	line := ctx.GetStart().GetLine()
	colStart := ctx.GetStart().GetColumn()
	colEnd := colStart + len(ctx.GetText())

	if _, inLoopScope := l.ScopeManager.SearchScopeByType(SCOPE_TYPES.LOOP); !inLoopScope {
		l.AddError(line, colStart, colEnd, "'continue' statement out loop scope.")
	}
}

func (l Listener) ExitBreakStatement(ctx *p.BreakStatementContext) {
	line := ctx.GetStart().GetLine()
	colStart := ctx.GetStart().GetColumn()
	colEnd := colStart + len(ctx.GetText())

	if _, inLoopScope := l.ScopeManager.SearchScopeByType(SCOPE_TYPES.LOOP); !inLoopScope {
		l.AddError(line, colStart, colEnd, "'break' statement out loop scope.")
	}
}

// ====================
// WHILE
// ====================

func (l Listener) EnterWhileBody(ctx *p.WhileBodyContext) {
	whileScope := NewScope("WHILE", SCOPE_TYPES.LOOP)
	l.ScopeManager.AddToCurrent(whileScope)
	l.ScopeManager.ReplaceCurrent(whileScope)
}

func (l Listener) ExitWhileBody(ctx *p.WhileBodyContext) {
	l.ScopeManager.ReplaceWithParent()
}

// ====================
// DO-WHILE
// ====================

func (l Listener) EnterDoWhileBody(ctx *p.DoWhileBodyContext) {
	doWhileScope := NewScope("DO-WHILE", SCOPE_TYPES.LOOP)
	l.ScopeManager.AddToCurrent(doWhileScope)
	l.ScopeManager.ReplaceCurrent(doWhileScope)
}

func (l Listener) ExitDoWhileBody(ctx *p.DoWhileBodyContext) {
	l.ScopeManager.ReplaceWithParent()
}

// ====================
// FOR
// ====================

func (l Listener) EnterForStatement(ctx *p.ForStatementContext) {
	forScope := NewScope("FOR", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(forScope)
	l.ScopeManager.ReplaceCurrent(forScope)
	// TODO: ADD EXTRA CONTEXT AGAIN
	blockScope := NewScope("FOR-BLOCK", SCOPE_TYPES.LOOP)
	l.ScopeManager.AddToCurrent(blockScope)
	l.ScopeManager.ReplaceCurrent(blockScope)
}

func (l Listener) ExitForStatement(ctx *p.ForStatementContext) {
	// Exit of block scope
	l.ScopeManager.ReplaceWithParent()
	// Exit of for scope
	l.ScopeManager.ReplaceWithParent()

}

// ===========================
// SWITCH
// ===========================

func (l Listener) EnterSwitchStatement(ctx *p.SwitchStatementContext) {
	forScope := NewScope("SWITCH", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(forScope)
	l.ScopeManager.ReplaceCurrent(forScope)
}

func (l Listener) ExitSwitchStatement(ctx *p.SwitchStatementContext) {
	l.ScopeManager.ReplaceWithParent()
}

func (l Listener) ExitSwitchValue(ctx *p.SwitchValueContext) {
	switchValue, _ := l.ScopeManager.CurrentScope.GetExpressionType(ctx.ConditionalExpr().GetText())
	l.ScopeManager.CurrentScope.UpsertExpressionType("$switch", switchValue)
}

func (l Listener) ExitCaseValue(ctx *p.CaseValueContext) {
	line := ctx.GetStart().GetLine()
	colStartI := ctx.GetStart().GetColumn()
	colEndI := ctx.GetStop().GetColumn()
	switchValue, _ := l.ScopeManager.CurrentScope.GetExpressionType("$switch")
	caseValue, _ := l.ScopeManager.CurrentScope.GetExpressionType(ctx.PrimaryExpr().GetText())

	if switchValue != caseValue {
		l.AddError(line, colStartI, colEndI,
			fmt.Sprintf("Exit CaseValue: expected type `%s` but `%s` found", switchValue, caseValue))
	}
}

func (l Listener) EnterCaseBody(ctx *p.CaseBodyContext) {
	forScope := NewScope("CASE", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(forScope)
	l.ScopeManager.ReplaceCurrent(forScope)
}

func (l Listener) ExitCaseBody(ctx *p.CaseBodyContext) {
	// Exit of block scope
	l.ScopeManager.ReplaceWithParent()
}

// ===========================
// TRY CATCH
// ===========================

func (l Listener) EnterTryStatement(ctx *p.TryStatementContext) {
	tryScope := NewScope("TRY", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(tryScope)
	l.ScopeManager.ReplaceCurrent(tryScope)
}

func (l Listener) EnterCatchStatement(ctx *p.CatchStatementContext) {
	catchScope := NewScope("CATCH", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(catchScope)
	l.ScopeManager.ReplaceCurrent(catchScope)
	l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.Identifier().GetText(), BASE_TYPES.STRING)
}

// ===========================
// BOOLEAN EXPRESIONS
// ===========================

func (l Listener) ExitConditionalExpr(ctx *p.ConditionalExprContext) {
	line := ctx.GetStart().GetLine()
	colStartI := ctx.GetStart().GetColumn()
	colEndI := ctx.GetStop().GetColumn()

	exprs := ctx.AllConditionalExpr()
	if len(exprs) == 0 {
		return
	}

	areSame, commonType := expresionsOfTheSameType(l.ScopeManager, exprs[0], exprs[1])

	if !areSame {
		l.AddError(line, colStartI, colEndI, fmt.Sprintf("Exit ConditionalExpr: ternary branches should be of the same type"))
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
		return
	} else if commonType == BASE_TYPES.UNKNOWN ||
		commonType == BASE_TYPES.INVALID {
		l.AddError(line, colStartI, colEndI, fmt.Sprintf("Exit ConditionalExpr: Une of the branches has invalid/unkown values"))
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
		return
	}

	l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), commonType)

}

func (l Listener) ExitLogicalOrExpr(ctx *p.LogicalOrExprContext) {
	line := ctx.GetStart().GetLine()
	exprs := ctx.AllLogicalAndExpr()
	colStartI := ctx.GetStart().GetColumn()
	colEndI := ctx.GetStop().GetColumn()
	exprs[0].GetStart()

	if len(exprs) == 1 {

		return
	}

	errors := expresionsOfTheRequiredType([]TypeIdentifier{BASE_TYPES.BOOLEAN}, l.ScopeManager, exprs[0], exprs[1])

	for _, e := range errors {
		l.AddError(line, colStartI, colEndI, "Exit RelationalExpr: "+e)
	}

	if len(errors) > 0 {
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
	} else {
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.BOOLEAN)
	}
}

func (l Listener) ExitLogicalAndExpr(ctx *p.LogicalAndExprContext) {
	line := ctx.GetStart().GetLine()
	exprs := ctx.AllEqualityExpr()
	colStartI := ctx.GetStart().GetColumn()
	colEndI := ctx.GetStop().GetColumn()
	exprs[0].GetStart()

	if len(exprs) == 1 {
		return
	}

	errors := expresionsOfTheRequiredType([]TypeIdentifier{BASE_TYPES.BOOLEAN}, l.ScopeManager, exprs[0], exprs[1])

	for _, e := range errors {
		l.AddError(line, colStartI, colEndI, "Exit RelationalExpr: "+e)
	}

	if len(errors) > 0 {
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
	} else {
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.BOOLEAN)
	}
}

func (l Listener) ExitEqualityExpr(ctx *p.EqualityExprContext) {
	line := ctx.GetStart().GetLine()
	exprs := ctx.AllRelationalExpr()
	colStartI := ctx.GetStart().GetColumn()
	colEndI := ctx.GetStop().GetColumn()
	exprs[0].GetStart()

	if len(exprs) == 1 {
		return
	}

	errors := expresionsOfTheRequiredType([]TypeIdentifier{
		BASE_TYPES.INTEGER,
		BASE_TYPES.BOOLEAN,
		BASE_TYPES.STRING,
		BASE_TYPES.NULL}, l.ScopeManager, exprs[0], exprs[1])

	for _, e := range errors {
		l.AddError(line, colStartI, colEndI, "Exit RelationalExpr: "+e)
	}

	if len(errors) > 0 {
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
	} else {
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.BOOLEAN)
	}
}

func (l Listener) ExitRelationalExpr(ctx *p.RelationalExprContext) {
	line := ctx.GetStart().GetLine()
	exprs := ctx.AllAdditiveExpr()
	colStartI := ctx.GetStart().GetColumn()
	colEndI := ctx.GetStop().GetColumn()
	exprs[0].GetStart()

	if len(exprs) == 1 {
		return
	}

	errors := expresionsOfTheRequiredType([]TypeIdentifier{BASE_TYPES.INTEGER, BASE_TYPES.BOOLEAN}, l.ScopeManager, exprs[0], exprs[1])

	for _, e := range errors {
		l.AddError(line, colStartI, colEndI, "Exit RelationalExpr: "+e)
	}

	if len(errors) > 0 {
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
	} else {
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.BOOLEAN)
	}
}

func (l Listener) ExitUnaryExpr(ctx *p.UnaryExprContext) {
	unary := ctx.UnaryExpr()

	if unary == nil {
		return
	}

	referenceType, _ := l.ScopeManager.CurrentScope.GetExpressionType(unary.GetText())

	if referenceType != BASE_TYPES.BOOLEAN {
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
	}

	l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.BOOLEAN)
}
