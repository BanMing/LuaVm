package vm

import "luavm/go/api"

// 加载指令 4条

// 用于给连续n个寄存器放置nil值
// 起始索引由操作数A指定，寄存器数量则由操作数B指定 iABC模式
func loadNil(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	vm.PushNil()
	for i := a; i <= a+b; i++ {
		vm.Copy(-1, i)
	}
	vm.Pop(1)
}

// 给单个寄存器设置布尔值。
// iABC模式）寄存器索引由操作数A指定，布尔值由寄存器B指定（0代表false，非0代表true，如果寄存器C非0则跳过下一条指令
func loadBoolean(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	vm.PushBoolean(b != 0)
	vm.Replace(a)
	if c == 0 {
		vm.AddPC(1)
	}
}

// 将常量表里的某个常量加载到指定寄存器
// iABx模式 寄存器索引由操作数A指定，常量表索引由操作数Bx指定
func loadK(i Instruction, vm api.LuaVM) {
	a, bx := i.ABx()
	a += 1
	vm.GetConst(bx)
	vm.Replace(a)
}

// iABx模式 需要和EXTRAARG指令（iAx模式）搭配使用，用后者的Ax操作数来指定常量索引
func loadKx(i Instruction, vm api.LuaVM) {
	a, _ := i.ABx()
	a += 1
	ax := Instruction(vm.Fetch()).Ax()
	vm.GetConst(ax)
	vm.Replace(a)
}


