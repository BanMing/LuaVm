package vm

import (
	"fmt"
	"luavm/go/api"
)

const MAXARG_Bx = 1<<18 - 1       // 2^18-1 = 262143
const MAXARG_sBx = MAXARG_Bx >> 1 // 262143 / 2 = 131071

type Instruction uint32

// 从指令中提取操作码
func (self Instruction) Opcode() int {
	return int(self & 0x3F)
}

// 从iABC模式指令中提取参数
func (self Instruction) ABC() (a, b, c int) {
	a = int(self >> 6 & 0xFF)
	c = int(self >> 14 & 0x1FF)
	b = int(self >> 23 & 0x1FF)
	return
}

// 从iABx模式指令中提取参数
func (self Instruction) ABx() (a, bx int) {
	a = int(self >> 6 & 0xFF)
	bx = int(self >> 14)
	return
}

// 从iAsBx模式指令中提取参数
func (self Instruction) AsBx() (a, sbx int) {
	a, bx := self.ABx()
	return a, bx - MAXARG_sBx
}

// 从iAx模式指令中提取参数
func (self Instruction) Ax() int {
	return int(self >> 6)
}

func (self Instruction) OpName() string {
	return opcodes[self.Opcode()].name
}

func (self Instruction) OpMode() byte {
	return opcodes[self.Opcode()].opMode
}

func (self Instruction) BMode() byte {
	return opcodes[self.Opcode()].argBMode
}

func (self Instruction) CMode() byte {
	return opcodes[self.Opcode()].argCMode
}

// 先从指令里提取操作码，然后根据操作码从指令表里查找对应的指令实现方法，最后调用指令实现方法执行指令
func (self Instruction) Execute(vm api.LuaVM) {
	action := opcodes[self.Opcode()].action
	if vm.PC()==13 {
		fmt.Sprintf("")
	}
	if action != nil {
		action(self, vm)
	} else {
		panic(self.OpName())
	}
}
