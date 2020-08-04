package api

type LuaType = int
type ArithOp = int   //新增的类型别名
type CompareOp = int //新增的类型别名

type LuaState interface {
	// basic stack manipulation
	GetTop() int
	AbsIndex(idx int) int
	CheckStack(n int) bool
	Pop(n int)
	Copy(formIdx, toIdx int)
	PushValue(idx int)
	Replace(idx int)
	Insert(idx int)
	Remove(idx int)
	Rotate(idx, n int)
	SetTop(idx int)

	//	access functions (stack - > Go
	TypeName(tp LuaType) string
	Type(idx int) LuaType
	IsNone(idx int) bool
	IsNil(idx int) bool
	IsNoneOrNil(idx int) bool
	IsBoolean(idx int) bool
	IsInteger(idx int) bool
	IsNumber(idx int) bool
	IsString(idx int) bool
	ToBoolean(idx int) bool
	ToInteger(idx int) int64
	ToIntegerX(idx int) (int64, bool)
	ToNumber(idx int) float64
	ToString(idx int) string
	ToStringX(idx int) (string, bool)

	//	push functions (Go ->stack)
	PushNil()
	PushBoolean(b bool)
	PushInteger(n int64)
	PushNumber(n float64)
	PushString(s string)

	Arith(op ArithOp)                          // 位运算
	Compare(idx1, idx2 int, op CompareOp) bool // 比较运算
	Len(idx int)                               // 取长度运算
	Concat(n int)                              // 字符串拼接运算

	//	get functions (Lua -> stack)
	NewTable()
	CreateTable(nArr, nRec int)
	GetTable(idx int) LuaType
	GetField(idx int, k string) LuaType
	GetI(idx int, i int64) LuaType

	//	set functions (Lua ->stack)
	SetTable(idx int)
	SetField(idx int, k string)
	SetI(idx int, n int64)

	//	lua function
	Load(chunk []byte, chunkName, mode string) int
	Call(nArgs, nResults int)
}
