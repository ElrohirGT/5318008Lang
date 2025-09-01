package listener

import (
	"fmt"
	"log"
	"strings"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

// Typing system scanner, responsable of the semantic during compiscript code.
// Handles the notion of types, definitions and scope management.
type Listener struct {
	*p.BaseCompiscriptListener
	KnownTypes   *TypeTable
	ScopeManager *ScopeManager
	Errors       *[]string
}

func NewListener() Listener {
	baseTypes := make(TypeTable)
	baseTypes[TypeIdentifier(BASE_TYPES.INTEGER)] = NewTypeInfo_Base()
	baseTypes[TypeIdentifier(BASE_TYPES.BOOLEAN)] = NewTypeInfo_Base()
	baseTypes[TypeIdentifier(BASE_TYPES.STRING)] = NewTypeInfo_Base()
	baseTypes[TypeIdentifier(BASE_TYPES.NULL)] = NewTypeInfo_Base()
	baseTypes[TypeIdentifier(BASE_TYPES.UNKNOWN)] = NewTypeInfo_Base()
	baseTypes[TypeIdentifier(BASE_TYPES.INVALID)] = NewTypeInfo_Base()

	for _, v := range BASE_TYPE_ARRAY {
		baseTypes[v] = NewTypeInfo_Base()
		baseTypes[NewArrayTypeIdentifier(v)] = NewTypeInfo_Base()
	}

	errors := []string{}
	currentScope := NewScope("GLOBAL", SCOPE_TYPES.GLOBAL)
	scopeManager := NewScopeManager(currentScope, currentScope)

	// ADD BUILTINS
	scopeManager.CurrentScope.UpsertFunctionDef("print",
		MethodInfo{ParameterList: []ParameterInfo{{"s", BASE_TYPES.STRING}}, ReturnType: BASE_TYPES.NULL})
	scopeManager.CurrentScope.UpsertFunctionDef("parseInt",
		MethodInfo{ParameterList: []ParameterInfo{{"v", BASE_TYPES.INTEGER}}, ReturnType: BASE_TYPES.STRING})
	scopeManager.CurrentScope.UpsertFunctionDef("parseBool",
		MethodInfo{ParameterList: []ParameterInfo{{"v", BASE_TYPES.BOOLEAN}}, ReturnType: BASE_TYPES.STRING})

	// FIXME: A function should not be registered in tye typesExpresion register
	// scopeManager.CurrentScope.UpsertExpressionType("print", BASE_TYPES.NULL)
	// scopeManager.CurrentScope.UpsertExpressionType("parseInt", BASE_TYPES.STRING)
	// scopeManager.CurrentScope.UpsertExpressionType("parseBool", BASE_TYPES.STRING)

	// DEBUG: Verify functions are registered
	log.Printf("Registered functions in GLOBAL scope: %v", scopeManager.CurrentScope.functions)
	for name, info := range scopeManager.CurrentScope.functions {
		log.Printf("Function: %s, ReturnType: %s, Params: %v", name, info.ReturnType, info.ParameterList)
	}

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

func (l Listener) TypeExists(identifier TypeIdentifier) bool {
	_, found := l.GetTypeInfo(identifier)
	return found
}

// ERROR LOGGING
// ====================

// FIXME: Improve error handling:
// GOAL: Achieve something similar to ELM errors (whatever it takes).
func (l Listener) AddError(line int, columnStart int, columnEnd int, content string, details ...string) {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(Red+"* Error: (line: %d, column: %d-%d) %s"+Reset, line, columnStart, columnEnd, content))

	for _, v := range details {
		b.WriteString("\n * " + v)
	}

	*l.Errors = append(*l.Errors, b.String())
}

func (l Listener) AddWarning(content string, line string, details ...string) {
	*l.Errors = append(*l.Errors, "Warning: "+content)
}

func (l Listener) HasErrors() bool {
	return len(*l.Errors) > 0
}
