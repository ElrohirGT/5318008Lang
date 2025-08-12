package listener

import (
	"fmt"

	"github.com/ElrohirGT/5318008Lang/lib"
	p "github.com/ElrohirGT/5318008Lang/parser"
)

type DefinitionInfo struct {
	Name string
	Type string
}

var SCOPE_TYPES = struct {
	GLOBAL   string
	CLASS    string
	FUNCTION string
	BLOCK    string
}{
	GLOBAL:   "GLOBAL",
	CLASS:    "CLASS",
	FUNCTION: "FUNCTION",
	BLOCK:    "BLOCK",
}

type Scope struct {
	// Obtained from SCOPE_TYPES
	Type        string
	Name        string
	Definitions map[string]DefinitionInfo
}

func NewScope(name string, _type string) Scope {
	return Scope{
		Type:        _type,
		Name:        name,
		Definitions: map[string]DefinitionInfo{},
	}
}

type Listener struct {
	*p.BaseCompiscriptListener
	KnownTypes        lib.Set[string]
	TypesByExpression *map[string]string
	Errors            *[]string
	Scopes            *lib.Stack[Scope]
}

func NewListener() Listener {
	baseTypes := lib.NewSet[string]()
	baseTypes.Add(BASE_TYPES.INTEGER)
	baseTypes.Add(BASE_TYPES.FLOAT)
	baseTypes.Add(BASE_TYPES.BOOLEAN)
	baseTypes.Add(BASE_TYPES.STRING)
	baseTypes.Add(BASE_TYPES.NULL)

	errors := []string{}
	scopes := lib.NewStack[Scope]()
	typesByExpr := make(map[string]string)

	return Listener{
		KnownTypes:        baseTypes,
		Errors:            &errors,
		TypesByExpression: &typesByExpr,
		Scopes:            &scopes,
	}
}

func (l Listener) AddError(content string) {
	*l.Errors = append(*l.Errors, fmt.Sprintf("ERROR: %s", content))
}

func (l Listener) AddWarning(content string) {
	*l.Errors = append(*l.Errors, fmt.Sprintf("WARNING: %s", content))
}

func (l Listener) AddTypeByExpr(expr string, t string) {
	(*l.TypesByExpression)[expr] = t
}

func (l Listener) HasErrors() bool {
	return len(*l.Errors) > 0
}
