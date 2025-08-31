grammar Compiscript;

// ------------------
// Parser Rules
// ------------------

program: statement* EOF;

statement
  : variableDeclaration
  | constantDeclaration
  | assignment
  | functionDeclaration
  | classDeclaration
  | printStatement
  | blockStatement
  | ifStatement
  | whileStatement
  | doWhileStatement
  | forStatement
  | foreachStatement
  | tryCatchStatement
  | switchStatement
  | breakStatement
  | continueStatement
  | returnStatement
  ;

block: '{' statement* '}';

variableDeclaration
  : ('let' | 'var') Identifier typeAnnotation? initializer? ';'
  ;

constantDeclaration
  : 'const' Identifier typeAnnotation? '=' conditionalExpr ';'
  ;

typeAnnotation: ':' type;
initializer: '=' conditionalExpr;

assignment
  : Identifier '=' conditionalExpr ';'
  // TODO: First expresion should just allow identifiers something like
  //    identifier ('.' Identifier)+ = conditionalExpr
  // Might work, not tested yet.
  | expression '.' Identifier '=' expression ';'
  ;

// expressionStatement: expression ';'; // Standalone expresions are not allowed
printStatement: 'print' '(' conditionalExpr ')' ';';

ifStatement: 'if' '(' conditionalExpr ')' block ('else' block)?;
whileStatement: 'while' '(' conditionalExpr ')' block;
doWhileStatement: 'do' block 'while' '(' conditionalExpr ')' ';';
forStatement: 'for' '(' (variableDeclaration | assignment | ';') conditionalExpr? ';' expression? ')' block;
foreachStatement: 'foreach' '(' Identifier 'in' conditionalExpr ')' block;
breakStatement: 'break' ';';
continueStatement: 'continue' ';';
returnStatement: 'return' conditionalExpr? ';';
blockStatement: block;

tryCatchStatement: 'try' block catchStatement;
catchStatement : 'catch' '(' Identifier ')' block;


switchStatement: 'switch' '(' conditionalExpr ')' '{' switchCase* defaultCase? '}';
switchCase: 'case' conditionalExpr ':' statement*;
defaultCase: 'default' ':' statement*;

functionDeclaration: 'function' Identifier '(' parameters? ')' (':' type)? block;
parameters: parameter (',' parameter)*;
parameter: Identifier (':' type)?;

classDeclaration: 'class' Identifier (':' Identifier)? '{' classMember* '}';
classMember: functionDeclaration | variableDeclaration | constantDeclaration;

// ------------------
// Expression Rules — Operator Precedence
// ------------------

expression: assignmentExpr;

assignmentExpr
  : lhs=leftHandSide '=' conditionalExpr
  | lhs=leftHandSide '.' Identifier '=' assignmentExpr
  | conditionalExpr
  ;

// Ternary operators just work for assignments
// Not as standalone expresion, in that case use if-else
conditionalExpr
  : logicalOrExpr ('?' conditionalExpr ':' conditionalExpr)? 
  ;

logicalOrExpr
  : logicalAndExpr ( '||' logicalAndExpr )*
  ;

logicalAndExpr
  : equalityExpr ( '&&' equalityExpr )*
  ;

// Values should be of the same type and just primary
equalityExpr
  : relationalExpr ( ('==' | '!=') relationalExpr )* 
  ;

relationalExpr
  : additiveExpr ( ('<' | '<=' | '>' | '>=') additiveExpr )*
  ;

additiveExpr
  : multiplicativeExpr ( ('+' | '-') multiplicativeExpr )*
  ;

multiplicativeExpr
  : unaryExpr ( ('*' | '/' | '%') unaryExpr )*
  ;

unaryExpr
  : ('-' | '!') unaryExpr
  | primaryExpr
  ;

primaryExpr
  : literalExpr
  | leftHandSide
  | '(' conditionalExpr ')'
  ;

literalExpr
  : Literal
  | arrayLiteral
  | 'null'
  | 'true'
  | 'false'
  ;

leftHandSide
  : primaryAtom (suffixOp)*
  ;

primaryAtom
  : Identifier                                 # IdentifierExpr
  | 'new' Identifier '(' arguments? ')'        # NewExpr
  | 'this'                                     # ThisExpr
  ;

suffixOp
  : '(' arguments? ')'                        # CallExpr
  | '[' conditionalExpr ']'                   # IndexExpr
  | '.' Identifier                            # PropertyAccessExpr
  ;

arguments: conditionalExpr (',' conditionalExpr)*;

arrayLiteral: '[' (conditionalExpr (',' conditionalExpr)*)? ']';

// ------------------
// Types
// ------------------

type: baseType ('[' ']')*;
baseType: 'boolean' | 'integer' | 'string' | Identifier;

// ------------------
// Lexer Rules
// ------------------

Literal
  : IntegerLiteral
  | StringLiteral
  ;

IntegerLiteral: [0-9]+;
StringLiteral: '"' (~["\r\n])* '"';

Identifier: [a-zA-Z_][a-zA-Z0-9_]*;

WS: [ \t\r\n]+ -> skip;
COMMENT: '//' ~[\r\n]* -> skip;
MULTILINE_COMMENT: '/*' .*? '*/' -> skip;
