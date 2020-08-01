package vm

import (
	"luavm/go/api"
)

const LFIEFDS_PER_FLUSH = 50

func newTable(i Instruction, vm api.LuaVM) {
	// R(A) := {} (size = B,C)
	a, b, c := i.ABC()
	a += 1
	// 操作数B和C只有9个比特，需要转为浮点字节的编码方式
	vm.CreateTable(Fb2int(b), Fb2int(c))
	vm.Replace(a)
}

// R(A) :=R(B)[RK(C)]
func getTable(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	vm.GetRK(c)
	vm.GetTable(b)
	vm.Replace(a)
}

// R(A)[RK(B)] :=Rk(C)
func setTable(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	vm.GetRK(b)
	vm.GetRK(c)
	vm.SetTable(a)
}

// R(A)[(C-1)*FPF+i] := R(A+i),1<=i<=Blu
func setList(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1

	if c > 0 {
		c = c - 1
	} else {
		c = Instruction(vm.Fetch()).Ax()
	}

	idx := int64(c * LFIEFDS_PER_FLUSH)
	for j := 1; j <= b; j++ {
		idx++
		vm.PushValue(a + j)
		vm.SetI(a, idx)
	}
}
