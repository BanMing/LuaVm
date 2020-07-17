package binchunk

// 相关常量
const (
	LUA_SINGATURE = "\x1bLua"
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
	luacNum         byte
}

// lua chunk
type binaryChunk struct {
	header            //头部
	sizeUpvalues byte //主函数upvalue数量
	//mainFunc     *Prototype //主函数原型
}
