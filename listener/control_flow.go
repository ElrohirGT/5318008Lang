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
	if _, inLoopScope := l.ScopeManager.SearchScopeByType(SCOPE_TYPES.LOOP); !inLoopScope {
		l.AddError(line, "'continue' statement out loop scope.")
	}
}

func (l Listener) ExitBreakStatement(ctx *p.BreakStatementContext) {
	line := ctx.GetStart().GetLine()
	if _, inLoopScope := l.ScopeManager.SearchScopeByType(SCOPE_TYPES.LOOP); !inLoopScope {
		l.AddError(line, "'break' statement out loop scope.")
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
