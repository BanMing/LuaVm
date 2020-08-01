package state

import (
	"fmt"
	"luavm/go/api"
)

// 基础栈操纵方法

//f返回栈顶
func (self *luaState) GetTop() int {
	return self.stack.top
}

// 返回绝对索引
func (self *luaState) AbsIndex(idx int) int {
	return self.stack.absIndex(idx)
}

// 检查栈剩余空间
func (self *luaState) CheckStack(n int) bool {
	self.stack.check(n)
	return true
}

func (self *luaState) Pop(n int) {
	//for i := 0; i < n; i++ {
	//	self.stack.pop()
	//}
	self.SetTop(-n - 1)
}

func (self *luaState) Copy(fromIdx, toIdx int) {
	val := self.stack.get(fromIdx)
	self.stack.set(toIdx, val)
}

// 将指定位置值 复制到栈顶
func (self *luaState) PushValue(idx int) {
	val := self.stack.get(idx)
	self.stack.push(val)
}

// 将栈顶值弹出，写入指定位置
func (self *luaState) Replace(idx int) {
	val := self.stack.pop()
	self.stack.set(idx, val)
}

// [idx,top]索引区间内的值朝栈顶方向旋转n个位置
func (self *luaState) Rotate(idx, n int) {
	t := self.stack.top - 1
	p := self.stack.absIndex(idx) - 1
	var m int
	if n > 0 {
		m = t - n
	} else {
		m = p - n - 1
	}

	self.stack.reverse(p, m)
	self.stack.reverse(m+1, t)
	self.stack.reverse(p, t)
}

// 将栈顶值插入任意位置
func (self *luaState) Insert(idx int) {
	self.Rotate(idx, 1)
}

// 删除任意位置
func (self *luaState) Remove(idx int) {
	self.Rotate(idx, -1)
	self.Pop(1)
}

func (self *luaState) SetTop(idx int) {
	newTop := self.stack.absIndex(idx)
	if newTop < 0 {
		panic("stack underflow!")
	}

	n := self.stack.top - newTop
	if n > 0 {
		for i := 0; i < n; i++ {
			self.stack.pop()
		}
	} else if n < 0 {
		for j := 0; j > n; j-- {
			self.stack.push(nil)
		}
	}
}

// 打印数据
func (self *luaState) PrintStackInfo() string {
	i := 0
	str := ""
	for i < self.stack.top {
		switch typeOf(self.stack.slots[i]) {
		case api.LUA_TTABLE:
			str += "[table]"
		case api.LUA_TSTRING:
			str += fmt.Sprintf("[\"%s\"]", self.stack.slots[i])
		case api.LUA_TNUMBER:
			num, _ := convertToFloat(self.stack.slots[i])
			str += fmt.Sprintf("[%f]", num)
		default:
			str += "[nil]"
		}
		i++
	}
	return str
}
