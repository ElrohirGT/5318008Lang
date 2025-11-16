package assgenerator

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ElrohirGT/5318008Lang/lib"
	"github.com/ElrohirGT/5318008Lang/tac_generator"
	"github.com/ElrohirGT/5318008Lang/type_checker"
)

type Mips32Generator struct {
	listener         *tac_generator.Listener
	source           *bytes.Buffer
	program          *Mips32Program
	lastUsedByte     *uint
	paramCount       *uint
	stackSizeByScope *map[string]uint
}

func (m *Mips32Generator) UpsertSizeByScope(scopeName string, stackSize uint) {
	log.Printf("Setting scope `%s` size: %d", scopeName, stackSize)
	(*m.stackSizeByScope)[scopeName] = stackSize
}

func (m *Mips32Generator) RegisterParamToCount() uint {
	count := *m.paramCount
	*m.paramCount += 1
	return count
}

func (m *Mips32Generator) ResetParamCount() {
	*m.paramCount = 0
}

func (m *Mips32Generator) GetNextFreeByteIdx() uint {
	*m.lastUsedByte++
	return *m.lastUsedByte
}

func (m *Mips32Generator) StackAddress(idx uint) string {
	return fmt.Sprintf("%d%s", idx, "($sp)")
}

type Mips32DataDeclaration struct {
	Name  string
	Type  string
	Value string
}

type Mips32Instruction struct {
	Operation lib.Optional[Mips32Operation]
	Comment   lib.Optional[string]
	Section   lib.Optional[string]
}

func NewMips32OperationInstruction(op Mips32Operation) Mips32Instruction {
	return Mips32Instruction{
		Operation: lib.NewOpValue(op),
		Comment:   lib.NewOpEmpty[string](),
		Section:   lib.NewOpEmpty[string](),
	}
}

type Mips32Operation struct {
	OpCode string
	Params []string
}

func NewMips32OperationParams(params ...string) []string {
	return params
}

type Mips32Program struct {
	Data          []Mips32DataDeclaration
	Text          []Mips32Instruction
	TemporaryVars map[string]string
	RAMVars       map[string]string
	FreeRegisters lib.Stack[string]
}

func NewMips32Program() *Mips32Program {
	freeRegisters := lib.NewStack[string]()
	for i := 10; i >= 0; i-- {
		freeRegisters.Push(
			fmt.Sprintf("$s%d", i),
		)
	}
	for i := 10; i >= 0; i-- {
		freeRegisters.Push(
			fmt.Sprintf("$t%d", i),
		)
	}

	return &Mips32Program{
		Data:          []Mips32DataDeclaration{},
		Text:          []Mips32Instruction{},
		TemporaryVars: map[string]string{},
		RAMVars:       map[string]string{},
		FreeRegisters: freeRegisters,
	}
}

func (p *Mips32Program) UpsertTemporary(varName, register string) {
	p.TemporaryVars[varName] = register
}

func (p *Mips32Program) UpsertRAM(varName, address string) {
	p.RAMVars[varName] = address
}

func (p *Mips32Program) PopFreeRegister() string {
	register := p.FreeRegisters.Pop()
	if !register.HasValue() {
		log.Panicf("No free registers available!")
	}

	return register.GetValue()
}

func (p *Mips32Program) PushFreeRegister(reg string) {
	p.FreeRegisters.Push(reg)
}

func (p *Mips32Program) LoadOrDefault(param string) (string, bool) {

	if _, err := strconv.Atoi(param); err == nil {
		freeRegister := p.PopFreeRegister()
		log.Printf("Loading literal (`%s`) on register: `%s`...\n", param, freeRegister)
		p.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "li",
			Params: NewMips32OperationParams(freeRegister, param),
		}))
		return freeRegister, true
	} else if address, found := p.TemporaryVars[param]; found {
		log.Printf("Temporary var `%s` found on register `%s`!\n", param, address)
		return address, true
	} else if address, found := p.RAMVars[param]; found {
		freeRegister := p.PopFreeRegister()
		log.Printf("RAM var `%s` found! Address: (%s)! Saving on `%s`...\n", param, address, freeRegister)
		p.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "lw",
			Params: NewMips32OperationParams(freeRegister, address),
		}))
		return freeRegister, false
	} else {
		log.Panicf("Variable `%s` not found in registers nor RAM!", param)
	}

	return "**Invalid**", false
}

