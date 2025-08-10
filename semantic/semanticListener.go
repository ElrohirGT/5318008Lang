package semantic

import (
	"fmt"
	"strconv"

	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/antlr4-go/antlr/v4"
)

type SemanticListener struct {
  *p.BaseCompiscriptListener
  symbols *SymbolTable
  Errors []string
}

func NewSemanticListener(symbols *SymbolTable) *SemanticListener {
  return &SemanticListener{
    symbols: symbols,
    Errors: []string{},
  }
}

func (l *SemanticListener) ExitIfStatement(ctx *p.IfStatementContext) {
    expr := ctx.Expression()
    exprType := l.getTypeFromNode(expr)
    if exprType != BOOLEAN {
        l.Errors = append(l.Errors, "The if condition has to be boolean")
    }
}

func (l *SemanticListener) getTypeFromNode(node antlr.Tree) int {
    if leaf, ok := node.(antlr.TerminalNode); ok {
        text := leaf.GetText()

        // Literal boolean
        if text == "true" || text == "false" {
            return BOOLEAN
        }

        // Literal integer
        if _, err := strconv.Atoi(text); err == nil {
            return INTEGER
        }

        // Variable
        if tipo, ok := l.symbols.GetType(text); ok {
            return tipo
        }

        return NULL
    }

    if node.GetChildCount() == 3 {
        leftType := l.getTypeFromNode(node.GetChild(0))
        var opText string
        if opNode, ok := node.GetChild(1).(antlr.TerminalNode); ok {
            opText = opNode.GetText()
        } else {
            return l.getTypeFromNode(node.GetChild(1))
        }
        rightType := l.getTypeFromNode(node.GetChild(2))

        if opText == "==" || opText == "!=" || opText == "<" || opText == "<=" || opText == ">" || opText == ">=" {
            if leftType != rightType {
                l.Errors = append(l.Errors,
                    fmt.Sprintf("Operands of '%s' must be same type", opText))
                return NULL
            }
            if (opText == "<" || opText == "<=" || opText == ">" || opText == ">=") &&
                leftType != INTEGER {
                l.Errors = append(l.Errors,
                    fmt.Sprintf("Operator '%s' only allowed on integers", opText))
                return NULL
            }
            return BOOLEAN
        }
    }

    for i := 0; i < node.GetChildCount(); i++ {
        t := l.getTypeFromNode(node.GetChild(i))
        if t != NULL {
             return t
        }
    }

    return NULL
}
