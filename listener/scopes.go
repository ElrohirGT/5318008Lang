package listener

import (
	"github.com/ElrohirGT/5318008Lang/lib"
)

type ScopeType string

// Available types of scopes the compiler can handle.
var SCOPE_TYPES = struct {
	GLOBAL   ScopeType
	CLASS    ScopeType
	FUNCTION ScopeType
	BLOCK    ScopeType
}{
	GLOBAL:   "GLOBAL",
	CLASS:    "CLASS",
	FUNCTION: "FUNCTION",
	BLOCK:    "BLOCK",
}

// Representation of a scope: an isolated environment that stores
//   - local variables
//   - method definitions
//   - expresion types
//
// Scopes follow a tree structure, when looking for the existence of a definition, we start
// from the current scope all the way up until the global scope :
//
//	   GLOBAL
//	 ┌────────┐
//	 ▼        ▼
//	FUN1    CLASS1
//	          ┌──────┐
//	          ▼      ▼
//	        FUN3   FUN3
type Scope struct {
	Children []*Scope
	Father   *Scope

	Type              ScopeType
	Name              string
	definitions       map[string]MethodInfo
	typesByExpression map[string]TypeIdentifier
	constants         lib.Set[string]
}

// Manges a scope tree. Providing helpful function to handle and move arount the tree.
type ScopeManager struct {
	// Current scope at time of writing
	CurrentScope *Scope
	// Quick reference to the global scope
	// In case anyone needs it
	GlobaScope *Scope
}

// =======================
// SCOPE MANAGER
// =======================

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

func (sc *ScopeManager) SearchClassScope() (*Scope, bool) {
	if sc.CurrentScope.Type == SCOPE_TYPES.CLASS {
		return sc.CurrentScope, true
	}

	if sc.CurrentScope.Father != nil {
		return sc.CurrentScope.Father.SearchClassScope()
	}

	return nil, false
}

// =======================
// SCOPES
// =======================

func NewScope(name string, _type ScopeType) *Scope {
	return &Scope{
		Type:              _type,
		Name:              name,
		definitions:       map[string]MethodInfo{},
		typesByExpression: map[string]TypeIdentifier{},
		constants:         lib.NewSet[string](),
	}
}

func (s *Scope) AddChildScope(childScope *Scope) {
	s.Children = append(s.Children, childScope)
	childScope.Father = s
}

func (s *Scope) UpsertExpressionType(expr string, _type TypeIdentifier) {
	s.typesByExpression[expr] = _type
}

func (s *Scope) GetExpressionType(expr string) (TypeIdentifier, bool) {
	t, found := s.typesByExpression[expr]
	if !found && s.Father != nil {
		return s.Father.GetExpressionType(expr)
	}

	return t, found
}

func (s *Scope) ContainsExpression(expr string) bool {
	_, found := s.typesByExpression[expr]
	return found
}

// Returns true if the constant is a new constant, otherwise it returns false.
func (s *Scope) AddConstant(exprName string) bool {
	return s.constants.Add(exprName)
}

func (s *Scope) SearchClassScope() (*Scope, bool) {
	if s.Type == SCOPE_TYPES.CLASS {
		return s, true
	}

	if s.Father != nil {
		return s.Father.SearchClassScope()
	}

	return nil, false
}

func (s *Scope) UpsertFunctionDef(funcName string, info MethodInfo) {
	s.definitions[funcName] = info
}
