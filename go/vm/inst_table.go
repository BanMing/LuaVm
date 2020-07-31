package vm

import "luavm/go/api"

func newTable(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	vm.CreateTable(Fb2int(b), Fb2int(c))
	vm.Replace(a)
}
