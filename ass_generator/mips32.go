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

const MIPS32_WORD_BYTE_SIZE uint = 4

type Mips32Generator struct {
	listener            *tac_generator.Listener
	source              *bytes.Buffer
	program             *Mips32Program
	freeStackIdx        *uint
	nextScopeParamCount *uint
	stackSizeByScope    *map[string]uint
	currentParams       *[]string
}

func (m *Mips32Generator) UpsertSizeByScope(scopeName string, stackSize uint) {
	log.Printf("Setting scope `%s` size: %d", scopeName, stackSize)
	(*m.stackSizeByScope)[scopeName] = stackSize
}

func (m *Mips32Generator) RegisterParamToCount() uint {
	count := *m.nextScopeParamCount
	*m.nextScopeParamCount += 1
	return count
}

func (m *Mips32Generator) ResetNextScopeParamCount() {
	*m.nextScopeParamCount = 0
}

func (m *Mips32Generator) AllocateOnStack(size uint) uint {
	freeIdx := *m.freeStackIdx
	*m.freeStackIdx += size
	return freeIdx
}

func (m *Mips32Generator) ResetStackIdx() {
	*m.freeStackIdx = 0
}

func (m *Mips32Generator) StackAddress(idx int) string {
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

func NewMips32CommentInstruction(comment string) Mips32Instruction {
	return Mips32Instruction{
		Operation: lib.NewOpEmpty[Mips32Operation](),
		Section:   lib.NewOpEmpty[string](),
		Comment:   lib.NewOpValue(comment),
	}
}

func NewMips32SectionInstruction(section string) Mips32Instruction {
	return Mips32Instruction{
		Operation: lib.NewOpEmpty[Mips32Operation](),
		Section:   lib.NewOpValue(section),
		Comment:   lib.NewOpEmpty[string](),
	}
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
	TemporaryVars map[string]string // All variables t0-t9
	StackVars     map[string]string // All values inside the stack
	ParamMap      map[string]string // All regs s0-s7
	FreeRegisters lib.Stack[string]
	ParamVars     lib.Stack[string]
}

func NewMips32Program() *Mips32Program {
	freeRegisters := lib.NewStack[string]()

	for i := 9; i >= 0; i-- {
		freeRegisters.Push(
			fmt.Sprintf("$t%d", i),
		)
	}

	paramVars := lib.NewStack[string]()
	for i := 7; i >= 0; i-- {
		paramVars.Push(
			fmt.Sprintf("$s%d", i),
		)
	}

	return &Mips32Program{
		Data:          []Mips32DataDeclaration{},
		Text:          []Mips32Instruction{},
		TemporaryVars: map[string]string{},
		StackVars:     map[string]string{},
		ParamMap:      map[string]string{},
		FreeRegisters: freeRegisters,
		ParamVars:     paramVars,
	}
}

func (p *Mips32Program) UpsertParam(varName, reg string) {
	p.ParamMap[varName] = reg
}

func (p *Mips32Program) UpsertTemporary(varName, register string) {
	p.TemporaryVars[varName] = register
}

func (p *Mips32Program) UpsertRAM(varName, address string) {
	p.StackVars[varName] = address
}

func (p *Mips32Program) PopParam() string {
	register := p.ParamVars.Pop()
	if !register.HasValue() {
		log.Panicf("No free registers available!")
	}

	return register.GetValue()
}

func (p *Mips32Program) PushParam(reg string) {
	p.ParamVars.Push(reg)
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
		log.Printf("Loading literal `%s` on register: `%s`...\n", param, freeRegister)
		p.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "li",
			Params: NewMips32OperationParams(freeRegister, param),
		}))
		return freeRegister, true
	} else if address, found := p.TemporaryVars[param]; found {
		log.Printf("Temporary var `%s` found on register `%s`!\n", param, address)
		return address, true
	} else if paramReg, found := p.ParamMap[param]; found {
		log.Printf("Param var `%s` found on register `%s`!\n", param, paramReg)
		return paramReg, true
	} else if address, found := p.StackVars[param]; found {
		freeRegister := p.PopFreeRegister()
		log.Printf("RAM var `%s` found! Address: `%s`! Saving on `%s`...\n", param, address, freeRegister)
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
	var freeStackIdx uint
	var nextScopeParamCount uint

	return Mips32Generator{
		listener:            listener,
		source:              source,
		program:             NewMips32Program(),
		freeStackIdx:        &freeStackIdx,
		nextScopeParamCount: &nextScopeParamCount,
		currentParams:       &[]string{},
		stackSizeByScope:    &map[string]uint{},
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
	var callsAnotherProcedure bool

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
		case "=", "LOAD":
			stackSize += MIPS32_WORD_BYTE_SIZE
		case "ALLOC":
			customSize, err := strconv.ParseUint(parts[2], 10, 32)
			if err != nil {
				log.Panicf("Failed to parse `%s` as an integer on ALLOC!", parts[2])
			}
			stackSize += uint(customSize)
		case "CALL", "CALLRET":
			if !callsAnotherProcedure {
				stackSize += MIPS32_WORD_BYTE_SIZE
				callsAnotherProcedure = true
			}
		case "SEC":
			if stackSize%4 != 0 {
				stackSize += stackSize % 4
			}
			m.UpsertSizeByScope(scopeName, stackSize)

			scopeName = parts[1]
			stackSize = 0
		default:
			continue
		}
	}

	if _, found := (*m.stackSizeByScope)[scopeName]; !found {
		if stackSize%4 != 0 {
			stackSize += stackSize % 4
		}
		m.UpsertSizeByScope(scopeName, stackSize)
	}
}

