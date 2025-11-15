package tac_generator

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ElrohirGT/5318008Lang/lib"
	"github.com/ElrohirGT/5318008Lang/type_checker"
)

type ScopeName string

type TagName string

type VariableName string

type LiteralOrVariable string

// Represents the operation to assign a value to a variable, for example:
//
// = t1 i32 5
//
// = t2 u1 1.
type AssignmentInstruction struct {
	Target VariableName
	Type   VariableType
	Value  LiteralOrVariable
}

func NewAssignmentInstruction(instruction AssignmentInstruction) Instruction {
	return Instruction{
		Assignment: lib.NewOpValue(instruction),
	}
}

type VariableType string

var VARIABLE_TYPES = struct {
	U8  VariableType
	U16 VariableType
	U32 VariableType
	I8  VariableType
	I16 VariableType
	I32 VariableType
}{
	U8:  "u8",
	U16: "u16",
	U32: "u32",
	I8:  "i8",
	I16: "i16",
	I32: "i32",
}

// Represents the operation to copy a variable.
//
// Example:
//
// cp t1 t2
//
// Copies t2 into t1.
type CopyInstruction struct {
	Target VariableName
	Source VariableName
}

func NewCopyInstruction(instruction CopyInstruction) Instruction {
	return Instruction{
		Copy: lib.NewOpValue(instruction),
	}
}

type LogicOperationType string

var BOOLEAN_OPERATION_TYPES = struct {
	And            LogicOperationType
	Or             LogicOperationType
	Equal          LogicOperationType
	NotEqual       LogicOperationType
	Less           LogicOperationType
	LessOrEqual    LogicOperationType
	Greater        LogicOperationType
	GreaterOrEqual LogicOperationType
}{
	And:            "AND",
	Or:             "OR",
	Equal:          "EQ",
	NotEqual:       "NEQ",
	Less:           "LT",
	LessOrEqual:    "LTE",
	Greater:        "GT",
	GreaterOrEqual: "GTE",
}

// Represents an logic instruction like and & or.
type LogicOpInstruction struct {
	Type   LogicOperationType
	Target VariableName
	P1     VariableName
	P2     VariableName
}

func NewLogicOpInstruction(instruction LogicOpInstruction) Instruction {
	return Instruction{
		Logic: lib.NewOpValue(instruction),
	}
}

// Represents string concatenation
type ConcatInstruction struct {
	Target  VariableName
	String1 LiteralOrVariable
	String2 LiteralOrVariable
}

func NewConcatInstruction(instruction ConcatInstruction) Instruction {
	return Instruction{
		Concat: lib.NewOpValue(instruction),
	}
}

// Represents a tag, used to mark sections where the program can jump to.
type SECInstruction struct {
	Name TagName
}

func NewSecInstruction(instruction SECInstruction) Instruction {
	return Instruction{
		Sec: lib.NewOpValue(instruction),
	}
}

// Represents a jump instruction, it can be conditional or unconditional.
//
// Valid formats are:
//
// * IF variable GOTO sección/tag
//
// * IF NOT variable GOTO sección/tag
//
// * IF relop variable 1 variable 2 GOTO sección/tag.
type JumpInstruction struct {
	Condition lib.Optional[JumpCondition]
	Target    TagName
}

func NewJumpInstruction(instruction JumpInstruction) Instruction {
	return Instruction{
		Jump: lib.NewOpValue(instruction),
	}
}

type CondJumpOperation struct {
	// If we're dealing with signed integers
	Signed bool
	Type   LogicOperationType
	P1     VariableName
	P2     VariableName
}

type JumpCondition struct {
	Simple        lib.Optional[VariableName]
	SimpleNegated lib.Optional[VariableName]
	Relation      lib.Optional[CondJumpOperation]
}

// Represents the instruction to create a parameter for a procedure.
type ParamInstruction struct {
	Parameter LiteralOrVariable
}

func NewParamInstruction(instruction ParamInstruction) Instruction {
	return Instruction{
		Param: lib.NewOpValue(instruction),
	}
}

// Loads an argument into a variable for later use.
type LoadInstruction struct {
	Variable VariableName
}

func NewLoadInstruction(instruction LoadInstruction) Instruction {
	return Instruction{
		Load: lib.NewOpValue(instruction),
	}
}

