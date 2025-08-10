package semantic

import "fmt"

const (
  INTEGER = iota
  STRING
  BOOLEAN
  NULL
)

type SymbolTable struct {
  Variables map[string]Content
}

type Content struct {
  DataType int 
  Class string
  Scope string
}

func NewSymbolTable() *SymbolTable {
  return &SymbolTable{
    Variables: make(map[string]Content),
  }
}

func (st *SymbolTable) Add(name string, content Content) error {
  if _, exists := st.Variables[name]; exists {
    return fmt.Errorf("variable '%s' already declared", name)
  }
  st.Variables[name] = content
  return nil
}

func (st *SymbolTable) Get(name string) (Content, bool) {
  val, exists := st.Variables[name]
  return val, exists
}

func (st *SymbolTable) GetType(name string) (int, bool) {
  t, ok := st.Variables[name]
  return t.DataType, ok
}
