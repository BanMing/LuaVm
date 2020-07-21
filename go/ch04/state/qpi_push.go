package state
// 将5种基础类型的值推入栈顶
func (self *luaState) PushNil() { self.stack.push(nil) }

func (self *luaState) PushBoolean(b bool) { self.stack.push(b) }

func (self *luaState) PushInteger(n int64) { self.stack.push(n) }

func (self *luaState) PushNumber(n float64) { self.stack.push(n) }

func (self *luaState) PushString(s string) { self.stack.push(s) }