// Represents the instruction to call a procedure, you may need to use ParamInstruction before calling this.
type CallInstruction struct {
	SaveReturnOn   lib.Optional[VariableName]
	ProcedureName  ScopeName
	NumberOfParams uint
}

func NewCallInstruction(instruction CallInstruction) Instruction {
	return Instruction{
		Call: lib.NewOpValue(instruction),
	}
}

// Represents the instruction to return from a procedure, optionally returning a variable.
type ReturnInstruction struct {
	Value LiteralOrVariable
}

func NewReturnInstruction(instruction ReturnInstruction) Instruction {
	return Instruction{
		Return: lib.NewOpValue(instruction),
	}
}

// Represents the instruction to allocate a certain amount of memory.
type AllocInstruction struct {
	Target VariableName
	Size   uint
}

func NewAllocInstruction(instruction AllocInstruction) Instruction {
	return Instruction{
		Alloc: lib.NewOpValue(instruction),
	}
}

// Get's from a specified source some memory pointer, and assigns it to target.
//
// Useful for arrays:
//
// b = a[0]
//
// which translates to:
//
// lwo b a 0
//
// And objects!
type LoadWithOffsetInstruction struct {
	Target VariableName
	Source VariableName
	Offset LiteralOrVariable
}

func NewLoadWithOffsetInstruction(instruction LoadWithOffsetInstruction) Instruction {
	return Instruction{
		LoadWithOffset: lib.NewOpValue(instruction),
	}
}

// Get's from a specified source some memory pointer, and assigns it to target.
//
// Useful for arrays:
//
// b[0] = a
//
// which translates to:
//
// swo b 0 a
//
// And objects!
type SetWithOffsetInstruction struct {
	Target VariableName
	Offset LiteralOrVariable
	Value  LiteralOrVariable
}

func NewSetWithOffsetInstruction(instruction SetWithOffsetInstruction) Instruction {
	return Instruction{
		SetWithOffset: lib.NewOpValue(instruction),
	}
}

// Get's the memory address from the specified variable target.
// Equivalent to: &.
type ReferenceInstruction struct {
	Target VariableName
}

func NewReferenceInstruction(instruction ReferenceInstruction) Instruction {
	return Instruction{
		Reference: lib.NewOpValue(instruction),
	}
}

// Get's the value from a memory address in the target variable.
// Equivalent to: @.
type DereferenceInstruction struct {
	Target VariableName
}

func NewDereferenceInstruction(instruction DereferenceInstruction) Instruction {
	return Instruction{
		Dereference: lib.NewOpValue(instruction),
	}
}

// Represents an arithmethic instruction like adding or multiplying.
type ArithmethicInstruction struct {
	// If we're dealing with signed integers.
	Signed bool
	Type   ArithmethicOpType
	Target VariableName
	P1     LiteralOrVariable
	P2     LiteralOrVariable
}

func NewArithmethicInstruction(instruction ArithmethicInstruction) Instruction {
	return Instruction{
		Arithmethic: lib.NewOpValue(instruction),
	}
}

type ArithmethicOpType string

var ARITHMETHIC_OPERATION_TYPES = struct {
	Add            ArithmethicOpType
	Subtract       ArithmethicOpType
	Multiplication ArithmethicOpType
	Divide         ArithmethicOpType
	// TODO: Are there any others?
}{
	Add:            "ADD",
	Subtract:       "SUB",
	Multiplication: "MULT",
	Divide:         "DIV",
}

// Represents a TAC instruction.
//
// This element is an enumeration, which means only one of the optionals is valid!
type Instruction struct {
	Comment        string
	Assignment     lib.Optional[AssignmentInstruction]
	Copy           lib.Optional[CopyInstruction]
	Sec            lib.Optional[SECInstruction]
	Jump           lib.Optional[JumpInstruction]
	Param          lib.Optional[ParamInstruction]
	Load           lib.Optional[LoadInstruction]
	Call           lib.Optional[CallInstruction]
	Return         lib.Optional[ReturnInstruction]
	Alloc          lib.Optional[AllocInstruction]
	LoadWithOffset lib.Optional[LoadWithOffsetInstruction]
	SetWithOffset  lib.Optional[SetWithOffsetInstruction]
	Free           lib.Optional[VariableName]
	Reference      lib.Optional[ReferenceInstruction]
	Dereference    lib.Optional[DereferenceInstruction]
	Arithmethic    lib.Optional[ArithmethicInstruction]
	Logic          lib.Optional[LogicOpInstruction]
	Concat         lib.Optional[ConcatInstruction]
}

