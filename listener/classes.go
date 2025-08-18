package listener

import (
	"fmt"
	"log"
	"slices"

	"github.com/ElrohirGT/5318008Lang/lib"
	p "github.com/ElrohirGT/5318008Lang/parser"
)

const CONSTRUCTOR_NAME = "constructor"

type ArrayTypeInfo struct {
	Type   TypeIdentifier
	Length uint
}

type ParameterInfo struct {
	Name string
	Type TypeIdentifier
}

type MethodInfo struct {
	ParameterList []ParameterInfo
	ReturnType    TypeIdentifier
}

type ClassTypeInfo struct {
	Name         TypeIdentifier
	InheritsFrom TypeIdentifier
	Fields       map[string]TypeIdentifier
	Methods      map[string]MethodInfo
	Constructor  MethodInfo
}

func (c *ClassTypeInfo) UpsertField(name string, _type TypeIdentifier) {
	c.Fields[name] = _type
}

func (c *ClassTypeInfo) UpsertMethod(name string, info MethodInfo) {
	c.Methods[name] = info
}

func NewClassTypeInfo(className string) ClassTypeInfo {
	return ClassTypeInfo{
		Name:    TypeIdentifier(className),
		Fields:  make(map[string]TypeIdentifier),
		Methods: make(map[string]MethodInfo),
	}
}

type TypeInfo struct {
	BaseType  bool
	ArrayType lib.Optional[ArrayTypeInfo]
	ClassType lib.Optional[ClassTypeInfo]
}

func NewTypeInfo_Base() TypeInfo {
	return TypeInfo{
		BaseType:  true,
		ArrayType: lib.NewOpEmpty[ArrayTypeInfo](),
		ClassType: lib.NewOpEmpty[ClassTypeInfo](),
	}
}

func NewTypeInfo_Array(arrInfo ArrayTypeInfo) TypeInfo {
	return TypeInfo{
		BaseType:  false,
		ArrayType: lib.NewOpValue(arrInfo),
		ClassType: lib.NewOpEmpty[ClassTypeInfo](),
	}
}

func NewTypeInfo_Class(classInfo ClassTypeInfo) TypeInfo {
	return TypeInfo{
		BaseType:  false,
		ArrayType: lib.NewOpEmpty[ArrayTypeInfo](),
		ClassType: lib.NewOpValue(classInfo),
	}
}

func (l Listener) ExitClassDeclaration(ctx *p.ClassDeclarationContext) {
	if l.ScopeManager.CurrentScope.Type != SCOPE_TYPES.CLASS {
		panic("Trying to exit a class declaration but the scope is not of type class!")
	}

	log.Printf("Escaping class declaration: %s", ctx.AllIdentifier()[0].GetText())
	l.ScopeManager.CurrentScope = l.ScopeManager.CurrentScope.Father
}

