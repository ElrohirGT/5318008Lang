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
	listener         *tac_generator.Listener
	source           *bytes.Buffer
	program          *Mips32Program
	freeStackIdx     *uint
	stackSizeByScope *map[string]uint
}

func (m *Mips32Generator) UpsertSizeByScope(scopeName string, stackSize uint) {
	log.Printf("Setting scope `%s` size: %d", scopeName, stackSize)
	(*m.stackSizeByScope)[scopeName] = stackSize
}

func (m *Mips32Generator) AllocateOnStack(size uint) uint {
	freeIdx := *m.freeStackIdx
	*m.freeStackIdx += size
	return freeIdx
}

func (m *Mips32Generator) ResetStackIdx() {
	*m.freeStackIdx = 0
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

type TACVariableOrValue string
type Mips32Register string
type StackAddress int

func (s StackAddress) String() string {
	return fmt.Sprintf("%d($sp)", s)
}

func (s StackAddress) Add(i int) StackAddress {
	return StackAddress(int(s) + i)
}

type Mips32Program struct {
	Data               []Mips32DataDeclaration
	Text               []Mips32Instruction
	TemporaryVars      map[TACVariableOrValue]Mips32Register // All variables to t0-t9
	StackVars          map[TACVariableOrValue]StackAddress   // All vars to values inside the stack
	ParamMap           map[TACVariableOrValue]Mips32Register // All vars to regs s0-s7
	FreeTemporaryRegs  lib.Stack[Mips32Register]             // Currently free t0-t9 regs
	SavedParamsAddress []StackAddress
	LoadCount          uint // How many `LOAD` have we passed yet?
}

func NewMips32Program() *Mips32Program {
	freeRegisters := lib.NewStack[Mips32Register]()

	for i := 9; i >= 0; i-- {
		freeRegisters.Push(
			Mips32Register(fmt.Sprintf("$t%d", i)),
		)
	}

	program := &Mips32Program{
		Data:              []Mips32DataDeclaration{},
		Text:              []Mips32Instruction{},
		TemporaryVars:     map[TACVariableOrValue]Mips32Register{},
		StackVars:         map[TACVariableOrValue]StackAddress{},
		ParamMap:          map[TACVariableOrValue]Mips32Register{},
		FreeTemporaryRegs: freeRegisters,
	}
	program.ResetParams()

	return program
}

func (p *Mips32Program) UpsertParam(varName TACVariableOrValue, reg Mips32Register) {
	p.ParamMap[varName] = reg
}

func (p *Mips32Program) UpsertTemporary(varName TACVariableOrValue, register Mips32Register) {
	p.TemporaryVars[varName] = register
}

func (p *Mips32Program) UpsertRAM(varName TACVariableOrValue, address StackAddress) {
	p.StackVars[varName] = address
}

func (p *Mips32Program) ReserveParam(address StackAddress) Mips32Register {
	if len(p.SavedParamsAddress) == 8 {
		log.Panicf("No empty `$sN` registers available!")
	}

	count := len(p.SavedParamsAddress)
	p.SavedParamsAddress = append(p.SavedParamsAddress, address)
	return Mips32Register(fmt.Sprintf("$s%d", count))
}

func (p *Mips32Program) ResetParams() {
	p.SavedParamsAddress = []StackAddress{}
}

func (p *Mips32Program) ReserveLoadParam() Mips32Register {
	if p.LoadCount == 8 {
		log.Panicf("No empty `$sN` registers available!")
	}

	count := p.LoadCount
	p.LoadCount++

	return Mips32Register(fmt.Sprintf("$s%d", count))
}

func (p *Mips32Program) ResetLoadParams() {
	p.LoadCount = 0
}

func (p *Mips32Program) PopFreeTemporary() Mips32Register {
	register := p.FreeTemporaryRegs.Pop()
	if !register.HasValue() {
		log.Panicf("No free registers available!")
	}

	return register.GetValue()
}

func (p *Mips32Program) PushFreeTemporary(reg Mips32Register) {
	p.FreeTemporaryRegs.Push(reg)
}

func (p *Mips32Program) LoadOrDefault(param TACVariableOrValue) (Mips32Register, bool) {

	if _, err := strconv.Atoi(string(param)); err == nil {
		freeRegister := p.PopFreeTemporary()
		log.Printf("Loading literal `%s` on register: `%s`...\n", param, freeRegister)
		p.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "li",
			Params: NewMips32OperationParams(string(freeRegister), string(param)),
		}))
		return freeRegister, true
	} else if address, found := p.TemporaryVars[param]; found {
		log.Printf("Temporary var `%s` found on register `%s`!\n", param, address)
		return address, true
	} else if paramReg, found := p.ParamMap[param]; found {
		log.Printf("Param var `%s` found on register `%s`!\n", param, paramReg)
		return paramReg, false
	} else if address, found := p.StackVars[param]; found {
		freeRegister := p.PopFreeTemporary()
		log.Printf("RAM var `%s` found! Address: `%s`! Saving on `%s`...\n", param, address, freeRegister)
		p.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "lw",
			Params: NewMips32OperationParams(string(freeRegister), address.String()),
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

	return Mips32Generator{
		listener:         listener,
		source:           source,
		program:          NewMips32Program(),
		freeStackIdx:     &freeStackIdx,
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
	var stackSize uint
	var maxParamCount uint64
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
			stackSize += MIPS32_WORD_BYTE_SIZE // Space reserved for saving the ref to this memory.
		case "CALL", "CALLRET":
			idx := 3
			if opCode == "CALL" {
				idx = 2
			}
			paramCount, err := strconv.ParseUint(parts[idx], 10, 0)
			if err != nil {
				log.Panicf("Failed to parse CALL/CALLRET param quantity `%s`", parts[idx])
			}
			maxParamCount = max(maxParamCount, paramCount)

			if !callsAnotherProcedure {
				stackSize += MIPS32_WORD_BYTE_SIZE
				callsAnotherProcedure = true
			}
		case "FUNC":
			stackSize += uint(maxParamCount) * MIPS32_WORD_BYTE_SIZE
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
		stackSize += uint(maxParamCount) * MIPS32_WORD_BYTE_SIZE
		if stackSize%4 != 0 {
			stackSize += stackSize % 4
		}
		m.UpsertSizeByScope(scopeName, stackSize)
	}
}