func (p *Mips32Program) AppendDataDeclaration(dc Mips32DataDeclaration) {
	p.Data = append(p.Data, dc)
}

func (p *Mips32Program) AppendInstruction(inst Mips32Instruction) {
	p.Text = append(p.Text, inst)
}

func NewMips32Generator(listener *tac_generator.Listener, source *bytes.Buffer) Mips32Generator {
	var lastUsedByte uint
	return Mips32Generator{
		listener:         listener,
		source:           source,
		program:          NewMips32Program(),
		lastUsedByte:     &lastUsedByte,
		stackSizeByScope: &map[string]uint{},
	}
}

func (m Mips32Generator) GetStackSize(secName string) uint {
	stackSize, found := (*m.stackSizeByScope)[secName]
	if !found {
		log.Panicf("Failed to find size of stack for section: %s", secName)
	}

	return stackSize
}

func (m Mips32Generator) ComputeScopeStackSizes() {
	scopeName := m.listener.TypeChecker.ScopeManager.GlobaScope.Name
	log.Printf("Global scope Name: %s\n", scopeName)
	var stackSize uint = 0

	var sourceCopy bytes.Buffer
	_, err := sourceCopy.Write(m.source.Bytes())
	if err != nil {
		log.Panicf("Failed to copy TAC bytes: %s", err)
	}

	scan := bufio.NewScanner(&sourceCopy)
	for scan.Scan() {
		line := scan.Text()

		parts := strings.Split(strings.TrimSpace(line), " ")
		opCode := parts[0]
		switch opCode {
		case "=":
			stackSize += 4
		case "ALLOC":
			customSize, err := strconv.ParseUint(parts[2], 10, 32)
			if err != nil {
				log.Panicf("Failed to parse `%s` as an integer on ALLOC!", parts[2])
			}
			stackSize += uint(customSize)
		case "LOAD":
			stackSize += 4
		case "SEC":
			m.UpsertSizeByScope(scopeName, stackSize)
			scopeName = parts[1]
			stackSize = 0
		default:
			continue
		}
	}

	if _, found := (*m.stackSizeByScope)[scopeName]; !found {
		m.UpsertSizeByScope(scopeName, stackSize)
	}
}

func (m Mips32Generator) GenerateTo(buff *bytes.Buffer) error {
	m.ComputeScopeStackSizes()
	scopeName := type_checker.GLOBAL_SCOPE_NAME
	stackSize := m.GetStackSize(scopeName)

	m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
		OpCode: "addiu",
		Params: []string{"$sp", "$sp", fmt.Sprintf("-%d", stackSize)},
	}))

	scan := bufio.NewScanner(m.source)
	for scan.Scan() {
		line := scan.Text()
		parts := strings.Split(strings.TrimSpace(line), " ")
		opCode := parts[0]

		log.Printf("Parsing line: %s", line)
		err := m.translate(&scopeName, opCode, parts[1:])
		if err != nil {
			return err
		}
	}

	buff.WriteString(`
# Code generated by 5318008Lang compiler.
# Checkout our repository here: https://github.com/ElrohirGT/5318008Lang
# The following code is designed to run on the compiler: https://cpulator.01xz.net/?sys=mipsr5-spim
`)

	_, err := buff.WriteString(".data")
	if err != nil {
		return err
	}

	for _, decl := range m.program.Data {
		_, err = fmt.Fprintf(buff,
			"%s:\t\t%s\t\t%s\n", decl.Name, decl.Type, decl.Value)
		if err != nil {
			return err
		}
	}

	_, err = buff.WriteString(`
# Main start of our program
.global _start
_start:
	jal main
	li $v0, 10
	syscall		# Use syscall 10 to stop simulation

main:
`)
	if err != nil {
		return err
	}

	for _, inst := range m.program.Text {
		switch {
		case inst.Comment.HasValue():
			_, err = fmt.Fprintf(buff, "# %s", inst.Comment.GetValue())
			if err != nil {
				return err
			}

		case inst.Section.HasValue():
			_, err = fmt.Fprintf(buff, "%s:", inst.Section.GetValue())
			if err != nil {
				return err
			}

		case inst.Operation.HasValue():
			op := inst.Operation.GetValue()
			_, err = fmt.Fprintf(buff, "\t%s", op.OpCode)
			if err != nil {
				return err
			}

			for i, param := range op.Params {
				_, err = fmt.Fprintf(buff, " %s", param)
				if err != nil {
					return err
				}

				if i+1 < len(op.Params) {
					_, err = buff.WriteRune(',')
					if err != nil {
						return err
					}
				}
			}
		}

		_, err = buff.WriteRune('\n')
		if err != nil {
			return err
		}
	}

	return err
}

