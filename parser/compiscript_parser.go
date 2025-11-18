// Code generated from compiscript/program/Compiscript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Compiscript

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CompiscriptParser struct {
	*antlr.BaseParser
}

var CompiscriptParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func compiscriptParserInit() {
	staticData := &CompiscriptParserStaticData
	staticData.LiteralNames = []string{
		"", "'{'", "'}'", "'let'", "'var'", "';'", "'const'", "'='", "':'",
		"'this'", "'.'", "'['", "']'", "'if'", "'('", "')'", "'else'", "'while'",
		"'do'", "'for'", "'in'", "'foreach'", "'break'", "'continue'", "'return'",
		"'try'", "'catch'", "'switch'", "'case'", "'default'", "'function'",
		"','", "'class'", "'?'", "'||'", "'&&'", "'=='", "'!='", "'<'", "'<='",
		"'>'", "'>='", "'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'null'", "'true'",
		"'false'", "'new'", "'boolean'", "'integer'", "'string'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "Literal", "IntegerLiteral", "StringLiteral", "Identifier",
		"WS", "COMMENT", "MULTILINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "statement", "block", "variableDeclaration", "constantDeclaration",
		"typeAnnotation", "initializer", "assignment", "thisAssignment", "variableAssignment",
		"assignmentPart", "mustBoolExpr", "ifStatement", "ifCondition", "ifBody",
		"elseBody", "whileStatement", "whileCondition", "whileBody", "doWhileStatement",
		"doWhileBody", "forStatement", "forBody", "forCondition", "forUpdate",
		"foreachValue", "foreachStatement", "breakStatement", "continueStatement",
		"returnStatement", "blockStatement", "exceptionStatement", "tryStatement",
		"catchStatement", "switchValue", "caseValue", "switchStatement", "switchCase",
		"defaultCase", "caseBody", "functionDeclaration", "parameters", "parameter",
		"classDeclaration", "classMember", "expression", "conditionalExpr",
		"logicalOrExpr", "logicalAndExpr", "equalityExpr", "relationalExpr",
		"additiveExpr", "multiplicativeExpr", "unaryExpr", "primaryExpr", "literalExpr",
		"leftHandSide", "standaloneExpresion", "primaryAtom", "standaloneAtom",
		"suffixOp", "arguments", "arrayLiteral", "type", "baseType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 61, 565, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 1, 0, 5, 0, 132, 8, 0, 10, 0, 12, 0, 135, 9, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 156, 8, 1, 1, 2, 1, 2,
		5, 2, 160, 8, 2, 10, 2, 12, 2, 163, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		3, 3, 170, 8, 3, 1, 3, 3, 3, 173, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3,
		4, 180, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 3, 7, 194, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 201,
		8, 8, 10, 8, 12, 8, 204, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9,
		212, 8, 9, 10, 9, 12, 9, 215, 9, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 227, 8, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 238, 8, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 271, 8,
		21, 1, 21, 1, 21, 1, 21, 3, 21, 276, 8, 21, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28,
		1, 29, 1, 29, 3, 29, 305, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 5,
		36, 333, 8, 36, 10, 36, 12, 36, 336, 9, 36, 1, 36, 3, 36, 339, 8, 36, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 39, 5, 39, 353, 8, 39, 10, 39, 12, 39, 356, 9, 39, 1, 40, 1, 40, 1,
		40, 1, 40, 3, 40, 362, 8, 40, 1, 40, 1, 40, 1, 40, 3, 40, 367, 8, 40, 1,
		40, 1, 40, 1, 41, 1, 41, 1, 41, 5, 41, 374, 8, 41, 10, 41, 12, 41, 377,
		9, 41, 1, 42, 1, 42, 1, 42, 3, 42, 382, 8, 42, 1, 43, 1, 43, 1, 43, 1,
		43, 3, 43, 388, 8, 43, 1, 43, 1, 43, 5, 43, 392, 8, 43, 10, 43, 12, 43,
		395, 9, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 3, 44, 402, 8, 44, 1, 45,
		1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 3, 46, 412, 8, 46, 1,
		47, 1, 47, 1, 47, 5, 47, 417, 8, 47, 10, 47, 12, 47, 420, 9, 47, 1, 48,
		1, 48, 1, 48, 5, 48, 425, 8, 48, 10, 48, 12, 48, 428, 9, 48, 1, 49, 1,
		49, 1, 49, 3, 49, 433, 8, 49, 1, 50, 1, 50, 1, 50, 3, 50, 438, 8, 50, 1,
		51, 1, 51, 1, 51, 5, 51, 443, 8, 51, 10, 51, 12, 51, 446, 9, 51, 1, 52,
		1, 52, 1, 52, 5, 52, 451, 8, 52, 10, 52, 12, 52, 454, 9, 52, 1, 53, 1,
		53, 1, 53, 3, 53, 459, 8, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54,
		3, 54, 467, 8, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 3, 55, 474, 8, 55,
		1, 56, 1, 56, 5, 56, 478, 8, 56, 10, 56, 12, 56, 481, 9, 56, 1, 57, 1,
		57, 5, 57, 485, 8, 57, 10, 57, 12, 57, 488, 9, 57, 1, 57, 1, 57, 1, 58,
		1, 58, 1, 58, 1, 58, 1, 58, 3, 58, 497, 8, 58, 1, 58, 1, 58, 3, 58, 501,
		8, 58, 1, 59, 1, 59, 1, 59, 1, 59, 1, 59, 3, 59, 508, 8, 59, 1, 59, 1,
		59, 3, 59, 512, 8, 59, 1, 60, 1, 60, 1, 60, 1, 60, 3, 60, 518, 8, 60, 1,
		60, 1, 60, 1, 60, 3, 60, 523, 8, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60,
		1, 60, 1, 60, 3, 60, 532, 8, 60, 1, 61, 1, 61, 1, 61, 5, 61, 537, 8, 61,
		10, 61, 12, 61, 540, 9, 61, 1, 62, 1, 62, 1, 62, 1, 62, 5, 62, 546, 8,
		62, 10, 62, 12, 62, 549, 9, 62, 3, 62, 551, 8, 62, 1, 62, 1, 62, 1, 63,
		1, 63, 1, 63, 5, 63, 558, 8, 63, 10, 63, 12, 63, 561, 9, 63, 1, 64, 1,
		64, 1, 64, 0, 0, 65, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
		64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98,
		100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128,
		0, 7, 1, 0, 3, 4, 1, 0, 36, 37, 1, 0, 38, 41, 1, 0, 42, 43, 1, 0, 44, 46,
		2, 0, 43, 43, 47, 47, 2, 0, 52, 54, 58, 58, 571, 0, 133, 1, 0, 0, 0, 2,
		155, 1, 0, 0, 0, 4, 157, 1, 0, 0, 0, 6, 166, 1, 0, 0, 0, 8, 176, 1, 0,
		0, 0, 10, 185, 1, 0, 0, 0, 12, 188, 1, 0, 0, 0, 14, 193, 1, 0, 0, 0, 16,
		195, 1, 0, 0, 0, 18, 209, 1, 0, 0, 0, 20, 226, 1, 0, 0, 0, 22, 228, 1,
		0, 0, 0, 24, 230, 1, 0, 0, 0, 26, 239, 1, 0, 0, 0, 28, 241, 1, 0, 0, 0,
		30, 243, 1, 0, 0, 0, 32, 245, 1, 0, 0, 0, 34, 251, 1, 0, 0, 0, 36, 253,
		1, 0, 0, 0, 38, 255, 1, 0, 0, 0, 40, 263, 1, 0, 0, 0, 42, 265, 1, 0, 0,
		0, 44, 280, 1, 0, 0, 0, 46, 282, 1, 0, 0, 0, 48, 284, 1, 0, 0, 0, 50, 286,
		1, 0, 0, 0, 52, 290, 1, 0, 0, 0, 54, 296, 1, 0, 0, 0, 56, 299, 1, 0, 0,
		0, 58, 302, 1, 0, 0, 0, 60, 308, 1, 0, 0, 0, 62, 310, 1, 0, 0, 0, 64, 313,
		1, 0, 0, 0, 66, 316, 1, 0, 0, 0, 68, 322, 1, 0, 0, 0, 70, 324, 1, 0, 0,
		0, 72, 326, 1, 0, 0, 0, 74, 342, 1, 0, 0, 0, 76, 347, 1, 0, 0, 0, 78, 354,
		1, 0, 0, 0, 80, 357, 1, 0, 0, 0, 82, 370, 1, 0, 0, 0, 84, 378, 1, 0, 0,
		0, 86, 383, 1, 0, 0, 0, 88, 401, 1, 0, 0, 0, 90, 403, 1, 0, 0, 0, 92, 405,
		1, 0, 0, 0, 94, 413, 1, 0, 0, 0, 96, 421, 1, 0, 0, 0, 98, 429, 1, 0, 0,
		0, 100, 434, 1, 0, 0, 0, 102, 439, 1, 0, 0, 0, 104, 447, 1, 0, 0, 0, 106,
		458, 1, 0, 0, 0, 108, 466, 1, 0, 0, 0, 110, 473, 1, 0, 0, 0, 112, 475,
		1, 0, 0, 0, 114, 482, 1, 0, 0, 0, 116, 500, 1, 0, 0, 0, 118, 511, 1, 0,
		0, 0, 120, 531, 1, 0, 0, 0, 122, 533, 1, 0, 0, 0, 124, 541, 1, 0, 0, 0,
		126, 554, 1, 0, 0, 0, 128, 562, 1, 0, 0, 0, 130, 132, 3, 2, 1, 0, 131,
		130, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134,
		1, 0, 0, 0, 134, 136, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 137, 5, 0,
		0, 1, 137, 1, 1, 0, 0, 0, 138, 156, 3, 114, 57, 0, 139, 156, 3, 6, 3, 0,
		140, 156, 3, 8, 4, 0, 141, 156, 3, 14, 7, 0, 142, 156, 3, 80, 40, 0, 143,
		156, 3, 86, 43, 0, 144, 156, 3, 60, 30, 0, 145, 156, 3, 24, 12, 0, 146,
		156, 3, 32, 16, 0, 147, 156, 3, 38, 19, 0, 148, 156, 3, 42, 21, 0, 149,
		156, 3, 52, 26, 0, 150, 156, 3, 62, 31, 0, 151, 156, 3, 72, 36, 0, 152,
		156, 3, 54, 27, 0, 153, 156, 3, 56, 28, 0, 154, 156, 3, 58, 29, 0, 155,
		138, 1, 0, 0, 0, 155, 139, 1, 0, 0, 0, 155, 140, 1, 0, 0, 0, 155, 141,
		1, 0, 0, 0, 155, 142, 1, 0, 0, 0, 155, 143, 1, 0, 0, 0, 155, 144, 1, 0,
		0, 0, 155, 145, 1, 0, 0, 0, 155, 146, 1, 0, 0, 0, 155, 147, 1, 0, 0, 0,
		155, 148, 1, 0, 0, 0, 155, 149, 1, 0, 0, 0, 155, 150, 1, 0, 0, 0, 155,
		151, 1, 0, 0, 0, 155, 152, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 154,
		1, 0, 0, 0, 156, 3, 1, 0, 0, 0, 157, 161, 5, 1, 0, 0, 158, 160, 3, 2, 1,
		0, 159, 158, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161,
		162, 1, 0, 0, 0, 162, 164, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 165,
		5, 2, 0, 0, 165, 5, 1, 0, 0, 0, 166, 167, 7, 0, 0, 0, 167, 169, 5, 58,
		0, 0, 168, 170, 3, 10, 5, 0, 169, 168, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0,
		170, 172, 1, 0, 0, 0, 171, 173, 3, 12, 6, 0, 172, 171, 1, 0, 0, 0, 172,
		173, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 5, 5, 0, 0, 175, 7, 1,
		0, 0, 0, 176, 177, 5, 6, 0, 0, 177, 179, 5, 58, 0, 0, 178, 180, 3, 10,
		5, 0, 179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0,
		181, 182, 5, 7, 0, 0, 182, 183, 3, 92, 46, 0, 183, 184, 5, 5, 0, 0, 184,
		9, 1, 0, 0, 0, 185, 186, 5, 8, 0, 0, 186, 187, 3, 126, 63, 0, 187, 11,
		1, 0, 0, 0, 188, 189, 5, 7, 0, 0, 189, 190, 3, 92, 46, 0, 190, 13, 1, 0,
		0, 0, 191, 194, 3, 16, 8, 0, 192, 194, 3, 18, 9, 0, 193, 191, 1, 0, 0,
		0, 193, 192, 1, 0, 0, 0, 194, 15, 1, 0, 0, 0, 195, 196, 5, 9, 0, 0, 196,
		197, 5, 10, 0, 0, 197, 198, 5, 58, 0, 0, 198, 202, 1, 0, 0, 0, 199, 201,
		3, 20, 10, 0, 200, 199, 1, 0, 0, 0, 201, 204, 1, 0, 0, 0, 202, 200, 1,
		0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 205, 1, 0, 0, 0, 204, 202, 1, 0, 0,
		0, 205, 206, 5, 7, 0, 0, 206, 207, 3, 92, 46, 0, 207, 208, 5, 5, 0, 0,
		208, 17, 1, 0, 0, 0, 209, 213, 5, 58, 0, 0, 210, 212, 3, 20, 10, 0, 211,
		210, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214,
		1, 0, 0, 0, 214, 216, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 217, 5, 7,
		0, 0, 217, 218, 3, 92, 46, 0, 218, 219, 5, 5, 0, 0, 219, 19, 1, 0, 0, 0,
		220, 221, 5, 10, 0, 0, 221, 227, 5, 58, 0, 0, 222, 223, 5, 11, 0, 0, 223,
		224, 3, 92, 46, 0, 224, 225, 5, 12, 0, 0, 225, 227, 1, 0, 0, 0, 226, 220,
		1, 0, 0, 0, 226, 222, 1, 0, 0, 0, 227, 21, 1, 0, 0, 0, 228, 229, 3, 92,
		46, 0, 229, 23, 1, 0, 0, 0, 230, 231, 5, 13, 0, 0, 231, 232, 5, 14, 0,
		0, 232, 233, 3, 26, 13, 0, 233, 234, 5, 15, 0, 0, 234, 237, 3, 28, 14,
		0, 235, 236, 5, 16, 0, 0, 236, 238, 3, 30, 15, 0, 237, 235, 1, 0, 0, 0,
		237, 238, 1, 0, 0, 0, 238, 25, 1, 0, 0, 0, 239, 240, 3, 22, 11, 0, 240,
		27, 1, 0, 0, 0, 241, 242, 3, 4, 2, 0, 242, 29, 1, 0, 0, 0, 243, 244, 3,
		4, 2, 0, 244, 31, 1, 0, 0, 0, 245, 246, 5, 17, 0, 0, 246, 247, 5, 14, 0,
		0, 247, 248, 3, 34, 17, 0, 248, 249, 5, 15, 0, 0, 249, 250, 3, 36, 18,
		0, 250, 33, 1, 0, 0, 0, 251, 252, 3, 22, 11, 0, 252, 35, 1, 0, 0, 0, 253,
		254, 3, 4, 2, 0, 254, 37, 1, 0, 0, 0, 255, 256, 5, 18, 0, 0, 256, 257,
		3, 40, 20, 0, 257, 258, 5, 17, 0, 0, 258, 259, 5, 14, 0, 0, 259, 260, 3,
		34, 17, 0, 260, 261, 5, 15, 0, 0, 261, 262, 5, 5, 0, 0, 262, 39, 1, 0,
		0, 0, 263, 264, 3, 4, 2, 0, 264, 41, 1, 0, 0, 0, 265, 266, 5, 19, 0, 0,
		266, 270, 5, 14, 0, 0, 267, 271, 3, 6, 3, 0, 268, 271, 3, 14, 7, 0, 269,
		271, 5, 5, 0, 0, 270, 267, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 270, 269,
		1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 273, 3, 46, 23, 0, 273, 275, 5,
		5, 0, 0, 274, 276, 3, 48, 24, 0, 275, 274, 1, 0, 0, 0, 275, 276, 1, 0,
		0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 5, 15, 0, 0, 278, 279, 3, 44, 22,
		0, 279, 43, 1, 0, 0, 0, 280, 281, 3, 4, 2, 0, 281, 45, 1, 0, 0, 0, 282,
		283, 3, 22, 11, 0, 283, 47, 1, 0, 0, 0, 284, 285, 3, 14, 7, 0, 285, 49,
		1, 0, 0, 0, 286, 287, 5, 58, 0, 0, 287, 288, 5, 20, 0, 0, 288, 289, 3,
		92, 46, 0, 289, 51, 1, 0, 0, 0, 290, 291, 5, 21, 0, 0, 291, 292, 5, 14,
		0, 0, 292, 293, 3, 50, 25, 0, 293, 294, 5, 15, 0, 0, 294, 295, 3, 4, 2,
		0, 295, 53, 1, 0, 0, 0, 296, 297, 5, 22, 0, 0, 297, 298, 5, 5, 0, 0, 298,
		55, 1, 0, 0, 0, 299, 300, 5, 23, 0, 0, 300, 301, 5, 5, 0, 0, 301, 57, 1,
		0, 0, 0, 302, 304, 5, 24, 0, 0, 303, 305, 3, 92, 46, 0, 304, 303, 1, 0,
		0, 0, 304, 305, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 307, 5, 5, 0, 0,
		307, 59, 1, 0, 0, 0, 308, 309, 3, 4, 2, 0, 309, 61, 1, 0, 0, 0, 310, 311,
		3, 64, 32, 0, 311, 312, 3, 66, 33, 0, 312, 63, 1, 0, 0, 0, 313, 314, 5,
		25, 0, 0, 314, 315, 3, 4, 2, 0, 315, 65, 1, 0, 0, 0, 316, 317, 5, 26, 0,
		0, 317, 318, 5, 14, 0, 0, 318, 319, 5, 58, 0, 0, 319, 320, 5, 15, 0, 0,
		320, 321, 3, 4, 2, 0, 321, 67, 1, 0, 0, 0, 322, 323, 3, 92, 46, 0, 323,
		69, 1, 0, 0, 0, 324, 325, 3, 108, 54, 0, 325, 71, 1, 0, 0, 0, 326, 327,
		5, 27, 0, 0, 327, 328, 5, 14, 0, 0, 328, 329, 3, 68, 34, 0, 329, 330, 5,
		15, 0, 0, 330, 334, 5, 1, 0, 0, 331, 333, 3, 74, 37, 0, 332, 331, 1, 0,
		0, 0, 333, 336, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0,
		335, 338, 1, 0, 0, 0, 336, 334, 1, 0, 0, 0, 337, 339, 3, 76, 38, 0, 338,
		337, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 341,
		5, 2, 0, 0, 341, 73, 1, 0, 0, 0, 342, 343, 5, 28, 0, 0, 343, 344, 3, 70,
		35, 0, 344, 345, 5, 8, 0, 0, 345, 346, 3, 78, 39, 0, 346, 75, 1, 0, 0,
		0, 347, 348, 5, 29, 0, 0, 348, 349, 5, 8, 0, 0, 349, 350, 3, 78, 39, 0,
		350, 77, 1, 0, 0, 0, 351, 353, 3, 2, 1, 0, 352, 351, 1, 0, 0, 0, 353, 356,
		1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 79, 1, 0,
		0, 0, 356, 354, 1, 0, 0, 0, 357, 358, 5, 30, 0, 0, 358, 359, 5, 58, 0,
		0, 359, 361, 5, 14, 0, 0, 360, 362, 3, 82, 41, 0, 361, 360, 1, 0, 0, 0,
		361, 362, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 366, 5, 15, 0, 0, 364,
		365, 5, 8, 0, 0, 365, 367, 3, 126, 63, 0, 366, 364, 1, 0, 0, 0, 366, 367,
		1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 369, 3, 4, 2, 0, 369, 81, 1, 0,
		0, 0, 370, 375, 3, 84, 42, 0, 371, 372, 5, 31, 0, 0, 372, 374, 3, 84, 42,
		0, 373, 371, 1, 0, 0, 0, 374, 377, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375,
		376, 1, 0, 0, 0, 376, 83, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 378, 381, 5,
		58, 0, 0, 379, 380, 5, 8, 0, 0, 380, 382, 3, 126, 63, 0, 381, 379, 1, 0,
		0, 0, 381, 382, 1, 0, 0, 0, 382, 85, 1, 0, 0, 0, 383, 384, 5, 32, 0, 0,
		384, 387, 5, 58, 0, 0, 385, 386, 5, 8, 0, 0, 386, 388, 5, 58, 0, 0, 387,
		385, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 393,
		5, 1, 0, 0, 390, 392, 3, 88, 44, 0, 391, 390, 1, 0, 0, 0, 392, 395, 1,
		0, 0, 0, 393, 391, 1, 0, 0, 0, 393, 394, 1, 0, 0, 0, 394, 396, 1, 0, 0,
		0, 395, 393, 1, 0, 0, 0, 396, 397, 5, 2, 0, 0, 397, 87, 1, 0, 0, 0, 398,
		402, 3, 80, 40, 0, 399, 402, 3, 6, 3, 0, 400, 402, 3, 8, 4, 0, 401, 398,
		1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 401, 400, 1, 0, 0, 0, 402, 89, 1, 0,
		0, 0, 403, 404, 3, 92, 46, 0, 404, 91, 1, 0, 0, 0, 405, 411, 3, 94, 47,
		0, 406, 407, 5, 33, 0, 0, 407, 408, 3, 92, 46, 0, 408, 409, 5, 8, 0, 0,
		409, 410, 3, 92, 46, 0, 410, 412, 1, 0, 0, 0, 411, 406, 1, 0, 0, 0, 411,
		412, 1, 0, 0, 0, 412, 93, 1, 0, 0, 0, 413, 418, 3, 96, 48, 0, 414, 415,
		5, 34, 0, 0, 415, 417, 3, 96, 48, 0, 416, 414, 1, 0, 0, 0, 417, 420, 1,
		0, 0, 0, 418, 416, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 95, 1, 0, 0,
		0, 420, 418, 1, 0, 0, 0, 421, 426, 3, 98, 49, 0, 422, 423, 5, 35, 0, 0,
		423, 425, 3, 98, 49, 0, 424, 422, 1, 0, 0, 0, 425, 428, 1, 0, 0, 0, 426,
		424, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 97, 1, 0, 0, 0, 428, 426, 1,
		0, 0, 0, 429, 432, 3, 100, 50, 0, 430, 431, 7, 1, 0, 0, 431, 433, 3, 100,
		50, 0, 432, 430, 1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 99, 1, 0, 0, 0,
		434, 437, 3, 102, 51, 0, 435, 436, 7, 2, 0, 0, 436, 438, 3, 102, 51, 0,
		437, 435, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 101, 1, 0, 0, 0, 439,
		444, 3, 104, 52, 0, 440, 441, 7, 3, 0, 0, 441, 443, 3, 104, 52, 0, 442,
		440, 1, 0, 0, 0, 443, 446, 1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 444, 445,
		1, 0, 0, 0, 445, 103, 1, 0, 0, 0, 446, 444, 1, 0, 0, 0, 447, 452, 3, 106,
		53, 0, 448, 449, 7, 4, 0, 0, 449, 451, 3, 106, 53, 0, 450, 448, 1, 0, 0,
		0, 451, 454, 1, 0, 0, 0, 452, 450, 1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 453,
		105, 1, 0, 0, 0, 454, 452, 1, 0, 0, 0, 455, 456, 7, 5, 0, 0, 456, 459,
		3, 106, 53, 0, 457, 459, 3, 108, 54, 0, 458, 455, 1, 0, 0, 0, 458, 457,
		1, 0, 0, 0, 459, 107, 1, 0, 0, 0, 460, 467, 3, 110, 55, 0, 461, 467, 3,
		112, 56, 0, 462, 463, 5, 14, 0, 0, 463, 464, 3, 92, 46, 0, 464, 465, 5,
		15, 0, 0, 465, 467, 1, 0, 0, 0, 466, 460, 1, 0, 0, 0, 466, 461, 1, 0, 0,
		0, 466, 462, 1, 0, 0, 0, 467, 109, 1, 0, 0, 0, 468, 474, 5, 55, 0, 0, 469,
		474, 3, 124, 62, 0, 470, 474, 5, 48, 0, 0, 471, 474, 5, 49, 0, 0, 472,
		474, 5, 50, 0, 0, 473, 468, 1, 0, 0, 0, 473, 469, 1, 0, 0, 0, 473, 470,
		1, 0, 0, 0, 473, 471, 1, 0, 0, 0, 473, 472, 1, 0, 0, 0, 474, 111, 1, 0,
		0, 0, 475, 479, 3, 116, 58, 0, 476, 478, 3, 120, 60, 0, 477, 476, 1, 0,
		0, 0, 478, 481, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 479, 480, 1, 0, 0, 0,
		480, 113, 1, 0, 0, 0, 481, 479, 1, 0, 0, 0, 482, 486, 3, 118, 59, 0, 483,
		485, 3, 120, 60, 0, 484, 483, 1, 0, 0, 0, 485, 488, 1, 0, 0, 0, 486, 484,
		1, 0, 0, 0, 486, 487, 1, 0, 0, 0, 487, 489, 1, 0, 0, 0, 488, 486, 1, 0,
		0, 0, 489, 490, 5, 5, 0, 0, 490, 115, 1, 0, 0, 0, 491, 501, 5, 58, 0, 0,
		492, 493, 5, 51, 0, 0, 493, 494, 5, 58, 0, 0, 494, 496, 5, 14, 0, 0, 495,
		497, 3, 122, 61, 0, 496, 495, 1, 0, 0, 0, 496, 497, 1, 0, 0, 0, 497, 498,
		1, 0, 0, 0, 498, 501, 5, 15, 0, 0, 499, 501, 5, 9, 0, 0, 500, 491, 1, 0,
		0, 0, 500, 492, 1, 0, 0, 0, 500, 499, 1, 0, 0, 0, 501, 117, 1, 0, 0, 0,
		502, 512, 5, 58, 0, 0, 503, 504, 5, 51, 0, 0, 504, 505, 5, 58, 0, 0, 505,
		507, 5, 14, 0, 0, 506, 508, 3, 122, 61, 0, 507, 506, 1, 0, 0, 0, 507, 508,
		1, 0, 0, 0, 508, 509, 1, 0, 0, 0, 509, 512, 5, 15, 0, 0, 510, 512, 5, 9,
		0, 0, 511, 502, 1, 0, 0, 0, 511, 503, 1, 0, 0, 0, 511, 510, 1, 0, 0, 0,
		512, 119, 1, 0, 0, 0, 513, 514, 5, 10, 0, 0, 514, 515, 5, 58, 0, 0, 515,
		517, 5, 14, 0, 0, 516, 518, 3, 122, 61, 0, 517, 516, 1, 0, 0, 0, 517, 518,
		1, 0, 0, 0, 518, 519, 1, 0, 0, 0, 519, 532, 5, 15, 0, 0, 520, 522, 5, 14,
		0, 0, 521, 523, 3, 122, 61, 0, 522, 521, 1, 0, 0, 0, 522, 523, 1, 0, 0,
		0, 523, 524, 1, 0, 0, 0, 524, 532, 5, 15, 0, 0, 525, 526, 5, 11, 0, 0,
		526, 527, 3, 92, 46, 0, 527, 528, 5, 12, 0, 0, 528, 532, 1, 0, 0, 0, 529,
		530, 5, 10, 0, 0, 530, 532, 5, 58, 0, 0, 531, 513, 1, 0, 0, 0, 531, 520,
		1, 0, 0, 0, 531, 525, 1, 0, 0, 0, 531, 529, 1, 0, 0, 0, 532, 121, 1, 0,
		0, 0, 533, 538, 3, 92, 46, 0, 534, 535, 5, 31, 0, 0, 535, 537, 3, 92, 46,
		0, 536, 534, 1, 0, 0, 0, 537, 540, 1, 0, 0, 0, 538, 536, 1, 0, 0, 0, 538,
		539, 1, 0, 0, 0, 539, 123, 1, 0, 0, 0, 540, 538, 1, 0, 0, 0, 541, 550,
		5, 11, 0, 0, 542, 547, 3, 92, 46, 0, 543, 544, 5, 31, 0, 0, 544, 546, 3,
		92, 46, 0, 545, 543, 1, 0, 0, 0, 546, 549, 1, 0, 0, 0, 547, 545, 1, 0,
		0, 0, 547, 548, 1, 0, 0, 0, 548, 551, 1, 0, 0, 0, 549, 547, 1, 0, 0, 0,
		550, 542, 1, 0, 0, 0, 550, 551, 1, 0, 0, 0, 551, 552, 1, 0, 0, 0, 552,
		553, 5, 12, 0, 0, 553, 125, 1, 0, 0, 0, 554, 559, 3, 128, 64, 0, 555, 556,
		5, 11, 0, 0, 556, 558, 5, 12, 0, 0, 557, 555, 1, 0, 0, 0, 558, 561, 1,
		0, 0, 0, 559, 557, 1, 0, 0, 0, 559, 560, 1, 0, 0, 0, 560, 127, 1, 0, 0,
		0, 561, 559, 1, 0, 0, 0, 562, 563, 7, 6, 0, 0, 563, 129, 1, 0, 0, 0, 47,
		133, 155, 161, 169, 172, 179, 193, 202, 213, 226, 237, 270, 275, 304, 334,
		338, 354, 361, 366, 375, 381, 387, 393, 401, 411, 418, 426, 432, 437, 444,
		452, 458, 466, 473, 479, 486, 496, 500, 507, 511, 517, 522, 531, 538, 547,
		550, 559,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CompiscriptParserInit initializes any static state used to implement CompiscriptParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCompiscriptParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CompiscriptParserInit() {
	staticData := &CompiscriptParserStaticData
	staticData.once.Do(compiscriptParserInit)
}

// NewCompiscriptParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCompiscriptParser(input antlr.TokenStream) *CompiscriptParser {
	CompiscriptParserInit()
	this := new(CompiscriptParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CompiscriptParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Compiscript.g4"

	return this
}

// CompiscriptParser tokens.
const (
	CompiscriptParserEOF               = antlr.TokenEOF
	CompiscriptParserT__0              = 1
	CompiscriptParserT__1              = 2
	CompiscriptParserT__2              = 3
	CompiscriptParserT__3              = 4
	CompiscriptParserT__4              = 5
	CompiscriptParserT__5              = 6
	CompiscriptParserT__6              = 7
	CompiscriptParserT__7              = 8
	CompiscriptParserT__8              = 9
	CompiscriptParserT__9              = 10
	CompiscriptParserT__10             = 11
	CompiscriptParserT__11             = 12
	CompiscriptParserT__12             = 13
	CompiscriptParserT__13             = 14
	CompiscriptParserT__14             = 15
	CompiscriptParserT__15             = 16
	CompiscriptParserT__16             = 17
	CompiscriptParserT__17             = 18
	CompiscriptParserT__18             = 19
	CompiscriptParserT__19             = 20
	CompiscriptParserT__20             = 21
	CompiscriptParserT__21             = 22
	CompiscriptParserT__22             = 23
	CompiscriptParserT__23             = 24
	CompiscriptParserT__24             = 25
	CompiscriptParserT__25             = 26
	CompiscriptParserT__26             = 27
	CompiscriptParserT__27             = 28
	CompiscriptParserT__28             = 29
	CompiscriptParserT__29             = 30
	CompiscriptParserT__30             = 31
	CompiscriptParserT__31             = 32
	CompiscriptParserT__32             = 33
	CompiscriptParserT__33             = 34
	CompiscriptParserT__34             = 35
	CompiscriptParserT__35             = 36
	CompiscriptParserT__36             = 37
	CompiscriptParserT__37             = 38
	CompiscriptParserT__38             = 39
	CompiscriptParserT__39             = 40
	CompiscriptParserT__40             = 41
	CompiscriptParserT__41             = 42
	CompiscriptParserT__42             = 43
	CompiscriptParserT__43             = 44
	CompiscriptParserT__44             = 45
	CompiscriptParserT__45             = 46
	CompiscriptParserT__46             = 47
	CompiscriptParserT__47             = 48
	CompiscriptParserT__48             = 49
	CompiscriptParserT__49             = 50
	CompiscriptParserT__50             = 51
	CompiscriptParserT__51             = 52
	CompiscriptParserT__52             = 53
	CompiscriptParserT__53             = 54
	CompiscriptParserLiteral           = 55
	CompiscriptParserIntegerLiteral    = 56
	CompiscriptParserStringLiteral     = 57
	CompiscriptParserIdentifier        = 58
	CompiscriptParserWS                = 59
	CompiscriptParserCOMMENT           = 60
	CompiscriptParserMULTILINE_COMMENT = 61
)

// CompiscriptParser rules.
const (
	CompiscriptParserRULE_program             = 0
	CompiscriptParserRULE_statement           = 1
	CompiscriptParserRULE_block               = 2
	CompiscriptParserRULE_variableDeclaration = 3
	CompiscriptParserRULE_constantDeclaration = 4
	CompiscriptParserRULE_typeAnnotation      = 5
	CompiscriptParserRULE_initializer         = 6
	CompiscriptParserRULE_assignment          = 7
	CompiscriptParserRULE_thisAssignment      = 8
	CompiscriptParserRULE_variableAssignment  = 9
	CompiscriptParserRULE_assignmentPart      = 10
	CompiscriptParserRULE_mustBoolExpr        = 11
	CompiscriptParserRULE_ifStatement         = 12
	CompiscriptParserRULE_ifCondition         = 13
	CompiscriptParserRULE_ifBody              = 14
	CompiscriptParserRULE_elseBody            = 15
	CompiscriptParserRULE_whileStatement      = 16
	CompiscriptParserRULE_whileCondition      = 17
	CompiscriptParserRULE_whileBody           = 18
	CompiscriptParserRULE_doWhileStatement    = 19
	CompiscriptParserRULE_doWhileBody         = 20
	CompiscriptParserRULE_forStatement        = 21
	CompiscriptParserRULE_forBody             = 22
	CompiscriptParserRULE_forCondition        = 23
	CompiscriptParserRULE_forUpdate           = 24
	CompiscriptParserRULE_foreachValue        = 25
	CompiscriptParserRULE_foreachStatement    = 26
	CompiscriptParserRULE_breakStatement      = 27
	CompiscriptParserRULE_continueStatement   = 28
	CompiscriptParserRULE_returnStatement     = 29
	CompiscriptParserRULE_blockStatement      = 30
	CompiscriptParserRULE_exceptionStatement  = 31
	CompiscriptParserRULE_tryStatement        = 32
	CompiscriptParserRULE_catchStatement      = 33
	CompiscriptParserRULE_switchValue         = 34
	CompiscriptParserRULE_caseValue           = 35
	CompiscriptParserRULE_switchStatement     = 36
	CompiscriptParserRULE_switchCase          = 37
	CompiscriptParserRULE_defaultCase         = 38
	CompiscriptParserRULE_caseBody            = 39
	CompiscriptParserRULE_functionDeclaration = 40
	CompiscriptParserRULE_parameters          = 41
	CompiscriptParserRULE_parameter           = 42
	CompiscriptParserRULE_classDeclaration    = 43
	CompiscriptParserRULE_classMember         = 44
	CompiscriptParserRULE_expression          = 45
	CompiscriptParserRULE_conditionalExpr     = 46
	CompiscriptParserRULE_logicalOrExpr       = 47
	CompiscriptParserRULE_logicalAndExpr      = 48
	CompiscriptParserRULE_equalityExpr        = 49
	CompiscriptParserRULE_relationalExpr      = 50
	CompiscriptParserRULE_additiveExpr        = 51
	CompiscriptParserRULE_multiplicativeExpr  = 52
	CompiscriptParserRULE_unaryExpr           = 53
	CompiscriptParserRULE_primaryExpr         = 54
	CompiscriptParserRULE_literalExpr         = 55
	CompiscriptParserRULE_leftHandSide        = 56
	CompiscriptParserRULE_standaloneExpresion = 57
	CompiscriptParserRULE_primaryAtom         = 58
	CompiscriptParserRULE_standaloneAtom      = 59
	CompiscriptParserRULE_suffixOp            = 60
	CompiscriptParserRULE_arguments           = 61
	CompiscriptParserRULE_arrayLiteral        = 62
	CompiscriptParserRULE_type                = 63
	CompiscriptParserRULE_baseType            = 64
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *CompiscriptParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CompiscriptParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&290482181534261850) != 0 {
		{
			p.SetState(130)
			p.Statement()
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
		p.Match(CompiscriptParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StandaloneExpresion() IStandaloneExpresionContext
	VariableDeclaration() IVariableDeclarationContext
	ConstantDeclaration() IConstantDeclarationContext
	Assignment() IAssignmentContext
	FunctionDeclaration() IFunctionDeclarationContext
	ClassDeclaration() IClassDeclarationContext
	BlockStatement() IBlockStatementContext
	IfStatement() IIfStatementContext
	WhileStatement() IWhileStatementContext
	DoWhileStatement() IDoWhileStatementContext
	ForStatement() IForStatementContext
	ForeachStatement() IForeachStatementContext
	ExceptionStatement() IExceptionStatementContext
	SwitchStatement() ISwitchStatementContext
	BreakStatement() IBreakStatementContext
	ContinueStatement() IContinueStatementContext
	ReturnStatement() IReturnStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) StandaloneExpresion() IStandaloneExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStandaloneExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStandaloneExpresionContext)
}

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) ConstantDeclaration() IConstantDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantDeclarationContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *StatementContext) ClassDeclaration() IClassDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *StatementContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) DoWhileStatement() IDoWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoWhileStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) ForeachStatement() IForeachStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForeachStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForeachStatementContext)
}

