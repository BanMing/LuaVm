package vm

import "luavm/go/api"

//CLOSURE指令（iBx模式）把当前Lua函数的子函数原型实例化为闭包，
//放入由操作数A指定的寄存器中
//子函数原型来自于当前函数原型的子函数原型表，索引由操作数Bx指定
//R(A) := closure(KPROTO[Bx])
func closure(i Instruction, vm api.LuaVM) {
	a, bx := i.ABx()
	a += 1

	vm.LoadProto(bx)
	vm.Replace(a)
}

// CALL指令（iABC模式）调用Lua函数。其中被调函数位于寄存器中，索引由操作数A指定
// 需要传递给被调函数的参数值也在寄存器中，紧挨着被调函数，数量由操作数B指定
// 具体有多少个返回值则由操作数C指定
//
func call(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	nArgs := _pushFuncAndArgs(a, b, vm)
	vm.Call(nArgs, c-1)
	_popResults(a, c, vm)
}

//把存放在连续多个寄存器里的值返回给主调函数。
//其中第一个寄存器的索引由操作数A指定，
//寄存器数量由操作数B指定，操作数C没用
func _return(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	if b == 1 {
		// no return values
	} else if b > 1 {
		// b-1 return values
		vm.CheckStack(b - 1)
		for i := a; i <= a+b-2; i++ {
			vm.PushValue(i)
		}
	} else {
		_fixStack(a, vm)
	}
}

//把传递给当前函数的变长参数加载到连续多个寄存器中。
//其中第一个寄存器的索引由操作数A指定，
//寄存器数量由操作数B指定，操作数C没有用
func vararg(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1

	if b != 1 {
		vm.LoadVararg(b - 1)
		_popResults(a, b, vm)
	}
}

func tailCall(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	c := 0
	nArgs := _pushFuncAndArgs(a, b, vm)
	vm.Call(nArgs, c-1)
	_popResults(a, c, vm)
}

// 把对象和方法拷贝到相邻的两个目标寄存器中。
// 对象在寄存器中，索引由操作数B指定。
// 方法名在常量表里，索引由操作数C指定。
// 目标寄存器索引由操作数A指定
func self(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1

	vm.Copy(b, a+1)
	vm.GetRK(c)
	vm.GetTable(b)
	vm.Replace(a)
}

// 把函数和参数压入栈
func _pushFuncAndArgs(a, b int, vm api.LuaVM) (nArgs int) {
	if b >= 1 {
		// 这里保证栈顶有足够空间
		vm.CheckStack(b)
		for i := a; i < a+b; i++ {
			vm.PushValue(i)
		}
		return b - 1
	} else {
		_fixStack(a, vm)
		return vm.GetTop() - vm.RegisterCount() - 1
	}
}

func _fixStack(a int, vm api.LuaVM) {
	x := int(vm.ToInteger(-1))
	vm.Pop(1)

	vm.CheckStack(x - a)
	for i := a; i < x; i++ {
		vm.PushValue(i)
	}
	vm.Rotate(vm.RegisterCount()+1, x-a)
}

// 弹出返回值
func _popResults(a, c int, vm api.LuaVM) {
	if c == 1 {
		// no results
	} else if c > 1 {
		for i := a + c - 2; i >= a; i-- {
			vm.Replace(i)
		}
	} else {
		// leave results on stack
		vm.CheckStack(1)
		vm.PushInteger(int64(a))
	}
}
