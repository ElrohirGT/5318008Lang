grammar Compiscript;

// ------------------
// Parser Rules
// ------------------

program: statement* EOF;

statement
  : standaloneExpresion // DONE: Prince
  | variableDeclaration // DONE: Flavio
  | constantDeclaration // DONE: Flavio
  | assignment          // DONE: Flavio
  | functionDeclaration // DONE: Flavio, Prince
												// DONE: Flavio: Arrays
												// DONE: Flavio: Duplicate variable/constant declarations
  | classDeclaration    // DONE: Flavio
  | printStatement      // DONE: Rayo, Primera (y talvez única) builtin del compilador
  | blockStatement      // DONE: Rayo, tiene que crear un nuevo scope.
  | ifStatement         // DONE: Rayo
  | whileStatement      // DONE: Rayo
  | doWhileStatement    // DONE: Rayo
  | forStatement        // DONE: Rayo
  | foreachStatement    // TODO: Rayo
  | tryStatement        // DONE: Rayo
  | switchStatement     // DONE: Rayo
  | breakStatement      // DONE: Rayo
  | continueStatement   // DONE: Rayo
  | returnStatement     // TODO: Prince, Rayo
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
	: thisAssignment
	| variableAssignment
  ;

thisAssignment: 'this' ('.' Identifier) assignmentPart* '=' conditionalExpr ';' ;
variableAssignment: Identifier assignmentPart* '=' conditionalExpr ';' ;

assignmentPart
  : '.' Identifier
  | '[' conditionalExpr ']'
  ;

// expressionStatement: expression ';'; // Standalone expresions are not allowed
printStatement: 'print' '(' conditionalExpr ')' ';';

mustBoolExpr : conditionalExpr ;

ifStatement: 'if' '(' mustBoolExpr ')' ifBody ('else' elseBody)?;
ifBody: block;
elseBody: block;
whileStatement: 'while' '(' mustBoolExpr ')' whileBody;
whileBody: block;
doWhileStatement: 'do' doWhileBody 'while' '(' mustBoolExpr ')' ';';
doWhileBody: block;
forStatement: 'for' '(' (variableDeclaration | assignment | ';') mustBoolExpr ';' assignment? ')' block;
foreachValue: Identifier 'in' conditionalExpr;
foreachStatement: 'foreach' '(' foreachValue ')' block;
breakStatement: 'break' ';';
continueStatement: 'continue' ';';
returnStatement: 'return' conditionalExpr? ';';
blockStatement: block;

tryStatement: 'try' block catchStatement;
catchStatement : 'catch' '(' Identifier ')' block;

switchValue : conditionalExpr;
caseValue : primaryExpr;
switchStatement: 'switch' '(' switchValue ')' '{' switchCase* defaultCase? '}';
switchCase: 'case' caseValue ':' caseBody;
defaultCase: 'default' ':' caseBody;
caseBody: statement*;

functionDeclaration: 'function' Identifier '(' parameters? ')' (':' type)? block;
parameters: parameter (',' parameter)*;
parameter: Identifier (':' type)?;

classDeclaration: 'class' Identifier (':' Identifier)? '{' classMember* '}';
classMember: functionDeclaration | variableDeclaration | constantDeclaration;

// ------------------
// Expression Rules — Operator Precedence
// ------------------

expression: conditionalExpr;

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
  : relationalExpr ( ('==' | '!=') relationalExpr )?
  ;

relationalExpr
  : additiveExpr ( ('<' | '<=' | '>' | '>=') additiveExpr )?
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

standaloneExpresion
  : standaloneAtom (suffixOp)* ';'
  ;

primaryAtom
  : Identifier                                 # IdentifierExpr
  | 'new' Identifier '(' arguments? ')'        # NewExpr
  | 'this'                                     # ThisExpr
  ;

standaloneAtom
  : Identifier                                 # StandaloneIdentifierExpr
  | 'new' Identifier '(' arguments? ')'        # StandaloneNewExpr        // Add custom error for this case
  | 'this'                                     # StandaloneThisExpr
  ;

suffixOp
  : '.' Identifier '(' arguments? ')'         # MethodCallExpr 
  | '(' arguments? ')'                        # CallExpr
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