func (s *StatementContext) ExceptionStatement() IExceptionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExceptionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExceptionStatementContext)
}

func (s *StatementContext) SwitchStatement() ISwitchStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) ContinueStatement() IContinueStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *CompiscriptParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CompiscriptParserRULE_statement)
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			p.StandaloneExpresion()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.VariableDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(140)
			p.ConstantDeclaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(141)
			p.Assignment()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(142)
			p.FunctionDeclaration()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(143)
			p.ClassDeclaration()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(144)
			p.BlockStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(145)
			p.IfStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(146)
			p.WhileStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(147)
			p.DoWhileStatement()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(148)
			p.ForStatement()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(149)
			p.ForeachStatement()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(150)
			p.ExceptionStatement()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(151)
			p.SwitchStatement()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(152)
			p.BreakStatement()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(153)
			p.ContinueStatement()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(154)
			p.ReturnStatement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *CompiscriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CompiscriptParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(CompiscriptParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&290482181534261850) != 0 {
		{
			p.SetState(158)
			p.Statement()
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(164)
		p.Match(CompiscriptParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	TypeAnnotation() ITypeAnnotationContext
	Initializer() IInitializerContext

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_variableDeclaration
	return p
}

func InitEmptyVariableDeclarationContext(p *VariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_variableDeclaration
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *VariableDeclarationContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *VariableDeclarationContext) Initializer() IInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitializerContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *CompiscriptParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CompiscriptParserRULE_variableDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CompiscriptParserT__2 || _la == CompiscriptParserT__3) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(167)
		p.Match(CompiscriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CompiscriptParserT__7 {
		{
			p.SetState(168)
			p.TypeAnnotation()
		}

	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CompiscriptParserT__6 {
		{
			p.SetState(171)
			p.Initializer()
		}

	}
	{
		p.SetState(174)
		p.Match(CompiscriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantDeclarationContext is an interface to support dynamic dispatch.
type IConstantDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	ConditionalExpr() IConditionalExprContext
	TypeAnnotation() ITypeAnnotationContext

	// IsConstantDeclarationContext differentiates from other interfaces.
	IsConstantDeclarationContext()
}

type ConstantDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantDeclarationContext() *ConstantDeclarationContext {
	var p = new(ConstantDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_constantDeclaration
	return p
}

func InitEmptyConstantDeclarationContext(p *ConstantDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_constantDeclaration
}

func (*ConstantDeclarationContext) IsConstantDeclarationContext() {}

func NewConstantDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantDeclarationContext {
	var p = new(ConstantDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_constantDeclaration

	return p
}

func (s *ConstantDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *ConstantDeclarationContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *ConstantDeclarationContext) TypeAnnotation() ITypeAnnotationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAnnotationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *ConstantDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterConstantDeclaration(s)
	}
}

func (s *ConstantDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitConstantDeclaration(s)
	}
}

func (p *CompiscriptParser) ConstantDeclaration() (localctx IConstantDeclarationContext) {
	localctx = NewConstantDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CompiscriptParserRULE_constantDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(CompiscriptParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Match(CompiscriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CompiscriptParserT__7 {
		{
			p.SetState(178)
			p.TypeAnnotation()
		}

	}
	{
		p.SetState(181)
		p.Match(CompiscriptParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.ConditionalExpr()
	}
	{
		p.SetState(183)
		p.Match(CompiscriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeAnnotationContext is an interface to support dynamic dispatch.
type ITypeAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsTypeAnnotationContext differentiates from other interfaces.
	IsTypeAnnotationContext()
}

type TypeAnnotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAnnotationContext() *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_typeAnnotation
	return p
}

func InitEmptyTypeAnnotationContext(p *TypeAnnotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_typeAnnotation
}

func (*TypeAnnotationContext) IsTypeAnnotationContext() {}

func NewTypeAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_typeAnnotation

	return p
}

func (s *TypeAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAnnotationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitTypeAnnotation(s)
	}
}

func (p *CompiscriptParser) TypeAnnotation() (localctx ITypeAnnotationContext) {
	localctx = NewTypeAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CompiscriptParserRULE_typeAnnotation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(CompiscriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInitializerContext is an interface to support dynamic dispatch.
type IInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionalExpr() IConditionalExprContext

	// IsInitializerContext differentiates from other interfaces.
	IsInitializerContext()
}

type InitializerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitializerContext() *InitializerContext {
	var p = new(InitializerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_initializer
	return p
}

func InitEmptyInitializerContext(p *InitializerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_initializer
}

func (*InitializerContext) IsInitializerContext() {}

func NewInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitializerContext {
	var p = new(InitializerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_initializer

	return p
}

func (s *InitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *InitializerContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *InitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterInitializer(s)
	}
}

func (s *InitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitInitializer(s)
	}
}

func (p *CompiscriptParser) Initializer() (localctx IInitializerContext) {
	localctx = NewInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CompiscriptParserRULE_initializer)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(CompiscriptParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.ConditionalExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ThisAssignment() IThisAssignmentContext
	VariableAssignment() IVariableAssignmentContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) ThisAssignment() IThisAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThisAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThisAssignmentContext)
}

func (s *AssignmentContext) VariableAssignment() IVariableAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableAssignmentContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *CompiscriptParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CompiscriptParserRULE_assignment)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CompiscriptParserT__8:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)
			p.ThisAssignment()
		}

	case CompiscriptParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.VariableAssignment()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IThisAssignmentContext is an interface to support dynamic dispatch.
type IThisAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionalExpr() IConditionalExprContext
	Identifier() antlr.TerminalNode
	AllAssignmentPart() []IAssignmentPartContext
	AssignmentPart(i int) IAssignmentPartContext

	// IsThisAssignmentContext differentiates from other interfaces.
	IsThisAssignmentContext()
}

type ThisAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThisAssignmentContext() *ThisAssignmentContext {
	var p = new(ThisAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_thisAssignment
	return p
}

func InitEmptyThisAssignmentContext(p *ThisAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_thisAssignment
}

func (*ThisAssignmentContext) IsThisAssignmentContext() {}

func NewThisAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThisAssignmentContext {
	var p = new(ThisAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_thisAssignment

	return p
}

func (s *ThisAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *ThisAssignmentContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *ThisAssignmentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *ThisAssignmentContext) AllAssignmentPart() []IAssignmentPartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentPartContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentPartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentPartContext); ok {
			tst[i] = t.(IAssignmentPartContext)
			i++
		}
	}

	return tst
}

func (s *ThisAssignmentContext) AssignmentPart(i int) IAssignmentPartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentPartContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentPartContext)
}

func (s *ThisAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThisAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThisAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterThisAssignment(s)
	}
}

func (s *ThisAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitThisAssignment(s)
	}
}

func (p *CompiscriptParser) ThisAssignment() (localctx IThisAssignmentContext) {
	localctx = NewThisAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CompiscriptParserRULE_thisAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(CompiscriptParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(196)
		p.Match(CompiscriptParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Match(CompiscriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CompiscriptParserT__9 || _la == CompiscriptParserT__10 {
		{
			p.SetState(199)
			p.AssignmentPart()
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(205)
		p.Match(CompiscriptParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.ConditionalExpr()
	}
	{
		p.SetState(207)
		p.Match(CompiscriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableAssignmentContext is an interface to support dynamic dispatch.
type IVariableAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	ConditionalExpr() IConditionalExprContext
	AllAssignmentPart() []IAssignmentPartContext
	AssignmentPart(i int) IAssignmentPartContext

	// IsVariableAssignmentContext differentiates from other interfaces.
	IsVariableAssignmentContext()
}

type VariableAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableAssignmentContext() *VariableAssignmentContext {
	var p = new(VariableAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_variableAssignment
	return p
}

func InitEmptyVariableAssignmentContext(p *VariableAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_variableAssignment
}

func (*VariableAssignmentContext) IsVariableAssignmentContext() {}

func NewVariableAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableAssignmentContext {
	var p = new(VariableAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_variableAssignment

	return p
}

func (s *VariableAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableAssignmentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *VariableAssignmentContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *VariableAssignmentContext) AllAssignmentPart() []IAssignmentPartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentPartContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentPartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentPartContext); ok {
			tst[i] = t.(IAssignmentPartContext)
			i++
		}
	}

	return tst
}

func (s *VariableAssignmentContext) AssignmentPart(i int) IAssignmentPartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentPartContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentPartContext)
}

func (s *VariableAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterVariableAssignment(s)
	}
}

func (s *VariableAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitVariableAssignment(s)
	}
}

func (p *CompiscriptParser) VariableAssignment() (localctx IVariableAssignmentContext) {
	localctx = NewVariableAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CompiscriptParserRULE_variableAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(CompiscriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CompiscriptParserT__9 || _la == CompiscriptParserT__10 {
		{
			p.SetState(210)
			p.AssignmentPart()
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
		p.Match(CompiscriptParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.ConditionalExpr()
	}
	{
		p.SetState(218)
		p.Match(CompiscriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentPartContext is an interface to support dynamic dispatch.
type IAssignmentPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignmentPartContext differentiates from other interfaces.
	IsAssignmentPartContext()
}

type AssignmentPartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentPartContext() *AssignmentPartContext {
	var p = new(AssignmentPartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_assignmentPart
	return p
}

func InitEmptyAssignmentPartContext(p *AssignmentPartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_assignmentPart
}

func (*AssignmentPartContext) IsAssignmentPartContext() {}

func NewAssignmentPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentPartContext {
	var p = new(AssignmentPartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_assignmentPart

	return p
}

func (s *AssignmentPartContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentPartContext) CopyAll(ctx *AssignmentPartContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignmentPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IndexAssignmentPartExprContext struct {
	AssignmentPartContext
}

func NewIndexAssignmentPartExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexAssignmentPartExprContext {
	var p = new(IndexAssignmentPartExprContext)

	InitEmptyAssignmentPartContext(&p.AssignmentPartContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentPartContext))

	return p
}

func (s *IndexAssignmentPartExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexAssignmentPartExprContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *IndexAssignmentPartExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterIndexAssignmentPartExpr(s)
	}
}

func (s *IndexAssignmentPartExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitIndexAssignmentPartExpr(s)
	}
}

type FieldAssignmentPartExprContext struct {
	AssignmentPartContext
}

func NewFieldAssignmentPartExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAssignmentPartExprContext {
	var p = new(FieldAssignmentPartExprContext)

	InitEmptyAssignmentPartContext(&p.AssignmentPartContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentPartContext))

	return p
}

func (s *FieldAssignmentPartExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAssignmentPartExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *FieldAssignmentPartExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterFieldAssignmentPartExpr(s)
	}
}

func (s *FieldAssignmentPartExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitFieldAssignmentPartExpr(s)
	}
}

func (p *CompiscriptParser) AssignmentPart() (localctx IAssignmentPartContext) {
	localctx = NewAssignmentPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CompiscriptParserRULE_assignmentPart)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CompiscriptParserT__9:
		localctx = NewFieldAssignmentPartExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.Match(CompiscriptParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.Match(CompiscriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CompiscriptParserT__10:
		localctx = NewIndexAssignmentPartExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.Match(CompiscriptParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.ConditionalExpr()
		}
		{
			p.SetState(224)
			p.Match(CompiscriptParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMustBoolExprContext is an interface to support dynamic dispatch.
type IMustBoolExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionalExpr() IConditionalExprContext

	// IsMustBoolExprContext differentiates from other interfaces.
	IsMustBoolExprContext()
}

type MustBoolExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMustBoolExprContext() *MustBoolExprContext {
	var p = new(MustBoolExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_mustBoolExpr
	return p
}

func InitEmptyMustBoolExprContext(p *MustBoolExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_mustBoolExpr
}

func (*MustBoolExprContext) IsMustBoolExprContext() {}

func NewMustBoolExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MustBoolExprContext {
	var p = new(MustBoolExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_mustBoolExpr

	return p
}

func (s *MustBoolExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MustBoolExprContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *MustBoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MustBoolExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MustBoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterMustBoolExpr(s)
	}
}

func (s *MustBoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitMustBoolExpr(s)
	}
}

func (p *CompiscriptParser) MustBoolExpr() (localctx IMustBoolExprContext) {
	localctx = NewMustBoolExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CompiscriptParserRULE_mustBoolExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.ConditionalExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfCondition() IIfConditionContext
	IfBody() IIfBodyContext
	ElseBody() IElseBodyContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IfCondition() IIfConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfConditionContext)
}

func (s *IfStatementContext) IfBody() IIfBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBodyContext)
}

func (s *IfStatementContext) ElseBody() IElseBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBodyContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *CompiscriptParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CompiscriptParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(CompiscriptParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.Match(CompiscriptParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.IfCondition()
	}
	{
		p.SetState(233)
		p.Match(CompiscriptParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.IfBody()
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CompiscriptParserT__15 {
		{
			p.SetState(235)
			p.Match(CompiscriptParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.ElseBody()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfConditionContext is an interface to support dynamic dispatch.
type IIfConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MustBoolExpr() IMustBoolExprContext

	// IsIfConditionContext differentiates from other interfaces.
	IsIfConditionContext()
}

type IfConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfConditionContext() *IfConditionContext {
	var p = new(IfConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_ifCondition
	return p
}

func InitEmptyIfConditionContext(p *IfConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_ifCondition
}

func (*IfConditionContext) IsIfConditionContext() {}

func NewIfConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfConditionContext {
	var p = new(IfConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_ifCondition

	return p
}

func (s *IfConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *IfConditionContext) MustBoolExpr() IMustBoolExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMustBoolExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMustBoolExprContext)
}

func (s *IfConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterIfCondition(s)
	}
}

func (s *IfConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitIfCondition(s)
	}
}

func (p *CompiscriptParser) IfCondition() (localctx IIfConditionContext) {
	localctx = NewIfConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CompiscriptParserRULE_ifCondition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.MustBoolExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfBodyContext is an interface to support dynamic dispatch.
type IIfBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsIfBodyContext differentiates from other interfaces.
	IsIfBodyContext()
}

type IfBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBodyContext() *IfBodyContext {
	var p = new(IfBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_ifBody
	return p
}

func InitEmptyIfBodyContext(p *IfBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_ifBody
}

func (*IfBodyContext) IsIfBodyContext() {}

func NewIfBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBodyContext {
	var p = new(IfBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_ifBody

	return p
}

func (s *IfBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBodyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterIfBody(s)
	}
}

func (s *IfBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitIfBody(s)
	}
}

func (p *CompiscriptParser) IfBody() (localctx IIfBodyContext) {
	localctx = NewIfBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CompiscriptParserRULE_ifBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseBodyContext is an interface to support dynamic dispatch.
type IElseBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsElseBodyContext differentiates from other interfaces.
	IsElseBodyContext()
}

type ElseBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBodyContext() *ElseBodyContext {
	var p = new(ElseBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_elseBody
	return p
}

func InitEmptyElseBodyContext(p *ElseBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_elseBody
}

func (*ElseBodyContext) IsElseBodyContext() {}

func NewElseBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBodyContext {
	var p = new(ElseBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_elseBody

	return p
}

func (s *ElseBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBodyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterElseBody(s)
	}
}

func (s *ElseBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitElseBody(s)
	}
}

func (p *CompiscriptParser) ElseBody() (localctx IElseBodyContext) {
	localctx = NewElseBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CompiscriptParserRULE_elseBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WhileCondition() IWhileConditionContext
	WhileBody() IWhileBodyContext

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_whileStatement
	return p
}

func InitEmptyWhileStatementContext(p *WhileStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_whileStatement
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) WhileCondition() IWhileConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileConditionContext)
}

func (s *WhileStatementContext) WhileBody() IWhileBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileBodyContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (p *CompiscriptParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CompiscriptParserRULE_whileStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(CompiscriptParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Match(CompiscriptParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.WhileCondition()
	}
	{
		p.SetState(248)
		p.Match(CompiscriptParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.WhileBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileConditionContext is an interface to support dynamic dispatch.
type IWhileConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MustBoolExpr() IMustBoolExprContext

	// IsWhileConditionContext differentiates from other interfaces.
	IsWhileConditionContext()
}

type WhileConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileConditionContext() *WhileConditionContext {
	var p = new(WhileConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_whileCondition
	return p
}

func InitEmptyWhileConditionContext(p *WhileConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_whileCondition
}

func (*WhileConditionContext) IsWhileConditionContext() {}

func NewWhileConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileConditionContext {
	var p = new(WhileConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_whileCondition

	return p
}

func (s *WhileConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileConditionContext) MustBoolExpr() IMustBoolExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMustBoolExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMustBoolExprContext)
}

func (s *WhileConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterWhileCondition(s)
	}
}

func (s *WhileConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitWhileCondition(s)
	}
}

func (p *CompiscriptParser) WhileCondition() (localctx IWhileConditionContext) {
	localctx = NewWhileConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CompiscriptParserRULE_whileCondition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.MustBoolExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileBodyContext is an interface to support dynamic dispatch.
type IWhileBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsWhileBodyContext differentiates from other interfaces.
	IsWhileBodyContext()
}

type WhileBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileBodyContext() *WhileBodyContext {
	var p = new(WhileBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_whileBody
	return p
}

func InitEmptyWhileBodyContext(p *WhileBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_whileBody
}

func (*WhileBodyContext) IsWhileBodyContext() {}

func NewWhileBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileBodyContext {
	var p = new(WhileBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_whileBody

	return p
}

func (s *WhileBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileBodyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterWhileBody(s)
	}
}

func (s *WhileBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitWhileBody(s)
	}
}

func (p *CompiscriptParser) WhileBody() (localctx IWhileBodyContext) {
	localctx = NewWhileBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CompiscriptParserRULE_whileBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDoWhileStatementContext is an interface to support dynamic dispatch.
type IDoWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DoWhileBody() IDoWhileBodyContext
	WhileCondition() IWhileConditionContext

	// IsDoWhileStatementContext differentiates from other interfaces.
	IsDoWhileStatementContext()
}

type DoWhileStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoWhileStatementContext() *DoWhileStatementContext {
	var p = new(DoWhileStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_doWhileStatement
	return p
}

func InitEmptyDoWhileStatementContext(p *DoWhileStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_doWhileStatement
}

func (*DoWhileStatementContext) IsDoWhileStatementContext() {}

func NewDoWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoWhileStatementContext {
	var p = new(DoWhileStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_doWhileStatement

	return p
}

func (s *DoWhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DoWhileStatementContext) DoWhileBody() IDoWhileBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoWhileBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoWhileBodyContext)
}

func (s *DoWhileStatementContext) WhileCondition() IWhileConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileConditionContext)
}

func (s *DoWhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoWhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoWhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterDoWhileStatement(s)
	}
}

func (s *DoWhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitDoWhileStatement(s)
	}
}

func (p *CompiscriptParser) DoWhileStatement() (localctx IDoWhileStatementContext) {
	localctx = NewDoWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CompiscriptParserRULE_doWhileStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(CompiscriptParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		p.DoWhileBody()
	}
	{
		p.SetState(257)
		p.Match(CompiscriptParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(CompiscriptParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)
		p.WhileCondition()
	}
	{
		p.SetState(260)
		p.Match(CompiscriptParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Match(CompiscriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDoWhileBodyContext is an interface to support dynamic dispatch.
type IDoWhileBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsDoWhileBodyContext differentiates from other interfaces.
	IsDoWhileBodyContext()
}

type DoWhileBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoWhileBodyContext() *DoWhileBodyContext {
	var p = new(DoWhileBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_doWhileBody
	return p
}

func InitEmptyDoWhileBodyContext(p *DoWhileBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_doWhileBody
}

func (*DoWhileBodyContext) IsDoWhileBodyContext() {}

func NewDoWhileBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoWhileBodyContext {
	var p = new(DoWhileBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_doWhileBody

	return p
}

func (s *DoWhileBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *DoWhileBodyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DoWhileBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoWhileBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoWhileBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterDoWhileBody(s)
	}
}

func (s *DoWhileBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitDoWhileBody(s)
	}
}

func (p *CompiscriptParser) DoWhileBody() (localctx IDoWhileBodyContext) {
	localctx = NewDoWhileBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CompiscriptParserRULE_doWhileBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ForCondition() IForConditionContext
	ForBody() IForBodyContext
	VariableDeclaration() IVariableDeclarationContext
	Assignment() IAssignmentContext
	ForUpdate() IForUpdateContext

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) ForCondition() IForConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForConditionContext)
}

func (s *ForStatementContext) ForBody() IForBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForBodyContext)
}

func (s *ForStatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ForStatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ForStatementContext) ForUpdate() IForUpdateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForUpdateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForUpdateContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (p *CompiscriptParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CompiscriptParserRULE_forStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Match(CompiscriptParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
		p.Match(CompiscriptParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CompiscriptParserT__2, CompiscriptParserT__3:
		{
			p.SetState(267)
			p.VariableDeclaration()
		}

	case CompiscriptParserT__8, CompiscriptParserIdentifier:
		{
			p.SetState(268)
			p.Assignment()
		}

	case CompiscriptParserT__4:
		{
			p.SetState(269)
			p.Match(CompiscriptParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(272)
		p.ForCondition()
	}
	{
		p.SetState(273)
		p.Match(CompiscriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CompiscriptParserT__8 || _la == CompiscriptParserIdentifier {
		{
			p.SetState(274)
			p.ForUpdate()
		}

	}
	{
		p.SetState(277)
		p.Match(CompiscriptParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.ForBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForBodyContext is an interface to support dynamic dispatch.
type IForBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsForBodyContext differentiates from other interfaces.
	IsForBodyContext()
}

type ForBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForBodyContext() *ForBodyContext {
	var p = new(ForBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_forBody
	return p
}

func InitEmptyForBodyContext(p *ForBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_forBody
}

func (*ForBodyContext) IsForBodyContext() {}

func NewForBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForBodyContext {
	var p = new(ForBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_forBody

	return p
}

func (s *ForBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ForBodyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterForBody(s)
	}
}

func (s *ForBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitForBody(s)
	}
}

func (p *CompiscriptParser) ForBody() (localctx IForBodyContext) {
	localctx = NewForBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CompiscriptParserRULE_forBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForConditionContext is an interface to support dynamic dispatch.
type IForConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MustBoolExpr() IMustBoolExprContext

	// IsForConditionContext differentiates from other interfaces.
	IsForConditionContext()
}

type ForConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForConditionContext() *ForConditionContext {
	var p = new(ForConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_forCondition
	return p
}

func InitEmptyForConditionContext(p *ForConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_forCondition
}

func (*ForConditionContext) IsForConditionContext() {}

func NewForConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForConditionContext {
	var p = new(ForConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_forCondition

	return p
}

func (s *ForConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ForConditionContext) MustBoolExpr() IMustBoolExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMustBoolExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMustBoolExprContext)
}

func (s *ForConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterForCondition(s)
	}
}

func (s *ForConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitForCondition(s)
	}
}

func (p *CompiscriptParser) ForCondition() (localctx IForConditionContext) {
	localctx = NewForConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CompiscriptParserRULE_forCondition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.MustBoolExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForUpdateContext is an interface to support dynamic dispatch.
type IForUpdateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext

	// IsForUpdateContext differentiates from other interfaces.
	IsForUpdateContext()
}

type ForUpdateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForUpdateContext() *ForUpdateContext {
	var p = new(ForUpdateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_forUpdate
	return p
}

func InitEmptyForUpdateContext(p *ForUpdateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_forUpdate
}

func (*ForUpdateContext) IsForUpdateContext() {}

func NewForUpdateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForUpdateContext {
	var p = new(ForUpdateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_forUpdate

	return p
}

func (s *ForUpdateContext) GetParser() antlr.Parser { return s.parser }

func (s *ForUpdateContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ForUpdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForUpdateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForUpdateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterForUpdate(s)
	}
}

func (s *ForUpdateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitForUpdate(s)
	}
}

func (p *CompiscriptParser) ForUpdate() (localctx IForUpdateContext) {
	localctx = NewForUpdateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CompiscriptParserRULE_forUpdate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Assignment()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForeachValueContext is an interface to support dynamic dispatch.
type IForeachValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	ConditionalExpr() IConditionalExprContext

	// IsForeachValueContext differentiates from other interfaces.
	IsForeachValueContext()
}

type ForeachValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForeachValueContext() *ForeachValueContext {
	var p = new(ForeachValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_foreachValue
	return p
}

func InitEmptyForeachValueContext(p *ForeachValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_foreachValue
}

func (*ForeachValueContext) IsForeachValueContext() {}

func NewForeachValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForeachValueContext {
	var p = new(ForeachValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_foreachValue

	return p
}

func (s *ForeachValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ForeachValueContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *ForeachValueContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *ForeachValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForeachValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForeachValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterForeachValue(s)
	}
}

func (s *ForeachValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitForeachValue(s)
	}
}

func (p *CompiscriptParser) ForeachValue() (localctx IForeachValueContext) {
	localctx = NewForeachValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CompiscriptParserRULE_foreachValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Match(CompiscriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Match(CompiscriptParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.ConditionalExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForeachStatementContext is an interface to support dynamic dispatch.
type IForeachStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ForeachValue() IForeachValueContext
	Block() IBlockContext

	// IsForeachStatementContext differentiates from other interfaces.
	IsForeachStatementContext()
}

type ForeachStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForeachStatementContext() *ForeachStatementContext {
	var p = new(ForeachStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_foreachStatement
	return p
}

func InitEmptyForeachStatementContext(p *ForeachStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_foreachStatement
}

func (*ForeachStatementContext) IsForeachStatementContext() {}

func NewForeachStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForeachStatementContext {
	var p = new(ForeachStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_foreachStatement

	return p
}

func (s *ForeachStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForeachStatementContext) ForeachValue() IForeachValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForeachValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForeachValueContext)
}

func (s *ForeachStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForeachStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForeachStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForeachStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterForeachStatement(s)
	}
}

func (s *ForeachStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitForeachStatement(s)
	}
}

func (p *CompiscriptParser) ForeachStatement() (localctx IForeachStatementContext) {
	localctx = NewForeachStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CompiscriptParserRULE_foreachStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(CompiscriptParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Match(CompiscriptParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.ForeachValue()
	}
	{
		p.SetState(293)
		p.Match(CompiscriptParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_breakStatement
	return p
}

func InitEmptyBreakStatementContext(p *BreakStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_breakStatement
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (p *CompiscriptParser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CompiscriptParserRULE_breakStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(CompiscriptParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.Match(CompiscriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_continueStatement
	return p
}

func InitEmptyContinueStatementContext(p *ContinueStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_continueStatement
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (p *CompiscriptParser) ContinueStatement() (localctx IContinueStatementContext) {
	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CompiscriptParserRULE_continueStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(CompiscriptParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Match(CompiscriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionalExpr() IConditionalExprContext

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *CompiscriptParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CompiscriptParserRULE_returnStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Match(CompiscriptParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&328630831402732032) != 0 {
		{
			p.SetState(303)
			p.ConditionalExpr()
		}

	}
	{
		p.SetState(306)
		p.Match(CompiscriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_blockStatement
	return p
}

func InitEmptyBlockStatementContext(p *BlockStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_blockStatement
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

func (p *CompiscriptParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CompiscriptParserRULE_blockStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExceptionStatementContext is an interface to support dynamic dispatch.
type IExceptionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TryStatement() ITryStatementContext
	CatchStatement() ICatchStatementContext

	// IsExceptionStatementContext differentiates from other interfaces.
	IsExceptionStatementContext()
}

type ExceptionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExceptionStatementContext() *ExceptionStatementContext {
	var p = new(ExceptionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_exceptionStatement
	return p
}

func InitEmptyExceptionStatementContext(p *ExceptionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_exceptionStatement
}

func (*ExceptionStatementContext) IsExceptionStatementContext() {}

func NewExceptionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExceptionStatementContext {
	var p = new(ExceptionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_exceptionStatement

	return p
}

func (s *ExceptionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExceptionStatementContext) TryStatement() ITryStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITryStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITryStatementContext)
}

func (s *ExceptionStatementContext) CatchStatement() ICatchStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatchStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatchStatementContext)
}

func (s *ExceptionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExceptionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExceptionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterExceptionStatement(s)
	}
}

func (s *ExceptionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitExceptionStatement(s)
	}
}

func (p *CompiscriptParser) ExceptionStatement() (localctx IExceptionStatementContext) {
	localctx = NewExceptionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CompiscriptParserRULE_exceptionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.TryStatement()
	}
	{
		p.SetState(311)
		p.CatchStatement()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITryStatementContext is an interface to support dynamic dispatch.
type ITryStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext

	// IsTryStatementContext differentiates from other interfaces.
	IsTryStatementContext()
}

type TryStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTryStatementContext() *TryStatementContext {
	var p = new(TryStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_tryStatement
	return p
}

func InitEmptyTryStatementContext(p *TryStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_tryStatement
}

func (*TryStatementContext) IsTryStatementContext() {}

func NewTryStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TryStatementContext {
	var p = new(TryStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_tryStatement

	return p
}

func (s *TryStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TryStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TryStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TryStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterTryStatement(s)
	}
}

func (s *TryStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitTryStatement(s)
	}
}

func (p *CompiscriptParser) TryStatement() (localctx ITryStatementContext) {
	localctx = NewTryStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CompiscriptParserRULE_tryStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Match(CompiscriptParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICatchStatementContext is an interface to support dynamic dispatch.
type ICatchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Block() IBlockContext

	// IsCatchStatementContext differentiates from other interfaces.
	IsCatchStatementContext()
}

type CatchStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCatchStatementContext() *CatchStatementContext {
	var p = new(CatchStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_catchStatement
	return p
}

func InitEmptyCatchStatementContext(p *CatchStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_catchStatement
}

func (*CatchStatementContext) IsCatchStatementContext() {}

func NewCatchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatchStatementContext {
	var p = new(CatchStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_catchStatement

	return p
}

func (s *CatchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CatchStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *CatchStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CatchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterCatchStatement(s)
	}
}

func (s *CatchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitCatchStatement(s)
	}
}

func (p *CompiscriptParser) CatchStatement() (localctx ICatchStatementContext) {
	localctx = NewCatchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CompiscriptParserRULE_catchStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.Match(CompiscriptParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(317)
		p.Match(CompiscriptParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(318)
		p.Match(CompiscriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)
		p.Match(CompiscriptParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(320)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchValueContext is an interface to support dynamic dispatch.
type ISwitchValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionalExpr() IConditionalExprContext

	// IsSwitchValueContext differentiates from other interfaces.
	IsSwitchValueContext()
}

type SwitchValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchValueContext() *SwitchValueContext {
	var p = new(SwitchValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_switchValue
	return p
}

func InitEmptySwitchValueContext(p *SwitchValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_switchValue
}

func (*SwitchValueContext) IsSwitchValueContext() {}

func NewSwitchValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchValueContext {
	var p = new(SwitchValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_switchValue

	return p
}

func (s *SwitchValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchValueContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *SwitchValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterSwitchValue(s)
	}
}

func (s *SwitchValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitSwitchValue(s)
	}
}

func (p *CompiscriptParser) SwitchValue() (localctx ISwitchValueContext) {
	localctx = NewSwitchValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CompiscriptParserRULE_switchValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.ConditionalExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseValueContext is an interface to support dynamic dispatch.
type ICaseValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpr() IPrimaryExprContext

	// IsCaseValueContext differentiates from other interfaces.
	IsCaseValueContext()
}

type CaseValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseValueContext() *CaseValueContext {
	var p = new(CaseValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_caseValue
	return p
}

func InitEmptyCaseValueContext(p *CaseValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_caseValue
}

func (*CaseValueContext) IsCaseValueContext() {}

func NewCaseValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseValueContext {
	var p = new(CaseValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_caseValue

	return p
}

func (s *CaseValueContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseValueContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *CaseValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterCaseValue(s)
	}
}

func (s *CaseValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitCaseValue(s)
	}
}

func (p *CompiscriptParser) CaseValue() (localctx ICaseValueContext) {
	localctx = NewCaseValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, CompiscriptParserRULE_caseValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.PrimaryExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SwitchValue() ISwitchValueContext
	AllSwitchCase() []ISwitchCaseContext
	SwitchCase(i int) ISwitchCaseContext
	DefaultCase() IDefaultCaseContext

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_switchStatement
	return p
}

func InitEmptySwitchStatementContext(p *SwitchStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_switchStatement
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) SwitchValue() ISwitchValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchValueContext)
}

func (s *SwitchStatementContext) AllSwitchCase() []ISwitchCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			len++
		}
	}

	tst := make([]ISwitchCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitchCaseContext); ok {
			tst[i] = t.(ISwitchCaseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStatementContext) SwitchCase(i int) ISwitchCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchCaseContext)
}

func (s *SwitchStatementContext) DefaultCase() IDefaultCaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultCaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultCaseContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (p *CompiscriptParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CompiscriptParserRULE_switchStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Match(CompiscriptParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(327)
		p.Match(CompiscriptParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(328)
		p.SwitchValue()
	}
	{
		p.SetState(329)
		p.Match(CompiscriptParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(330)
		p.Match(CompiscriptParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CompiscriptParserT__27 {
		{
			p.SetState(331)
			p.SwitchCase()
		}

		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CompiscriptParserT__28 {
		{
			p.SetState(337)
			p.DefaultCase()
		}

	}
	{
		p.SetState(340)
		p.Match(CompiscriptParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchCaseContext is an interface to support dynamic dispatch.
type ISwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CaseValue() ICaseValueContext
	CaseBody() ICaseBodyContext

	// IsSwitchCaseContext differentiates from other interfaces.
	IsSwitchCaseContext()
}

type SwitchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchCaseContext() *SwitchCaseContext {
	var p = new(SwitchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_switchCase
	return p
}

func InitEmptySwitchCaseContext(p *SwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_switchCase
}

func (*SwitchCaseContext) IsSwitchCaseContext() {}

func NewSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_switchCase

	return p
}

func (s *SwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchCaseContext) CaseValue() ICaseValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseValueContext)
}

func (s *SwitchCaseContext) CaseBody() ICaseBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseBodyContext)
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterSwitchCase(s)
	}
}

func (s *SwitchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitSwitchCase(s)
	}
}

func (p *CompiscriptParser) SwitchCase() (localctx ISwitchCaseContext) {
	localctx = NewSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, CompiscriptParserRULE_switchCase)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.Match(CompiscriptParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.CaseValue()
	}
	{
		p.SetState(344)
		p.Match(CompiscriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.CaseBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultCaseContext is an interface to support dynamic dispatch.
type IDefaultCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CaseBody() ICaseBodyContext

	// IsDefaultCaseContext differentiates from other interfaces.
	IsDefaultCaseContext()
}

type DefaultCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultCaseContext() *DefaultCaseContext {
	var p = new(DefaultCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_defaultCase
	return p
}

func InitEmptyDefaultCaseContext(p *DefaultCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_defaultCase
}

func (*DefaultCaseContext) IsDefaultCaseContext() {}

func NewDefaultCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultCaseContext {
	var p = new(DefaultCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_defaultCase

	return p
}

func (s *DefaultCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultCaseContext) CaseBody() ICaseBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseBodyContext)
}

func (s *DefaultCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterDefaultCase(s)
	}
}

func (s *DefaultCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitDefaultCase(s)
	}
}

func (p *CompiscriptParser) DefaultCase() (localctx IDefaultCaseContext) {
	localctx = NewDefaultCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, CompiscriptParserRULE_defaultCase)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(CompiscriptParserT__28)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.Match(CompiscriptParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)
		p.CaseBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseBodyContext is an interface to support dynamic dispatch.
type ICaseBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsCaseBodyContext differentiates from other interfaces.
	IsCaseBodyContext()
}

type CaseBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseBodyContext() *CaseBodyContext {
	var p = new(CaseBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_caseBody
	return p
}

func InitEmptyCaseBodyContext(p *CaseBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_caseBody
}

func (*CaseBodyContext) IsCaseBodyContext() {}

func NewCaseBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseBodyContext {
	var p = new(CaseBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_caseBody

	return p
}

func (s *CaseBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseBodyContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *CaseBodyContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CaseBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterCaseBody(s)
	}
}

func (s *CaseBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitCaseBody(s)
	}
}

func (p *CompiscriptParser) CaseBody() (localctx ICaseBodyContext) {
	localctx = NewCaseBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, CompiscriptParserRULE_caseBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&290482181534261850) != 0 {
		{
			p.SetState(351)
			p.Statement()
		}

		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Block() IBlockContext
	Parameters() IParametersContext
	Type_() ITypeContext

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_functionDeclaration
	return p
}

func InitEmptyFunctionDeclarationContext(p *FunctionDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_functionDeclaration
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclarationContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (p *CompiscriptParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, CompiscriptParserRULE_functionDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(CompiscriptParserT__29)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
		p.Match(CompiscriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)
		p.Match(CompiscriptParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CompiscriptParserIdentifier {
		{
			p.SetState(360)
			p.Parameters()
		}

	}
	{
		p.SetState(363)
		p.Match(CompiscriptParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CompiscriptParserT__7 {
		{
			p.SetState(364)
			p.Match(CompiscriptParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.Type_()
		}

	}
	{
		p.SetState(368)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_parameters
	return p
}

func InitEmptyParametersContext(p *ParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_parameters
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ParametersContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *CompiscriptParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, CompiscriptParserRULE_parameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Parameter()
	}
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CompiscriptParserT__30 {
		{
			p.SetState(371)
			p.Match(CompiscriptParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.Parameter()
		}

		p.SetState(377)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Type_() ITypeContext

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *ParameterContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *CompiscriptParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, CompiscriptParserRULE_parameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(CompiscriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CompiscriptParserT__7 {
		{
			p.SetState(379)
			p.Match(CompiscriptParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.Type_()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClassDeclarationContext is an interface to support dynamic dispatch.
type IClassDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllClassMember() []IClassMemberContext
	ClassMember(i int) IClassMemberContext

	// IsClassDeclarationContext differentiates from other interfaces.
	IsClassDeclarationContext()
}

type ClassDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDeclarationContext() *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_classDeclaration
	return p
}

func InitEmptyClassDeclarationContext(p *ClassDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_classDeclaration
}

func (*ClassDeclarationContext) IsClassDeclarationContext() {}

func NewClassDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_classDeclaration

	return p
}

func (s *ClassDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDeclarationContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(CompiscriptParserIdentifier)
}

func (s *ClassDeclarationContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, i)
}

func (s *ClassDeclarationContext) AllClassMember() []IClassMemberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IClassMemberContext); ok {
			len++
		}
	}

	tst := make([]IClassMemberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IClassMemberContext); ok {
			tst[i] = t.(IClassMemberContext)
			i++
		}
	}

	return tst
}

func (s *ClassDeclarationContext) ClassMember(i int) IClassMemberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassMemberContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassMemberContext)
}

func (s *ClassDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitClassDeclaration(s)
	}
}

func (p *CompiscriptParser) ClassDeclaration() (localctx IClassDeclarationContext) {
	localctx = NewClassDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, CompiscriptParserRULE_classDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(CompiscriptParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(384)
		p.Match(CompiscriptParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CompiscriptParserT__7 {
		{
			p.SetState(385)
			p.Match(CompiscriptParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.Match(CompiscriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(389)
		p.Match(CompiscriptParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1073741912) != 0 {
		{
			p.SetState(390)
			p.ClassMember()
		}

		p.SetState(395)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(396)
		p.Match(CompiscriptParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClassMemberContext is an interface to support dynamic dispatch.
type IClassMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionDeclaration() IFunctionDeclarationContext
	VariableDeclaration() IVariableDeclarationContext
	ConstantDeclaration() IConstantDeclarationContext

	// IsClassMemberContext differentiates from other interfaces.
	IsClassMemberContext()
}

type ClassMemberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassMemberContext() *ClassMemberContext {
	var p = new(ClassMemberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_classMember
	return p
}

func InitEmptyClassMemberContext(p *ClassMemberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_classMember
}

func (*ClassMemberContext) IsClassMemberContext() {}

func NewClassMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassMemberContext {
	var p = new(ClassMemberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_classMember

	return p
}

func (s *ClassMemberContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassMemberContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *ClassMemberContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ClassMemberContext) ConstantDeclaration() IConstantDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantDeclarationContext)
}

func (s *ClassMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassMemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterClassMember(s)
	}
}

func (s *ClassMemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitClassMember(s)
	}
}

func (p *CompiscriptParser) ClassMember() (localctx IClassMemberContext) {
	localctx = NewClassMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, CompiscriptParserRULE_classMember)
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CompiscriptParserT__29:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)
			p.FunctionDeclaration()
		}

	case CompiscriptParserT__2, CompiscriptParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(399)
			p.VariableDeclaration()
		}

	case CompiscriptParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(400)
			p.ConstantDeclaration()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionalExpr() IConditionalExprContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CompiscriptParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, CompiscriptParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(403)
		p.ConditionalExpr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalExprContext is an interface to support dynamic dispatch.
type IConditionalExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOrExpr() ILogicalOrExprContext
	AllConditionalExpr() []IConditionalExprContext
	ConditionalExpr(i int) IConditionalExprContext

	// IsConditionalExprContext differentiates from other interfaces.
	IsConditionalExprContext()
}

type ConditionalExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalExprContext() *ConditionalExprContext {
	var p = new(ConditionalExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_conditionalExpr
	return p
}

func InitEmptyConditionalExprContext(p *ConditionalExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_conditionalExpr
}

func (*ConditionalExprContext) IsConditionalExprContext() {}

func NewConditionalExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalExprContext {
	var p = new(ConditionalExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_conditionalExpr

	return p
}

func (s *ConditionalExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalExprContext) LogicalOrExpr() ILogicalOrExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExprContext)
}

func (s *ConditionalExprContext) AllConditionalExpr() []IConditionalExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionalExprContext); ok {
			len++
		}
	}

	tst := make([]IConditionalExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionalExprContext); ok {
			tst[i] = t.(IConditionalExprContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalExprContext) ConditionalExpr(i int) IConditionalExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *ConditionalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterConditionalExpr(s)
	}
}

func (s *ConditionalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitConditionalExpr(s)
	}
}

func (p *CompiscriptParser) ConditionalExpr() (localctx IConditionalExprContext) {
	localctx = NewConditionalExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, CompiscriptParserRULE_conditionalExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(405)
		p.LogicalOrExpr()
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CompiscriptParserT__32 {
		{
			p.SetState(406)
			p.Match(CompiscriptParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.ConditionalExpr()
		}
		{
			p.SetState(408)
			p.Match(CompiscriptParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.ConditionalExpr()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalOrExprContext is an interface to support dynamic dispatch.
type ILogicalOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLogicalAndExpr() []ILogicalAndExprContext
	LogicalAndExpr(i int) ILogicalAndExprContext

	// IsLogicalOrExprContext differentiates from other interfaces.
	IsLogicalOrExprContext()
}

type LogicalOrExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOrExprContext() *LogicalOrExprContext {
	var p = new(LogicalOrExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_logicalOrExpr
	return p
}

func InitEmptyLogicalOrExprContext(p *LogicalOrExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_logicalOrExpr
}

func (*LogicalOrExprContext) IsLogicalOrExprContext() {}

func NewLogicalOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrExprContext {
	var p = new(LogicalOrExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_logicalOrExpr

	return p
}

func (s *LogicalOrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrExprContext) AllLogicalAndExpr() []ILogicalAndExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalAndExprContext); ok {
			len++
		}
	}

	tst := make([]ILogicalAndExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalAndExprContext); ok {
			tst[i] = t.(ILogicalAndExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrExprContext) LogicalAndExpr(i int) ILogicalAndExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalAndExprContext)
}

func (s *LogicalOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterLogicalOrExpr(s)
	}
}

func (s *LogicalOrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitLogicalOrExpr(s)
	}
}

func (p *CompiscriptParser) LogicalOrExpr() (localctx ILogicalOrExprContext) {
	localctx = NewLogicalOrExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, CompiscriptParserRULE_logicalOrExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(413)
		p.LogicalAndExpr()
	}
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CompiscriptParserT__33 {
		{
			p.SetState(414)
			p.Match(CompiscriptParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.LogicalAndExpr()
		}

		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalAndExprContext is an interface to support dynamic dispatch.
type ILogicalAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEqualityExpr() []IEqualityExprContext
	EqualityExpr(i int) IEqualityExprContext

	// IsLogicalAndExprContext differentiates from other interfaces.
	IsLogicalAndExprContext()
}

type LogicalAndExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAndExprContext() *LogicalAndExprContext {
	var p = new(LogicalAndExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_logicalAndExpr
	return p
}

func InitEmptyLogicalAndExprContext(p *LogicalAndExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_logicalAndExpr
}

func (*LogicalAndExprContext) IsLogicalAndExprContext() {}

func NewLogicalAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndExprContext {
	var p = new(LogicalAndExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_logicalAndExpr

	return p
}

func (s *LogicalAndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndExprContext) AllEqualityExpr() []IEqualityExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualityExprContext); ok {
			len++
		}
	}

	tst := make([]IEqualityExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualityExprContext); ok {
			tst[i] = t.(IEqualityExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndExprContext) EqualityExpr(i int) IEqualityExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExprContext)
}

func (s *LogicalAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterLogicalAndExpr(s)
	}
}

func (s *LogicalAndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitLogicalAndExpr(s)
	}
}

func (p *CompiscriptParser) LogicalAndExpr() (localctx ILogicalAndExprContext) {
	localctx = NewLogicalAndExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, CompiscriptParserRULE_logicalAndExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.EqualityExpr()
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CompiscriptParserT__34 {
		{
			p.SetState(422)
			p.Match(CompiscriptParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.EqualityExpr()
		}

		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualityExprContext is an interface to support dynamic dispatch.
type IEqualityExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRelationalExpr() []IRelationalExprContext
	RelationalExpr(i int) IRelationalExprContext

	// IsEqualityExprContext differentiates from other interfaces.
	IsEqualityExprContext()
}

type EqualityExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExprContext() *EqualityExprContext {
	var p = new(EqualityExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_equalityExpr
	return p
}

func InitEmptyEqualityExprContext(p *EqualityExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_equalityExpr
}

func (*EqualityExprContext) IsEqualityExprContext() {}

func NewEqualityExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExprContext {
	var p = new(EqualityExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_equalityExpr

	return p
}

func (s *EqualityExprContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExprContext) AllRelationalExpr() []IRelationalExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationalExprContext); ok {
			len++
		}
	}

	tst := make([]IRelationalExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationalExprContext); ok {
			tst[i] = t.(IRelationalExprContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExprContext) RelationalExpr(i int) IRelationalExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalExprContext)
}

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterEqualityExpr(s)
	}
}

func (s *EqualityExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitEqualityExpr(s)
	}
}

func (p *CompiscriptParser) EqualityExpr() (localctx IEqualityExprContext) {
	localctx = NewEqualityExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, CompiscriptParserRULE_equalityExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)
		p.RelationalExpr()
	}
	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CompiscriptParserT__35 || _la == CompiscriptParserT__36 {
		{
			p.SetState(430)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CompiscriptParserT__35 || _la == CompiscriptParserT__36) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(431)
			p.RelationalExpr()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationalExprContext is an interface to support dynamic dispatch.
type IRelationalExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAdditiveExpr() []IAdditiveExprContext
	AdditiveExpr(i int) IAdditiveExprContext

	// IsRelationalExprContext differentiates from other interfaces.
	IsRelationalExprContext()
}

type RelationalExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExprContext() *RelationalExprContext {
	var p = new(RelationalExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_relationalExpr
	return p
}

func InitEmptyRelationalExprContext(p *RelationalExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_relationalExpr
}

func (*RelationalExprContext) IsRelationalExprContext() {}

func NewRelationalExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExprContext {
	var p = new(RelationalExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_relationalExpr

	return p
}

func (s *RelationalExprContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExprContext) AllAdditiveExpr() []IAdditiveExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveExprContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveExprContext); ok {
			tst[i] = t.(IAdditiveExprContext)
			i++
		}
	}

	return tst
}

func (s *RelationalExprContext) AdditiveExpr(i int) IAdditiveExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExprContext)
}

func (s *RelationalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterRelationalExpr(s)
	}
}

func (s *RelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitRelationalExpr(s)
	}
}

func (p *CompiscriptParser) RelationalExpr() (localctx IRelationalExprContext) {
	localctx = NewRelationalExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, CompiscriptParserRULE_relationalExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		p.AdditiveExpr()
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4123168604160) != 0 {
		{
			p.SetState(435)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4123168604160) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(436)
			p.AdditiveExpr()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditiveExprContext is an interface to support dynamic dispatch.
type IAdditiveExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMultiplicativeExpr() []IMultiplicativeExprContext
	MultiplicativeExpr(i int) IMultiplicativeExprContext

	// IsAdditiveExprContext differentiates from other interfaces.
	IsAdditiveExprContext()
}

type AdditiveExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExprContext() *AdditiveExprContext {
	var p = new(AdditiveExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_additiveExpr
	return p
}

func InitEmptyAdditiveExprContext(p *AdditiveExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_additiveExpr
}

func (*AdditiveExprContext) IsAdditiveExprContext() {}

func NewAdditiveExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExprContext {
	var p = new(AdditiveExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_additiveExpr

	return p
}

func (s *AdditiveExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExprContext) AllMultiplicativeExpr() []IMultiplicativeExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExprContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExprContext); ok {
			tst[i] = t.(IMultiplicativeExprContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExprContext) MultiplicativeExpr(i int) IMultiplicativeExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExprContext)
}

func (s *AdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitAdditiveExpr(s)
	}
}

func (p *CompiscriptParser) AdditiveExpr() (localctx IAdditiveExprContext) {
	localctx = NewAdditiveExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, CompiscriptParserRULE_additiveExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)
		p.MultiplicativeExpr()
	}
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CompiscriptParserT__41 || _la == CompiscriptParserT__42 {
		{
			p.SetState(440)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CompiscriptParserT__41 || _la == CompiscriptParserT__42) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(441)
			p.MultiplicativeExpr()
		}

		p.SetState(446)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicativeExprContext is an interface to support dynamic dispatch.
type IMultiplicativeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnaryExpr() []IUnaryExprContext
	UnaryExpr(i int) IUnaryExprContext

	// IsMultiplicativeExprContext differentiates from other interfaces.
	IsMultiplicativeExprContext()
}

type MultiplicativeExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExprContext() *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_multiplicativeExpr
	return p
}

func InitEmptyMultiplicativeExprContext(p *MultiplicativeExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_multiplicativeExpr
}

func (*MultiplicativeExprContext) IsMultiplicativeExprContext() {}

func NewMultiplicativeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_multiplicativeExpr

	return p
}

func (s *MultiplicativeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExprContext) AllUnaryExpr() []IUnaryExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryExprContext); ok {
			len++
		}
	}

	tst := make([]IUnaryExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryExprContext); ok {
			tst[i] = t.(IUnaryExprContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExprContext) UnaryExpr(i int) IUnaryExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *MultiplicativeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterMultiplicativeExpr(s)
	}
}

func (s *MultiplicativeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitMultiplicativeExpr(s)
	}
}

func (p *CompiscriptParser) MultiplicativeExpr() (localctx IMultiplicativeExprContext) {
	localctx = NewMultiplicativeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, CompiscriptParserRULE_multiplicativeExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(447)
		p.UnaryExpr()
	}
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&123145302310912) != 0 {
		{
			p.SetState(448)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&123145302310912) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(449)
			p.UnaryExpr()
		}

		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpr() IUnaryExprContext
	PrimaryExpr() IPrimaryExprContext

	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_unaryExpr
	return p
}

