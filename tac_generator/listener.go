package tac_generator

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	p "github.com/ElrohirGT/5318008Lang/parser"
	"github.com/ElrohirGT/5318008Lang/type_checker"
	"github.com/antlr4-go/antlr/v4"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

// Typing system scanner, responsable of the semantic during compiscript code.
// Handles the notion of types, definitions and scope management.
type Listener struct {
	*p.BaseCompiscriptListener
	Source      *antlr.CommonTokenStream
	TypeChecker *type_checker.Listener
	Program     *Program
	Errors      *[]string
}

func NewListener(typeChecker *type_checker.Listener, source *antlr.CommonTokenStream) Listener {
	typeChecker.ScopeManager.ReplaceCurrent(typeChecker.ScopeManager.GlobaScope)
	return Listener{
		Program:     NewProgram(),
		TypeChecker: typeChecker,
		Errors:      &[]string{},
		Source:      source,
	}
}

// Generates the final TAC contents.
func (l *Listener) Generate(buff *bytes.Buffer) error {
	// Write main scope
	mainScope, found := l.Program.Scopes[l.Program.MainScope]
	if !found {
		log.Panicf("SKILL ISSUE:\nSomeone forgot to include the main scope inside the program!")
	}

	for _, inst := range mainScope.Instructions {
		err := instructionToBuffer(&inst, buff, "")
		if err != nil {
			return err
		}
	}

	buff.WriteString("\n")

	var scopeNames []string
	for name := range l.Program.Scopes {
		scopeNames = append(scopeNames, string(name))
	}
	// Sort Alphabetically
	sort.Strings(scopeNames)

	for _, scopeName := range scopeNames {
		scope := l.Program.Scopes[ScopeName(scopeName)]
		if ScopeName(scopeName) == l.Program.MainScope {
			continue
		}

		// _, err := fmt.Fprintf(buff, "SEC %s:\n", scopeName)
		// if err != nil {
		// 	return err
		// }

		log.Printf("Entering scope with name: %s\n", scopeName)

		for _, inst := range scope.Instructions {
			err := instructionToBuffer(&inst, buff, "\t")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func instructionToBuffer(inst *Instruction, buff *bytes.Buffer, tab string) error {
	defer buff.WriteString("\n")

	if inst.Comment != "" {
		fmt.Fprintf(buff, "%s// %s\n", tab, inst.Comment)
	}

	// FIXME: Keep implementing branches
	var err error
	switch {
	case inst.Assignment.HasValue():
		assignment := inst.Assignment.GetValue()
		_, err = fmt.Fprintf(buff, "%s= %s %s %s",
			tab, assignment.Target, assignment.Type, assignment.Value)

	case inst.Copy.HasValue():
		copy := inst.Copy.GetValue()
		_, err = fmt.Fprintf(buff, "%sCP %s %s", tab, copy.Target, copy.Source)

	case inst.Sec.HasValue():
		tag := inst.Sec.GetValue()
		_, err = fmt.Fprintf(buff, "%sSEC %s:", tab, tag.Name)

	case inst.Func.HasValue():
		tag := inst.Func.GetValue()
		_, err = fmt.Fprintf(buff, "FUNC %s:", tag.Name)

	case inst.Jump.HasValue():
		jump := inst.Jump.GetValue()
		if jump.Condition.HasValue() {
			// Conditional jumps
			cond := jump.Condition.GetValue()

			if cond.Simple.HasValue() {
				// Conditional simple jump
				_, err = fmt.Fprintf(buff, "%sIF %s GOTO %s",
					tab, cond.Simple.GetValue(), jump.Target)
			} else if cond.SimpleNegated.HasValue() {
				// Conditional negation jump
				_, err = fmt.Fprintf(buff, "%sIF NOT %s GOTO %s",
					tab, cond.SimpleNegated.GetValue(), jump.Target)
			} else if cond.Relation.HasValue() {
				// Conditional multivariable jump
				rel := cond.Relation.GetValue()
				_, err = fmt.Fprintf(buff, "%sIF %s %s %s GOTO %s",
					tab, rel.Type, rel.P1, rel.P2, jump.Target)
			}
		} else {
			// Unconditional jump
			_, err = fmt.Fprintf(buff, "%sGOTO %s",
				tab, jump.Target)
		}

	case inst.Param.HasValue():
		param := inst.Param.GetValue()
		_, err = fmt.Fprintf(buff, "%sPARAM %s",
			tab, param.Parameter)

	case inst.Call.HasValue():
		call := inst.Call.GetValue()
		if call.SaveReturnOn.HasValue() {
			// call function with return
			_, err = fmt.Fprintf(buff, "%sCALLRET %s %s %d",
				tab, call.SaveReturnOn.GetValue(), call.ProcedureName, call.NumberOfParams)
		} else {
			_, err = fmt.Fprintf(buff, "%sCALL %s %d",
				tab, call.ProcedureName, call.NumberOfParams)
		}

	case inst.Return.HasValue():
		ret := inst.Return.GetValue()
		_, err = fmt.Fprintf(buff, "%sRETURN %s",
			tab, ret.Value)

	case inst.Alloc.HasValue():
		alloc := inst.Alloc.GetValue()
		_, err = fmt.Fprintf(buff, "%sALLOC %s %d",
			tab, alloc.Target, alloc.Size)

	case inst.LoadWithOffset.HasValue():
		lwo := inst.LoadWithOffset.GetValue()
		_, err = fmt.Fprintf(buff, "%sLWO %s %s %s",
			tab, lwo.Target, lwo.Source, lwo.Offset)

	case inst.SetWithOffset.HasValue():
		swo := inst.SetWithOffset.GetValue()
		_, err = fmt.Fprintf(buff, "%sSWO %s %s %s",
			tab, swo.Target, swo.Offset, swo.Value)

	case inst.Free.HasValue():
		free := inst.Free.GetValue()
		_, err = fmt.Fprintf(buff, "%sFREE %s",
			tab, free)

	case inst.Reference.HasValue():
		ref := inst.Reference.GetValue()
		_, err = fmt.Fprintf(buff, "%s&%s",
			tab, ref.Target)

	case inst.Dereference.HasValue():
		def := inst.Dereference.GetValue()
		_, err = fmt.Fprintf(buff, "%s@%s",
			tab, def.Target)

	case inst.Arithmethic.HasValue():
		arith := inst.Arithmethic.GetValue()
		_, err = fmt.Fprintf(buff, "%s%s %s %s %s",
			tab, arith.Type, arith.Target, arith.P1, arith.P2)

	case inst.Logic.HasValue():
		log := inst.Logic.GetValue()
		_, err = fmt.Fprintf(buff, "%s%s %s %s %s",
			tab, log.Type, log.Target, log.P1, log.P2)
	case inst.Load.HasValue():
		def := inst.Load.GetValue()
		_, err = fmt.Fprintf(buff, "%sLOAD %s",
			tab, def.Variable)

	case inst.Concat.HasValue():
		concat := inst.Concat.GetValue()
		_, err = fmt.Fprintf(buff, "%sCONCAT %s %s %s",
			tab, concat.Target, concat.String1, concat.String2)

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

func (l *Listener) GetParentScopeName() ScopeName {
	parentScope := l.GetCurrentScope().Father
	name := ScopeName("")
	if parentScope != nil {
		name = ScopeName(parentScope.Name)
	}
	return ScopeName(name)
}

func (l *Listener) AppendInstruction(scopeName ScopeName, inst Instruction) {
	scopeInfo := l.Program.Scopes[scopeName]
	scopeInfo.Instructions = append(scopeInfo.Instructions, inst)
	log.Printf("Appending to scope `%s`: %s", scopeName, inst)
	l.Program.Scopes[scopeName] = scopeInfo
}

// Used to force de creation of a new temporal varialbe for assignment.
func (l *Listener) CreateVariableDeclaration(scopeName ScopeName, varName string, varType VariableType, rawValue string) {
	target := l.Program.GetNextVariable()

	l.Program.UpsertTranslation(scopeName, varName, target)

	l.AppendInstruction(scopeName, NewAssignmentInstruction(AssignmentInstruction{
		Target: target,
		Type:   varType,
		Value:  LiteralOrVariable(rawValue),
	}))
}

// Searches for the given variable in previus scopes to do the assignment to.
func (l *Listener) CreateVariableAssignment(scopeName ScopeName, varName string, varType VariableType, rawValue string) {
	target, found := l.Program.GetVariableFor(varName, scopeName)

	if !found {
		panic("Variable " + varName + " not found in any scope")
	}

	l.AppendInstruction(scopeName, NewAssignmentInstruction(AssignmentInstruction{
		Target: target,
		Type:   varType,
		Value:  LiteralOrVariable(rawValue),
	}))
}

func (l *Listener) MapBaseTypeToTacType(_type type_checker.TypeIdentifier) (VariableType, bool) {
	switch _type {
	case type_checker.BASE_TYPES.INTEGER:
		return VARIABLE_TYPES.I32, true
	case type_checker.BASE_TYPES.BOOLEAN:
		return VARIABLE_TYPES.I8, true
	}

	return "", false
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

func (l Listener) processStringLiteral(literal string, scopeName ScopeName) VariableName {
	// Remove quotes from the string literal
	unquoted := strings.Trim(literal, "\"")

	// Check if we already have a variable for this exact string
	if existingVar, found := l.Program.GetVariableFor(literal, scopeName); found {
		return existingVar
	}

	// Create a new variable to hold the string reference
	stringVar := l.Program.GetNextVariable()

	// Calculate the size needed (length + 1 for null terminator)
	stringSize := uint(len(unquoted) + 1)

	// Allocate memory for the string
	l.AppendInstruction(scopeName, NewAllocInstruction(AllocInstruction{
		Target: stringVar,
		Size:   stringSize,
	}).AddComment("Allocate string: "+literal))

	// Store each character in the allocated memory
	for i, char := range unquoted {
		l.AppendInstruction(scopeName, NewSetWithOffsetInstruction(SetWithOffsetInstruction{
			Target: stringVar,
			Offset: LiteralOrVariable(strconv.Itoa(i)),
			Value:  LiteralOrVariable(strconv.Itoa(int(char))),
		}).AddComment("Set char: "+string(char)))
	}

	// Add null terminator
	l.AppendInstruction(scopeName, NewSetWithOffsetInstruction(SetWithOffsetInstruction{
		Target: stringVar,
		Offset: LiteralOrVariable(strconv.Itoa(len(unquoted))),
		Value:  "0",
	}).AddComment("Null terminator"))

	// Store the translation for future use
	l.Program.UpsertTranslation(scopeName, literal, stringVar)

	return stringVar
}

func (l Listener) getOrCreateExpressionVariable(exprText string, scopeName ScopeName) VariableName {
	// Check if already computed
	if result, found := l.Program.GetVariableFor(exprText, scopeName); found {
		return result
	}

	// Check for literal
	if _, isLiteral := l.TypeChecker.GetLiteralType(exprText); isLiteral {
		return VariableName(exprText)
	}

	return l.Program.GetOrGenerateVariable(exprText, scopeName)
}

func getTACScope(scope *type_checker.Scope) ScopeName {
	scopeName := ScopeName(scope.Name)

	if scope.Type == type_checker.SCOPE_TYPES.CLASS {
		scopeName = ScopeName(scope.Name + "_" + type_checker.CONSTRUCTOR_NAME)
	}

	return scopeName
}
