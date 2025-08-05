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

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterDoWhileStatement is called when entering the doWhileStatement production.
	EnterDoWhileStatement(c *DoWhileStatementContext)

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

	// EnterTryCatchStatement is called when entering the tryCatchStatement production.
	EnterTryCatchStatement(c *TryCatchStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterSwitchCase is called when entering the switchCase production.
	EnterSwitchCase(c *SwitchCaseContext)

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

	// EnterAssignExpr is called when entering the AssignExpr production.
	EnterAssignExpr(c *AssignExprContext)

	// EnterPropertyAssignExpr is called when entering the PropertyAssignExpr production.
	EnterPropertyAssignExpr(c *PropertyAssignExprContext)

	// EnterExprNoAssign is called when entering the ExprNoAssign production.
	EnterExprNoAssign(c *ExprNoAssignContext)

	// EnterTernaryExpr is called when entering the TernaryExpr production.
	EnterTernaryExpr(c *TernaryExprContext)

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

	// EnterIdentifierExpr is called when entering the IdentifierExpr production.
	EnterIdentifierExpr(c *IdentifierExprContext)

	// EnterNewExpr is called when entering the NewExpr production.
	EnterNewExpr(c *NewExprContext)

	// EnterThisExpr is called when entering the ThisExpr production.
	EnterThisExpr(c *ThisExprContext)

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

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitDoWhileStatement is called when exiting the doWhileStatement production.
	ExitDoWhileStatement(c *DoWhileStatementContext)

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

	// ExitTryCatchStatement is called when exiting the tryCatchStatement production.
	ExitTryCatchStatement(c *TryCatchStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitSwitchCase is called when exiting the switchCase production.
	ExitSwitchCase(c *SwitchCaseContext)

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

	// ExitAssignExpr is called when exiting the AssignExpr production.
	ExitAssignExpr(c *AssignExprContext)

	// ExitPropertyAssignExpr is called when exiting the PropertyAssignExpr production.
	ExitPropertyAssignExpr(c *PropertyAssignExprContext)

	// ExitExprNoAssign is called when exiting the ExprNoAssign production.
	ExitExprNoAssign(c *ExprNoAssignContext)

	// ExitTernaryExpr is called when exiting the TernaryExpr production.
	ExitTernaryExpr(c *TernaryExprContext)

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

	// ExitIdentifierExpr is called when exiting the IdentifierExpr production.
	ExitIdentifierExpr(c *IdentifierExprContext)

	// ExitNewExpr is called when exiting the NewExpr production.
	ExitNewExpr(c *NewExprContext)

	// ExitThisExpr is called when exiting the ThisExpr production.
	ExitThisExpr(c *ThisExprContext)

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
