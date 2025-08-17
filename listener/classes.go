package listener

import (
	"fmt"
	"log"

	"github.com/ElrohirGT/5318008Lang/lib"
	p "github.com/ElrohirGT/5318008Lang/parser"
)

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

func (l Listener) EnterClassDeclaration(ctx *p.ClassDeclarationContext) {
	identifiers := ctx.AllIdentifier()
	className := identifiers[0]
	log.Println("Declaring", className)
	line := ctx.GetStart().GetLine()

	// FIXME: Check if we're currently on the global scope, if not throw error!
	onGlobaScope := true
	if !onGlobaScope {
		l.AddError(fmt.Sprintf(
			"(line: %d) Can't define class `%s` inside scope! Classes can only be defined on global scope!",
			line,
			className.GetText(),
		))
	}

	if _, found := l.GetTypeInfo(TypeIdentifier(className.GetText())); found {
		l.AddError(fmt.Sprintf("(line: %d) Can't redefine existing class! `%s` already exists!", line, className.GetText()))
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

func (l Listener) EnterNewExpr(ctx *p.NewExprContext) {
	className := ctx.Identifier()
	log.Println("Instantiating class", className.GetText())

	// FIXME: We assume the constructor is called correctly!
	expr := ctx.GetText()
	exprType := className.GetText()
	log.Println("Adding", expr, "as an expresion of type", exprType)
	l.AddTypeByExpr(expr, TypeIdentifier(exprType))
}
