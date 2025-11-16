package tac_generator

import (
	"fmt"
	"log"
	"strings"

	"github.com/ElrohirGT/5318008Lang/lib"
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
		ctx.AdditiveExpr(0).GetText(), scopeName, l.TypeChecker)
	p2, _ := l.Program.GetVariableOrLiteral(
		ctx.AdditiveExpr(1).GetText(), scopeName, l.TypeChecker)

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
		ctx.RelationalExpr(0).GetText(), scopeName, l.TypeChecker)
	p2, _ := l.Program.GetVariableOrLiteral(
		ctx.RelationalExpr(1).GetText(), scopeName, l.TypeChecker)

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
		ctx.EqualityExpr(0).GetText(), scopeName, l.TypeChecker)
	p2, _ := l.Program.GetVariableOrLiteral(
		ctx.EqualityExpr(1).GetText(), scopeName, l.TypeChecker)

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
			ctx.EqualityExpr(i).GetText(), scopeName, l.TypeChecker)

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
		ctx.LogicalAndExpr(0).GetText(), scopeName, l.TypeChecker)
	p2, _ := l.Program.GetVariableOrLiteral(
		ctx.LogicalAndExpr(1).GetText(), scopeName, l.TypeChecker)

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
			ctx.LogicalAndExpr(i).GetText(), scopeName, l.TypeChecker)

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
	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(scope.Name), parentName)

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

func (l Listener) EnterIfStatement(ctx *p.IfStatementContext) {
	scope := l.GetCurrentScope()
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	ifScope := l.GetCurrentScope()
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(ifScope.Name),
	}))
	l.AppendInstruction(ScopeName(scope.Name), NewSecInstruction(SecInstruction{
		Name: TagName(ifScope.Name + "_RETURN"),
	}))

	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(ifScope.Name), parentName)
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitIfStatement(ctx *p.IfStatementContext) {
	scope := l.GetCurrentScope()
	condVar, _ := l.Program.GetVariableOrLiteral(
		ctx.IfCondition().GetText(), ScopeName(scope.Name), l.TypeChecker)
	// if !found {
	// 	panic("Unable to found condition expresion for if statement")
	// }
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.NewOpValue(JumpCondition{
			Simple: lib.NewOpValue(condVar),
		}),
		Target: TagName(scope.Name + "_BODY"),
	}))
	if ctx.ElseBody() != nil {
		l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
			Condition: lib.NewOpValue(JumpCondition{
				SimpleNegated: lib.NewOpValue(condVar),
			}),
			Target: TagName(scope.Name + "_ELSE"),
		}))
	} else {
		l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
			Condition: lib.NewOpEmpty[JumpCondition](),
			Target:    TagName(scope.Name + "_RETURN"),
		}))
	}
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

func (l Listener) ExitIfCondition(ctx *p.IfConditionContext) {
}

func (l Listener) EnterIfBody(ctx *p.IfBodyContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(scope.Name), parentName)
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitIfBody(ctx *p.IfBodyContext) {
	scope := l.GetCurrentScope()
	returnTag := strings.TrimSuffix(scope.Name, "_BODY") + "_RETURN"
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(returnTag),
	}))
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterElseBody(ctx *p.ElseBodyContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(scope.Name), parentName)
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitElseBody(ctx *p.ElseBodyContext) {
	scope := l.GetCurrentScope()
	returnTag := strings.TrimSuffix(scope.Name, "_ELSE") + "_RETURN"
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(returnTag),
	}))
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

// ================
// 	WHILE
// ================

func (l Listener) EnterWhileStatement(ctx *p.WhileStatementContext) {
	scope := l.GetCurrentScope()
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	loopScope := l.GetCurrentScope()
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(loopScope.Name),
	}))
	l.AppendInstruction(ScopeName(scope.Name), NewSecInstruction(SecInstruction{
		Name: TagName(loopScope.Name + "_RETURN"),
	}))

	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(loopScope.Name), parentName)
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}
func (l Listener) ExitWhileStatement(ctx *p.WhileStatementContext) {
	scope := l.GetCurrentScope()
	l.AppendInstruction(ScopeName(scope.Name), NewSecInstruction(SecInstruction{
		Name: TagName(scope.Name + "_UPDATE"),
	}))
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(scope.Name + "_CONDITION"),
	}))
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterWhileCondition(ctx *p.WhileConditionContext) {
	scope := l.GetCurrentScope()
	l.AppendInstruction(ScopeName(scope.Name), NewSecInstruction(SecInstruction{
		Name: TagName(scope.Name + "_CONDITION"),
	}))
}

func (l Listener) ExitWhileCondition(ctx *p.WhileConditionContext) {
	scope := l.GetCurrentScope()
	condVar, _ := l.Program.GetVariableOrLiteral(
		ctx.MustBoolExpr().GetText(), ScopeName(scope.Name), l.TypeChecker)
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.NewOpValue(JumpCondition{
			Simple: lib.NewOpValue(condVar),
		}),
		Target: TagName(scope.Name + "_BODY"),
	}))
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.NewOpValue(JumpCondition{
			SimpleNegated: lib.NewOpValue(condVar),
		}),
		Target: TagName(scope.Name + "_RETURN"),
	}))
}

