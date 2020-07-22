package vm

import "luavm/go/api"

// 其他类型 2条

// 移动值
// 把源寄存器（索引由操作数B指定）里的值移动到目标寄存器（索引由操作数A指定）里
func move(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.Copy(b, a)
}

// 执行无条件跳转 goto
func jmp(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	vm.AddPC(sBx)
	if a != 0 {
		panic("todo!")
	}
}
