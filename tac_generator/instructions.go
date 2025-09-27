package tac_generator

import (
	"strconv"

	"github.com/ElrohirGT/5318008Lang/lib"
)

type ScopeName string

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
	Target    ScopeName
}

func NewJumpInstruction(instruction JumpInstruction) Instruction {
	return Instruction{
		Jump: lib.NewOpValue(instruction),
	}
}

type BooleanOperationType string

var BOOLEAN_OPERATION_TYPES = struct {
	Equal          BooleanOperationType
	NotEqual       BooleanOperationType
	Less           BooleanOperationType
	LessOrEqual    BooleanOperationType
	Greater        BooleanOperationType
	GreaterOrEqual BooleanOperationType
}{
	Equal:          "==",
	NotEqual:       "!=",
	Less:           "<",
	LessOrEqual:    "<=",
	Greater:        ">",
	GreaterOrEqual: ">=",
}

type BooleanOperation struct {
	// If we're dealing with signed integers
	Signed bool
	Type   BooleanOperationType
	P1     VariableName
	P2     VariableName
}

type JumpCondition struct {
	Simple        lib.Optional[VariableName]
	SimpleNegated lib.Optional[VariableName]
	Relation      lib.Optional[BooleanOperation]
}

// Represents the instruction to create a parameter for a procedure.
type ParamInstruction struct {
	Parameter VariableName
}

func NewParamInstruction(instruction ParamInstruction) Instruction {
	return Instruction{
		Param: lib.NewOpValue(instruction),
	}
}

// Represents the instruction to call a procedure, you may need to use ParamInstruction before calling this.
type CallInstruction struct {
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
	Value VariableName
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
	Offset uint
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
	Offset uint
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
	Type   ArithmethicOperationType
	P1     VariableName
	P2     LiteralOrVariable
}

func NewArithmethicInstruction(instruction ArithmethicInstruction) Instruction {
	return Instruction{
		Arithmethic: lib.NewOpValue(instruction),
	}
}

type ArithmethicOperationType string

var ARITHMETHIC_OPERATION_TYPES = struct {
	Add            ArithmethicOperationType
	Subtract       ArithmethicOperationType
	Multiplication ArithmethicOperationType
	Divide         ArithmethicOperationType
	// TODO: Are there any others?
}{
	Add:            "+",
	Subtract:       "-",
	Multiplication: "*",
	Divide:         "/",
}

// Represents a TAC instruction.
//
// This element is an enumeration, which means only one of the optionals is valid!
type Instruction struct {
	Assignment     lib.Optional[AssignmentInstruction]
	Copy           lib.Optional[CopyInstruction]
	Jump           lib.Optional[JumpInstruction]
	Param          lib.Optional[ParamInstruction]
	Call           lib.Optional[CallInstruction]
	Return         lib.Optional[ReturnInstruction]
	Alloc          lib.Optional[AllocInstruction]
	LoadWithOffset lib.Optional[LoadWithOffsetInstruction]
	SetWithOffset  lib.Optional[SetWithOffsetInstruction]
	Free           lib.Optional[VariableName]
	Reference      lib.Optional[ReferenceInstruction]
	Dereference    lib.Optional[DereferenceInstruction]
	Arithmethic    lib.Optional[ArithmethicInstruction]
}

// A program represents a complete set of scopes and instructions to execute.
type Program struct {
	variableCounter uint
	Scopes          map[ScopeName][]Instruction
}

func NewProgram() *Program {
	return &Program{
		Scopes: make(map[ScopeName][]Instruction),
	}
}

func (p *Program) GetNextVariableName() VariableName {
	p.variableCounter += 1
	return VariableName("T" + strconv.FormatUint(uint64(p.variableCounter), 10))
}
