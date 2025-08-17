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

type TypeIdentifier string

type Listener struct {
	*p.BaseCompiscriptListener
	KnownTypes        *map[TypeIdentifier]TypeInfo
	TypesByExpression *map[string]TypeIdentifier
	Errors            *[]string
	Scopes            *lib.Stack[Scope]
}

func NewListener() Listener {
	baseTypes := make(map[TypeIdentifier]TypeInfo)
	baseTypes[TypeIdentifier(BASE_TYPES.INTEGER)] = NewTypeInfo_Base()
	baseTypes[TypeIdentifier(BASE_TYPES.BOOLEAN)] = NewTypeInfo_Base()
	baseTypes[TypeIdentifier(BASE_TYPES.STRING)] = NewTypeInfo_Base()
	baseTypes[TypeIdentifier(BASE_TYPES.NULL)] = NewTypeInfo_Base()

	errors := []string{}
	scopes := lib.NewStack[Scope]()
	typesByExpr := make(map[string]TypeIdentifier)

	return Listener{
		KnownTypes:        &baseTypes,
		Errors:            &errors,
		TypesByExpression: &typesByExpr,
		Scopes:            &scopes,
	}
}

func (l Listener) AddTypeInfo(identifier TypeIdentifier, info TypeInfo) {
	(*l.KnownTypes)[identifier] = info
}

func (l Listener) GetTypeInfo(identifier TypeIdentifier) (TypeInfo, bool) {
	info, found := (*l.KnownTypes)[identifier]
	return info, found
}

func (l Listener) ModifyClassTypeInfo(identifier TypeIdentifier, exe func(*ClassTypeInfo)) {
	info := (*l.KnownTypes)[identifier]
	classInfo := info.ClassType.GetValue()
	exe(&classInfo)
	(*l.KnownTypes)[identifier] = NewTypeInfo_Class(classInfo)
}

func (l Listener) AddError(content string) {
	*l.Errors = append(*l.Errors, fmt.Sprintf("ERROR: %s", content))
}

func (l Listener) AddWarning(content string) {
	*l.Errors = append(*l.Errors, fmt.Sprintf("WARNING: %s", content))
}

func (l Listener) AddTypeByExpr(expr string, t TypeIdentifier) {
	(*l.TypesByExpression)[expr] = t
}

func (l Listener) HasErrors() bool {
	return len(*l.Errors) > 0
}

func (l Listener) TypeExists(identifier TypeIdentifier) bool {
	_, found := l.GetTypeInfo(identifier)
	return found
}
