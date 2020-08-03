package state

import "luavm/go/binchunk"

// 闭包
type luaClosure struct {
	proto *binchunk.Prototype // 函数原型
}

func newLuaClosure(proto *binchunk.Prototype) *luaClosure {
	return &luaClosure{proto: proto}
}
