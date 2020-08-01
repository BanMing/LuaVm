package main

import (
	"fmt"
	"io/ioutil"
	"luavm/go/binchunk"
	"luavm/go/state"
	"luavm/go/vm"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		proto := binchunk.Undump(data)
		luaMain(proto)
	}
}

// 调用Lua主函数
func luaMain(proto *binchunk.Prototype) {
	nRegs := int(proto.MaxStackSize)
	ls := state.New(nRegs+8, proto)
	ls.SetTop(nRegs)
	for {
		pc := ls.PC()
		inst := vm.Instruction(ls.Fetch())
		if inst.Opcode() != vm.OP_RETURN {
			inst.Execute(ls)
			fmt.Printf("[%02d] %s %s \n", pc+1, inst.OpName(), ls.PrintStackInfo())
		} else {
			break
		}
	}
}
