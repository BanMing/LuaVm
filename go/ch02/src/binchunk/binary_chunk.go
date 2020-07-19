package binchunk

// header相关常量
const (
	LUA_SINGATURE    = "\x1bLua" // 通用签名
	LUAC_VERSION     = 0X53      // 版本号5.3.4.对应计算 5 * 16 + 3 = 83 对应16进制
	LUAC_FORMAT      = 0
	LUAC_DATA        = "\x19\x93\r\n\\xla\n"
	CINT_SIZE        = 4
	CSIZET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT         = 0X5678
	LUAC_NUM         = 370.5
)

//chunk 头部
type header struct {
	singnature      [4]byte // 签名 - 魔数 "\x1bLua"
	version         byte    // 对应lua版本号
	format          byte    // 格式号 与虚拟机匹配
	luacData        [6]byte // 检测文件是否损坏 "\x19\x93\r\n\xla\n"
	cintSize        byte    // 数据类型 占用字节数
	sizetSize       byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         byte // 检测chunk的大小端方式
	luacNum         byte // 浮点数检测
}

// 函数原型 定义tag常量
const (
	TAG_NIL       = 0X00 // 不存在
	TAG_BOOLEAN   = 0X01 // bool 字节（0、1）
	TAG_NUMBER    = 0X03 // lua浮点数
	TAG_INTERGER  = 0X13 // lua整数
	TAG_SHORT_STR = 0X04 // 短字符串
	TAG_LONG_STR  = 0X14 // 长字符串
)

// Upvalue表
type Upvalue struct {
	Instack byte
	Idx     byte
}

// 局部变量表
type LocVar struct {
	VarName string // 变量名
	StartPC uint32 // 起始指令索引
	EndPC   uint32 // 终止指令索引
}

//函数原型
type Prototype struct {
	Source          string        // 源文件名
	LineDefined     uint32        // 起始行号
	LastLineDefined uint32        // 结束行号
	NumParams       byte          // 固定参数个数
	IsVararg        byte          // 是否是Vararg函数
	MaxStackSize    byte          // 寄存器数量
	Code            []uint32      // 指令表
	Constants       []interface{} // 常量表
	Upvalues        []Upvalue     // Upvalue表
	Protos          []*Prototype  // 子函数原型表
	LineInfo        []uint32      // 行号
	LocVars         []LocVar      // 局部变量表
	UpvalueNames    []string      // Upvalue名列表
}

// lua chunk
type binaryChunk struct {
	header            //头部
	sizeUpvalues byte //主函数upvalue数量
	//mainFunc     *Prototype //主函数原型
}

// 解析二进制chunk文件
func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()        // 校验头部
	reader.readByte()           // 跳过Upvalue数量
	return reader.readProto("") // 读取函数原型
}
