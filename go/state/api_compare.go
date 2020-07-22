package state

import (
	"luavm/go/api"
)

// 相等情况
func _eq(a, b luaValue) bool {
	switch x := a.(type) {
	case nil:
		return b == nil
	case bool:
		y, ok := b.(bool)
		return ok && y == x
	case string:
		y, ok := b.(string)
		return ok && y == x
	case int64:
		switch y := b.(type) {
		case int64:
			return x == y
		case float64:
			return float64(x) == y
		default:
			return false
		}
	case float64:
		switch y := b.(type) {
		case float64:
			return x == y
		case int64:
			return int64(x) == y
		default:
			return false
		}
	default:
		return a == b
	}
}

// 小于
func _lt(a, b luaValue) bool {
	switch x := a.(type) {
	case string:
		if y, ok := b.(string); ok {
			return x < y
		}
	case int64:
		switch y := b.(type) {
		case int64:
			return x < y
		case float64:
			return float64(x) < y
		}
	case float64:
		switch y := b.(type) {
		case float64:
			return x < y
		case int64:
			return x < float64(y)

		}
	}
	panic("comparison error!")
}

// 小于等于
func _le(a, b luaValue) bool {
	switch x := a.(type) {
	case string:
		if y, ok := b.(string); ok {
			return x <= y
		}
	case int64:
		switch y := b.(type) {
		case int64:
			return x <= y
		case float64:
			return float64(x) <= y
		}
	case float64:
		switch y := b.(type) {
		case float64:
			return x <= y
		case int64:
			return x <= float64(y)

		}
	}
	panic("comparison error!")
}

// 比较函数
func (self *luaState) Compare(idx1, idx2 int, op api.CompareOp) bool {
	a := self.stack.get(idx1)
	b := self.stack.get(idx2)
	switch op {
	case api.LUA_OPEQ:
		return _eq(a, b)
	case api.LUA_OPLE:
		return _le(a, b)
	case api.LUA_OPLT:
		return _lt(a, b)
	default:
		panic("invalid compare op!")
	}
}
