package semantic

import "fmt"

const (
  INTEGER = iota
  STRING
  BOOLEAN
  NULL
)

type SymbolTable struct {
  variables map[string]Content
}

type Content struct {
  dataType int 
  class string
  scope string
}

func NewSymbolTable() *SymbolTable {
  return &SymbolTable{
    variables: make(map[string]Content),
  }
}

func (st *SymbolTable) Add(name string, content Content) error {
  if _, exists := st.variables[name]; exists {
    return fmt.Errorf("variable '%s' already declared", name)
  }
  st.variables[name] = content
  return nil
}

func (st *SymbolTable) Get(name string) (Content, bool) {
  val, exists := st.variables[name]
  return val, exists
}