func (m Mips32Generator) GenerateTo(buff *bytes.Buffer) error {
	m.ComputeScopeStackSizes()
	scopeName := type_checker.GLOBAL_SCOPE_NAME
	stackSize := m.GetStackSize(scopeName)

	m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
		OpCode: "addiu",
		Params: NewMips32OperationParams("$sp", "$sp", fmt.Sprintf("-%d", stackSize)),
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
.text
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

	onlyGlobalScope := len(*m.stackSizeByScope) == 1
	if onlyGlobalScope { //
		// Return $sp to normal
		m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "addiu",
			Params: NewMips32OperationParams("$sp", "$sp", strconv.FormatUint(uint64(stackSize), 10)),
		}))

		// Add:
		// jr $ra
		// at the end of main
		m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "jr",
			Params: NewMips32OperationParams("$ra"),
		}))
	}

	for _, inst := range m.program.Text {
		switch {
		case inst.Comment.HasValue():
			_, err = fmt.Fprintf(buff, "# %s", inst.Comment.GetValue())
			if err != nil {
				return err
			}

		case inst.Section.HasValue():
			_, err = fmt.Fprintf(buff, "%s", inst.Section.GetValue())
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

		stackAddress := m.StackAddress(int(m.AllocateOnStack(MIPS32_WORD_BYTE_SIZE)))

		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "sw",
			Params: NewMips32OperationParams(valueParam, stackAddress),
		}))
		program.UpsertRAM(varName, stackAddress)

	case "PARAM":
		value := params[0]
		valueParam, shouldFreeRegister := program.LoadOrDefault(value)
		if shouldFreeRegister {
			defer program.PushFreeRegister(valueParam)
		}

		// Back up previous value on stack and modify to actual param
		paramReg := program.PopParam()
		stackBackupAddr := m.StackAddress(int(m.AllocateOnStack(MIPS32_WORD_BYTE_SIZE)))
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "sw",
			Params: NewMips32OperationParams(paramReg, stackBackupAddr),
		}))
		*m.currentParams = append(*m.currentParams, stackBackupAddr)
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "move",
			Params: NewMips32OperationParams(paramReg, valueParam),
		}))

	case "LOAD":
		varName := params[0]
		paramReg := program.PopParam()
		program.UpsertParam(varName, paramReg)
		log.Printf("Variable `%s` retrieved from register: %s", varName, paramReg)

	case "SEC":
		// Add:
		// jr $ra
		// at the end of main (if programs contains more sections)
		if *secName == type_checker.GLOBAL_SCOPE_NAME {
			// Return $sp to normal
			m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "addiu",
				Params: NewMips32OperationParams("$sp", "$sp", strconv.FormatUint(uint64(stackSize), 10)),
			}))

			program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "jr",
				Params: NewMips32OperationParams("$ra"),
			}))
		}

		*secName = params[0]
		program.AppendInstruction(NewMips32SectionInstruction(*secName))
		stackSize = m.GetStackSize(*secName)
		m.ResetStackIdx()
		m.ResetNextScopeParamCount()

		m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "addiu",
			Params: NewMips32OperationParams("$sp", "$sp", fmt.Sprintf("-%d", stackSize)),
		}))

	case "RETURN":
		if len(params) > 0 {
			varName := params[0]
			varParam, shouldFreeRegister := m.program.LoadOrDefault(varName)
			if shouldFreeRegister {
				defer m.program.PushFreeRegister(varParam)
			}

			m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "move",
				Params: NewMips32OperationParams("$v0", varParam),
			}))
		}

		// Restore stack and jump back
		m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "addiu",
			Params: NewMips32OperationParams("$sp", "$sp", strconv.FormatUint(uint64(stackSize), 10)),
		}))
		m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "jr",
			Params: NewMips32OperationParams("$ra"),
		}))

		// Free param registers
		for i := len(*m.currentParams) - 1; i >= 0; i-- {
			paramReg := fmt.Sprintf("$s%d", i)
			m.program.PushParam(paramReg)
		}

	case "CALL":
		m.ResetNextScopeParamCount()
		newSecName := params[1]
		argCount := params[2]

		// Save $ra in the stack
		stackAddress := m.StackAddress(int(stackSize - MIPS32_WORD_BYTE_SIZE)) // Last reserved space from the stack
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

		// Restore current scope params
		for i := range argCount {
			paramReg := fmt.Sprintf("$s%d", i)
			stackAddress := (*m.currentParams)[i]
			program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "lw",
				Params: NewMips32OperationParams(paramReg, stackAddress),
			}))
			program.PopParam()
		}
	case "CALLRET":
		m.ResetNextScopeParamCount()
		varName := params[0]
		newSecName := params[1]
		argCount := params[2]

		// Save $ra in the stack
		stackAddress := m.StackAddress(int(stackSize - MIPS32_WORD_BYTE_SIZE)) // Last reserved space from the stack
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
			OpCode: "move",
			Params: NewMips32OperationParams(temporary, "$v0"),
		}))
		program.UpsertTemporary(varName, temporary)

		// Restore $ra from the stack
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "lw",
			Params: NewMips32OperationParams("$ra", stackAddress),
		}))

		// Restore current scope params
		for i := range argCount {
			paramReg := fmt.Sprintf("$s%d", i)
			stackAddress := (*m.currentParams)[i]
			program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "lw",
				Params: NewMips32OperationParams(paramReg, stackAddress),
			}))
			program.PopParam()
		}
	}

	return nil
}
