package binchunk

import (
	"encoding/binary"
	"math"
)

// 解析二进制chunk读取类型
type reader struct {
	data []byte
}

// 从字节流中读取一个字节
func (self *reader) readByte() byte {
	b := self.data[0]
	self.data = self.data[1:]
	return b
}

// 按小端方式出去cint类型 4字节
func (self *reader) readUint32() uint32 {
	i := binary.LittleEndian.Uint32(self.data)
	self.data = self.data[4:]
	return i
}

// 按小端方式读出size_t类型 8字节
func (self *reader) readUnit64() uint64 {
	i := binary.LittleEndian.Uint64(self.data)
	self.data = self.data[8:]
	return i
}

// 读取lua整数
func (self *reader) readLuaInteger() int64 {
	return int64(self.readUnit64())
}

// 读取lua浮点数 8字节
func (self *reader) readLuaNumber() float64 {
	return math.Float64frombits(self.readUnit64())
}

// 读取n个字节
func (self *reader) readBytes(n uint) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}

// 读取string
func (self *reader) readString() string {
	size := uint(self.readByte()) //短字符串？
	if size == 0 {                //NULL字符串
		return ""
	}
	if size == 0xFF {
		size = uint(self.readUnit64())
	}
	bytes := self.readBytes(size - 1)
	return string(bytes);
}

// 检查头部
func (self *reader) checkHeader() {
	if string(self.readBytes(4)) != LUA_SINGATURE {
		panic("not a precompiled chunk!")
	} else if self.readByte() != LUAC_VERSION {
		panic("version mismatch!")
	} else if self.readByte() != LUAC_FORMAT {
		panic("format mismatch!")
	} else if string(self.readBytes(6)) != LUAC_DATA {
		panic("corrupted!")
	} else if self.readByte() != CINT_SIZE {
		panic("int size mismatch!")
	} else if self.readByte() != CSIZET_SIZE {
		panic("size_t size mismatch!")
	} else if self.readByte() != INSTRUCTION_SIZE {
		panic("instruction size mismatch!")
	} else if self.readByte() != LUA_INTEGER_SIZE {
		panic("lua_Integer size mismatch!")
	} else if self.readByte() != LUA_NUMBER_SIZE {
		panic("lua_number size mismatch!")
	} else if self.readLuaInteger() != LUAC_INT {
		panic("endianness mismatch!")
	} else if self.readLuaNumber() != LUAC_NUM {
		panic("float format mismatch!")
	}
}

// 从字节流里读取函数原型
func (self *reader) readProto(parentSource string) *Prototype {
	source := self.readString()
	if source == "" {
		source = parentSource
	}
	return &Prototype{
		Source: source,
		LineDefined:     self.readUint32(),
		LastLineDefined: self.readUint32(),
		NumParams:       self.readByte(),
		IsVararg:        self.readByte(),
		MaxStackSize:    self.readByte(),
		Code:            self.readco,
		Constants:       nil,
		Upvalues:        nil,
		Protos:          nil,
		LineInfo:        nil,
		LocVars:         nil,
		UpvalueNames:    nil,
	}
}