func (m *Mips32Generator) translate(secName *string, opCode string, params []string) error {
	program := m.program
	stackSize := m.GetStackSize(*secName)

	manageAddSubOp := func(opCode string) {
		varName := params[0]
		a := params[1]
		b := params[2]

		aParam, shouldFreeRegister := program.LoadOrDefault(a)
		if shouldFreeRegister {
			defer program.PushFreeRegister(aParam)
		}

		bParam, shouldFreeRegister := program.LoadOrDefault(b)
		if shouldFreeRegister {
			defer program.PushFreeRegister(bParam)
		}

		freeRegister := program.PopFreeRegister()
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: opCode,
			Params: NewMips32OperationParams(freeRegister, aParam, bParam),
		}))

		program.UpsertTemporary(varName, freeRegister)
	}

	manageDivMultOp := func(opCode string) {
		varName := params[0]
		a := params[1]
		b := params[2]

		aParam, shouldFreeRegister := program.LoadOrDefault(a)
		if shouldFreeRegister {
			defer program.PushFreeRegister(aParam)
		}

		bParam, shouldFreeRegister := program.LoadOrDefault(b)
		if shouldFreeRegister {
			defer program.PushFreeRegister(bParam)
		}

		freeRegister := program.PopFreeRegister()
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: opCode,
			Params: NewMips32OperationParams(aParam, bParam),
		}))

		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "mflo",
			Params: NewMips32OperationParams(freeRegister),
		}))

		program.UpsertTemporary(varName, freeRegister)
	}

	switch opCode {
	case "ADD": // a+b
		log.Println("Adding ADD operation")
		manageAddSubOp("add")
	case "SUB": // a-b
		log.Println("Adding SUB operation")
		manageAddSubOp("sub")

	case "MULT": // a*b
		manageDivMultOp("mult")
	case "DIV": // a/b
		manageDivMultOp("div")

	case "=":
		varName := params[0]
		value := params[2]

		valueParam, shouldFreeRegister := program.LoadOrDefault(value)
		if shouldFreeRegister {
			defer program.PushFreeRegister(valueParam)
		}

		stackAddress := m.StackAddress(m.GetNextFreeByteIdx())

		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "sw",
			Params: NewMips32OperationParams(valueParam, stackAddress),
		}))
		program.UpsertRAM(varName, stackAddress)

	case "PARAM":
		value := params[1]
		valueParam, shouldFreeRegister := program.LoadOrDefault(value)
		if shouldFreeRegister {
			defer program.PushFreeRegister(valueParam)
		}

		stackAddress := m.StackAddress(stackSize + m.RegisterParamToCount())
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "sw",
			Params: NewMips32OperationParams(valueParam, stackAddress),
		}))

	case "LOAD":
		varName := params[1]
		stackAddress := m.StackAddress(m.RegisterParamToCount())
		program.UpsertRAM(varName, stackAddress)

	case "CALL":
		newSecName := params[1]
		// argCount := params[2]

		stackAddress := m.StackAddress(stackSize - 4) // Last reserved space from the stack
		// Save $ra in the stack
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "sw",
			Params: NewMips32OperationParams("$ra", stackAddress),
		}))

		// Call method
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "jal",
			Params: NewMips32OperationParams(newSecName),
		}))

		// Restore $ra from the stack
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "lw",
			Params: NewMips32OperationParams("$ra", stackAddress),
		}))
	case "CALLRET":
		varName := params[1]
		newSecName := params[2]
		// argCount := params[3]

		stackAddress := m.StackAddress(stackSize - 4) // Last reserved space from the stack
		// Save $ra in the stack
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "sw",
			Params: NewMips32OperationParams("$ra", stackAddress),
		}))

		// Call method
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "jal",
			Params: NewMips32OperationParams(newSecName),
		}))

		// Save return value on temporary
		temporary := program.PopFreeRegister()
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "mv",
			Params: NewMips32OperationParams(temporary),
		}))
		program.UpsertTemporary(varName, temporary)

		// Restore $ra from the stack
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "lw",
			Params: NewMips32OperationParams("$ra", stackAddress),
		}))
	}

	return nil
}
