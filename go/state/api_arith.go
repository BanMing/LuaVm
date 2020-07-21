package state

import (
	"luavm/go/api"
	"luavm/go/number"
	"math"
)

// 算术和位移运算
var (
	iadd = func(a, b int64) int64 { return a + b }
	fadd = func(a, b float64) float64 { return a + b }

	isub = func(a, b int64) int64 { return a - b }
	fsub = func(a, b float64) float64 { return a - b }

	imul = func(a, b int64) int64 { return a * b }
	fmul = func(a, b float64) float64 { return a * b }

	div = func(a, b float64) float64 { return a / b }

	imod = number.IMod
	fmod = number.FMod

	pow = math.Pow

	iidiv = number.IFloorDiv
	fidiv = number.FFloorDiv

	band = func(a, b int64) int64 { return a & b }
	bor  = func(a, b int64) int64 { return a | b }
	bxor = func(a, b int64) int64 { return a ^ b }
	bnot = func(a, _ int64) int64 { return ^a }

	shl = number.ShiftLeft
	shr = number.ShifRight

	iunm = func(a, _ int64) int64 { return -a }
	funm = func(a, _ float64) float64 { return -a }
)

// 存储运算类型
type operator struct {
	integerFunc func(int64, int64) int64
	floatFunc   func(float64, float64) float64
}

// 运算类型映射
var operators = []operator{
	{iadd, fadd},
	{isub, fsub},
	{imul, fmul},
	{iunm, funm},
	{imod, fmod},
	{iidiv, fidiv},
	{nil, div},
	{nil, pow},
	{band, nil},
	{bor, nil},
	{bxor, nil},
	{shl, nil},
	{shr, nil},
	{bnot, nil},
}

// 执行计算
func _arith(a, b luaValue, op operator) luaValue {
	if op.floatFunc == nil {
		if x, ok := convertToInteger(a); ok {
			if y, ok := convertToInteger(b); ok {
				return op.integerFunc(x, y)
			}
		}
	} else {
		if op.integerFunc != nil {
			if x, ok := a.(int64); ok {
				if y, ok := b.(int64); ok {
					return op.integerFunc(x, y)
				}
			}
		}

		if x, ok := convertToFloat(a); ok {
			if y, ok := convertToFloat(b); ok {
				return op.floatFunc(x, y)
			}
		}
	}
	return nil
}

// 四则运算与位运算
func (self *luaState) Arith(op api.ArithOp) {
	var a, b luaValue // operands
	b = self.stack.pop()
	if op != api.LUA_OPUNM && op != api.LUA_OPBNOT {
		a = self.stack.pop()
	} else {
		a = b
	}
	operator := operators[op]
	if result := _arith(a, b, operator); result != nil {
		self.stack.push(result)
	} else {
		panic("arithmetic error!")
	}
}
