package state

import (
	"luavm/go/api"
	"luavm/go/number"
)

type luaValue interface {
}

func typeOf(val luaValue) api.LuaType {
	switch val.(type) {
	case nil:
		return api.LUA_TNIL
	case bool:
		return api.LUA_TBOOLEAN
	case int64:
		return api.LUA_TNUMBER
	case float64:
		return api.LUA_TNUMBER
	case string:
		return api.LUA_TSTRING
	default:
		panic("todo!")
	}
}

func convertToBoolean(val luaValue) bool {
	switch x := val.(type) {
	case nil:
		return false
	case bool:
		return x
	default:
		return true
	}
}

func _floatToInteger(key luaValue) luaValue {
	if f, ok := key.(float64); ok {
		if i, ok := number.FloatToInteger(f); ok {
			return i
		}
	}
	return key
}

// 转成浮点数
func convertToFloat(val luaValue) (float64, bool) {
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	case string:
		return number.ParserFloat(x)
	default:
		return 0, false
	}
}

// 先转为整数，不行，则先转为浮点数，再转Wie整数
func _stringToInteger(s string) (int64, bool) {
	if i, ok := number.ParserInteger(s); ok {
		return i, true
	}
	if f, ok := number.ParserFloat(s); ok {
		return number.FloatToInteger(f)
	}
	return 0, false
}

// 转为整数
func convertToInteger(val luaValue) (int64, bool) {
	switch x := val.(type) {
	case int64:
		return x, true
	case float64:
		return number.FloatToInteger(x)
	case string:
		return _stringToInteger(x)
	default:
		return 0, false
	}
}
