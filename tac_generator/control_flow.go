package tac_generator

import (
	"log"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

// ================
// 	BLOCK
// ================

func (l Listener) EnterBlockStatement(ctx *p.BlockStatementContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
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
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitSwitchStatement(ctx *p.SwitchStatementContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterCaseBody(ctx *p.CaseBodyContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
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
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) EnterCatchStatement(ctx *p.CatchStatementContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
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
