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

func (c *ClassTypeInfo) UpsertField(name string, _type TypeIdentifier) {
	c.Fields[name] = _type
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

func (s Listener) ExitPropertyAssignExpr(ctx *p.PropertyAssignExprContext) {
	// FIXME: Where the fok is this used!?
	log.Println("Property Assignment!", ctx.GetText())
}

func (l Listener) ExitAssignment(ctx *p.AssignmentContext) {
	log.Println("General Assignment!", ctx.GetText())
	line := ctx.GetStart().GetLine()

	isPropertyAssignment := len(ctx.AllExpression()) > 1
	classScope, _ := l.ScopeManager.SearchClassScope()
	if isPropertyAssignment {
		firstExpr := ctx.Expression(0)
		t, found := l.ScopeManager.CurrentScope.GetExpressionType(firstExpr.GetText())
		if !found {
			l.AddError(fmt.Sprintf(
				"(line: %d) Undeclared variable `%s`",
				line,
				firstExpr.GetText(),
			))
			return
		}

		info, found := l.GetTypeInfo(t)
		if !found {
			l.AddError(fmt.Sprintf(
				"(line: %d) Undeclared type `%s` for variable `%s`",
				line,
				t,
				firstExpr.GetText(),
			))
			return
		}

		if !info.ClassType.HasValue() {
			l.AddError(fmt.Sprintf(
				"(line: %d) Trying to access a field `%s` from type `%s` but `%s` is not a class!",
				line,
				firstExpr.GetText(),
				t,
				firstExpr.GetText(),
			))
			return
		}

		identifier := ctx.Identifier()
		classInfo := info.ClassType.GetValue()
		fieldType, hasField := classInfo.Fields[identifier.GetText()]
		if !hasField {
			l.AddError(fmt.Sprintf(
				"(line: %d) Trying to access field `%s` not defined in class `%s`!",
				line,
				identifier.GetText(),
				classInfo.Name,
			))
			return
		}

		assignExpr := ctx.Expression(1)
		assignType, found := l.ScopeManager.CurrentScope.GetExpressionType(assignExpr.GetText())
		if !found {
			l.AddError(fmt.Sprintf(
				"(line: %d) Type of expression `%s` not found!",
				line,
				assignExpr.GetText(),
			))
			return
		}

		if fieldType == BASE_TYPES.UNKNOWN && assignType == BASE_TYPES.UNKNOWN {
			l.AddError(fmt.Sprintf(
				"(line: %d) Trying to assign `%s` into `%s` but I don't know the types of both! Please give me hints!",
				line,
				assignExpr.GetText(),
				identifier.GetText(),
			))
			return
		}

		if fieldType == BASE_TYPES.UNKNOWN {
			log.Printf("Inferring `%s` as type `%s`", "this."+identifier.GetText(), assignType)
			classScope.UpsertExpressionType("this."+identifier.GetText(), assignType)
			l.ModifyClassTypeInfo(TypeIdentifier(classScope.Name), func(cti *ClassTypeInfo) {
				cti.UpsertField(identifier.GetText(), assignType)
			})
			fieldType = assignType
		}

		if fieldType != assignType {
			l.AddError(fmt.Sprintf(
				"(line: %d) Trying to assign `%s` to field `%s` but types don't match! (`%s` != `%s`)",
				line,
				assignExpr.GetText(),
				identifier.GetText(),
				fieldType,
				assignType,
			))
			return
		}
	} else {
		identifier := ctx.Identifier()
		identifierType, found := l.ScopeManager.CurrentScope.GetExpressionType(identifier.GetText())
		if !found {
			l.AddError(fmt.Sprintf(
				"(line: %d) Undeclared variable `%s`!",
				line,
				identifier.GetText(),
			))
			return
		}

		assignExpr := ctx.Expression(0)
		assignType, found := l.ScopeManager.CurrentScope.GetExpressionType(assignExpr.GetText())
		if !found {
			l.AddError(fmt.Sprintf(
				"(line: %d) Expression `%s` doesn't have a type!",
				line,
				assignExpr,
			))
		}

		if identifierType == BASE_TYPES.UNKNOWN && assignType == BASE_TYPES.UNKNOWN {
			l.AddError(fmt.Sprintf(
				"(line: %d) Can't assign `%s` = `%s` because both type are unknown! Use one of them first or write type hints!",
				line,
				identifier.GetText(),
				assignExpr.GetText(),
			))
			return
		}

		if identifierType == BASE_TYPES.UNKNOWN {
			log.Printf("Inferring type of `%s` as `%s`\n", identifier.GetText(), assignType)
			l.ScopeManager.CurrentScope.UpsertExpressionType(identifier.GetText(), assignType)
			identifierType = assignType
		}

		if identifierType != assignType {
			l.AddError(fmt.Sprintf(
				"(line: %d) Trying to assign `%s` to variable `%s` but types don't match! (`%s` != `%s`)",
				line,
				assignExpr.GetText(),
				identifier.GetText(),
				identifierType,
				assignType,
			))
			return
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
	l.ScopeManager.CurrentScope.UpsertExpressionType(expr, TypeIdentifier(exprType))
}