func InitEmptyUnaryExprContext(p *UnaryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_unaryExpr
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *UnaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

func (p *CompiscriptParser) UnaryExpr() (localctx IUnaryExprContext) {
	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, CompiscriptParserRULE_unaryExpr)
	var _la int

	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CompiscriptParserT__42, CompiscriptParserT__46:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(455)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CompiscriptParserT__42 || _la == CompiscriptParserT__46) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(456)
			p.UnaryExpr()
		}

	case CompiscriptParserT__8, CompiscriptParserT__10, CompiscriptParserT__13, CompiscriptParserT__47, CompiscriptParserT__48, CompiscriptParserT__49, CompiscriptParserT__50, CompiscriptParserLiteral, CompiscriptParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(457)
			p.PrimaryExpr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LiteralExpr() ILiteralExprContext
	LeftHandSide() ILeftHandSideContext
	ConditionalExpr() IConditionalExprContext

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_primaryExpr
	return p
}

func InitEmptyPrimaryExprContext(p *PrimaryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_primaryExpr
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) LiteralExpr() ILiteralExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralExprContext)
}

func (s *PrimaryExprContext) LeftHandSide() ILeftHandSideContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftHandSideContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftHandSideContext)
}

func (s *PrimaryExprContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (p *CompiscriptParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, CompiscriptParserRULE_primaryExpr)
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CompiscriptParserT__10, CompiscriptParserT__47, CompiscriptParserT__48, CompiscriptParserT__49, CompiscriptParserLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(460)
			p.LiteralExpr()
		}

	case CompiscriptParserT__8, CompiscriptParserT__50, CompiscriptParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(461)
			p.LeftHandSide()
		}

	case CompiscriptParserT__13:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(462)
			p.Match(CompiscriptParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)
			p.ConditionalExpr()
		}
		{
			p.SetState(464)
			p.Match(CompiscriptParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralExprContext is an interface to support dynamic dispatch.
type ILiteralExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() antlr.TerminalNode
	ArrayLiteral() IArrayLiteralContext

	// IsLiteralExprContext differentiates from other interfaces.
	IsLiteralExprContext()
}

type LiteralExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralExprContext() *LiteralExprContext {
	var p = new(LiteralExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_literalExpr
	return p
}

func InitEmptyLiteralExprContext(p *LiteralExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_literalExpr
}

func (*LiteralExprContext) IsLiteralExprContext() {}

func NewLiteralExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralExprContext {
	var p = new(LiteralExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_literalExpr

	return p
}

func (s *LiteralExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralExprContext) Literal() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserLiteral, 0)
}

func (s *LiteralExprContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *LiteralExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterLiteralExpr(s)
	}
}

