// Code generated from compiscript/program/Compiscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Compiscript

import "github.com/antlr4-go/antlr/v4"

// CompiscriptListener is a complete listener for a parse tree produced by CompiscriptParser.
type CompiscriptListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterConstantDeclaration is called when entering the constantDeclaration production.
	EnterConstantDeclaration(c *ConstantDeclarationContext)

	// EnterTypeAnnotation is called when entering the typeAnnotation production.
	EnterTypeAnnotation(c *TypeAnnotationContext)

	// EnterInitializer is called when entering the initializer production.
	EnterInitializer(c *InitializerContext)

	// EnterThisAssignment is called when entering the ThisAssignment production.
	EnterThisAssignment(c *ThisAssignmentContext)

	// EnterVariableAssignment is called when entering the VariableAssignment production.
	EnterVariableAssignment(c *VariableAssignmentContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterMustBoolExpr is called when entering the mustBoolExpr production.
	EnterMustBoolExpr(c *MustBoolExprContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterIfBody is called when entering the ifBody production.
	EnterIfBody(c *IfBodyContext)

	// EnterElseBody is called when entering the elseBody production.
	EnterElseBody(c *ElseBodyContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterWhileBody is called when entering the whileBody production.
	EnterWhileBody(c *WhileBodyContext)

	// EnterDoWhileStatement is called when entering the doWhileStatement production.
	EnterDoWhileStatement(c *DoWhileStatementContext)

	// EnterDoWhileBody is called when entering the doWhileBody production.
	EnterDoWhileBody(c *DoWhileBodyContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForeachStatement is called when entering the foreachStatement production.
	EnterForeachStatement(c *ForeachStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterTryCatchStatement is called when entering the tryCatchStatement production.
	EnterTryCatchStatement(c *TryCatchStatementContext)

	// EnterCatchStatement is called when entering the catchStatement production.
	EnterCatchStatement(c *CatchStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterSwitchCase is called when entering the switchCase production.
	EnterSwitchCase(c *SwitchCaseContext)

	// EnterCaseBody is called when entering the caseBody production.
	EnterCaseBody(c *CaseBodyContext)

	// EnterDefaultCase is called when entering the defaultCase production.
	EnterDefaultCase(c *DefaultCaseContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassMember is called when entering the classMember production.
	EnterClassMember(c *ClassMemberContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAssignmentExpr is called when entering the assignmentExpr production.
	EnterAssignmentExpr(c *AssignmentExprContext)

	// EnterConditionalExpr is called when entering the conditionalExpr production.
	EnterConditionalExpr(c *ConditionalExprContext)

	// EnterLogicalOrExpr is called when entering the logicalOrExpr production.
	EnterLogicalOrExpr(c *LogicalOrExprContext)

	// EnterLogicalAndExpr is called when entering the logicalAndExpr production.
	EnterLogicalAndExpr(c *LogicalAndExprContext)

	// EnterEqualityExpr is called when entering the equalityExpr production.
	EnterEqualityExpr(c *EqualityExprContext)

	// EnterRelationalExpr is called when entering the relationalExpr production.
	EnterRelationalExpr(c *RelationalExprContext)

	// EnterAdditiveExpr is called when entering the additiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterMultiplicativeExpr is called when entering the multiplicativeExpr production.
	EnterMultiplicativeExpr(c *MultiplicativeExprContext)

	// EnterUnaryExpr is called when entering the unaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterLiteralExpr is called when entering the literalExpr production.
	EnterLiteralExpr(c *LiteralExprContext)

	// EnterLeftHandSide is called when entering the leftHandSide production.
	EnterLeftHandSide(c *LeftHandSideContext)

	// EnterStandaloneExpresion is called when entering the standaloneExpresion production.
	EnterStandaloneExpresion(c *StandaloneExpresionContext)

	// EnterIdentifierExpr is called when entering the IdentifierExpr production.
	EnterIdentifierExpr(c *IdentifierExprContext)

	// EnterNewExpr is called when entering the NewExpr production.
	EnterNewExpr(c *NewExprContext)

	// EnterThisExpr is called when entering the ThisExpr production.
	EnterThisExpr(c *ThisExprContext)

	// EnterStandaloneIdentifierExpr is called when entering the StandaloneIdentifierExpr production.
	EnterStandaloneIdentifierExpr(c *StandaloneIdentifierExprContext)

	// EnterStandaloneNewExpr is called when entering the StandaloneNewExpr production.
	EnterStandaloneNewExpr(c *StandaloneNewExprContext)

	// EnterStandaloneThisExpr is called when entering the StandaloneThisExpr production.
	EnterStandaloneThisExpr(c *StandaloneThisExprContext)

	// EnterMethodCallExpr is called when entering the MethodCallExpr production.
	EnterMethodCallExpr(c *MethodCallExprContext)

	// EnterCallExpr is called when entering the CallExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterIndexExpr is called when entering the IndexExpr production.
	EnterIndexExpr(c *IndexExprContext)

	// EnterPropertyAccessExpr is called when entering the PropertyAccessExpr production.
	EnterPropertyAccessExpr(c *PropertyAccessExprContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterBaseType is called when entering the baseType production.
	EnterBaseType(c *BaseTypeContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitConstantDeclaration is called when exiting the constantDeclaration production.
	ExitConstantDeclaration(c *ConstantDeclarationContext)

	// ExitTypeAnnotation is called when exiting the typeAnnotation production.
	ExitTypeAnnotation(c *TypeAnnotationContext)

	// ExitInitializer is called when exiting the initializer production.
	ExitInitializer(c *InitializerContext)

	// ExitThisAssignment is called when exiting the ThisAssignment production.
	ExitThisAssignment(c *ThisAssignmentContext)

	// ExitVariableAssignment is called when exiting the VariableAssignment production.
	ExitVariableAssignment(c *VariableAssignmentContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitMustBoolExpr is called when exiting the mustBoolExpr production.
	ExitMustBoolExpr(c *MustBoolExprContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitIfBody is called when exiting the ifBody production.
	ExitIfBody(c *IfBodyContext)

	// ExitElseBody is called when exiting the elseBody production.
	ExitElseBody(c *ElseBodyContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitWhileBody is called when exiting the whileBody production.
	ExitWhileBody(c *WhileBodyContext)

	// ExitDoWhileStatement is called when exiting the doWhileStatement production.
	ExitDoWhileStatement(c *DoWhileStatementContext)

	// ExitDoWhileBody is called when exiting the doWhileBody production.
	ExitDoWhileBody(c *DoWhileBodyContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForeachStatement is called when exiting the foreachStatement production.
	ExitForeachStatement(c *ForeachStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitTryCatchStatement is called when exiting the tryCatchStatement production.
	ExitTryCatchStatement(c *TryCatchStatementContext)

	// ExitCatchStatement is called when exiting the catchStatement production.
	ExitCatchStatement(c *CatchStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitSwitchCase is called when exiting the switchCase production.
	ExitSwitchCase(c *SwitchCaseContext)

	// ExitCaseBody is called when exiting the caseBody production.
	ExitCaseBody(c *CaseBodyContext)

	// ExitDefaultCase is called when exiting the defaultCase production.
	ExitDefaultCase(c *DefaultCaseContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassMember is called when exiting the classMember production.
	ExitClassMember(c *ClassMemberContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAssignmentExpr is called when exiting the assignmentExpr production.
	ExitAssignmentExpr(c *AssignmentExprContext)

	// ExitConditionalExpr is called when exiting the conditionalExpr production.
	ExitConditionalExpr(c *ConditionalExprContext)

	// ExitLogicalOrExpr is called when exiting the logicalOrExpr production.
	ExitLogicalOrExpr(c *LogicalOrExprContext)

	// ExitLogicalAndExpr is called when exiting the logicalAndExpr production.
	ExitLogicalAndExpr(c *LogicalAndExprContext)

	// ExitEqualityExpr is called when exiting the equalityExpr production.
	ExitEqualityExpr(c *EqualityExprContext)

	// ExitRelationalExpr is called when exiting the relationalExpr production.
	ExitRelationalExpr(c *RelationalExprContext)

	// ExitAdditiveExpr is called when exiting the additiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitMultiplicativeExpr is called when exiting the multiplicativeExpr production.
	ExitMultiplicativeExpr(c *MultiplicativeExprContext)

	// ExitUnaryExpr is called when exiting the unaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitLiteralExpr is called when exiting the literalExpr production.
	ExitLiteralExpr(c *LiteralExprContext)

	// ExitLeftHandSide is called when exiting the leftHandSide production.
	ExitLeftHandSide(c *LeftHandSideContext)

	// ExitStandaloneExpresion is called when exiting the standaloneExpresion production.
	ExitStandaloneExpresion(c *StandaloneExpresionContext)

	// ExitIdentifierExpr is called when exiting the IdentifierExpr production.
	ExitIdentifierExpr(c *IdentifierExprContext)

	// ExitNewExpr is called when exiting the NewExpr production.
	ExitNewExpr(c *NewExprContext)

	// ExitThisExpr is called when exiting the ThisExpr production.
	ExitThisExpr(c *ThisExprContext)

	// ExitStandaloneIdentifierExpr is called when exiting the StandaloneIdentifierExpr production.
	ExitStandaloneIdentifierExpr(c *StandaloneIdentifierExprContext)

	// ExitStandaloneNewExpr is called when exiting the StandaloneNewExpr production.
	ExitStandaloneNewExpr(c *StandaloneNewExprContext)

	// ExitStandaloneThisExpr is called when exiting the StandaloneThisExpr production.
	ExitStandaloneThisExpr(c *StandaloneThisExprContext)

	// ExitMethodCallExpr is called when exiting the MethodCallExpr production.
	ExitMethodCallExpr(c *MethodCallExprContext)

	// ExitCallExpr is called when exiting the CallExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitIndexExpr is called when exiting the IndexExpr production.
	ExitIndexExpr(c *IndexExprContext)

	// ExitPropertyAccessExpr is called when exiting the PropertyAccessExpr production.
	ExitPropertyAccessExpr(c *PropertyAccessExprContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitBaseType is called when exiting the baseType production.
	ExitBaseType(c *BaseTypeContext)
}