func (l Listener) EnterWhileBody(ctx *p.WhileBodyContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(scope.Name), parentName)
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitWhileBody(ctx *p.WhileBodyContext) {
	scope := l.GetCurrentScope()
	returnTag := strings.TrimSuffix(scope.Name, "_BODY") + "_UPDATE"
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(returnTag),
	}))
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

// ================
// 	DO-WHILE
// ================

func (l Listener) EnterDoWhileStatement(ctx *p.DoWhileStatementContext) {
	scope := l.GetCurrentScope()
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	loopScope := l.GetCurrentScope()
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(loopScope.Name),
	}))
	l.AppendInstruction(ScopeName(scope.Name), NewSecInstruction(SecInstruction{
		Name: TagName(loopScope.Name + "_RETURN"),
	}))
	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(loopScope.Name), parentName)
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}
func (l Listener) ExitDoWhileStatement(ctx *p.DoWhileStatementContext) {
	scope := l.GetCurrentScope()
	l.AppendInstruction(ScopeName(scope.Name), NewSecInstruction(SecInstruction{
		Name: TagName(scope.Name + "_UPDATE"),
	}))
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(scope.Name + "_CONDITION"),
	}))
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterDoWhileBody(ctx *p.DoWhileBodyContext) {
	scope := l.GetCurrentScope()
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(scope.Name + "_BODY"),
	}))
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope = l.GetCurrentScope()
	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(scope.Name), parentName)
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitDoWhileBody(ctx *p.DoWhileBodyContext) {
	scope := l.GetCurrentScope()
	returnTag := strings.TrimSuffix(scope.Name, "_BODY") + "_UPDATE"
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(returnTag),
	}))
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

// ================
// 	FOR
// ================

func (l Listener) EnterForStatement(ctx *p.ForStatementContext) {
	scope := l.GetCurrentScope()
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	loopScope := l.GetCurrentScope()
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(loopScope.Name),
	}))
	l.AppendInstruction(ScopeName(scope.Name), NewSecInstruction(SecInstruction{
		Name: TagName(loopScope.Name + "_RETURN"),
	}))

	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(loopScope.Name), parentName)
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitForStatement(ctx *p.ForStatementContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterForCondition(ctx *p.ForConditionContext) {
	scope := l.GetCurrentScope()
	l.AppendInstruction(ScopeName(scope.Name), NewSecInstruction(SecInstruction{
		Name: TagName(scope.Name + "_CONDITION"),
	}))
}

func (l Listener) ExitForCondition(ctx *p.ForConditionContext) {
	scope := l.GetCurrentScope()
	condVar, _ := l.Program.GetVariableOrLiteral(
		ctx.MustBoolExpr().GetText(), ScopeName(scope.Name), l.TypeChecker)
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.NewOpValue(JumpCondition{
			Simple: lib.NewOpValue(condVar),
		}),
		Target: TagName(scope.Name + "_BODY"),
	}))
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.NewOpValue(JumpCondition{
			SimpleNegated: lib.NewOpValue(condVar),
		}),
		Target: TagName(scope.Name + "_RETURN"),
	}))
}

func (l Listener) EnterForUpdate(ctx *p.ForUpdateContext) {
	scope := l.GetCurrentScope()
	l.AppendInstruction(ScopeName(scope.Name), NewSecInstruction(SecInstruction{
		Name: TagName(scope.Name + "_UPDATE"),
	}))
}

func (l Listener) ExitForUpdate(ctx *p.ForUpdateContext) {
	scope := l.GetCurrentScope()
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.NewOpEmpty[JumpCondition](),
		Target:    TagName(scope.Name + "_CONDITION"),
	}))
}

func (l Listener) EnterForBody(ctx *p.ForBodyContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(scope.Name), parentName)
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitForBody(ctx *p.ForBodyContext) {
	scope := l.GetCurrentScope()
	returnTag := strings.TrimSuffix(scope.Name, "_BODY") + "_UPDATE"
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(returnTag),
	}))
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

// ================
// 	FOR-EACH
// ================

func (l Listener) EnterForeachStatement(ctx *p.ForeachStatementContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(scope.Name), parentName)
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitForeachStatement(ctx *p.ForeachStatementContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

// ====================
// CONTROL KEYWORDS
// ====================

func (l Listener) ExitContinueStatement(ctx *p.ContinueStatementContext) {
	scope := l.GetCurrentScope()
	loopScope, inLoopScope := l.TypeChecker.ScopeManager.SearchScopeByType(type_checker.SCOPE_TYPES.LOOP)
	if !inLoopScope {
		panic("Continue should be on loop statement")
	}
	returnTag := loopScope.Name + "_UPDATE"
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(returnTag),
	}))
}
func (l Listener) ExitBreakStatement(ctx *p.BreakStatementContext) {
	scope := l.GetCurrentScope()
	loopScope, inLoopScope := l.TypeChecker.ScopeManager.SearchScopeByType(type_checker.SCOPE_TYPES.LOOP)
	if !inLoopScope {
		panic("Continue should be on loop statement")
	}
	returnTag := loopScope.Name + "_RETURN"
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(returnTag),
	}))
}