func (s *LiteralExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitLiteralExpr(s)
	}
}

func (p *CompiscriptParser) LiteralExpr() (localctx ILiteralExprContext) {
	localctx = NewLiteralExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, CompiscriptParserRULE_literalExpr)
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CompiscriptParserLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(468)
			p.Match(CompiscriptParserLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CompiscriptParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(469)
			p.ArrayLiteral()
		}

	case CompiscriptParserT__47:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(470)
			p.Match(CompiscriptParserT__47)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CompiscriptParserT__48:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(471)
			p.Match(CompiscriptParserT__48)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CompiscriptParserT__49:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(472)
			p.Match(CompiscriptParserT__49)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILeftHandSideContext is an interface to support dynamic dispatch.
type ILeftHandSideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryAtom() IPrimaryAtomContext
	AllSuffixOp() []ISuffixOpContext
	SuffixOp(i int) ISuffixOpContext

	// IsLeftHandSideContext differentiates from other interfaces.
	IsLeftHandSideContext()
}

type LeftHandSideContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeftHandSideContext() *LeftHandSideContext {
	var p = new(LeftHandSideContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_leftHandSide
	return p
}

func InitEmptyLeftHandSideContext(p *LeftHandSideContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_leftHandSide
}

func (*LeftHandSideContext) IsLeftHandSideContext() {}

func NewLeftHandSideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftHandSideContext {
	var p = new(LeftHandSideContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_leftHandSide

	return p
}

func (s *LeftHandSideContext) GetParser() antlr.Parser { return s.parser }

func (s *LeftHandSideContext) PrimaryAtom() IPrimaryAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryAtomContext)
}

