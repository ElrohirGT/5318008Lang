package tac_generator

import (
	"bytes"
	"fmt"
	"log"
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
	// Write main scope
	mainScope, found := l.Program.Scopes[l.Program.MainScope]
	if !found {
		log.Panicf("SKILL ISSUE:\nSomeone forgot to include the main scope inside the program!")
	}

	_, err := fmt.Fprintf(buff, "%s:\n", l.Program.MainScope)
	if err != nil {
		return err
	}

	for _, inst := range mainScope.Instructions {
		err := instructionToBuffer(&inst, buff)
		if err != nil {
			return err
		}
	}

	buff.WriteString("\n")

	for scopeName, instructions := range l.Program.Scopes {
		if scopeName == l.Program.MainScope {
			continue
		}

		_, err := fmt.Fprintf(buff, "%s:\n", scopeName)
		if err != nil {
			return err
		}

		for _, inst := range instructions.Instructions {
			err := instructionToBuffer(&inst, buff)
			if err != nil {
				return err
			}
		}

		buff.WriteString("\n")
	}

	return nil
}

func instructionToBuffer(inst *Instruction, buff *bytes.Buffer) error {
	defer buff.WriteString("\n")

	// FIXME: Keep implementing branches
	var err error
	switch {
	case inst.Assignment.HasValue():
		assignment := inst.Assignment.GetValue()
		_, err = fmt.Fprintf(buff, "= %s %s %s",
			assignment.Target, assignment.Type, assignment.Value)

	case inst.Copy.HasValue():
		copy := inst.Copy.GetValue()
		_, err = fmt.Fprintf(buff, "CP %s %s",
			copy.Target, copy.Source)

	case inst.Jump.HasValue():
		jump := inst.Jump.GetValue()
		if jump.Condition.HasValue() {
			// Conditional jumps
			cond := jump.Condition.GetValue()

			if cond.Simple.HasValue() {
				// Conditional simple jump
				_, err = fmt.Fprintf(buff, "IF %s GOTO %s",
					cond.Simple.GetValue(), jump.Target)
			} else if cond.SimpleNegated.HasValue() {
				// Conditional negation jump
				_, err = fmt.Fprintf(buff, "IF NOT %s GOTO %s",
					cond.SimpleNegated.GetValue(), jump.Target)
			} else if cond.Relation.HasValue() {
				// Conditional multivariable jump
				rel := cond.Relation.GetValue()
				_, err = fmt.Fprintf(buff, "IF %s %s %s GOTO %s",
					rel.Type, rel.P1, rel.P2, jump.Target)
			}
		} else {
			// Unconditional jump
			_, err = fmt.Fprintf(buff, "GOTO %s",
				jump.Target)
		}

	case inst.Param.HasValue():
		param := inst.Param.GetValue()
		_, err = fmt.Fprintf(buff, "PARAM %s",
			param.Parameter)

	case inst.Call.HasValue():
		call := inst.Call.GetValue()
		if call.SaveReturnOn.HasValue() {
			// call function with return
			_, err = fmt.Fprintf(buff, "CALLRET %s %s %d",
				call.SaveReturnOn.GetValue(), call.ProcedureName, call.NumberOfParams)
		} else {
			_, err = fmt.Fprintf(buff, "CALL %s %d",
				call.ProcedureName, call.NumberOfParams)
		}

	case inst.Return.HasValue():
		ret := inst.Return.GetValue()
		_, err = fmt.Fprintf(buff, "RETURN %s",
			ret.Value)

	case inst.Alloc.HasValue():
		alloc := inst.Alloc.GetValue()
		_, err = fmt.Fprintf(buff, "ALLOC %s %d",
			alloc.Target, alloc.Size)

	case inst.LoadWithOffset.HasValue():
		lwo := inst.LoadWithOffset.GetValue()
		_, err = fmt.Fprintf(buff, "LWO %s %s %s",
			lwo.Target, lwo.Source, lwo.Offset)

	case inst.SetWithOffset.HasValue():
		swo := inst.SetWithOffset.GetValue()
		_, err = fmt.Fprintf(buff, "SWO %s %s %s",
			swo.Target, swo.Offset, swo.Value)

	case inst.Free.HasValue():
		free := inst.Free.GetValue()
		_, err = fmt.Fprintf(buff, "FREE %s",
			free)

	case inst.Reference.HasValue():
		ref := inst.Reference.GetValue()
		_, err = fmt.Fprintf(buff, "&%s",
			ref.Target)

	case inst.Dereference.HasValue():
		def := inst.Dereference.GetValue()
		_, err = fmt.Fprintf(buff, "@%s",
			def.Target)

	case inst.Arithmethic.HasValue():
		// FIXME: implement me
	case inst.Logic.HasValue():
		// FIXME: implement me
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

func (l *Listener) CreateAssignment(varName string, literalType type_checker.TypeIdentifier, rawValue string) {
	currentScope := l.GetCurrentScope()
	target := l.Program.GetOrGenerateVariable(varName, ScopeName(currentScope.Name))
	varType, literalValue := literalToTAC(rawValue, literalType)

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
