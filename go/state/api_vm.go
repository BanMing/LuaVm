package state

func (self *luaState) PC() int {
	return self.pc
}

func (self *luaState) AddPC(n int) {
	self.pc += n
}

// 获取当前指令
func (self *luaState) Fetch() uint32 {
	i := self.proto.Code[self.pc]
	self.pc++
	return i
}

// 取出一个常量值 放置栈顶
func (self *luaState) GetConst(idx int) {
	c := self.proto.Constants[idx]
	self.stack.push(c)
}

// 将指定常量或栈值推入栈顶
func (self *luaState) GetRK(rk int) {
	if rk > 0xFF {
		self.GetConst(rk & 0xFF)
	} else {
		self.PushValue(rk + 1)
	}
}
