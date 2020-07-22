package vm

import "luavm/go/api"

// 运算相关指令 22条

// 对两个寄存器或常量值（索引由操作数B和C指定）进行运算，将结果放入另一个寄存器（索引由操作数A指定
func _binaryArith(i Instruction, vm api.LuaVM, op api.ArithOp) {
	a, b, c := i.ABC()
	a += 1
	vm.GetRK(b)
	vm.GetRK(c)
	vm.Arith(op)
	vm.Replace(c)
}

// 一元算术运算指令（iABC模式），对操作数B所指定的寄存器里的值进行运算，然后把结果放入操作数A所指定的寄存器中
func _unaryArith(i Instruction, vm api.LuaVM, op api.ArithOp) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.PushValue(a)
	vm.Arith(op)
	vm.Replace(b)
}

// +
func add(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPADD) }

// -
func sub(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPSUB) }

// *
func mul(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPMUL) }

// %
func mod(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPMOD) }

// ^
func pow(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPPOW) }

// /
func div(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPDIV) }

// //
func idiv(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPIDIV) }

// &
func band(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPBAND) }

// |
func bor(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPBOR) }

// ~
func bxor(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPBXOR) }

// <<
func shl(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPSHL) }

// >>
func shr(i Instruction, vm api.LuaVM) { _binaryArith(i, vm, api.LUA_OPSHR) }

// -
func unm(i Instruction, vm api.LuaVM) { _unaryArith(i, vm, api.LUA_OPUNM) }

// ~
func bnot(i Instruction, vm api.LuaVM) { _unaryArith(i, vm, api.LUA_OPBNOT) }

// 获得长度
func len(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.Len(b)
	vm.Replace(a)
}

// 连接 将连续n个寄存器（起止索引分别由操作数B和C指定）里的值拼接，将结果放入另一个寄存器（索引由操作数A指定）
func concat(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	c += 1
	n := c - b + 1
	vm.CheckStack(n)
	for i := b; i <= c; i++ {
		vm.PushValue(i)
	}
	vm.Concat(n)
	vm.Replace(a)
}

// 比较寄存器或常量表里的两个值（索引分别由操作数B和C指定），如果比较结果和操作数A（转换为布尔值）匹配，则跳过下一条指令
func _compare(i Instruction, vm api.LuaVM, op api.CompareOp) {
	a, b, c := i.ABC()
	vm.GetRK(b)
	vm.GetRK(c)
	if vm.Compare(-2, -1, op) != (a != 0) {
		vm.AddPC(1)
	}
	vm.Pop(2)
}

// ==
func eq(i Instruction, vm api.LuaVM) { _compare(i, vm, api.LUA_OPEQ) }

// <
func lt(i Instruction, vm api.LuaVM) { _compare(i, vm, api.LUA_OPLT) }

// <=
func le(i Instruction, vm api.LuaVM) { _compare(i, vm, api.LUA_OPLE) }

// NOT指令对应Lua语言里的逻辑非运算符
func not(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.PushBoolean(!vm.ToBoolean(b))
	vm.Replace(a)
}

//对应Lua语言里的逻辑与和逻辑或运算符
func testSet(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	if vm.ToBoolean(b) == (c != 0) {
		vm.Copy(b, a)
	} else {
		vm.AddPC(1)
	}
}

// 是TESTSET指令的特殊形式
func test(i Instruction, vm api.LuaVM) {
	a, _, c := i.ABC()
	a += 1
	if vm.ToBoolean(a) != (c != 0) {
		vm.AddPC(1)
	}
}

