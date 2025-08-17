package listener

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
}

func NewScope(name string, _type string) Scope {
	return Scope{
		Type:              _type,
		Name:              name,
		definitions:       map[string]FunctionInfo{},
		typesByExpression: map[string]TypeIdentifier{},
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
