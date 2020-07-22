package state

import "luavm/go/binchunk"

type luaState struct {
	stack *luaStack
	proto *binchunk.Prototype // 函数原型
	pc    int                 // 计数
}

func New(stackSize int, proto *binchunk.Prototype) *luaState {
	return &luaState{
		stack: newLuaStack(stackSize),
		proto: proto,
		pc:    0}
}
