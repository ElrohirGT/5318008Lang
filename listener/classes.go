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

func (c *ClassTypeInfo) AddField(name string, _type TypeIdentifier) {
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

	l.ScopeManager.CurrentScope = *l.ScopeManager.CurrentScope.Father
	fmt.Println("Escaping class declaration:", ctx.AllIdentifier()[0].GetText())
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
	l.ScopeManager.CurrentScope.AddChildScope(&classScope)
	l.ScopeManager.CurrentScope = classScope
	log.Printf("Scope %#v", l.ScopeManager)

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

func (l Listener) ExitClassMember(ctx *p.ClassMemberContext) {
	line := ctx.GetStart().GetLine()

	isFunctionDecl := ctx.FunctionDeclaration() != nil
	isVarDecl := ctx.VariableDeclaration() != nil
	isConstantDecl := ctx.ConstantDeclaration() != nil

	if l.ScopeManager.CurrentScope.Type != SCOPE_TYPES.CLASS {
		log.Panicf("Trying to declare a class member but not inside a class! %#v", l.ScopeManager)
	}
	classType := TypeIdentifier(l.ScopeManager.CurrentScope.Name)

	if isVarDecl {
		varDeclCtx := ctx.VariableDeclaration()
		name := varDeclCtx.Identifier()

		typeAnnot := varDeclCtx.TypeAnnotation()
		hasAnnotation := typeAnnot != nil

		declarationExpr := varDeclCtx.Initializer()
		hasInitialExpr := declarationExpr != nil

		if !hasAnnotation {
			log.Println("Field", name.GetText(), "does NOT have a type! We need to infer it...")
			if hasInitialExpr {
				declarationText := declarationExpr.Expression().GetText()
				inferedType, found := l.ScopeManager.CurrentScope.GetExpressionType(declarationText)
				if !found {
					l.AddError(fmt.Sprintf(
						"(line: %d) Couldn't infer the type of variable `%s`, initialized with: `%s`",
						line,
						name.GetText(),
						declarationText,
					))
				} else {
					l.ModifyClassTypeInfo(classType, func(info *ClassTypeInfo) {
						info.AddField(name.GetText(), inferedType)
					})
				}
			} else {
				l.ModifyClassTypeInfo(classType, func(info *ClassTypeInfo) {
					info.AddField(name.GetText(), BASE_TYPES.UNKNOWN)
				})
			}
		} else {
			declarationType := TypeIdentifier(typeAnnot.Type_().GetText())
			log.Println("Field", name.GetText(), "has type", declarationType)

			if !l.TypeExists(declarationType) {
				l.AddError(fmt.Sprintf(
					"(line: %d) %s doesn't exist!",
					line,
					declarationType,
				))
			}

			if hasInitialExpr {
				exprText := declarationExpr.Expression().GetText()
				log.Println("Known expressions", l.ScopeManager.CurrentScope.typesByExpression)

				initialExprType, exists := l.ScopeManager.CurrentScope.GetExpressionType(exprText)
				if !exists {
					l.AddError(fmt.Sprintf(
						"(line: %d) `%s` doesn't have a type!",
						line,
						exprText,
					))
				}

				if initialExprType != declarationType {
					l.AddError(fmt.Sprintf(
						"(line: %d) The declaration of `%s` specifies a type of `%s` but `%s` was given",
						line,
						name,
						declarationType,
						initialExprType,
					))
				}
			}

			l.ScopeManager.CurrentScope.AddExpressionType(name.GetText(), declarationType)
			l.ModifyClassTypeInfo(classType, func(cti *ClassTypeInfo) {
				cti.AddField(name.GetText(), declarationType)
			})
		}
	} else if isFunctionDecl {
		funCtx := ctx.FunctionDeclaration()
		funName := funCtx.Identifier()
		params := funCtx.Parameters()
		if params == nil {
			log.Println("No parameters supplied to function:", funName.GetText())
		} else {
			for _, param := range params.AllParameter() {
				log.Println(param.GetText())
			}
		}

	} else if isConstantDecl {

	} else {
		panic("Class member must be a function, variable or constant declaration!")
	}
}

func (l Listener) EnterNewExpr(ctx *p.NewExprContext) {
	className := ctx.Identifier()
	log.Println("Instantiating class", className.GetText())

	// FIXME: We assume the constructor is called correctly!
	expr := ctx.GetText()
	exprType := className.GetText()
	log.Println("Adding", expr, "as an expresion of type", exprType)
	l.ScopeManager.CurrentScope.AddExpressionType(expr, TypeIdentifier(exprType))
}
