package type_checker

import (
	"fmt"
	"strings"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

// ================
// 	BLOCK
// ================

func _buildUniqueName(sc *ScopeManager, scopeName string) string {
	functionScope, found := sc.SearchScopeByType(SCOPE_TYPES.FUNCTION)
	if !found {
		return fmt.Sprintf("GLOBAL_%s%d", scopeName, sc.GetUniqueID())
	}
	return fmt.Sprintf("%s_%s%d", functionScope.Name, scopeName, sc.GetUniqueID())
}

func (l Listener) EnterBlockStatement(ctx *p.BlockStatementContext) {

	blockScope := NewScope(_buildUniqueName(l.ScopeManager, "BLOCK"), SCOPE_TYPES.BLOCK)
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

func (l Listener) EnterIfStatement(ctx *p.IfStatementContext) {
	ifScope := NewScope(_buildUniqueName(l.ScopeManager, "IF_"), SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(ifScope)
	l.ScopeManager.ReplaceCurrent(ifScope)
}

func (l Listener) ExitIfStatement(ctx *p.IfStatementContext) {
	l.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterIfBody(ctx *p.IfBodyContext) {
	scopeName := l.ScopeManager.CurrentScope.Name
	ifScope := NewScope(scopeName+"_BODY", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(ifScope)
	l.ScopeManager.ReplaceCurrent(ifScope)
}

func (l Listener) ExitIfBody(ctx *p.IfBodyContext) {
	l.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterElseBody(ctx *p.ElseBodyContext) {
	scopeName := l.ScopeManager.CurrentScope.Name
	ifScope := NewScope(scopeName+"_ELSE", SCOPE_TYPES.BLOCK)
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

	l.ScopeManager.CurrentScope.terminated = true

	if _, inLoopScope := l.ScopeManager.SearchScopeByType(SCOPE_TYPES.LOOP); !inLoopScope {
		l.AddError(line, colStart, colEnd, "'continue' statement out loop scope.")
	}
}

func (l Listener) ExitBreakStatement(ctx *p.BreakStatementContext) {
	line := ctx.GetStart().GetLine()
	colStart := ctx.GetStart().GetColumn()
	colEnd := colStart + len(ctx.GetText())

	l.ScopeManager.CurrentScope.terminated = true

	if _, inLoopScope := l.ScopeManager.SearchScopeByType(SCOPE_TYPES.LOOP); !inLoopScope {
		l.AddError(line, colStart, colEnd, "'break' statement out loop scope.")
	}
}

// ====================
// WHILE
// ====================

func (l Listener) EnterWhileStatement(ctx *p.WhileStatementContext) {
	whileScope := NewScope(_buildUniqueName(l.ScopeManager, "LOOP_"), SCOPE_TYPES.LOOP)
	l.ScopeManager.AddToCurrent(whileScope)
	l.ScopeManager.ReplaceCurrent(whileScope)
}

func (l Listener) ExitWhileStatement(ctx *p.WhileStatementContext) {
	l.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterWhileBody(ctx *p.WhileBodyContext) {
	scopeName := l.ScopeManager.CurrentScope.Name
	loopScope := NewScope(scopeName+"_BODY", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(loopScope)
	l.ScopeManager.ReplaceCurrent(loopScope)
}

func (l Listener) ExitWhileBody(ctx *p.WhileBodyContext) {
	l.ScopeManager.ReplaceWithParent()
}

// ====================
// DO-WHILE
// ====================

func (l Listener) EnterDoWhileStatement(ctx *p.DoWhileStatementContext) {
	whileScope := NewScope(_buildUniqueName(l.ScopeManager, "LOOP_"), SCOPE_TYPES.LOOP)
	l.ScopeManager.AddToCurrent(whileScope)
	l.ScopeManager.ReplaceCurrent(whileScope)
}

func (l Listener) ExitDoWhileStatement(ctx *p.DoWhileStatementContext) {
	l.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterDoWhileBody(ctx *p.DoWhileBodyContext) {
	scopeName := l.ScopeManager.CurrentScope.Name
	loopScope := NewScope(scopeName+"_BODY", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(loopScope)
	l.ScopeManager.ReplaceCurrent(loopScope)
}

func (l Listener) ExitDoWhileBody(ctx *p.DoWhileBodyContext) {
	l.ScopeManager.ReplaceWithParent()
}

// ====================
// FOR
// ====================

func (l Listener) EnterForStatement(ctx *p.ForStatementContext) {
	forScope := NewScope(_buildUniqueName(l.ScopeManager, "LOOP_"), SCOPE_TYPES.LOOP)
	l.ScopeManager.AddToCurrent(forScope)
	l.ScopeManager.ReplaceCurrent(forScope)
}

func (l Listener) ExitForStatement(ctx *p.ForStatementContext) {
	l.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterForBody(ctx *p.ForBodyContext) {
	scopeName := l.ScopeManager.CurrentScope.Name
	loopScope := NewScope(scopeName+"_BODY", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(loopScope)
	l.ScopeManager.ReplaceCurrent(loopScope)
}

func (l Listener) ExitForBody(ctx *p.ForBodyContext) {
	l.ScopeManager.ReplaceWithParent()
}

// ====================
// FOR-EACH
// ====================

func (l Listener) EnterForeachStatement(ctx *p.ForeachStatementContext) {
	forScope := NewScope(_buildUniqueName(l.ScopeManager, "FOR-EACH"), SCOPE_TYPES.LOOP)
	l.ScopeManager.AddToCurrent(forScope)
	l.ScopeManager.ReplaceCurrent(forScope)
}

func (l Listener) ExitForeachStatement(ctx *p.ForeachStatementContext) {
	l.ScopeManager.ReplaceWithParent()
}

// FIXME: HOW ARRAYS ARE HANDLED MUST BE CHANGE.
func (l Listener) ExitForeachValue(ctx *p.ForeachValueContext) {
	line := ctx.GetStart().GetLine()
	colStartI := ctx.GetStart().GetColumn()
	colEndI := ctx.GetStop().GetColumn()
	arrayType, available := l.ScopeManager.CurrentScope.GetExpressionType(ctx.ConditionalExpr().GetText())

	if !available {
		l.AddError(line, colStartI, colEndI, "Exit ForeachValue: Array type not defined!")
		return
	}

	name := string(arrayType)

	if !strings.HasSuffix(name, "[]") {
		l.AddError(line, colStartI, colEndI, "Exit ForeachValue: Value is not type array")
		return
	}

	trimmed := strings.TrimSuffix(name, "[]")

	l.ScopeManager.CurrentScope.UpsertExpressionType(
		ctx.Identifier().GetText(),
		TypeIdentifier(trimmed))
}

// ===========================
// SWITCH
// ===========================

func (l Listener) EnterSwitchStatement(ctx *p.SwitchStatementContext) {
	forScope := NewScope(_buildUniqueName(l.ScopeManager, "SWITCH_"), SCOPE_TYPES.BLOCK)
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

func (l Listener) EnterSwitchCase(ctx *p.SwitchCaseContext) {
	scopeName := l.ScopeManager.CurrentScope.Name
	caseScope := NewScope(
		fmt.Sprintf("%s_CASE_%d", scopeName, l.ScopeManager.GetUniqueID()),
		SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(caseScope)
	l.ScopeManager.ReplaceCurrent(caseScope)
}

func (l Listener) ExitSwitchCase(ctx *p.SwitchCaseContext) {
	l.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterDefaultCase(ctx *p.DefaultCaseContext) {
	scopeName := l.ScopeManager.CurrentScope.Name
	caseScope := NewScope(scopeName+"_DEFAULT", SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(caseScope)
	l.ScopeManager.ReplaceCurrent(caseScope)
}

func (l Listener) ExitDefaultCase(ctx *p.DefaultCaseContext) {
	l.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterCaseBody(ctx *p.CaseBodyContext) {
	// forScope := NewScope(_buildUniqueName(l.ScopeManager, "CASE"), SCOPE_TYPES.BLOCK)
	// l.ScopeManager.AddToCurrent(forScope)
	// l.ScopeManager.ReplaceCurrent(forScope)
}

func (l Listener) ExitCaseBody(ctx *p.CaseBodyContext) {
	// Exit of block scope
	// l.ScopeManager.ReplaceWithParent()
}

// ===========================
// TRY CATCH
// ===========================

func (l Listener) EnterTryStatement(ctx *p.TryStatementContext) {
	tryScope := NewScope(_buildUniqueName(l.ScopeManager, "TRY"), SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(tryScope)
	l.ScopeManager.ReplaceCurrent(tryScope)
}

func (l Listener) EnterCatchStatement(ctx *p.CatchStatementContext) {
	catchScope := NewScope(_buildUniqueName(l.ScopeManager, "CATCH"), SCOPE_TYPES.BLOCK)
	l.ScopeManager.AddToCurrent(catchScope)
	l.ScopeManager.ReplaceCurrent(catchScope)
	l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.Identifier().GetText(), BASE_TYPES.STRING)
}

func (l Listener) ExitTryStatement(ctx *p.TryStatementContext) {
	l.ScopeManager.ReplaceWithParent()
}

func (l Listener) ExitCatchStatement(ctx *p.CatchStatementContext) {
	l.ScopeManager.ReplaceWithParent()
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
		l.AddError(line, colStartI, colEndI, "Exit ConditionalExpr: ternary branches should be of the same type")
		l.ScopeManager.CurrentScope.UpsertExpressionType(ctx.GetText(), BASE_TYPES.INVALID)
		return
	} else if commonType == BASE_TYPES.UNKNOWN ||
		commonType == BASE_TYPES.INVALID {
		l.AddError(line, colStartI, colEndI, "Exit ConditionalExpr: Une of the branches has invalid/unkown values")
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

func (l Listener) ExitStatement(ctx *p.StatementContext) {
	scope := l.ScopeManager.CurrentScope

	if l.isTerminator(ctx) {
		scope.terminated = true
		return
	}

	if scope.terminated {
		line := ctx.GetStart().GetLine()
		startCol := ctx.GetStart().GetColumn()
		endCol := ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText())

		//FIXME: For death code may be a warning not an error
		l.AddError(line, startCol, endCol, "Code unusued after: break, continue or return")
	}

}

func (l Listener) isTerminator(ctx *p.StatementContext) bool {
	switch {
	case ctx.ReturnStatement() != nil:
		return true
	case ctx.BreakStatement() != nil:
		return true
	case ctx.ContinueStatement() != nil:
		return true
	}
	return false
}
