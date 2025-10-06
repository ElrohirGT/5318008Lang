package tac_generator

import (
	"fmt"
	"log"

	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/ElrohirGT/5318008Lang/type_checker"
	"github.com/antlr4-go/antlr/v4"
)

// ========================
// CONDITIONAL EXPRESIONS
// =======================

func (l Listener) ExitRelationalExpr(ctx *p.RelationalExprContext) {
	scopeName := ScopeName(l.GetCurrentScope().Name)

	if len(ctx.AllAdditiveExpr()) == 1 {
		return
	}

	opToken := ctx.GetChild(1).(antlr.TerminalNode).GetSymbol()
	operator := BOOLEAN_OPERATION_TYPES.Greater
	destiny := l.getOrCreateExpressionVariable(ctx.GetText(), scopeName)

	p1, _ := l.Program.GetVariableOrLiteral(
		ctx.AdditiveExpr(0).GetText(), scopeName, l.TypeChecker.ScopeManager, l.TypeChecker)
	p2, _ := l.Program.GetVariableOrLiteral(
		ctx.AdditiveExpr(1).GetText(), scopeName, l.TypeChecker.ScopeManager, l.TypeChecker)

	switch opToken.GetText() {
	case "<":
		operator = BOOLEAN_OPERATION_TYPES.Less
	case "<=":
		operator = BOOLEAN_OPERATION_TYPES.LessOrEqual
	case ">":
		operator = BOOLEAN_OPERATION_TYPES.Greater
	case ">=":
		operator = BOOLEAN_OPERATION_TYPES.GreaterOrEqual
	}

	l.AppendInstruction(scopeName, NewLogicOpInstruction(LogicOpInstruction{
		Type:   operator,
		Target: destiny,
		P1:     p1,
		P2:     p2,
	}))
}

func (l Listener) ExitEqualityExpr(ctx *p.EqualityExprContext) {
	scopeName := ScopeName(l.GetCurrentScope().Name)

	if len(ctx.AllRelationalExpr()) == 1 {
		return
	}

	opToken := ctx.GetChild(1).(antlr.TerminalNode).GetSymbol()
	operator := BOOLEAN_OPERATION_TYPES.Equal
	destiny := l.getOrCreateExpressionVariable(ctx.GetText(), scopeName)

	p1, exprType := l.Program.GetVariableOrLiteral(
		ctx.RelationalExpr(0).GetText(), scopeName, l.TypeChecker.ScopeManager, l.TypeChecker)
	p2, _ := l.Program.GetVariableOrLiteral(
		ctx.RelationalExpr(1).GetText(), scopeName, l.TypeChecker.ScopeManager, l.TypeChecker)

	if exprType == type_checker.BASE_TYPES.STRING {
		panic("Equality between strings not supported yet :p")
	}

	switch opToken.GetText() {
	case "==":
		operator = BOOLEAN_OPERATION_TYPES.Equal
	case "!=":
		operator = BOOLEAN_OPERATION_TYPES.NotEqual
	}

	l.AppendInstruction(scopeName, NewLogicOpInstruction(LogicOpInstruction{
		Type:   operator,
		Target: destiny,
		P1:     p1,
		P2:     p2,
	}))
}

func (l Listener) ExitLogicalAndExpr(ctx *p.LogicalAndExprContext) {

	if len(ctx.AllEqualityExpr()) == 1 {
		return
	}

	// Process first instruction
	firstExpr := ctx.EqualityExpr(0).GetText() + "&&" + ctx.EqualityExpr(1).GetText()
	fmt.Println(firstExpr)

	scopeName := ScopeName(l.GetCurrentScope().Name)
	destiny := l.getOrCreateExpressionVariable(firstExpr, scopeName)
	p1, _ := l.Program.GetVariableOrLiteral(
		ctx.EqualityExpr(0).GetText(), scopeName, l.TypeChecker.ScopeManager, l.TypeChecker)
	p2, _ := l.Program.GetVariableOrLiteral(
		ctx.EqualityExpr(1).GetText(), scopeName, l.TypeChecker.ScopeManager, l.TypeChecker)

	l.AppendInstruction(scopeName, NewLogicOpInstruction(LogicOpInstruction{
		Type:   BOOLEAN_OPERATION_TYPES.And,
		Target: destiny,
		P1:     p1,
		P2:     p2,
	}))

	secondExpr := ""
	p1 = destiny
	// Concat the rest of the expresions
	for i := 2; i < len(ctx.AllEqualityExpr()); i++ {
		secondExpr = firstExpr + "&&" + ctx.EqualityExpr(i).GetText()
		destiny = l.getOrCreateExpressionVariable(secondExpr, scopeName)
		p2, _ = l.Program.GetVariableOrLiteral(
			ctx.EqualityExpr(i).GetText(), scopeName, l.TypeChecker.ScopeManager, l.TypeChecker)

		l.AppendInstruction(scopeName, NewLogicOpInstruction(LogicOpInstruction{
			Type:   BOOLEAN_OPERATION_TYPES.And,
			Target: destiny,
			P1:     p1,
			P2:     p2,
		}))

		p1 = p2
		firstExpr = secondExpr
	}
}

