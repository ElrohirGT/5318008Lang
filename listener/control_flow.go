package listener

import (
	p "github.com/ElrohirGT/5318008Lang/parser"
)

// ===========================
// CONDITIONALS
// ===========================

func (l Listener) EnterIfStatement(ctx *p.IfStatementContext) {
	ifScope := NewScope("IF", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(ifScope)
	l.ScopeManager.ReplaceCurrent(ifScope)
}

func (l Listener) ExitIfStatement(ctx *p.IfStatementContext) {
	l.ScopeManager.ReplaceWithParent()
}

// TODO: Switch case statement

// ===========================
// LOOPS
// ===========================

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

func (l Listener) EnterWhileStatement(ctx *p.WhileStatementContext) {
	whileScope := NewScope("WHILE", SCOPE_TYPES.LOOP)
	l.ScopeManager.AddToCurrent(whileScope)
	l.ScopeManager.ReplaceCurrent(whileScope)

}

func (l Listener) ExitWhileStatement(ctx *p.WhileStatementContext) {
	l.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterDoWhileStatement(ctx *p.DoWhileStatementContext) {
	doWhileScope := NewScope("DO-WHILE", SCOPE_TYPES.LOOP)
	l.ScopeManager.AddToCurrent(doWhileScope)
	l.ScopeManager.ReplaceCurrent(doWhileScope)
}

func (l Listener) ExitDoWhileStatement(ctx *p.DoWhileStatementContext) {
	l.ScopeManager.ReplaceWithParent()
}

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
// BOOLEAN EXPRESIONS
// ===========================

func (l Listener) ExitLogicalOrExpr(ctx *p.LogicalOrExprContext) {
	line := ctx.GetStart().GetLine()
	exprs := ctx.AllLogicalAndExpr()
	colStartI := ctx.GetStart().GetColumn()
	colEndI := ctx.GetStop().GetColumn()
	exprs[0].GetStart()

	if len(exprs) == 1 {
		return
	}

	errors := expresionOfTheSameType([]TypeIdentifier{BASE_TYPES.BOOLEAN}, l.ScopeManager, exprs[0], exprs[1])

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

	errors := expresionOfTheSameType([]TypeIdentifier{BASE_TYPES.BOOLEAN}, l.ScopeManager, exprs[0], exprs[1])

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

	errors := expresionOfTheSameType([]TypeIdentifier{
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

	errors := expresionOfTheSameType([]TypeIdentifier{BASE_TYPES.INTEGER, BASE_TYPES.BOOLEAN}, l.ScopeManager, exprs[0], exprs[1])

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
