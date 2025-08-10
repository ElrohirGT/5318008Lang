// Generated from ./program/Compiscript.g4 by ANTLR 4.13.2
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link CompiscriptParser}.
 */
public interface CompiscriptListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(CompiscriptParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(CompiscriptParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(CompiscriptParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(CompiscriptParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(CompiscriptParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(CompiscriptParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterVariableDeclaration(CompiscriptParser.VariableDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitVariableDeclaration(CompiscriptParser.VariableDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#constantDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterConstantDeclaration(CompiscriptParser.ConstantDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#constantDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitConstantDeclaration(CompiscriptParser.ConstantDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#typeAnnotation}.
	 * @param ctx the parse tree
	 */
	void enterTypeAnnotation(CompiscriptParser.TypeAnnotationContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#typeAnnotation}.
	 * @param ctx the parse tree
	 */
	void exitTypeAnnotation(CompiscriptParser.TypeAnnotationContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#initializer}.
	 * @param ctx the parse tree
	 */
	void enterInitializer(CompiscriptParser.InitializerContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#initializer}.
	 * @param ctx the parse tree
	 */
	void exitInitializer(CompiscriptParser.InitializerContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#assignment}.
	 * @param ctx the parse tree
	 */
	void enterAssignment(CompiscriptParser.AssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#assignment}.
	 * @param ctx the parse tree
	 */
	void exitAssignment(CompiscriptParser.AssignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterExpressionStatement(CompiscriptParser.ExpressionStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitExpressionStatement(CompiscriptParser.ExpressionStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#printStatement}.
	 * @param ctx the parse tree
	 */
	void enterPrintStatement(CompiscriptParser.PrintStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#printStatement}.
	 * @param ctx the parse tree
	 */
	void exitPrintStatement(CompiscriptParser.PrintStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatement(CompiscriptParser.IfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatement(CompiscriptParser.IfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#whileStatement}.
	 * @param ctx the parse tree
	 */
	void enterWhileStatement(CompiscriptParser.WhileStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#whileStatement}.
	 * @param ctx the parse tree
	 */
	void exitWhileStatement(CompiscriptParser.WhileStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#doWhileStatement}.
	 * @param ctx the parse tree
	 */
	void enterDoWhileStatement(CompiscriptParser.DoWhileStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#doWhileStatement}.
	 * @param ctx the parse tree
	 */
	void exitDoWhileStatement(CompiscriptParser.DoWhileStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void enterForStatement(CompiscriptParser.ForStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void exitForStatement(CompiscriptParser.ForStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#foreachStatement}.
	 * @param ctx the parse tree
	 */
	void enterForeachStatement(CompiscriptParser.ForeachStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#foreachStatement}.
	 * @param ctx the parse tree
	 */
	void exitForeachStatement(CompiscriptParser.ForeachStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#breakStatement}.
	 * @param ctx the parse tree
	 */
	void enterBreakStatement(CompiscriptParser.BreakStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#breakStatement}.
	 * @param ctx the parse tree
	 */
	void exitBreakStatement(CompiscriptParser.BreakStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#continueStatement}.
	 * @param ctx the parse tree
	 */
	void enterContinueStatement(CompiscriptParser.ContinueStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#continueStatement}.
	 * @param ctx the parse tree
	 */
	void exitContinueStatement(CompiscriptParser.ContinueStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#returnStatement}.
	 * @param ctx the parse tree
	 */
	void enterReturnStatement(CompiscriptParser.ReturnStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#returnStatement}.
	 * @param ctx the parse tree
	 */
	void exitReturnStatement(CompiscriptParser.ReturnStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#tryCatchStatement}.
	 * @param ctx the parse tree
	 */
	void enterTryCatchStatement(CompiscriptParser.TryCatchStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#tryCatchStatement}.
	 * @param ctx the parse tree
	 */
	void exitTryCatchStatement(CompiscriptParser.TryCatchStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStatement(CompiscriptParser.SwitchStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStatement(CompiscriptParser.SwitchStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#switchCase}.
	 * @param ctx the parse tree
	 */
	void enterSwitchCase(CompiscriptParser.SwitchCaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#switchCase}.
	 * @param ctx the parse tree
	 */
	void exitSwitchCase(CompiscriptParser.SwitchCaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#defaultCase}.
	 * @param ctx the parse tree
	 */
	void enterDefaultCase(CompiscriptParser.DefaultCaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#defaultCase}.
	 * @param ctx the parse tree
	 */
	void exitDefaultCase(CompiscriptParser.DefaultCaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#functionDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterFunctionDeclaration(CompiscriptParser.FunctionDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#functionDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitFunctionDeclaration(CompiscriptParser.FunctionDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#parameters}.
	 * @param ctx the parse tree
	 */
	void enterParameters(CompiscriptParser.ParametersContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#parameters}.
	 * @param ctx the parse tree
	 */
	void exitParameters(CompiscriptParser.ParametersContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#parameter}.
	 * @param ctx the parse tree
	 */
	void enterParameter(CompiscriptParser.ParameterContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#parameter}.
	 * @param ctx the parse tree
	 */
	void exitParameter(CompiscriptParser.ParameterContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#classDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterClassDeclaration(CompiscriptParser.ClassDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#classDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitClassDeclaration(CompiscriptParser.ClassDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#classMember}.
	 * @param ctx the parse tree
	 */
	void enterClassMember(CompiscriptParser.ClassMemberContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#classMember}.
	 * @param ctx the parse tree
	 */
	void exitClassMember(CompiscriptParser.ClassMemberContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(CompiscriptParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(CompiscriptParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AssignExpr}
	 * labeled alternative in {@link CompiscriptParser#assignmentExpr}.
	 * @param ctx the parse tree
	 */
	void enterAssignExpr(CompiscriptParser.AssignExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AssignExpr}
	 * labeled alternative in {@link CompiscriptParser#assignmentExpr}.
	 * @param ctx the parse tree
	 */
	void exitAssignExpr(CompiscriptParser.AssignExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code PropertyAssignExpr}
	 * labeled alternative in {@link CompiscriptParser#assignmentExpr}.
	 * @param ctx the parse tree
	 */
	void enterPropertyAssignExpr(CompiscriptParser.PropertyAssignExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code PropertyAssignExpr}
	 * labeled alternative in {@link CompiscriptParser#assignmentExpr}.
	 * @param ctx the parse tree
	 */
	void exitPropertyAssignExpr(CompiscriptParser.PropertyAssignExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ExprNoAssign}
	 * labeled alternative in {@link CompiscriptParser#assignmentExpr}.
	 * @param ctx the parse tree
	 */
	void enterExprNoAssign(CompiscriptParser.ExprNoAssignContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ExprNoAssign}
	 * labeled alternative in {@link CompiscriptParser#assignmentExpr}.
	 * @param ctx the parse tree
	 */
	void exitExprNoAssign(CompiscriptParser.ExprNoAssignContext ctx);
	/**
	 * Enter a parse tree produced by the {@code TernaryExpr}
	 * labeled alternative in {@link CompiscriptParser#conditionalExpr}.
	 * @param ctx the parse tree
	 */
	void enterTernaryExpr(CompiscriptParser.TernaryExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code TernaryExpr}
	 * labeled alternative in {@link CompiscriptParser#conditionalExpr}.
	 * @param ctx the parse tree
	 */
	void exitTernaryExpr(CompiscriptParser.TernaryExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#logicalOrExpr}.
	 * @param ctx the parse tree
	 */
	void enterLogicalOrExpr(CompiscriptParser.LogicalOrExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#logicalOrExpr}.
	 * @param ctx the parse tree
	 */
	void exitLogicalOrExpr(CompiscriptParser.LogicalOrExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#logicalAndExpr}.
	 * @param ctx the parse tree
	 */
	void enterLogicalAndExpr(CompiscriptParser.LogicalAndExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#logicalAndExpr}.
	 * @param ctx the parse tree
	 */
	void exitLogicalAndExpr(CompiscriptParser.LogicalAndExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#equalityExpr}.
	 * @param ctx the parse tree
	 */
	void enterEqualityExpr(CompiscriptParser.EqualityExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#equalityExpr}.
	 * @param ctx the parse tree
	 */
	void exitEqualityExpr(CompiscriptParser.EqualityExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#relationalExpr}.
	 * @param ctx the parse tree
	 */
	void enterRelationalExpr(CompiscriptParser.RelationalExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#relationalExpr}.
	 * @param ctx the parse tree
	 */
	void exitRelationalExpr(CompiscriptParser.RelationalExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#additiveExpr}.
	 * @param ctx the parse tree
	 */
	void enterAdditiveExpr(CompiscriptParser.AdditiveExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#additiveExpr}.
	 * @param ctx the parse tree
	 */
	void exitAdditiveExpr(CompiscriptParser.AdditiveExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#multiplicativeExpr}.
	 * @param ctx the parse tree
	 */
	void enterMultiplicativeExpr(CompiscriptParser.MultiplicativeExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#multiplicativeExpr}.
	 * @param ctx the parse tree
	 */
	void exitMultiplicativeExpr(CompiscriptParser.MultiplicativeExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#unaryExpr}.
	 * @param ctx the parse tree
	 */
	void enterUnaryExpr(CompiscriptParser.UnaryExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#unaryExpr}.
	 * @param ctx the parse tree
	 */
	void exitUnaryExpr(CompiscriptParser.UnaryExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#primaryExpr}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpr(CompiscriptParser.PrimaryExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#primaryExpr}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpr(CompiscriptParser.PrimaryExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#literalExpr}.
	 * @param ctx the parse tree
	 */
	void enterLiteralExpr(CompiscriptParser.LiteralExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#literalExpr}.
	 * @param ctx the parse tree
	 */
	void exitLiteralExpr(CompiscriptParser.LiteralExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#leftHandSide}.
	 * @param ctx the parse tree
	 */
	void enterLeftHandSide(CompiscriptParser.LeftHandSideContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#leftHandSide}.
	 * @param ctx the parse tree
	 */
	void exitLeftHandSide(CompiscriptParser.LeftHandSideContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IdentifierExpr}
	 * labeled alternative in {@link CompiscriptParser#primaryAtom}.
	 * @param ctx the parse tree
	 */
	void enterIdentifierExpr(CompiscriptParser.IdentifierExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IdentifierExpr}
	 * labeled alternative in {@link CompiscriptParser#primaryAtom}.
	 * @param ctx the parse tree
	 */
	void exitIdentifierExpr(CompiscriptParser.IdentifierExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code NewExpr}
	 * labeled alternative in {@link CompiscriptParser#primaryAtom}.
	 * @param ctx the parse tree
	 */
	void enterNewExpr(CompiscriptParser.NewExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code NewExpr}
	 * labeled alternative in {@link CompiscriptParser#primaryAtom}.
	 * @param ctx the parse tree
	 */
	void exitNewExpr(CompiscriptParser.NewExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ThisExpr}
	 * labeled alternative in {@link CompiscriptParser#primaryAtom}.
	 * @param ctx the parse tree
	 */
	void enterThisExpr(CompiscriptParser.ThisExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ThisExpr}
	 * labeled alternative in {@link CompiscriptParser#primaryAtom}.
	 * @param ctx the parse tree
	 */
	void exitThisExpr(CompiscriptParser.ThisExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CallExpr}
	 * labeled alternative in {@link CompiscriptParser#suffixOp}.
	 * @param ctx the parse tree
	 */
	void enterCallExpr(CompiscriptParser.CallExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CallExpr}
	 * labeled alternative in {@link CompiscriptParser#suffixOp}.
	 * @param ctx the parse tree
	 */
	void exitCallExpr(CompiscriptParser.CallExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IndexExpr}
	 * labeled alternative in {@link CompiscriptParser#suffixOp}.
	 * @param ctx the parse tree
	 */
	void enterIndexExpr(CompiscriptParser.IndexExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IndexExpr}
	 * labeled alternative in {@link CompiscriptParser#suffixOp}.
	 * @param ctx the parse tree
	 */
	void exitIndexExpr(CompiscriptParser.IndexExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code PropertyAccessExpr}
	 * labeled alternative in {@link CompiscriptParser#suffixOp}.
	 * @param ctx the parse tree
	 */
	void enterPropertyAccessExpr(CompiscriptParser.PropertyAccessExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code PropertyAccessExpr}
	 * labeled alternative in {@link CompiscriptParser#suffixOp}.
	 * @param ctx the parse tree
	 */
	void exitPropertyAccessExpr(CompiscriptParser.PropertyAccessExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#arguments}.
	 * @param ctx the parse tree
	 */
	void enterArguments(CompiscriptParser.ArgumentsContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#arguments}.
	 * @param ctx the parse tree
	 */
	void exitArguments(CompiscriptParser.ArgumentsContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#arrayLiteral}.
	 * @param ctx the parse tree
	 */
	void enterArrayLiteral(CompiscriptParser.ArrayLiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#arrayLiteral}.
	 * @param ctx the parse tree
	 */
	void exitArrayLiteral(CompiscriptParser.ArrayLiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#type}.
	 * @param ctx the parse tree
	 */
	void enterType(CompiscriptParser.TypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#type}.
	 * @param ctx the parse tree
	 */
	void exitType(CompiscriptParser.TypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link CompiscriptParser#baseType}.
	 * @param ctx the parse tree
	 */
	void enterBaseType(CompiscriptParser.BaseTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link CompiscriptParser#baseType}.
	 * @param ctx the parse tree
	 */
	void exitBaseType(CompiscriptParser.BaseTypeContext ctx);
}