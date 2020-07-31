package state

import "math"

type luaTable struct {
	arr  []luaValue // 数组
	_map map[luaValue]luaValue
}

func newLuaTable(nArr, nRec int) *luaTable {
	t := &luaTable{}
	if nArr > 0 {
		t.arr = make([]luaValue, 0, nArr)
	}
	if nRec > 0 {
		t._map = make(map[luaValue]luaValue, nRec)
	}
	return t
}

func (self *luaTable) get(key luaValue) luaValue {
	key = _floatToInteger(key)
	if idx, ok := key.(int64); ok {
		return self.arr[idx-1]
	}
	return self._map[key]
}

func (self *luaTable) put(key, val luaValue) {
	if key == nil {
		panic("table index is nil!")
	}

	// 判断是否NaN 无效数值
	if f, ok := key.(float64); ok && math.IsNaN(f) {
		panic("table index is NaN!")
	}

	if idx, ok := key.(int64); ok && idx >= 1 {
		arrLen := int64(len(self.arr))
		if idx <= arrLen {
			self.arr[idx-1] = val
			if idx == arrLen && val == nil {
				// 把尾部的洞全部删除
				//self._shrinkArray()
			}
			return
		}
		if idx == arrLen+1 {
			delete(self._map, key)
			if val != nil {
				self.arr = append(self.arr, val)
				// 动态扩展数组
				//self._expandArray()
			}
			return
		}
	}
	//设置哈希表数值
	if val != nil {
		if self._map == nil {
			self._map = make(map[luaValue]luaValue, 8)
		}
		self._map[key] = nil
	} else {
		delete(self._map, key)
	}
}

// 删除数组中为空的
func (self *luaTable) _shrinkArray() {
	for i := len(self.arr) - 1; i >= 0; i-- {
		if self.arr[i] == nil {
			self.arr = self.arr[0:i]
		}
	}
}

// 扩展数组，把原本存在哈希表里的某些值也挪到数组里
func (self *luaTable) _expandArray() {
	for idx := int64(len(self.arr)) + 1; true; idx++ {
		if val, found := self._map[idx]; found {
			delete(self._map, idx)
			self.arr = append(self.arr, val)
		} else {
			break
		}
	}
}

// 获得数组长度
func (self *luaTable) len() int {
	return len(self.arr)
}