func (l Listener) EnterClassDeclaration(ctx *p.ClassDeclarationContext) {
	identifiers := ctx.AllIdentifier()
	className := identifiers[0]
	log.Println("Declaring", className)
	line := ctx.GetStart().GetLine()

	onGlobaScope := l.ScopeManager.CurrentScope.Type == SCOPE_TYPES.GLOBAL
	if !onGlobaScope {
		l.AddError(fmt.Sprintf(
			"(line: %d) Can't define class `%s` inside scope! Classes can only be defined on global scope!",
			line,
			className.GetText(),
		))
	}

	classScope := NewScope(className.GetText(), SCOPE_TYPES.CLASS)
	l.ScopeManager.AddToCurrent(classScope)
	l.ScopeManager.ReplaceCurrent(classScope)
	// Add this as a valid expression with type of this class
	classScope.UpsertExpressionType("this", TypeIdentifier(className.GetText()))

	if slices.Contains(BASE_TYPE_ARRAY, TypeIdentifier(className.GetText())) {
		l.AddError(fmt.Sprintf(
			"(line: %d) Can't define class `%s` because it collides with a builtin type!",
			line,
			className,
		))
		return
	}

	if _, found := l.GetTypeInfo(TypeIdentifier(className.GetText())); found {
		l.AddError(fmt.Sprintf(
			"(line: %d) Can't redefine existing class! `%s` already exists!",
			line,
			className.GetText(),
		))
	} else {
		classInfo := NewClassTypeInfo(className.GetText())
		l.AddTypeInfo(TypeIdentifier(className.GetText()), NewTypeInfo_Class(classInfo))
	}

	if len(identifiers) > 1 {
		fatherClassName := identifiers[1]
		log.Println("Class inherits from", fatherClassName)
		info, found := l.GetTypeInfo(TypeIdentifier(fatherClassName.GetText()))
		if !found {
			l.AddError(fmt.Sprintf(
				"(line: %d) Can't inherit from a type that doesn't exists! `%s` wants to inherit from `%s`!",
				line,
				className.GetText(),
				fatherClassName.GetText(),
			))
		} else {
			if !info.ClassType.HasValue() {
				l.AddError(fmt.Sprintf(
					"(line: %d) Can't make a nonexistent class inherit from another! `%s` wants to inherit from `%s` but `%s` is not a class!",
					line,
					className.GetText(),
					fatherClassName.GetText(),
					className.GetText(),
				))
			} else {
				l.ModifyClassTypeInfo(TypeIdentifier(className.GetText()), func(classInfo *ClassTypeInfo) {
					classInfo.InheritsFrom = TypeIdentifier(fatherClassName.GetText())
				})
			}
		}
	}
}

func (l Listener) ExitNewExpr(ctx *p.NewExprContext) {
	line := ctx.GetStart().GetLine()
	className := ctx.Identifier()
	log.Println("Trying to instantiate class:", className.GetText())

	exprArguments := []p.IExpressionContext{}
	if ctx.Arguments() != nil {
		exprArguments = ctx.Arguments().AllExpression()
	}

	classScope, found := l.ScopeManager.SearchClassScope()

	if found && classScope.Name == className.GetText() {
		l.AddError(fmt.Sprintf(
			"(line: %d) Can't create a new instance of `%s` while defining the same class!",
			line,
			className,
		))
		return
	}

	typeInfo, found := l.GetTypeInfo(TypeIdentifier(className.GetText()))
	if !found || !typeInfo.ClassType.HasValue() {
		l.AddError(fmt.Sprintf(
			"(line: %d) Can't create a new instance of undefined class `%s`",
			line,
			className.GetText(),
		))
		return
	}

	methodInfo := typeInfo.ClassType.GetValue().Constructor
	if len(methodInfo.ParameterList) != len(exprArguments) {
		l.AddError(fmt.Sprintf(
			"(line: %d) Constructor of `%s` asks for `%d` arguments but `%d` given!",
			line,
			className,
			len(methodInfo.ParameterList),
			len(exprArguments),
		))
		return
	}

	errorWithParams := false
	for i, consParam := range methodInfo.ParameterList {
		exprParam := exprArguments[i]
		exprType, found := l.ScopeManager.CurrentScope.GetExpressionType(exprParam.GetText())
		if !found {
			l.AddError(fmt.Sprintf(
				"(line: %d) Type of `%s` not found!",
				line,
				exprParam.GetText(),
			))
			errorWithParams = true
		}

		if consParam.Type != exprType {
			l.AddError(fmt.Sprintf(
				"(line: %d) Constructor for `%s` requires `%s` to be of type `%s` but it's `%s` instead!",
				line,
				className,
				exprParam.GetText(),
				consParam.Type,
				exprType,
			))
			errorWithParams = true
		}
	}

	if errorWithParams {
		return
	}

	expr := ctx.GetText()
	exprType := className.GetText()
	log.Println("Adding", expr, "as an expression of type", exprType)
	l.ScopeManager.CurrentScope.UpsertExpressionType(expr, TypeIdentifier(exprType))
}
