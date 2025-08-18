package listener

import (
	"fmt"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

type TypeIdentifier string

type Listener struct {
	*p.BaseCompiscriptListener
	KnownTypes   *map[TypeIdentifier]TypeInfo
	Errors       *[]string
	ScopeManager *ScopeManager
}

func NewListener() Listener {
	baseTypes := make(map[TypeIdentifier]TypeInfo)
	baseTypes[TypeIdentifier(BASE_TYPES.INTEGER)] = NewTypeInfo_Base()
	baseTypes[TypeIdentifier(BASE_TYPES.BOOLEAN)] = NewTypeInfo_Base()
	baseTypes[TypeIdentifier(BASE_TYPES.STRING)] = NewTypeInfo_Base()
	baseTypes[TypeIdentifier(BASE_TYPES.NULL)] = NewTypeInfo_Base()
	baseTypes[TypeIdentifier(BASE_TYPES.UNKNOWN)] = NewTypeInfo_Base()

	errors := []string{}
	currentScope := NewScope("GLOBAL", SCOPE_TYPES.GLOBAL)
	scopeManager := NewScopeManager(currentScope, currentScope)

	return Listener{
		KnownTypes:   &baseTypes,
		Errors:       &errors,
		ScopeManager: &scopeManager,
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

func (l Listener) HasErrors() bool {
	return len(*l.Errors) > 0
}

func (l Listener) TypeExists(identifier TypeIdentifier) bool {
	_, found := l.GetTypeInfo(identifier)
	return found
}
