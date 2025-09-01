// Code generated from compiscript/program/Compiscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Compiscript

import "github.com/antlr4-go/antlr/v4"

// BaseCompiscriptListener is a complete listener for a parse tree produced by CompiscriptParser.
type BaseCompiscriptListener struct{}

var _ CompiscriptListener = &BaseCompiscriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCompiscriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCompiscriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCompiscriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCompiscriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseCompiscriptListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseCompiscriptListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCompiscriptListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCompiscriptListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseCompiscriptListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseCompiscriptListener) ExitBlock(ctx *BlockContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseCompiscriptListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseCompiscriptListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterConstantDeclaration is called when production constantDeclaration is entered.
func (s *BaseCompiscriptListener) EnterConstantDeclaration(ctx *ConstantDeclarationContext) {}

// ExitConstantDeclaration is called when production constantDeclaration is exited.
func (s *BaseCompiscriptListener) ExitConstantDeclaration(ctx *ConstantDeclarationContext) {}

// EnterTypeAnnotation is called when production typeAnnotation is entered.
func (s *BaseCompiscriptListener) EnterTypeAnnotation(ctx *TypeAnnotationContext) {}

// ExitTypeAnnotation is called when production typeAnnotation is exited.
func (s *BaseCompiscriptListener) ExitTypeAnnotation(ctx *TypeAnnotationContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *BaseCompiscriptListener) EnterInitializer(ctx *InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *BaseCompiscriptListener) ExitInitializer(ctx *InitializerContext) {}

// EnterThisAssignment is called when production ThisAssignment is entered.
func (s *BaseCompiscriptListener) EnterThisAssignment(ctx *ThisAssignmentContext) {}

// ExitThisAssignment is called when production ThisAssignment is exited.
func (s *BaseCompiscriptListener) ExitThisAssignment(ctx *ThisAssignmentContext) {}

// EnterVariableAssignment is called when production VariableAssignment is entered.
func (s *BaseCompiscriptListener) EnterVariableAssignment(ctx *VariableAssignmentContext) {}

// ExitVariableAssignment is called when production VariableAssignment is exited.
func (s *BaseCompiscriptListener) ExitVariableAssignment(ctx *VariableAssignmentContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseCompiscriptListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseCompiscriptListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterMustBoolExpr is called when production mustBoolExpr is entered.
func (s *BaseCompiscriptListener) EnterMustBoolExpr(ctx *MustBoolExprContext) {}

// ExitMustBoolExpr is called when production mustBoolExpr is exited.
func (s *BaseCompiscriptListener) ExitMustBoolExpr(ctx *MustBoolExprContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseCompiscriptListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseCompiscriptListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterIfBody is called when production ifBody is entered.
func (s *BaseCompiscriptListener) EnterIfBody(ctx *IfBodyContext) {}

// ExitIfBody is called when production ifBody is exited.
func (s *BaseCompiscriptListener) ExitIfBody(ctx *IfBodyContext) {}

// EnterElseBody is called when production elseBody is entered.
func (s *BaseCompiscriptListener) EnterElseBody(ctx *ElseBodyContext) {}

// ExitElseBody is called when production elseBody is exited.
func (s *BaseCompiscriptListener) ExitElseBody(ctx *ElseBodyContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseCompiscriptListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseCompiscriptListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterWhileBody is called when production whileBody is entered.
func (s *BaseCompiscriptListener) EnterWhileBody(ctx *WhileBodyContext) {}

// ExitWhileBody is called when production whileBody is exited.
func (s *BaseCompiscriptListener) ExitWhileBody(ctx *WhileBodyContext) {}

// EnterDoWhileStatement is called when production doWhileStatement is entered.
func (s *BaseCompiscriptListener) EnterDoWhileStatement(ctx *DoWhileStatementContext) {}

// ExitDoWhileStatement is called when production doWhileStatement is exited.
func (s *BaseCompiscriptListener) ExitDoWhileStatement(ctx *DoWhileStatementContext) {}

// EnterDoWhileBody is called when production doWhileBody is entered.
func (s *BaseCompiscriptListener) EnterDoWhileBody(ctx *DoWhileBodyContext) {}

// ExitDoWhileBody is called when production doWhileBody is exited.
func (s *BaseCompiscriptListener) ExitDoWhileBody(ctx *DoWhileBodyContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseCompiscriptListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseCompiscriptListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForeachStatement is called when production foreachStatement is entered.
func (s *BaseCompiscriptListener) EnterForeachStatement(ctx *ForeachStatementContext) {}

// ExitForeachStatement is called when production foreachStatement is exited.
func (s *BaseCompiscriptListener) ExitForeachStatement(ctx *ForeachStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseCompiscriptListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseCompiscriptListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseCompiscriptListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseCompiscriptListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseCompiscriptListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseCompiscriptListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseCompiscriptListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseCompiscriptListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseCompiscriptListener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseCompiscriptListener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterCatchStatement is called when production catchStatement is entered.
func (s *BaseCompiscriptListener) EnterCatchStatement(ctx *CatchStatementContext) {}

// ExitCatchStatement is called when production catchStatement is exited.
func (s *BaseCompiscriptListener) ExitCatchStatement(ctx *CatchStatementContext) {}

// EnterSwitchValue is called when production switchValue is entered.
func (s *BaseCompiscriptListener) EnterSwitchValue(ctx *SwitchValueContext) {}

// ExitSwitchValue is called when production switchValue is exited.
func (s *BaseCompiscriptListener) ExitSwitchValue(ctx *SwitchValueContext) {}

// EnterCaseValue is called when production caseValue is entered.
func (s *BaseCompiscriptListener) EnterCaseValue(ctx *CaseValueContext) {}

// ExitCaseValue is called when production caseValue is exited.
func (s *BaseCompiscriptListener) ExitCaseValue(ctx *CaseValueContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseCompiscriptListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseCompiscriptListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterSwitchCase is called when production switchCase is entered.
func (s *BaseCompiscriptListener) EnterSwitchCase(ctx *SwitchCaseContext) {}

// ExitSwitchCase is called when production switchCase is exited.
func (s *BaseCompiscriptListener) ExitSwitchCase(ctx *SwitchCaseContext) {}

// EnterDefaultCase is called when production defaultCase is entered.
func (s *BaseCompiscriptListener) EnterDefaultCase(ctx *DefaultCaseContext) {}

// ExitDefaultCase is called when production defaultCase is exited.
func (s *BaseCompiscriptListener) ExitDefaultCase(ctx *DefaultCaseContext) {}

// EnterCaseBody is called when production caseBody is entered.
func (s *BaseCompiscriptListener) EnterCaseBody(ctx *CaseBodyContext) {}

// ExitCaseBody is called when production caseBody is exited.
func (s *BaseCompiscriptListener) ExitCaseBody(ctx *CaseBodyContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseCompiscriptListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseCompiscriptListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseCompiscriptListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseCompiscriptListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseCompiscriptListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseCompiscriptListener) ExitParameter(ctx *ParameterContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseCompiscriptListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseCompiscriptListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassMember is called when production classMember is entered.
func (s *BaseCompiscriptListener) EnterClassMember(ctx *ClassMemberContext) {}

// ExitClassMember is called when production classMember is exited.
func (s *BaseCompiscriptListener) ExitClassMember(ctx *ClassMemberContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCompiscriptListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCompiscriptListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignmentExpr is called when production assignmentExpr is entered.
func (s *BaseCompiscriptListener) EnterAssignmentExpr(ctx *AssignmentExprContext) {}

// ExitAssignmentExpr is called when production assignmentExpr is exited.
func (s *BaseCompiscriptListener) ExitAssignmentExpr(ctx *AssignmentExprContext) {}

// EnterConditionalExpr is called when production conditionalExpr is entered.
func (s *BaseCompiscriptListener) EnterConditionalExpr(ctx *ConditionalExprContext) {}

// ExitConditionalExpr is called when production conditionalExpr is exited.
func (s *BaseCompiscriptListener) ExitConditionalExpr(ctx *ConditionalExprContext) {}

// EnterLogicalOrExpr is called when production logicalOrExpr is entered.
func (s *BaseCompiscriptListener) EnterLogicalOrExpr(ctx *LogicalOrExprContext) {}

// ExitLogicalOrExpr is called when production logicalOrExpr is exited.
func (s *BaseCompiscriptListener) ExitLogicalOrExpr(ctx *LogicalOrExprContext) {}

// EnterLogicalAndExpr is called when production logicalAndExpr is entered.
func (s *BaseCompiscriptListener) EnterLogicalAndExpr(ctx *LogicalAndExprContext) {}

// ExitLogicalAndExpr is called when production logicalAndExpr is exited.
func (s *BaseCompiscriptListener) ExitLogicalAndExpr(ctx *LogicalAndExprContext) {}

// EnterEqualityExpr is called when production equalityExpr is entered.
func (s *BaseCompiscriptListener) EnterEqualityExpr(ctx *EqualityExprContext) {}

// ExitEqualityExpr is called when production equalityExpr is exited.
func (s *BaseCompiscriptListener) ExitEqualityExpr(ctx *EqualityExprContext) {}

// EnterRelationalExpr is called when production relationalExpr is entered.
func (s *BaseCompiscriptListener) EnterRelationalExpr(ctx *RelationalExprContext) {}

// ExitRelationalExpr is called when production relationalExpr is exited.
func (s *BaseCompiscriptListener) ExitRelationalExpr(ctx *RelationalExprContext) {}

// EnterAdditiveExpr is called when production additiveExpr is entered.
func (s *BaseCompiscriptListener) EnterAdditiveExpr(ctx *AdditiveExprContext) {}

// ExitAdditiveExpr is called when production additiveExpr is exited.
func (s *BaseCompiscriptListener) ExitAdditiveExpr(ctx *AdditiveExprContext) {}

// EnterMultiplicativeExpr is called when production multiplicativeExpr is entered.
func (s *BaseCompiscriptListener) EnterMultiplicativeExpr(ctx *MultiplicativeExprContext) {}

// ExitMultiplicativeExpr is called when production multiplicativeExpr is exited.
func (s *BaseCompiscriptListener) ExitMultiplicativeExpr(ctx *MultiplicativeExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseCompiscriptListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseCompiscriptListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseCompiscriptListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseCompiscriptListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterLiteralExpr is called when production literalExpr is entered.
func (s *BaseCompiscriptListener) EnterLiteralExpr(ctx *LiteralExprContext) {}

// ExitLiteralExpr is called when production literalExpr is exited.
func (s *BaseCompiscriptListener) ExitLiteralExpr(ctx *LiteralExprContext) {}

// EnterLeftHandSide is called when production leftHandSide is entered.
func (s *BaseCompiscriptListener) EnterLeftHandSide(ctx *LeftHandSideContext) {}

// ExitLeftHandSide is called when production leftHandSide is exited.
func (s *BaseCompiscriptListener) ExitLeftHandSide(ctx *LeftHandSideContext) {}

// EnterStandaloneExpresion is called when production standaloneExpresion is entered.
func (s *BaseCompiscriptListener) EnterStandaloneExpresion(ctx *StandaloneExpresionContext) {}

// ExitStandaloneExpresion is called when production standaloneExpresion is exited.
func (s *BaseCompiscriptListener) ExitStandaloneExpresion(ctx *StandaloneExpresionContext) {}

// EnterIdentifierExpr is called when production IdentifierExpr is entered.
func (s *BaseCompiscriptListener) EnterIdentifierExpr(ctx *IdentifierExprContext) {}

// ExitIdentifierExpr is called when production IdentifierExpr is exited.
func (s *BaseCompiscriptListener) ExitIdentifierExpr(ctx *IdentifierExprContext) {}

// EnterNewExpr is called when production NewExpr is entered.
func (s *BaseCompiscriptListener) EnterNewExpr(ctx *NewExprContext) {}

// ExitNewExpr is called when production NewExpr is exited.
func (s *BaseCompiscriptListener) ExitNewExpr(ctx *NewExprContext) {}

// EnterThisExpr is called when production ThisExpr is entered.
func (s *BaseCompiscriptListener) EnterThisExpr(ctx *ThisExprContext) {}

// ExitThisExpr is called when production ThisExpr is exited.
func (s *BaseCompiscriptListener) ExitThisExpr(ctx *ThisExprContext) {}

// EnterStandaloneIdentifierExpr is called when production StandaloneIdentifierExpr is entered.
func (s *BaseCompiscriptListener) EnterStandaloneIdentifierExpr(ctx *StandaloneIdentifierExprContext) {
}

// ExitStandaloneIdentifierExpr is called when production StandaloneIdentifierExpr is exited.
func (s *BaseCompiscriptListener) ExitStandaloneIdentifierExpr(ctx *StandaloneIdentifierExprContext) {
}

// EnterStandaloneNewExpr is called when production StandaloneNewExpr is entered.
func (s *BaseCompiscriptListener) EnterStandaloneNewExpr(ctx *StandaloneNewExprContext) {}

// ExitStandaloneNewExpr is called when production StandaloneNewExpr is exited.
func (s *BaseCompiscriptListener) ExitStandaloneNewExpr(ctx *StandaloneNewExprContext) {}

// EnterStandaloneThisExpr is called when production StandaloneThisExpr is entered.
func (s *BaseCompiscriptListener) EnterStandaloneThisExpr(ctx *StandaloneThisExprContext) {}

// ExitStandaloneThisExpr is called when production StandaloneThisExpr is exited.
func (s *BaseCompiscriptListener) ExitStandaloneThisExpr(ctx *StandaloneThisExprContext) {}

// EnterMethodCallExpr is called when production MethodCallExpr is entered.
func (s *BaseCompiscriptListener) EnterMethodCallExpr(ctx *MethodCallExprContext) {}

// ExitMethodCallExpr is called when production MethodCallExpr is exited.
func (s *BaseCompiscriptListener) ExitMethodCallExpr(ctx *MethodCallExprContext) {}

// EnterCallExpr is called when production CallExpr is entered.
func (s *BaseCompiscriptListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production CallExpr is exited.
func (s *BaseCompiscriptListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterIndexExpr is called when production IndexExpr is entered.
func (s *BaseCompiscriptListener) EnterIndexExpr(ctx *IndexExprContext) {}

// ExitIndexExpr is called when production IndexExpr is exited.
func (s *BaseCompiscriptListener) ExitIndexExpr(ctx *IndexExprContext) {}

// EnterPropertyAccessExpr is called when production PropertyAccessExpr is entered.
func (s *BaseCompiscriptListener) EnterPropertyAccessExpr(ctx *PropertyAccessExprContext) {}

// ExitPropertyAccessExpr is called when production PropertyAccessExpr is exited.
func (s *BaseCompiscriptListener) ExitPropertyAccessExpr(ctx *PropertyAccessExprContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseCompiscriptListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseCompiscriptListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseCompiscriptListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseCompiscriptListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterType is called when production type is entered.
func (s *BaseCompiscriptListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseCompiscriptListener) ExitType(ctx *TypeContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseCompiscriptListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseCompiscriptListener) ExitBaseType(ctx *BaseTypeContext) {}