func (l Listener) ExitLogicalOrExpr(ctx *p.LogicalOrExprContext) {
	if len(ctx.AllLogicalAndExpr()) == 1 {
		return
	}

	// Process first instruction
	firstExpr := ctx.LogicalAndExpr(0).GetText() + "||" + ctx.LogicalAndExpr(1).GetText()
	fmt.Println(firstExpr)

	scopeName := ScopeName(l.GetCurrentScope().Name)
	destiny := l.getOrCreateExpressionVariable(firstExpr, scopeName)
	p1, _ := l.Program.GetVariableOrLiteral(
		ctx.LogicalAndExpr(0).GetText(), scopeName, l.TypeChecker.ScopeManager, l.TypeChecker)
	p2, _ := l.Program.GetVariableOrLiteral(
		ctx.LogicalAndExpr(1).GetText(), scopeName, l.TypeChecker.ScopeManager, l.TypeChecker)

	l.AppendInstruction(scopeName, NewLogicOpInstruction(LogicOpInstruction{
		Type:   BOOLEAN_OPERATION_TYPES.Or,
		Target: destiny,
		P1:     p1,
		P2:     p2,
	}))

	secondExpr := ""
	p1 = destiny
	// Concat the rest of the expresions
	for i := 2; i < len(ctx.AllLogicalAndExpr()); i++ {
		secondExpr = firstExpr + "||" + ctx.LogicalAndExpr(i).GetText()
		destiny = l.getOrCreateExpressionVariable(secondExpr, scopeName)
		p2, _ = l.Program.GetVariableOrLiteral(
			ctx.LogicalAndExpr(i).GetText(), scopeName, l.TypeChecker.ScopeManager, l.TypeChecker)

		l.AppendInstruction(scopeName, NewLogicOpInstruction(LogicOpInstruction{
			Type:   BOOLEAN_OPERATION_TYPES.Or,
			Target: destiny,
			P1:     p1,
			P2:     p2,
		}))

		p1 = p2
		firstExpr = secondExpr
	}
}

// ================
// 	BLOCK
// ================

func (l Listener) EnterBlockStatement(ctx *p.BlockStatementContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()

	scope := l.GetCurrentScope()
	l.Program.UpsertScope(ScopeName(scope.Name))

	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitBlockStatement(ctx *p.BlockStatementContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

// ================
// 	IF - ELSE
// ================

func (l Listener) EnterIfBody(ctx *p.IfBodyContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	l.Program.UpsertScope(ScopeName(scope.Name))
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitIfBody(ctx *p.IfBodyContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

// ================
// 	WHILE
// ================

func (l Listener) EnterWhileBody(ctx *p.WhileBodyContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	l.Program.UpsertScope(ScopeName(scope.Name))
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitWhileBody(ctx *p.WhileBodyContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

// ================
// 	DO-WHILE
// ================

func (l Listener) EnterDoWhileBody(ctx *p.DoWhileBodyContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	l.Program.UpsertScope(ScopeName(scope.Name))
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitDoWhileBody(ctx *p.DoWhileBodyContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

// ================
// 	FOR
// ================

func (l Listener) EnterForStatement(ctx *p.ForStatementContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	l.Program.UpsertScope(ScopeName(scope.Name))
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitForStatement(ctx *p.ForStatementContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

// ================
// 	FOR-EACH
// ================

func (l Listener) EnterForeachStatement(ctx *p.ForeachStatementContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	l.Program.UpsertScope(ScopeName(scope.Name))
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitForeachStatement(ctx *p.ForeachStatementContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

// ================
// 	SWITCH
// ================

func (l Listener) EnterSwitchStatement(ctx *p.SwitchStatementContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	l.Program.UpsertScope(ScopeName(scope.Name))
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitSwitchStatement(ctx *p.SwitchStatementContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterCaseBody(ctx *p.CaseBodyContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	l.Program.UpsertScope(ScopeName(scope.Name))
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitCaseBody(ctx *p.CaseBodyContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

// ================
// 	TRY CATH
// ================

func (l Listener) EnterTryStatement(ctx *p.TryStatementContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	l.Program.UpsertScope(ScopeName(scope.Name))
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) EnterCatchStatement(ctx *p.CatchStatementContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	l.Program.UpsertScope(ScopeName(scope.Name))
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitTryStatement(ctx *p.TryStatementContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

func (l Listener) ExitCatchStatement(ctx *p.CatchStatementContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}