// Buffer to fill, addBuiltins = flag to add builin functions at the end.
func (m Mips32Generator) GenerateTo(buff *bytes.Buffer, addBuiltins bool) error {
	m.ComputeScopeStackSizes()
	functionName := type_checker.GLOBAL_SCOPE_NAME
	// Main should have a minimum stack size of 7 param registers
	stackSize := m.GetStackSize(functionName)

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
		err := m.translate(&functionName, opCode, parts[1:])
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

	// onlyGlobalScope := len(*m.stackSizeByScope) == 1
	// if onlyGlobalScope { //
	// 	// Return $sp to normal
	// 	m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
	// 		OpCode: "addiu",
	// 		Params: NewMips32OperationParams("$sp", "$sp", strconv.FormatUint(uint64(stackSize), 10)),
	// 	}))
	//
	// 	// Add:
	// 	// jr $ra
	// 	// at the end of main
	// 	m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
	// 		OpCode: "jr",
	// 		Params: NewMips32OperationParams("$ra"),
	// 	}))
	// }

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

	// Add builtins
	if addBuiltins {
		m.footer(buff)
	}

	return err
}

func (m *Mips32Generator) translate(functionName *string, opCode string, params []string) error {
	program := m.program
	stackSize := m.GetStackSize(*functionName)

	manageAddSubOp := func(opCode string) {
		varName := params[0]
		a := params[1]
		b := params[2]

		aParam, shouldFreeRegister := program.LoadOrDefault(TACVariableOrValue(a))
		if shouldFreeRegister {
			defer program.PushFreeTemporary(aParam)
		}

		bParam, shouldFreeRegister := program.LoadOrDefault(TACVariableOrValue(b))
		if shouldFreeRegister {
			defer program.PushFreeTemporary(bParam)
		}

		freeRegister := program.PopFreeTemporary()
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: opCode,
			Params: NewMips32OperationParams(string(freeRegister), string(aParam), string(bParam)),
		}))

		program.UpsertTemporary(TACVariableOrValue(varName), freeRegister)
	}

	manageDivMultOp := func(opCode string) {
		varName := params[0]
		a := params[1]
		b := params[2]

		aParam, shouldFreeRegister := program.LoadOrDefault(TACVariableOrValue(a))
		if shouldFreeRegister {
			defer program.PushFreeTemporary(aParam)
		}

		bParam, shouldFreeRegister := program.LoadOrDefault(TACVariableOrValue(b))
		if shouldFreeRegister {
			defer program.PushFreeTemporary(bParam)
		}

		freeRegister := program.PopFreeTemporary()
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: opCode,
			Params: NewMips32OperationParams(string(aParam), string(bParam)),
		}))

		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "mflo",
			Params: NewMips32OperationParams(string(freeRegister)),
		}))

		program.UpsertTemporary(TACVariableOrValue(varName), freeRegister)
	}

	manageSetWithOffset := func(opCode string) {
		varName := TACVariableOrValue(params[0])
		offset := TACVariableOrValue(params[1])
		value := TACVariableOrValue(params[2])

		offsetReg, shouldFreeRegister := m.program.LoadOrDefault(offset)
		if shouldFreeRegister {
			defer program.PushFreeTemporary(offsetReg)
		}

		vParam, shouldFreeRegister := m.program.LoadOrDefault(value)
		if shouldFreeRegister {
			defer program.PushFreeTemporary(vParam)
		}

		freeReg := program.PopFreeTemporary()
		defer program.PushFreeTemporary(freeReg)

		if stackAddress, found := program.StackVars[varName]; found {
			// Load reference to object from stack
			program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "lw",
				Params: NewMips32OperationParams(string(freeReg), stackAddress.String()),
			}))
		} else if reg, found := program.ParamMap[varName]; found {
			// Load reference to object from params
			program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "move",
				Params: NewMips32OperationParams(string(freeReg), string(reg)),
			}))
		} else {
			log.Panicf("Failed to find `%s` on stack or params!", varName)
		}

		// Add to reference
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "addu",
			Params: NewMips32OperationParams(string(freeReg), string(freeReg), string(offsetReg)),
		}))
		// Load value from reference
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: opCode,
			Params: NewMips32OperationParams(string(vParam), "0("+string(freeReg)+")"),
		}))
	}

	manageLoadWithOffset := func(opCode string) {
		varName := TACVariableOrValue(params[0])
		objVarName := TACVariableOrValue(params[1])
		offset := TACVariableOrValue(params[2])

		offsetReg, shouldFreeRegister := m.program.LoadOrDefault(offset)
		if shouldFreeRegister {
			defer program.PushFreeTemporary(offsetReg)
		}

		nameParam := program.PopFreeTemporary()
		program.UpsertTemporary(varName, nameParam)

		freeReg := program.PopFreeTemporary()
		defer program.PushFreeTemporary(freeReg)

		if stackAddress, found := program.StackVars[objVarName]; found {
			// Load reference to object from stack
			program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "lw",
				Params: NewMips32OperationParams(string(freeReg), stackAddress.String()),
			}))
		} else if reg, found := program.ParamMap[objVarName]; found {
			// Load reference to object from params
			program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "move",
				Params: NewMips32OperationParams(string(freeReg), string(reg)),
			}))
		} else {
			log.Panicf("Failed to find `%s` on stack or params!", varName)
		}

		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "addu",
			Params: NewMips32OperationParams(string(freeReg), string(freeReg), string(offsetReg)),
		}))
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: opCode,
			Params: NewMips32OperationParams(string(nameParam), "0("+string(freeReg)+")"),
		}))
	}

	manageComparisons := func(opCode string) {
		varName := TACVariableOrValue(params[0])
		a := TACVariableOrValue(params[1])
		b := TACVariableOrValue(params[2])

		aParam, shouldFreeRegister := program.LoadOrDefault(a)
		if shouldFreeRegister {
			defer program.PushFreeTemporary(aParam)
		}

		bParam, shouldFreeRegister := program.LoadOrDefault(b)
		if shouldFreeRegister {
			defer program.PushFreeTemporary(bParam)
		}

		freeRegister := program.PopFreeTemporary()
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: opCode,
			Params: NewMips32OperationParams(string(freeRegister), string(aParam), string(bParam)),
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

	case "GT": // a>b
		manageComparisons("sgt")
	case "GTE": // a<=b
		manageComparisons("sge")
	case "LT": // a<b
		manageComparisons("slt")
	case "LTE": // a<=b
		manageComparisons("sle")

	case "OR":
		varName := params[0]
		a := params[1]
		b := params[2]

		aParam, shouldFreeRegister := program.LoadOrDefault(TACVariableOrValue(a))
		if shouldFreeRegister {
			defer program.PushFreeTemporary(aParam)
		}

		bParam, shouldFreeRegister := program.LoadOrDefault(TACVariableOrValue(b))
		if shouldFreeRegister {
			defer program.PushFreeTemporary(bParam)
		}

		freeRegister := program.PopFreeTemporary()
		// OR
		// -1 + -1 	= -2
		// 0 + -1 	= -1
		// -1 + 0 	= -1
		// 0 + 0 		= 0
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "add",
			Params: NewMips32OperationParams(string(freeRegister), string(aParam), string(bParam)),
		}))
		program.UpsertTemporary(TACVariableOrValue(varName), freeRegister)

	case "AND": // a && b
		varName := params[0]
		a := params[1]
		b := params[2]

		aParam, shouldFreeRegister := program.LoadOrDefault(TACVariableOrValue(a))
		if shouldFreeRegister {
			defer program.PushFreeTemporary(aParam)
		}

		bParam, shouldFreeRegister := program.LoadOrDefault(TACVariableOrValue(b))
		if shouldFreeRegister {
			defer program.PushFreeTemporary(bParam)
		}

		freeRegister := program.PopFreeTemporary()
		// AND
		// -1 * -1 	= 1
		// 0 * -1 	= 0
		// -1 * 0 	= 0
		// 0 * 0 		= 0
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "mult",
			Params: NewMips32OperationParams(string(aParam), string(bParam)),
		}))
		// Extract result:
		// * 1: Both were true!
		// * 0: Both were false!
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "mflo",
			Params: NewMips32OperationParams(string(freeRegister)),
		}))
		program.UpsertTemporary(TACVariableOrValue(varName), freeRegister)

	case "=":
		varName := TACVariableOrValue(params[0])
		value := TACVariableOrValue(params[2])

		valueParam, shouldFreeRegister := program.LoadOrDefault(TACVariableOrValue(value))
		if shouldFreeRegister {
			defer program.PushFreeTemporary(valueParam)
		}

		stackAddress, found := program.StackVars[varName]
		if !found {
			stackAddress = StackAddress(int(m.AllocateOnStack(MIPS32_WORD_BYTE_SIZE)))
		}

		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "sw",
			Params: NewMips32OperationParams(string(valueParam), stackAddress.String()),
		}))
		program.UpsertRAM(varName, stackAddress)

	case "IF":
		condition := TACVariableOrValue(params[0])
		target := params[2]

		negates := params[0] == "NOT"
		if negates {
			condition = TACVariableOrValue(params[1])
			target = params[3]
		}

		conditionParam, shouldFreeRegister := program.LoadOrDefault(condition)
		if shouldFreeRegister {
			defer program.PushFreeTemporary(conditionParam)
		}

		if negates {
			// Since 0 is false.
			// IF NOT <condition> GOTO <target>, only goes to target if condition is false.
			// So we check if condition is 0.
			program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "beq",
				Params: NewMips32OperationParams(string(conditionParam), "$zero", target),
			}))
		} else {
			// By the same token.
			// IF <condition> GOTO <target>, only goes to target if condition is true.
			// So we check if condition is not 0.
			program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "bne",
				Params: NewMips32OperationParams(string(conditionParam), "$zero", target),
			}))
		}

	case "PARAM":
		value := TACVariableOrValue(params[0])
		valueParam, shouldFreeRegister := program.LoadOrDefault(value)
		if shouldFreeRegister {
			defer program.PushFreeTemporary(valueParam)
		}

		// Back up previous value on stack and modify to actual param
		stackBackupAddr := StackAddress(int(m.AllocateOnStack(MIPS32_WORD_BYTE_SIZE)))
		paramReg := program.ReserveParam(stackBackupAddr)
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "sw",
			Params: NewMips32OperationParams(string(paramReg), stackBackupAddr.String()),
		}))
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "move",
			Params: NewMips32OperationParams(string(paramReg), string(valueParam)),
		}))

	case "LOAD":
		varName := TACVariableOrValue(params[0])
		paramReg := program.ReserveLoadParam()
		program.UpsertParam(varName, paramReg)
		log.Printf("Variable `%s` retrieved from register: %s", varName, paramReg)

	case "SEC":
		sectionName := params[0]
		program.AppendInstruction(NewMips32SectionInstruction(sectionName))

	case "END": // Marks the end of a function
		m.ResetStackIdx()
		program.ResetLoadParams()
		// Return $sp to normal
		m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "addiu",
			Params: NewMips32OperationParams("$sp", "$sp", strconv.FormatUint(uint64(stackSize), 10)),
		}))

		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "jr",
			Params: NewMips32OperationParams("$ra"),
		}))

	case "ALLOC":
		varName := TACVariableOrValue(params[0])
		size, err := strconv.ParseUint(params[1], 10, 0)
		if err != nil {
			log.Panicf("Failed to parse `%s` as an uint for ALLOC instruction", params[1])
		}

		idx := m.AllocateOnStack(MIPS32_WORD_BYTE_SIZE)
		refStackAddress := StackAddress(int(idx))
		program.UpsertRAM(varName, refStackAddress)

		idx = m.AllocateOnStack(uint(size))
		contentStackAddress := StackAddress(int(idx))

		freeReg := program.PopFreeTemporary()
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "la",
			Params: NewMips32OperationParams(string(freeReg), contentStackAddress.String()),
		}))
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "sw",
			Params: NewMips32OperationParams(string(freeReg), refStackAddress.String()),
		}))

	case "SBO":
		manageSetWithOffset("sb")
	case "SWO":
		manageSetWithOffset("sw")

	case "LBO":
		manageLoadWithOffset("lb")
	case "LWO":
		manageLoadWithOffset("lw")

	case "FUNC":
		*functionName = params[0]
		program.AppendInstruction(NewMips32SectionInstruction(*functionName))
		stackSize = m.GetStackSize(*functionName)
		m.ResetStackIdx()
		program.ResetLoadParams()

		m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "addiu",
			Params: NewMips32OperationParams("$sp", "$sp", fmt.Sprintf("-%d", stackSize)),
		}))

	case "GOTO":
		target := params[0]
		log.Printf("GOTO -> b %s", target)
		m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "b",
			Params: NewMips32OperationParams(target),
		}))

	case "RETURN":
		if len(params) > 0 {
			varName := TACVariableOrValue(params[0])
			varParam, shouldFreeRegister := m.program.LoadOrDefault(varName)
			if shouldFreeRegister {
				defer m.program.PushFreeTemporary(varParam)
			}

			m.program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "move",
				Params: NewMips32OperationParams("$v0", string(varParam)),
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
		program.ResetLoadParams()

	case "CALL":
		newSecName := params[0]
		// argCount, err := strconv.ParseUint(params[1], 10, 0)
		// if err != nil {
		// 	log.Panicf("Failed to parse arg count `%s` as integer", params[1])
		// }

		// Save $ra in the stack
		stackAddress := StackAddress(stackSize - MIPS32_WORD_BYTE_SIZE) // Last reserved space from the stack
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "sw",
			Params: NewMips32OperationParams("$ra", stackAddress.String()),
		}))

		// Call method
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "jal",
			Params: NewMips32OperationParams(newSecName),
		}))

		// Restore $ra from the stack
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "lw",
			Params: NewMips32OperationParams("$ra", stackAddress.String()),
		}))

		// Restore current scope params
		for i, stackAddress := range program.SavedParamsAddress {
			paramReg := fmt.Sprintf("$s%d", i)
			log.Printf("Restoring param `%s` from `%s`", paramReg, stackAddress)
			program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "lw",
				Params: NewMips32OperationParams(paramReg, stackAddress.String()),
			}))
		}
		program.ResetParams()
	case "CALLRET":
		varName := TACVariableOrValue(params[0])
		newSecName := params[1]
		// argCount, err := strconv.ParseUint(params[2], 10, 0)
		// if err != nil {
		// 	log.Panicf("Failed to parse arg count `%s` as integer", params[2])
		// }

		// Save $ra in the stack
		stackAddress := StackAddress(int(stackSize - MIPS32_WORD_BYTE_SIZE)) // Last reserved space from the stack
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "sw",
			Params: NewMips32OperationParams("$ra", stackAddress.String()),
		}))

		// Call method
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "jal",
			Params: NewMips32OperationParams(newSecName),
		}))

		// Save return value on temporary
		temporary := program.PopFreeTemporary()
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "move",
			Params: NewMips32OperationParams(string(temporary), "$v0"),
		}))
		program.UpsertTemporary(varName, temporary)

		// Restore $ra from the stack
		program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
			OpCode: "lw",
			Params: NewMips32OperationParams("$ra", stackAddress.String()),
		}))

		// Restore current scope params
		for i, stackAddress := range program.SavedParamsAddress {
			paramReg := fmt.Sprintf("$s%d", i)
			log.Printf("Restoring param `%s` from `%s`", paramReg, stackAddress)
			program.AppendInstruction(NewMips32OperationInstruction(Mips32Operation{
				OpCode: "lw",
				Params: NewMips32OperationParams(paramReg, stackAddress.String()),
			}))
		}
		program.ResetParams()
	}

	return nil
}

