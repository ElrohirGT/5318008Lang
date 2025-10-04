package tac_generator

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/ElrohirGT/5318008Lang/type_checker"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

// Typing system scanner, responsable of the semantic during compiscript code.
// Handles the notion of types, definitions and scope management.
type Listener struct {
	*p.BaseCompiscriptListener
	TypeChecker *type_checker.Listener
	Program     *Program
	Errors      *[]string
}

func NewListener(typeChecker *type_checker.Listener) Listener {
	typeChecker.ScopeManager.ReplaceCurrent(typeChecker.ScopeManager.GlobaScope)
	return Listener{
		Program:     NewProgram(),
		TypeChecker: typeChecker,
		Errors:      &[]string{},
	}
}

// Generates the final TAC contents.
func (l *Listener) Generate(buff *bytes.Buffer) error {
	mainScope, found := l.Program.Scopes[l.Program.MainScope]
	if !found {
		log.Panicf("SKILL ISSUE:\nSomeone forgot to include the main scope inside the program!")
	}

	for _, inst := range mainScope.Instructions {
		err := instructionToBuffer(&inst, buff)
		if err != nil {
			return err
		}
	}
	// FIXME: Prince needs to fill this!
	return nil
}

func instructionToBuffer(inst *Instruction, buff *bytes.Buffer) error {
	// FIXME: Prince needs to fill this!
	defer buff.WriteString("\n")

	// FIXME: Keep implementing branches
	var err error
	switch {
	case inst.Assignment.HasValue():
		assignment := inst.Assignment.GetValue()
		_, err = buff.WriteString("= " + string(assignment.Target) + " " + string(assignment.Type) + " " + string(assignment.Value))
	case inst.Copy.HasValue():
	case inst.Jump.HasValue():
	case inst.Param.HasValue():
	case inst.Call.HasValue():
	case inst.Return.HasValue():
	case inst.Alloc.HasValue():
	case inst.LoadWithOffset.HasValue():
	case inst.SetWithOffset.HasValue():
	case inst.Free.HasValue():
	case inst.Reference.HasValue():
	case inst.Dereference.HasValue():
	case inst.Arithmethic.HasValue():
	case inst.Logic.HasValue():
	default:
		log.Panicf("Unrecognizable instruction type!\n%#v", *inst)
	}

	if err != nil {
		return err
	}

	return nil
}

func (l *Listener) GetCurrentScope() *type_checker.Scope {
	currentScope := l.TypeChecker.ScopeManager.CurrentScope
	if currentScope == nil {
		log.Panic("Failed to obtain current scope!")
	}
	return currentScope
}

func (l *Listener) AppendInstruction(inst Instruction) {
	currentScope := l.GetCurrentScope()
	scopeInfo := l.Program.Scopes[ScopeName(currentScope.Name)]
	scopeInfo.Instructions = append(scopeInfo.Instructions, inst)
	log.Printf("Appending to scope `%s`: %s", currentScope.Name, inst)
	l.Program.Scopes[ScopeName(currentScope.Name)] = scopeInfo
}

func (l *Listener) CreateLiteralAssignment(varName string, literalType type_checker.TypeIdentifier, rawValue string) {
	currentScope := l.GetCurrentScope()
	literalValue := "**SKILL ISSUE VALUE**"
	target := l.Program.GetOrGenerateVariable(varName, ScopeName(currentScope.Name))
	varType := VARIABLE_TYPES.I32

	switch literalType {
	case type_checker.BASE_TYPES.BOOLEAN:
		varType = VARIABLE_TYPES.I8
		switch rawValue {
		case type_checker.LITERAL_VALUES.False:
			literalValue = "0"
		case type_checker.LITERAL_VALUES.True:
			literalValue = strconv.FormatInt(^0, 10)
		default:
			log.Panicf(
				"Expression: `%s`\nis of type `%s`\nbut it isn't `%s` nor `%s`",
				rawValue,
				literalType,
				type_checker.LITERAL_VALUES.False,
				type_checker.LITERAL_VALUES.True,
			)
		}

	case type_checker.BASE_TYPES.INTEGER:
		literalValue = rawValue
	case type_checker.BASE_TYPES.STRING:
	case type_checker.BASE_TYPES.NULL, type_checker.BASE_TYPES.INVALID, type_checker.BASE_TYPES.UNKNOWN:
		log.Panicf(
			"Literal expression: `%s` is of invalid type! `%s`",
			varName,
			literalType,
		)
	default:
		// FIXME: It's an array! Handle array cases.
	}

	l.AppendInstruction(NewAssignmentInstruction(AssignmentInstruction{
		Target: target,
		Type:   varType,
		Value:  LiteralOrVariable(literalValue),
	}))
}

// FIXME: Improve error handling:
// Centralize error handling in lib maybe?
func (l Listener) AddError(line int, columnStart int, columnEnd int, content string, details ...string) {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(Red+"* Error: (line: %d, column: %d-%d) %s"+Reset, line, columnStart, columnEnd, content))

	for _, v := range details {
		b.WriteString("\n * " + v)
	}

	*l.Errors = append(*l.Errors, b.String())
}

func (l Listener) AddWarning(content string, line string, details ...string) {
	*l.Errors = append(*l.Errors, "Warning: "+content)
}

func (l Listener) HasErrors() bool {
	return len(*l.Errors) > 0
}