func (s *LeftHandSideContext) AllSuffixOp() []ISuffixOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISuffixOpContext); ok {
			len++
		}
	}

	tst := make([]ISuffixOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISuffixOpContext); ok {
			tst[i] = t.(ISuffixOpContext)
			i++
		}
	}

	return tst
}

func (s *LeftHandSideContext) SuffixOp(i int) ISuffixOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISuffixOpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISuffixOpContext)
}

func (s *LeftHandSideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftHandSideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeftHandSideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterLeftHandSide(s)
	}
}

func (s *LeftHandSideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitLeftHandSide(s)
	}
}

func (p *CompiscriptParser) LeftHandSide() (localctx ILeftHandSideContext) {
	localctx = NewLeftHandSideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, CompiscriptParserRULE_leftHandSide)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(475)
		p.PrimaryAtom()
	}
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&19456) != 0 {
		{
			p.SetState(476)
			p.SuffixOp()
		}

		p.SetState(481)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStandaloneExpresionContext is an interface to support dynamic dispatch.
type IStandaloneExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StandaloneAtom() IStandaloneAtomContext
	AllSuffixOp() []ISuffixOpContext
	SuffixOp(i int) ISuffixOpContext

	// IsStandaloneExpresionContext differentiates from other interfaces.
	IsStandaloneExpresionContext()
}

type StandaloneExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStandaloneExpresionContext() *StandaloneExpresionContext {
	var p = new(StandaloneExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_standaloneExpresion
	return p
}

func InitEmptyStandaloneExpresionContext(p *StandaloneExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_standaloneExpresion
}

func (*StandaloneExpresionContext) IsStandaloneExpresionContext() {}

func NewStandaloneExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StandaloneExpresionContext {
	var p = new(StandaloneExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_standaloneExpresion

	return p
}

func (s *StandaloneExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *StandaloneExpresionContext) StandaloneAtom() IStandaloneAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStandaloneAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStandaloneAtomContext)
}

func (s *StandaloneExpresionContext) AllSuffixOp() []ISuffixOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISuffixOpContext); ok {
			len++
		}
	}

	tst := make([]ISuffixOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISuffixOpContext); ok {
			tst[i] = t.(ISuffixOpContext)
			i++
		}
	}

	return tst
}

func (s *StandaloneExpresionContext) SuffixOp(i int) ISuffixOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISuffixOpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISuffixOpContext)
}

func (s *StandaloneExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StandaloneExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StandaloneExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterStandaloneExpresion(s)
	}
}

func (s *StandaloneExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitStandaloneExpresion(s)
	}
}

func (p *CompiscriptParser) StandaloneExpresion() (localctx IStandaloneExpresionContext) {
	localctx = NewStandaloneExpresionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, CompiscriptParserRULE_standaloneExpresion)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(482)
		p.StandaloneAtom()
	}
	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&19456) != 0 {
		{
			p.SetState(483)
			p.SuffixOp()
		}

		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(489)
		p.Match(CompiscriptParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryAtomContext is an interface to support dynamic dispatch.
type IPrimaryAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryAtomContext differentiates from other interfaces.
	IsPrimaryAtomContext()
}

type PrimaryAtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryAtomContext() *PrimaryAtomContext {
	var p = new(PrimaryAtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_primaryAtom
	return p
}

func InitEmptyPrimaryAtomContext(p *PrimaryAtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_primaryAtom
}

func (*PrimaryAtomContext) IsPrimaryAtomContext() {}

func NewPrimaryAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryAtomContext {
	var p = new(PrimaryAtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_primaryAtom

	return p
}

func (s *PrimaryAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryAtomContext) CopyAll(ctx *PrimaryAtomContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentifierExprContext struct {
	PrimaryAtomContext
}

func NewIdentifierExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExprContext {
	var p = new(IdentifierExprContext)

	InitEmptyPrimaryAtomContext(&p.PrimaryAtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryAtomContext))

	return p
}

func (s *IdentifierExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *IdentifierExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterIdentifierExpr(s)
	}
}

func (s *IdentifierExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitIdentifierExpr(s)
	}
}

type NewExprContext struct {
	PrimaryAtomContext
}

func NewNewExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewExprContext {
	var p = new(NewExprContext)

	InitEmptyPrimaryAtomContext(&p.PrimaryAtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryAtomContext))

	return p
}

func (s *NewExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *NewExprContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *NewExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterNewExpr(s)
	}
}

func (s *NewExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitNewExpr(s)
	}
}

type ThisExprContext struct {
	PrimaryAtomContext
}

func NewThisExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThisExprContext {
	var p = new(ThisExprContext)

	InitEmptyPrimaryAtomContext(&p.PrimaryAtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryAtomContext))

	return p
}

func (s *ThisExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThisExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterThisExpr(s)
	}
}

func (s *ThisExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitThisExpr(s)
	}
}

func (p *CompiscriptParser) PrimaryAtom() (localctx IPrimaryAtomContext) {
	localctx = NewPrimaryAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, CompiscriptParserRULE_primaryAtom)
	var _la int

	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CompiscriptParserIdentifier:
		localctx = NewIdentifierExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(491)
			p.Match(CompiscriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CompiscriptParserT__50:
		localctx = NewNewExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(492)
			p.Match(CompiscriptParserT__50)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(493)
			p.Match(CompiscriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)
			p.Match(CompiscriptParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(496)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&328630831402732032) != 0 {
			{
				p.SetState(495)
				p.Arguments()
			}

		}
		{
			p.SetState(498)
			p.Match(CompiscriptParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CompiscriptParserT__8:
		localctx = NewThisExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(499)
			p.Match(CompiscriptParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStandaloneAtomContext is an interface to support dynamic dispatch.
type IStandaloneAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStandaloneAtomContext differentiates from other interfaces.
	IsStandaloneAtomContext()
}

type StandaloneAtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStandaloneAtomContext() *StandaloneAtomContext {
	var p = new(StandaloneAtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_standaloneAtom
	return p
}

func InitEmptyStandaloneAtomContext(p *StandaloneAtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_standaloneAtom
}

func (*StandaloneAtomContext) IsStandaloneAtomContext() {}

func NewStandaloneAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StandaloneAtomContext {
	var p = new(StandaloneAtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_standaloneAtom

	return p
}

func (s *StandaloneAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *StandaloneAtomContext) CopyAll(ctx *StandaloneAtomContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StandaloneAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StandaloneAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StandaloneThisExprContext struct {
	StandaloneAtomContext
}

func NewStandaloneThisExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StandaloneThisExprContext {
	var p = new(StandaloneThisExprContext)

	InitEmptyStandaloneAtomContext(&p.StandaloneAtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*StandaloneAtomContext))

	return p
}

func (s *StandaloneThisExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StandaloneThisExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterStandaloneThisExpr(s)
	}
}

func (s *StandaloneThisExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitStandaloneThisExpr(s)
	}
}

type StandaloneNewExprContext struct {
	StandaloneAtomContext
}

func NewStandaloneNewExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StandaloneNewExprContext {
	var p = new(StandaloneNewExprContext)

	InitEmptyStandaloneAtomContext(&p.StandaloneAtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*StandaloneAtomContext))

	return p
}

func (s *StandaloneNewExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StandaloneNewExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *StandaloneNewExprContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *StandaloneNewExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterStandaloneNewExpr(s)
	}
}

func (s *StandaloneNewExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitStandaloneNewExpr(s)
	}
}

type StandaloneIdentifierExprContext struct {
	StandaloneAtomContext
}

func NewStandaloneIdentifierExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StandaloneIdentifierExprContext {
	var p = new(StandaloneIdentifierExprContext)

	InitEmptyStandaloneAtomContext(&p.StandaloneAtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*StandaloneAtomContext))

	return p
}

func (s *StandaloneIdentifierExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StandaloneIdentifierExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *StandaloneIdentifierExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterStandaloneIdentifierExpr(s)
	}
}

func (s *StandaloneIdentifierExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitStandaloneIdentifierExpr(s)
	}
}

func (p *CompiscriptParser) StandaloneAtom() (localctx IStandaloneAtomContext) {
	localctx = NewStandaloneAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, CompiscriptParserRULE_standaloneAtom)
	var _la int

	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CompiscriptParserIdentifier:
		localctx = NewStandaloneIdentifierExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(502)
			p.Match(CompiscriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CompiscriptParserT__50:
		localctx = NewStandaloneNewExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(503)
			p.Match(CompiscriptParserT__50)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(504)
			p.Match(CompiscriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(505)
			p.Match(CompiscriptParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(507)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&328630831402732032) != 0 {
			{
				p.SetState(506)
				p.Arguments()
			}

		}
		{
			p.SetState(509)
			p.Match(CompiscriptParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CompiscriptParserT__8:
		localctx = NewStandaloneThisExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(510)
			p.Match(CompiscriptParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISuffixOpContext is an interface to support dynamic dispatch.
type ISuffixOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSuffixOpContext differentiates from other interfaces.
	IsSuffixOpContext()
}

type SuffixOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuffixOpContext() *SuffixOpContext {
	var p = new(SuffixOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_suffixOp
	return p
}

func InitEmptySuffixOpContext(p *SuffixOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_suffixOp
}

func (*SuffixOpContext) IsSuffixOpContext() {}

func NewSuffixOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuffixOpContext {
	var p = new(SuffixOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_suffixOp

	return p
}

func (s *SuffixOpContext) GetParser() antlr.Parser { return s.parser }

func (s *SuffixOpContext) CopyAll(ctx *SuffixOpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SuffixOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuffixOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CallExprContext struct {
	SuffixOpContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	InitEmptySuffixOpContext(&p.SuffixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*SuffixOpContext))

	return p
}

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

type PropertyAccessExprContext struct {
	SuffixOpContext
}

func NewPropertyAccessExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyAccessExprContext {
	var p = new(PropertyAccessExprContext)

	InitEmptySuffixOpContext(&p.SuffixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*SuffixOpContext))

	return p
}

func (s *PropertyAccessExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyAccessExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *PropertyAccessExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterPropertyAccessExpr(s)
	}
}

func (s *PropertyAccessExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitPropertyAccessExpr(s)
	}
}

type IndexExprContext struct {
	SuffixOpContext
}

func NewIndexExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExprContext {
	var p = new(IndexExprContext)

	InitEmptySuffixOpContext(&p.SuffixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*SuffixOpContext))

	return p
}

func (s *IndexExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExprContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *IndexExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterIndexExpr(s)
	}
}

func (s *IndexExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitIndexExpr(s)
	}
}

type MethodCallExprContext struct {
	SuffixOpContext
}

func NewMethodCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethodCallExprContext {
	var p = new(MethodCallExprContext)

	InitEmptySuffixOpContext(&p.SuffixOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*SuffixOpContext))

	return p
}

func (s *MethodCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *MethodCallExprContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *MethodCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterMethodCallExpr(s)
	}
}

func (s *MethodCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitMethodCallExpr(s)
	}
}

func (p *CompiscriptParser) SuffixOp() (localctx ISuffixOpContext) {
	localctx = NewSuffixOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, CompiscriptParserRULE_suffixOp)
	var _la int

	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMethodCallExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(513)
			p.Match(CompiscriptParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(514)
			p.Match(CompiscriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(515)
			p.Match(CompiscriptParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(517)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&328630831402732032) != 0 {
			{
				p.SetState(516)
				p.Arguments()
			}

		}
		{
			p.SetState(519)
			p.Match(CompiscriptParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewCallExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(520)
			p.Match(CompiscriptParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(522)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&328630831402732032) != 0 {
			{
				p.SetState(521)
				p.Arguments()
			}

		}
		{
			p.SetState(524)
			p.Match(CompiscriptParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewIndexExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(525)
			p.Match(CompiscriptParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(526)
			p.ConditionalExpr()
		}
		{
			p.SetState(527)
			p.Match(CompiscriptParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewPropertyAccessExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(529)
			p.Match(CompiscriptParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(530)
			p.Match(CompiscriptParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConditionalExpr() []IConditionalExprContext
	ConditionalExpr(i int) IConditionalExprContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllConditionalExpr() []IConditionalExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionalExprContext); ok {
			len++
		}
	}

	tst := make([]IConditionalExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionalExprContext); ok {
			tst[i] = t.(IConditionalExprContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) ConditionalExpr(i int) IConditionalExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *CompiscriptParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, CompiscriptParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(533)
		p.ConditionalExpr()
	}
	p.SetState(538)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CompiscriptParserT__30 {
		{
			p.SetState(534)
			p.Match(CompiscriptParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(535)
			p.ConditionalExpr()
		}

		p.SetState(540)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConditionalExpr() []IConditionalExprContext
	ConditionalExpr(i int) IConditionalExprContext

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) AllConditionalExpr() []IConditionalExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionalExprContext); ok {
			len++
		}
	}

	tst := make([]IConditionalExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionalExprContext); ok {
			tst[i] = t.(IConditionalExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayLiteralContext) ConditionalExpr(i int) IConditionalExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (p *CompiscriptParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, CompiscriptParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(541)
		p.Match(CompiscriptParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(550)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&328630831402732032) != 0 {
		{
			p.SetState(542)
			p.ConditionalExpr()
		}
		p.SetState(547)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CompiscriptParserT__30 {
			{
				p.SetState(543)
				p.Match(CompiscriptParserT__30)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(544)
				p.ConditionalExpr()
			}

			p.SetState(549)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(552)
		p.Match(CompiscriptParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BaseType() IBaseTypeContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) BaseType() IBaseTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *CompiscriptParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, CompiscriptParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(554)
		p.BaseType()
	}
	p.SetState(559)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CompiscriptParserT__10 {
		{
			p.SetState(555)
			p.Match(CompiscriptParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(556)
			p.Match(CompiscriptParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(561)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBaseTypeContext is an interface to support dynamic dispatch.
type IBaseTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsBaseTypeContext differentiates from other interfaces.
	IsBaseTypeContext()
}

type BaseTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseTypeContext() *BaseTypeContext {
	var p = new(BaseTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_baseType
	return p
}

func InitEmptyBaseTypeContext(p *BaseTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CompiscriptParserRULE_baseType
}

func (*BaseTypeContext) IsBaseTypeContext() {}

func NewBaseTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseTypeContext {
	var p = new(BaseTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompiscriptParserRULE_baseType

	return p
}

func (s *BaseTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseTypeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CompiscriptParserIdentifier, 0)
}

func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.EnterBaseType(s)
	}
}

func (s *BaseTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompiscriptListener); ok {
		listenerT.ExitBaseType(s)
	}
}

func (p *CompiscriptParser) BaseType() (localctx IBaseTypeContext) {
	localctx = NewBaseTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, CompiscriptParserRULE_baseType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(562)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&319755573543305216) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
