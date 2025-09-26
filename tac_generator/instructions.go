package tac_generator

import (
	"strconv"

	"github.com/ElrohirGT/5318008Lang/lib"
)

type ScopeName string

type VariableName string

// Represents the operation to assign a value to a variable, for example:
//
// = t1 i32 5
//
// = t2 u1 1.
type AssignmentInstruction struct {
	Target VariableName
	Type   string
	Value  string
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
	Source string
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

// Represents the instruction to call a procedure, you may need to use ParamInstruction before calling this.
type CallInstruction struct {
	ProcedureName  ScopeName
	NumberOfParams uint
}

// Represents the instruction to return from a procedure, optionally returning a variable.
type ReturnInstruction struct {
	Value VariableName
}

// Represents the instruction to allocate a certain amount of memory.
type AllocInstruction struct {
	Target VariableName
	Size   uint
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
	Value  VariableName
}

// Get's the memory address from the specified variable target.
// Equivalent to: &.
type ReferenceInstruction struct {
	Target VariableName
}

// Get's the value from a memory address in the target variable.
// Equivalent to: @.
type DereferenceInstruction struct {
	Target VariableName
}

// Represents an arithmethic instruction like adding or multiplying.
type ArithmethicInstruction struct {
	// If we're dealing with signed integers.
	Signed bool
	Type   ArithmethicOperationType
	P1     VariableName
	P2     VariableName
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

func (p *Program) GetNextVariableName() VariableName {
	p.variableCounter += 1
	return VariableName("T" + strconv.FormatUint(uint64(p.variableCounter), 10))
}
