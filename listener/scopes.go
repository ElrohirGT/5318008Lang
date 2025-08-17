package listener

import (
	"github.com/ElrohirGT/5318008Lang/lib"
)

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

type FunctionInfo struct {
	MethodInfo
}

type Scope struct {
	Children []*Scope
	Father   *Scope

	Type              string
	Name              string
	definitions       map[string]FunctionInfo
	typesByExpression map[string]TypeIdentifier
	constants         lib.Set[string]
}

type ScopeManager struct {
	// Current scope at time of writing
	CurrentScope *Scope
	// Quick reference to the global scope
	GlobaScope *Scope
}

func NewScopeManager(current *Scope, globalScope *Scope) ScopeManager {
	return ScopeManager{
		CurrentScope: current,
		GlobaScope:   globalScope,
	}
}

func (sc *ScopeManager) AddToCurrent(child *Scope) {
	sc.CurrentScope.AddChildScope(child)
}

func (sc *ScopeManager) ReplaceCurrent(newScope *Scope) {
	sc.CurrentScope = newScope
}

func NewScope(name string, _type string) *Scope {
	return &Scope{
		Type:              _type,
		Name:              name,
		definitions:       map[string]FunctionInfo{},
		typesByExpression: map[string]TypeIdentifier{},
		constants:         lib.NewSet[string](),
	}
}

func (s *Scope) AddChildScope(childScope *Scope) {
	s.Children = append(s.Children, childScope)
	childScope.Father = s
}

func (s *Scope) AddExpressionType(expr string, _type TypeIdentifier) {
	s.typesByExpression[expr] = _type
}

func (s *Scope) GetExpressionType(expr string) (TypeIdentifier, bool) {
	t, found := s.typesByExpression[expr]
	if !found && s.Father != nil {
		return s.Father.GetExpressionType(expr)
	}

	return t, found
}

// Returns true if the constant is a new constant, otherwise it returns false
func (s *Scope) AddConstant(exprName string) bool {
	return s.constants.Add(exprName)
}