func (i Instruction) AddComment(comment string) Instruction {
	i.Comment = comment
	return i
}

// Debug representation of an instruction.
func (i Instruction) String() string {
	switch {
	case i.Assignment.HasValue():
		ass := i.Assignment.GetValue()
		return fmt.Sprintf("{%s: %s = %s}", ass.Target, ass.Type, ass.Value)
	case i.Copy.HasValue():
		ass := i.Copy.GetValue()
		return fmt.Sprintf("{copy %s into %s}", ass.Source, ass.Target)
	case i.Jump.HasValue():
		ass := i.Jump.GetValue()
		if ass.Condition.HasValue() {
			return fmt.Sprintf("{jmp %s (IF: %#v)}", ass.Target, ass.Condition.GetValue())
		} else {
			return fmt.Sprintf("{jmp %s}", ass.Target)
		}
	case i.Param.HasValue():
		ass := i.Param.GetValue()
		return fmt.Sprintf("{param %s}", ass.Parameter)
	case i.Call.HasValue():
		ass := i.Call.GetValue()
		if ass.SaveReturnOn.HasValue() {
			return fmt.Sprintf("{%s = call %s with %d params}", ass.SaveReturnOn.GetValue(), ass.ProcedureName, ass.NumberOfParams)
		} else {
			return fmt.Sprintf("{call %s with %d params}", ass.ProcedureName, ass.NumberOfParams)
		}
	case i.Return.HasValue():
		ass := i.Return.GetValue()
		return fmt.Sprintf("{return %s}", ass.Value)
	case i.Alloc.HasValue():
		ass := i.Alloc.GetValue()
		return fmt.Sprintf("{alloc %s with %d bytes}", ass.Target, ass.Size)
	case i.LoadWithOffset.HasValue():
		ass := i.LoadWithOffset.GetValue()
		return fmt.Sprintf("{%s = %s[%s]}", ass.Target, ass.Source, ass.Offset)
	case i.SetWithOffset.HasValue():
		ass := i.SetWithOffset.GetValue()
		return fmt.Sprintf("{%s[%s] = %s}", ass.Target, ass.Offset, ass.Value)
	case i.Free.HasValue():
		ass := i.Free.GetValue()
		return fmt.Sprintf("{free %s}", ass)
	case i.Reference.HasValue():
		ass := i.Reference.GetValue()
		return fmt.Sprintf("{&%s}", ass.Target)
	case i.Dereference.HasValue():
		ass := i.Dereference.GetValue()
		return fmt.Sprintf("{@%s}", ass.Target)
	case i.Arithmethic.HasValue():
		ass := i.Arithmethic.GetValue()
		return fmt.Sprintf("{%s = %s %s %s (signed? %t)}", ass.Target, ass.P1, ass.Type, ass.P2, ass.Signed)
	case i.Logic.HasValue():
		ass := i.Logic.GetValue()
		return fmt.Sprintf("{%s = %s %s %s}", ass.Target, ass.P1, ass.Type, ass.P2)
	case i.Load.HasValue():
		ass := i.Load.GetValue()
		return fmt.Sprintf("{LOAD %s}", ass.Variable)
	case i.Concat.HasValue():
		ass := i.Concat.GetValue()
		return fmt.Sprintf("{%s = concat(%s, %s)}", ass.Target, ass.String1, ass.String2)
	}

	return "{**invalid instruction**}"
}

type ScopeInformation struct {
	Parent       ScopeName // Set to empty when no available
	Instructions []Instruction
	// We need to translate between `localName` of a variable to a `t1` or `t2`.
	// This map aids on that.
	translations map[string]VariableName
}

// A program represents a complete set of scopes and instructions to execute.
type Program struct {
	variableCounter uint
	Scopes          map[ScopeName]ScopeInformation
	ScopesNames     []ScopeName
	MainScope       ScopeName
}

