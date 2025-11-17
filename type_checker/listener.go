package type_checker

import (
	"fmt"
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
	Literals     *LiteralTable
	Errors       *[]string
}

const GLOBAL_SCOPE_NAME = "GLOBAL"

func NewListener() Listener {
	literals := make(LiteralTable)

	baseTypes := make(TypeTable)
	for _, baseType := range BASE_TYPE_ARRAY {
		typeSize := uint(4) // 4 bytes is the base type size except on booleans
		if baseType == BASE_TYPES.BOOLEAN {
			typeSize = uint(1)
		}

		baseTypes[baseType] = NewTypeInfo_Base(typeSize)
		if baseType == BASE_TYPES.INTEGER || baseType == BASE_TYPES.BOOLEAN || baseType == BASE_TYPES.STRING {
			// string[]
			baseTypes[NewArrayTypeIdentifier(baseType)] = NewTypeInfo_Array(ArrayTypeInfo{Type: baseType})
		}
	}

	errors := []string{}
	currentScope := NewScope(GLOBAL_SCOPE_NAME, SCOPE_TYPES.GLOBAL)
	scopeManager := NewScopeManager(currentScope, currentScope)

	// ADD BUILTINS
	scopeManager.CurrentScope.UpsertFunctionDef("print",
		MethodInfo{ParameterList: []ParameterInfo{{"s", BASE_TYPES.STRING}}, ReturnType: BASE_TYPES.NULL})
	// FIXME: Rename this builtins, confusing AF!!!
	scopeManager.CurrentScope.UpsertFunctionDef("int_to_str",
		MethodInfo{ParameterList: []ParameterInfo{{"v", BASE_TYPES.INTEGER}}, ReturnType: BASE_TYPES.STRING})
	scopeManager.CurrentScope.UpsertFunctionDef("bool_to_str",
		MethodInfo{ParameterList: []ParameterInfo{{"v", BASE_TYPES.BOOLEAN}}, ReturnType: BASE_TYPES.STRING})
	scopeManager.CurrentScope.UpsertFunctionDef("strcmp",
		MethodInfo{ParameterList: []ParameterInfo{{"v", BASE_TYPES.STRING}, {"v2", BASE_TYPES.STRING}}, ReturnType: BASE_TYPES.BOOLEAN})

	// FIXME: A function should not be registered in tye typesExpresion register
	// scopeManager.CurrentScope.UpsertExpressionType("print", BASE_TYPES.NULL)
	// scopeManager.CurrentScope.UpsertExpressionType("int_to_str", BASE_TYPES.STRING)
	// scopeManager.CurrentScope.UpsertExpressionType("bool_to_str", BASE_TYPES.STRING)

	return Listener{
		KnownTypes:   &baseTypes,
		Errors:       &errors,
		ScopeManager: &scopeManager,
		Literals:     &literals,
	}
}

func (l Listener) UpsertTypeInfo(identifier TypeIdentifier, info TypeInfo) {
	(*l.KnownTypes)[identifier] = info
}

func (l Listener) UpsertLiteral(literalExpr string, id TypeIdentifier) {
	(*l.Literals)[literalExpr] = id
}

func (l Listener) GetLiteralType(literalExpr string) (TypeIdentifier, bool) {
	id, found := (*l.Literals)[literalExpr]
	return id, found
}

func (l Listener) CheckInheritanceTree(ctx *p.ClassDeclarationContext, baseClass TypeIdentifier, fatherClass TypeIdentifier) bool {
	if baseClass == fatherClass {
		l.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			ctx.GetStop().GetColumn(),
			"Can't declare a class that inherits from itself!",
		)
		return false
	}

	fatherInfo, found := l.GetTypeInfo(fatherClass)
	if !found {
		l.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			ctx.GetStop().GetColumn(),
			fmt.Sprintf("Can't inherit from an undefined class `%s`!", fatherClass),
		)
		return false
	}

	if !fatherInfo.ClassType.HasValue() {
		l.AddError(
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn(),
			ctx.GetStop().GetColumn(),
			"Can't inherit from a non class type!",
		)
		return false
	}

	if fatherInfo.ClassType.GetValue().InheritsFrom != TypeIdentifier("") {
		return l.CheckInheritanceTree(ctx, fatherClass, fatherInfo.ClassType.GetValue().InheritsFrom)
	} else {
		return true
	}
}

func (l Listener) GetTypeInfo(identifier TypeIdentifier) (TypeInfo, bool) {
	info, found := (*l.KnownTypes)[identifier]
	return info, found
}

func (l Listener) ModifyClassTypeInfo(identifier TypeIdentifier, exe func(*ClassTypeInfo)) {
	info := (*l.KnownTypes)[identifier]
	classInfo := info.ClassType.GetValue()
	exe(&classInfo)
	(*l.KnownTypes)[identifier] = NewTypeInfo_Class(classInfo, info.Size)
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
