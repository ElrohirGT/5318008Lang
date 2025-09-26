package type_checker

import (
	"fmt"
	"log"
	"slices"

	p "github.com/ElrohirGT/5318008Lang/parser"
)

// METHODS FOR HANDLING CLASSES TYPES

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
	colStart := className.GetSymbol().GetColumn()
	colEnd := colStart + len(className.GetText())

	onGlobaScope := l.ScopeManager.CurrentScope.Type == SCOPE_TYPES.GLOBAL
	if !onGlobaScope {
		l.AddError(line, colStart, colEnd, fmt.Sprintf(
			"Can't define class `%s` inside scope! Classes can only be defined on global scope!",
			className.GetText(),
		))
	}

	classScope := NewScope(className.GetText(), SCOPE_TYPES.CLASS)
	l.ScopeManager.AddToCurrent(classScope)
	l.ScopeManager.ReplaceCurrent(classScope)
	// Add this as a valid expression with type of this class
	classScope.UpsertExpressionType("this", TypeIdentifier(className.GetText()))

	if slices.Contains(BASE_TYPE_ARRAY, TypeIdentifier(className.GetText())) {
		l.AddError(line, colStart, colEnd, fmt.Sprintf(
			"Can't define class `%s` because it collides with a builtin type!",
			className,
		))
		return
	}

	if _, found := l.GetTypeInfo(TypeIdentifier(className.GetText())); found {
		l.AddError(line, colStart, colEnd, fmt.Sprintf(
			"Can't redefine existing class! `%s` already exists!",
			className.GetText(),
		))
	} else {
		classInfo := NewClassTypeInfo(className.GetText())
		classTypeId := TypeIdentifier(className.GetText())
		l.AddTypeInfo(classTypeId, NewTypeInfo_Class(classInfo))
		l.AddTypeInfo(NewArrayTypeIdentifier(classTypeId), NewTypeInfo_Base())
	}

	if len(identifiers) > 1 {
		fatherClassName := identifiers[1]
		fColStart := fatherClassName.GetSymbol().GetColumn()
		fColEnd := fColStart + len(fatherClassName.GetText())

		canInherit := l.CheckInheritanceTree(ctx, TypeIdentifier(className.GetText()), TypeIdentifier(fatherClassName.GetText()))
		if canInherit {
			log.Println("Class inherits from", fatherClassName)
			info, found := l.GetTypeInfo(TypeIdentifier(fatherClassName.GetText()))
			if !found {
				l.AddError(line, fColStart, fColEnd, fmt.Sprintf(
					"Can't inherit from a type that doesn't exists! `%s` wants to inherit from `%s`!",
					className.GetText(),
					fatherClassName.GetText(),
				))
			} else {
				if !info.ClassType.HasValue() {
					l.AddError(line, fColStart, fColEnd, fmt.Sprintf(
						"Can't make a nonexistent class inherit from another! `%s` wants to inherit from `%s` but `%s` is not a class!",
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
}

func (l Listener) ExitNewExpr(ctx *p.NewExprContext) {
	line := ctx.GetStart().GetLine()
	className := ctx.Identifier()
	colStart := className.GetSymbol().GetColumn()
	colEnd := colStart + len(className.GetText())
	log.Println("Trying to instantiate class:", className.GetText())

	exprArguments := []p.IConditionalExprContext{}
	if ctx.Arguments() != nil {
		exprArguments = ctx.Arguments().AllConditionalExpr()
	}

	classScope, found := l.ScopeManager.SearchClassScope()

	if found && classScope.Name == className.GetText() {
		l.AddError(line, colStart, colEnd, fmt.Sprintf(
			"Can't create a new instance of `%s` while defining the same class!",
			className,
		))
		return
	}

	typeInfo, found := l.GetTypeInfo(TypeIdentifier(className.GetText()))
	if !found || !typeInfo.ClassType.HasValue() {
		l.AddError(line, colStart, colEnd, fmt.Sprintf(
			"Can't create a new instance of undefined class `%s`",
			className.GetText(),
		))
		return
	}

	classTypeInfo := typeInfo.ClassType.GetValue()
	methodInfo := classTypeInfo.GetConstructor(&l)
	if len(methodInfo.ParameterList) != len(exprArguments) {
		l.AddError(line, colStart, colEnd, fmt.Sprintf(
			"Constructor of `%s` asks for `%d` arguments but `%d` given!",
			className,
			len(methodInfo.ParameterList),
			len(exprArguments),
		))
		return
	}

	errorWithParams := false
	for i, consParam := range methodInfo.ParameterList {
		exprParam := exprArguments[i]
		start := exprParam.GetStart().GetColumn()
		end := start + len(exprParam.GetText())

		exprType, found := l.ScopeManager.CurrentScope.GetExpressionType(exprParam.GetText())
		if !found {
			l.AddError(line, start, end, fmt.Sprintf(
				"Type of `%s` not found!",
				exprParam.GetText(),
			))
			errorWithParams = true
		}

		if consParam.Type != exprType {
			l.AddError(line, start, end, fmt.Sprintf(
				"Constructor for `%s` requires `%s` to be of type `%s` but it's `%s` instead!",
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
