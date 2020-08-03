package state

type luaState struct {
	stack *luaStack
}

// 构造
func New(stackSize int) *luaState {
	return &luaState{stack: newLuaStack(stackSize)}
}

//
func (self *luaState) pushLuaStack(stack *luaStack) {
	stack.prev = self.stack
	self.stack = stack
}

func (self *luaState) popLuaStack() {
	stack := self.stack
	self.stack = stack.prev
	stack.prev = nil
}
