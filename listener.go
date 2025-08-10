package main

import (
	"fmt"
	"log"

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
	baseTypes.Add("integer")
	baseTypes.Add("float")
	baseTypes.Add("boolean")
	baseTypes.Add("string")
	baseTypes.Add("null")

	errors := []string{}
	scopes := lib.NewStack[Scope]()
	typesByExpr := map[string]string{}

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

func (l Listener) EnterClassDeclaration(ctx *p.ClassDeclarationContext) {
	identifiers := ctx.AllIdentifier()
	className := identifiers[0]

	log.Println("Declaring", className)
	l.KnownTypes.Add(className.GetText())
	l.Scopes.Push(NewScope(fmt.Sprintf("Scope_%s", className), SCOPE_TYPES.CLASS))
}

func (l Listener) EnterNewExpr(ctx *p.NewExprContext) {
	className := ctx.Identifier()
	log.Println("Instantiating class", className.GetText())

	// FIXME: We assume the constructor is called correctly!
	expr := ctx.GetText()
	exprType := className.GetText()
	log.Println("Adding", expr, "as an expresion of type", exprType)
	(*l.TypesByExpression)[expr] = exprType
}

func (l Listener) ExitClassDeclaration(ctx *p.ClassDeclarationContext) {
	identifiers := ctx.AllIdentifier()
	className := identifiers[0]

	log.Println("Leaving declaration for", className)
	if op := l.Scopes.Peek(); op.GetValue().Type != SCOPE_TYPES.CLASS {
		panic("Leaving class declaration but last scope is not of type class!")
	}

	l.Scopes.Pop()
}

func (l Listener) EnterVariableDeclaration(ctx *p.VariableDeclarationContext) {
	name := ctx.Identifier()

	typeAnnot := ctx.TypeAnnotation()
	hasAnnotation := typeAnnot != nil

	declarationExpr := ctx.Initializer()
	hasInitialExpr := declarationExpr != nil

	if !hasAnnotation {
		log.Println("Variable", name.GetText(), "does NOT have a type! We need to infer it...")
		// TODO: Infer variable type and check if it exists
	} else {
		declarationType := typeAnnot.Type_().GetText()
		log.Println("Variable", name.GetText(), "has type", declarationType)

		if !l.KnownTypes.Exists(declarationType) {
			l.AddError(fmt.Sprintf("%s doesn't exist!", declarationType))
		}

		if hasInitialExpr {
			exprText := declarationExpr.Expression().GetText()
			log.Println("Known expressions", *l.TypesByExpression)
			initialExprType, exists := (*l.TypesByExpression)[exprText]
			if !exists {
				l.AddError(fmt.Sprintf("`%s` doesn't have a type!", exprText))
			}

			if initialExprType != declarationType {
				l.AddError(fmt.Sprintf("The declaration of `%s` specifies a type of `%s` but `%s` was given", name, declarationType, initialExprType))
			}
		}
	}
}

func (l Listener) HasErrors() bool {
	return len(*l.Errors) > 0
}