func (m *Mips32Generator) footer(buff *bytes.Buffer) {

	buff.WriteString("\n\n\n\n")
	buff.WriteString(`
# Print String
# ========================================
#   $a0 - pointer to null-terminated string
# ========================================
print:
    move $t0, $s0           # $t0 = pointer to string
print_loop:
    lb   $t1, 0($t0)        # Load byte from string
    beq  $t1, $zero, print_end  # If null terminator, exit
    # Print character using syscall 11
    move $a0, $t1           # $a0 = character to print
    li   $v0, 11            # syscall 11 = print character
    syscall
    addiu $t0, $t0, 1        # Move to next character
    j    print_loop
print_end:
	li $a0, 10
	li $v0, 11
    syscall
    jr	$ra

# String Comparison Function
# ========================================
#   $a0 - pointer to first string
#   $a1 - pointer to second string
# Returns:
#   $v0 - 1 if strings are equal, 0 if different
# ========================================
strcmp:
    li      $t0, 0              # index = 0
strcmp_loop:
    add     $t1, $s0, $t0       # address of str1[i]
    add     $t2, $s1, $t0       # address of str2[i]
    
    lb      $t3, 0($t1)         # load byte from str1[i]
    lb      $t4, 0($t2)         # load byte from str2[i]
    
    bne     $t3, $t4, strcmp_not_equal  # if str1[i] != str2[i], not equal
    
    beq     $t3, $zero, strcmp_equal    # if str1[i] == '\0', both ended -> equal
    
    addi    $t0, $t0, 1         # i++
    j       strcmp_loop
strcmp_equal:
    li      $v0, 1              # return true (1)
    jr      $ra
strcmp_not_equal:
    li      $v0, 0              # return false (0)
    jr      $ra

# String Concatenation
# ========================================
# Arguments:
#   $s0 - pointer to first string (str1)
#   $s1 - pointer to second string (str2)
#   $s2 - pointer to destination buffer
# Returns:
#   $v0 - pointer to destination buffer (same as $s2)
# ========================================
concat_str:
    move $t0, $s0               # $t0 = str1 pointer
    move $t1, $s1               # $t1 = str2 pointer
    move $t2, $s2               # $t2 = dest buffer pointer
    
    # Step 1: Copy first string to buffer
concat_str_copy_first:
    lb   $t3, 0($t0)            # Load byte from str1
    beq  $t3, $zero, concat_str_copy_second  # If null, done with first string
    sb   $t3, 0($t2)            # Store byte to dest
    addiu $t0, $t0, 1           # Move str1 pointer
    addiu $t2, $t2, 1           # Move dest pointer
    j    concat_str_copy_first
    
    # Step 2: Copy second string to buffer
concat_str_copy_second:
    lb   $t3, 0($t1)            # Load byte from str2
    beq  $t3, $zero, concat_str_done  # If null, done
    sb   $t3, 0($t2)            # Store byte to dest
    addiu $t1, $t1, 1           # Move str2 pointer
    addiu $t2, $t2, 1           # Move dest pointer
    j    concat_str_copy_second
    
    # Step 3: Add null terminator
concat_str_done:
    sb   $zero, 0($t2)          # Add '\0' at end
    move $v0, $s2               # Return pointer to dest buffer
    jr   $ra

# Boolean to String Conversion
# ========================================
#   $s0 - boolean value (word: 0 = false, non-zero = true)
#   $s1 - pointer to destination buffer (min 6 bytes)
# Returns:
#   $v0 - pointer to destination buffer (same as $s1)
# Note: Result is null-terminated
#       Buffer must have at least 6 bytes for "false\0"
# ========================================
bool_to_str:
    move $t0, $s1               # $t0 = buffer pointer
    
    # Check if zero (false) or non-zero (true)
    beqz $s0, bool_to_str_false
    
    # Build "true" in buffer
    li   $t1, 116               # 't'
    sb   $t1, 0($t0)
    li   $t1, 114               # 'r'
    sb   $t1, 1($t0)
    li   $t1, 117               # 'u'
    sb   $t1, 2($t0)
    li   $t1, 101               # 'e'
    sb   $t1, 3($t0)
    sb   $zero, 4($t0)          # '\0'
    move $v0, $s1               # Return buffer pointer
    jr   $ra
    
bool_to_str_false:
    # Build "false" in buffer
    li   $t1, 102               # 'f'
    sb   $t1, 0($t0)
    li   $t1, 97                # 'a'
    sb   $t1, 1($t0)
    li   $t1, 108               # 'l'
    sb   $t1, 2($t0)
    li   $t1, 115               # 's'
    sb   $t1, 3($t0)
    li   $t1, 101               # 'e'
    sb   $t1, 4($t0)
    sb   $zero, 5($t0)          # '\0'
    move $v0, $s1               # Return buffer pointer
    jr   $ra
	
# ========================================
# Integer to String Conversion
# ========================================
# Arguments:
#   $s0 - integer value to convert (word)
#   $s1 - pointer to destination buffer (min 12 bytes)
# Returns:
#   $v0 - pointer to destination buffer (same as $s1)
# Note: Result is null-terminated
#       Buffer must have at least 12 bytes
# ========================================
int_to_str:
    move $t0, $s1           # $t0 = buffer pointer
    move $t1, $s0           # $t1 = number to convert
    li   $t2, 0             # $t2 = is_negative flag
    
    # Special case: handle 0
    bnez $t1, int_to_str_check_negative
    li   $t3, 48            # ASCII '0'
    sb   $t3, 0($t0)
    sb   $zero, 1($t0)      # Null terminator
    move $v0, $s1
    jr   $ra
    
int_to_str_check_negative:
    # Check if negative
    bgez $t1, int_to_str_positive
    
    # Handle negative number
    li   $t2, 1             # Set negative flag
    
    # Special case: -2147483648 can't be negated in 32-bit
    li   $t3, -2147483648
    bne  $t1, $t3, int_to_str_regular_negative
    
    # Hardcode the most negative value
    li   $t3, 45            # '-'
    sb   $t3, 0($t0)
    li   $t3, 50            # '2'
    sb   $t3, 1($t0)
    li   $t3, 49            # '1'
    sb   $t3, 2($t0)
    li   $t3, 52            # '4'
    sb   $t3, 3($t0)
    li   $t3, 55            # '7'
    sb   $t3, 4($t0)
    li   $t3, 52            # '4'
    sb   $t3, 5($t0)
    li   $t3, 56            # '8'
    sb   $t3, 6($t0)
    li   $t3, 51            # '3'
    sb   $t3, 7($t0)
    li   $t3, 54            # '6'
    sb   $t3, 8($t0)
    li   $t3, 52            # '4'
    sb   $t3, 9($t0)
    li   $t3, 56            # '8'
    sb   $t3, 10($t0)
    sb   $zero, 11($t0)     # Null terminator
    move $v0, $s1
    jr   $ra
    
int_to_str_regular_negative:
    neg  $t1, $t1           # Make positive (t1 = -t1)
    
int_to_str_positive:
    # Convert digits (in reverse order)
    move $t3, $t0           # $t3 = current position in buffer
    
int_to_str_convert_loop:
    # Get last digit: digit = num % 10
    li   $t4, 10
    divu $t1, $t4           # Divide by 10
    mfhi $t5                # $t5 = remainder (digit)
    mflo $t1                # $t1 = quotient
    
    # Convert digit to ASCII: '0' + digit
    addiu $t5, $t5, 48      # 48 = ASCII '0'
    sb   $t5, 0($t3)        # Store digit
    addiu $t3, $t3, 1       # Move buffer pointer
    
    # Continue if quotient != 0
    bnez $t1, int_to_str_convert_loop
    
    # Add minus sign if negative
    beqz $t2, int_to_str_no_minus
    li   $t5, 45            # ASCII '-'
    sb   $t5, 0($t3)
    addiu $t3, $t3, 1
    
int_to_str_no_minus:
    # Add null terminator
    sb   $zero, 0($t3)
    
    # String is reversed, need to reverse it
    addiu $t3, $t3, -1      # Point to last character (before null)
    move $t4, $t0           # $t4 = start pointer
    
int_to_str_reverse_loop:
    bge  $t4, $t3, int_to_str_reverse_done
    
    # Swap characters at $t4 and $t3
    lb   $t5, 0($t4)        # Load char from start
    lb   $t6, 0($t3)        # Load char from end
    sb   $t6, 0($t4)        # Store end char at start
    sb   $t5, 0($t3)        # Store start char at end
    
    addiu $t4, $t4, 1       # Move start forward
    addiu $t3, $t3, -1      # Move end backward
    j    int_to_str_reverse_loop
    
int_to_str_reverse_done:
    move $v0, $s1           # Return pointer to buffer
    jr   $ra
	`)
}