// ================
// 	SWITCH
// ================

func (l Listener) EnterSwitchStatement(ctx *p.SwitchStatementContext) {
	scope := l.GetCurrentScope()
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	switchScope := l.GetCurrentScope()
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(switchScope.Name),
	}))
	l.AppendInstruction(ScopeName(scope.Name), NewSecInstruction(SecInstruction{
		Name: TagName(switchScope.Name + "_RETURN"),
	}))

	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(switchScope.Name), parentName)
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) ExitSwitchStatement(ctx *p.SwitchStatementContext) {
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

func (l Listener) ExitSwitchValue(ctx *p.SwitchValueContext) {
	scope := l.GetCurrentScope()
	variableName, _ := l.Program.GetVariableFor(ctx.GetText(), ScopeName(scope.Name))
	l.Program.UpsertTranslation(ScopeName(scope.Name), "$switch", "$switch")
	l.AppendInstruction(ScopeName(scope.Name), NewCopyInstruction(CopyInstruction{
		Target: "$switch",
		Source: variableName,
	}))
}

func (l Listener) EnterSwitchCase(ctx *p.SwitchCaseContext) {
	// parentScope := l.GetCurrentScope()
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	childScope := l.GetCurrentScope()

	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(childScope.Name), parentName)
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) EnterCaseBody(ctx *p.CaseBodyContext) {
	scope := l.GetCurrentScope()
	tagName := scope.Name + "_BODY"
	l.AppendInstruction(ScopeName(scope.Name), NewSecInstruction(SecInstruction{
		Name: TagName(tagName),
	}))
}

func (l Listener) ExitSwitchCase(ctx *p.SwitchCaseContext) {
	scope := l.GetCurrentScope()
	scopeWithoutId := scope.Name[:len(scope.Name)-1]

	tagName := strings.TrimSuffix(scopeWithoutId, "_CASE_") + "_RETURN"
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(tagName),
	}))

	tagName = scope.Name + "_END"
	l.AppendInstruction(ScopeName(scope.Name), NewSecInstruction(SecInstruction{
		Name: TagName(tagName),
	}))
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

func (l Listener) EnterDefaultCase(ctx *p.DefaultCaseContext) {
	// parentScope := l.GetCurrentScope()
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	childScope := l.GetCurrentScope()
	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(childScope.Name), parentName)

	if err != nil {
		log.Println("Something when wrong during Scope management")
	}

	// l.AppendInstruction(ScopeName(parentScope.Name), NewJumpInstruction(JumpInstruction{
	// 	Condition: lib.NewOpEmpty[JumpCondition](),
	// 	Target:    TagName(childScope.Name),
	// }))
}

func (l Listener) ExitDefaultCase(ctx *p.DefaultCaseContext) {
	scope := l.GetCurrentScope()
	tagName := strings.TrimSuffix(l.GetCurrentScope().Name, "_DEFAULT") + "_RETURN"
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(tagName),
	}))
	l.TypeChecker.ScopeManager.ReplaceWithParent()
}

func (l Listener) ExitCaseValue(ctx *p.CaseValueContext) {
	scope := l.GetCurrentScope()
	targetVar, _ := l.Program.GetVariableOrLiteral(
		ctx.GetText(), ScopeName(scope.Name), l.TypeChecker)
	tagName := l.GetCurrentScope().Name + "_BODY"
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.NewOpValue(
			JumpCondition{Relation: lib.NewOpValue(CondJumpOperation{
				Signed: false,
				Type:   BOOLEAN_OPERATION_TYPES.Equal,
				P1:     "$switch",
				P2:     targetVar,
			})}),
		Target: TagName(tagName),
	}))
	tagName = l.GetCurrentScope().Name + "_END"
	// tagName := strings.TrimSuffix(scopeWithoutId, "_CASE_") + "_RETURN"
	l.AppendInstruction(ScopeName(scope.Name), NewJumpInstruction(JumpInstruction{
		Condition: lib.Optional[JumpCondition]{},
		Target:    TagName(tagName),
	}))

}

// ================
// 	TRY CATH
// ================

func (l Listener) EnterTryStatement(ctx *p.TryStatementContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(scope.Name), parentName)
	if err != nil {
		log.Println("Something when wrong during Scope management")
	}
}

func (l Listener) EnterCatchStatement(ctx *p.CatchStatementContext) {
	err := l.TypeChecker.ScopeManager.ReplaceWithNextChild()
	scope := l.GetCurrentScope()
	parentName := l.GetParentScopeName()
	l.Program.UpsertScope(ScopeName(scope.Name), parentName)
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