func NewProgram() *Program {
	scopes := make(map[ScopeName]ScopeInformation)
	scopes[type_checker.GLOBAL_SCOPE_NAME] = ScopeInformation{
		Instructions: []Instruction{},
		translations: map[string]VariableName{},
	}

	return &Program{
		Scopes:    scopes,
		MainScope: type_checker.GLOBAL_SCOPE_NAME,
	}
}

func (p *Program) InsertIfNotExists(scopeName, parentScope ScopeName) {
	if _, found := p.Scopes[scopeName]; found {
		return
	}

	p.Scopes[scopeName] = ScopeInformation{
		Parent:       parentScope,
		Instructions: []Instruction{},
		translations: map[string]VariableName{},
	}
}

func (p *Program) UpsertScope(scopeName, parentScope ScopeName) {
	p.Scopes[scopeName] = ScopeInformation{
		Parent:       parentScope,
		Instructions: []Instruction{},
		translations: map[string]VariableName{},
	}
}

func (p *Program) UpsertTranslation(scope ScopeName, localName string, tacName VariableName) {
	scopeInfo, found := p.Scopes[scope]
	if !found {
		log.Panicf(
			"Failed to find scope `%s` to upsert translation for `%s` into `%s`",
			scope,
			localName,
			tacName,
		)
	}

	scopeInfo.translations[localName] = tacName
	p.Scopes[scope] = scopeInfo
}

func (p *Program) GetVariableFor(expr string, scope ScopeName) (VariableName, bool) {
	scopeInfo, found := p.Scopes[scope]
	if !found {
		log.Panicf(
			"Failed to find scope `%s` when trying to translate `%s` expression",
			scope,
			expr,
		)
	}

	tacName, found := scopeInfo.translations[expr]
	if found {
		return tacName, found
	}
	for scopeInfo.Parent != "" {
		scopeInfo = p.Scopes[scopeInfo.Parent]
		tacName, found := scopeInfo.translations[expr]
		if found {
			return tacName, found
		}
	}

	return "", false
}

func (p *Program) GetOrGenerateVariable(name string, scope ScopeName) VariableName {
	scopeInfo, found := p.Scopes[scope]
	if !found {
		log.Panicf(
			"Failed to find scope `%s` when trying to translate `%s` variable",
			scope,
			name,
		)
	}

	tacName, found := scopeInfo.translations[name]
	if found {
		// fmt.Println("FOUND FIRST TIME =" + varName)
		return tacName
	}
	// Search on parents
	for scopeInfo.Parent != "" {
		// parent := scopeInfo.Parent
		scopeInfo = p.Scopes[scopeInfo.Parent]
		varName, found := scopeInfo.translations[name]
		if found {
			// fmt.Println("FOUND ON PARENT =" + parent)
			return varName
		}
	}

	if !found {
		// fmt.Println("NOT FOUND :( CREATING")
		p.variableCounter += 1
		tacName = VariableName("t" + strconv.FormatUint(uint64(p.variableCounter), 10))
		p.UpsertTranslation(scope, name, tacName)
	}

	return tacName
}

func (p *Program) GetNextVariable() VariableName {
	p.variableCounter += 1
	return VariableName("t" + strconv.FormatUint(uint64(p.variableCounter), 10))
}

// Given and expresion and scope, if the expresion is a literal it returns
// its TAC representation, if not it attemps to find the variable TAC name.
func (p *Program) GetVariableOrLiteral(expresion string,
	scopeName ScopeName, l *type_checker.Listener) (VariableName, type_checker.TypeIdentifier) {
	if literalType, isLiteral := l.GetLiteralType(expresion); isLiteral {
		literalType, value := literalToTAC(expresion, literalType)
		return VariableName(value), type_checker.TypeIdentifier(literalType)
	} else {
		exprType, found := l.ScopeManager.CurrentScope.GetExpressionType(expresion)
		if !found {
			log.Panicf("Failed to find a variable for the expression:\n`%s`",
				expresion)
		}
		tacName, found := p.GetVariableFor(expresion, scopeName)
		if !found {
			log.Panicf("Failed to find a variable for the expression:\n`%s`",
				expresion)
		}
		return tacName, exprType
	}
}
