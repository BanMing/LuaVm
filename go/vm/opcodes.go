package vm

import "luavm/go/api"

// 指令编码模式
const (
	IABC = iota
	IABx
	IAsBx
	IAx
)

// 操作码
const (
	OP_MOVE     = iota // 在寄存器移动值 iABC模式
	OP_LOADK           // iABx模式）将常量表里的某个常量加载到指定寄存器，寄存器索引由操作数A指定，常量表索引由操作数Bx指定
	OP_LOADKX          // LOADKX指令（也是iABx模式）需要和EXTRAARG指令（iAx模式）搭配使用，用后者的Ax操作数来指定常量索引
	OP_LOADBOOL        // iABC模式）给单个寄存器设置布尔值。寄存器索引由操作数A指定，布尔值由寄存器B指定（0代表false，非0代表true，如果寄存器C非0则跳过下一条指令
	OP_LOADNIL         // 用于给连续n个寄存器放置nil值 起始索引由操作数A指定，寄存器数量则由操作数B指定 iABC模式
	OP_GETTUPVAL
	OP_GETTABUP
	OP_GETTABLE
	OP_SETTABUP
	OP_SETUPVAL
	OP_SETTABLE
	OP_NEWTABLE
	OP_SELF
	OP_ADD // 对两个寄存器或常量值（索引由操作数B和C指定）进行运算，将结果放入另一个寄存器（索引由操作数A指定
	OP_SUB
	OP_MUL
	OP_MOD
	OP_POW
	OP_DIV
	OP_IDIV
	OP_BAND // 一元算术运算指令（iABC模式），对操作数B所指定的寄存器里的值进行运算，然后把结果放入操作数A所指定的寄存器中
	OP_BOR
	OP_BXOR
	OP_SHL
	OP_SHR
	OP_UNM
	OP_BNOT
	OP_NOT
	OP_LEN
	OP_CONCAT
	OP_JMP // 执行无条件跳转 iAsBx模式
	OP_EQ
	OP_LT
	OP_LE
	OP_TEST
	OP_TESTSET
	OP_CALL
	OP_TAILCALL
	OP_RETURN
	OP_FORLOOP
	OP_FORPREP
	OP_TFORCALL
	OP_TFORLOOP
	OP_SETLIST
	OP_CLOSURE
	OP_VARARG
	OP_EXTRAARG
)

// 操作数
const (
	OpArgN = iota // argument is not used
	OpArgU        // argument is used
	OpArgR        // argument is a register or a jump offset
	OpArgK        // argument is a constant or register/constant
)

// 每条指令的基本信息
type opcode struct {
	testFlag byte // operator is a test (next instruction must be a jump)
	setAFlag byte // instruction set register A
	argBMode byte // B arg mode
	argCMode byte // C arg mode
	opMode   byte // op mode
	name     string
	action   func(i Instruction, vm api.LuaVM)
}

var opcodes = []opcode{
	{0, 1, OpArgR, OpArgN, IABC, "MOVE     ", move},
	{0, 1, OpArgK, OpArgN, IABx, "LOADK    ", loadK},
	{0, 1, OpArgN, OpArgN, IABx, "LOADKX   ", loadKx},
	{0, 1, OpArgU, OpArgU, IABC, "LOADBOOL ", loadBoolean},
	{0, 1, OpArgU, OpArgN, IABC, "LOADNIL  ", loadNil},
	{0, 1, OpArgU, OpArgN, IABC, "GETUPVAL ", nil},
	{0, 1, OpArgU, OpArgK, IABC, "GETTABUP ", nil},
	{0, 1, OpArgR, OpArgK, IABC, "GETTABLE ", getTable},
	{0, 0, OpArgK, OpArgK, IABC, "SETTABUP ", nil},
	{0, 0, OpArgU, OpArgN, IABC, "SETUPVAL ", nil},
	{0, 0, OpArgK, OpArgK, IABC, "SETTABLE ", setTable},
	{0, 1, OpArgU, OpArgU, IABC, "NEWTABLE ", newTable},
	{0, 1, OpArgR, OpArgK, IABC, "SELF     ", nil},
	{0, 1, OpArgK, OpArgK, IABC, "ADD      ", add},
	{0, 1, OpArgK, OpArgK, IABC, "SUB      ", sub},
	{0, 1, OpArgK, OpArgK, IABC, "MUL      ", mul},
	{0, 1, OpArgK, OpArgK, IABC, "MOD      ", mod},
	{0, 1, OpArgK, OpArgK, IABC, "POW      ", pow},
	{0, 1, OpArgK, OpArgK, IABC, "DIV      ", div},
	{0, 1, OpArgK, OpArgK, IABC, "IDIV     ", idiv},
	{0, 1, OpArgK, OpArgK, IABC, "BAND     ", band},
	{0, 1, OpArgK, OpArgK, IABC, "BOR      ", bor},
	{0, 1, OpArgK, OpArgK, IABC, "BXOR     ", bxor},
	{0, 1, OpArgK, OpArgK, IABC, "SHL      ", shl},
	{0, 1, OpArgK, OpArgK, IABC, "SHR      ", shr},
	{0, 1, OpArgR, OpArgN, IABC, "UNM      ", unm},
	{0, 1, OpArgR, OpArgN, IABC, "BNOT     ", bnot},
	{0, 1, OpArgR, OpArgN, IABC, "NOT      ", not},
	{0, 1, OpArgR, OpArgN, IABC, "LEN      ", len},
	{0, 1, OpArgR, OpArgR, IABC, "CONCAT   ", concat},
	{0, 0, OpArgR, OpArgN, IAsBx, "JMP      ", jmp},
	{0, 0, OpArgK, OpArgK, IABC, "EQ       ", eq},
	{0, 0, OpArgK, OpArgK, IABC, "LT       ", lt},
	{0, 0, OpArgK, OpArgK, IABC, "LE       ", le},
	{0, 0, OpArgN, OpArgU, IABC, "TEST     ", test},
	{0, 1, OpArgR, OpArgU, IABC, "TESTSET  ", testSet},
	{0, 1, OpArgU, OpArgU, IABC, "CALL     ", nil},
	{0, 1, OpArgU, OpArgU, IABC, "TAILCALL ", nil},
	{0, 0, OpArgU, OpArgN, IABC, "RETURN   ", nil},
	{0, 1, OpArgR, OpArgN, IAsBx, "FORLOOP ", forLoop},
	{0, 1, OpArgR, OpArgN, IAsBx, "FORPREP ", forPrep},
	{0, 0, OpArgN, OpArgU, IABC, "TFORCALL ", nil},
	{0, 1, OpArgR, OpArgN, IAsBx, "TFORLOOP", nil},
	{0, 0, OpArgU, OpArgU, IABC, "SETLIST  ", setList},
	{0, 1, OpArgU, OpArgN, IABx, "CLOSURE  ", nil},
	{0, 1, OpArgU, OpArgN, IABC, "VARARG   ", nil},
	{0, 0, OpArgU, OpArgU, IAx, "EXTRAARG  ", nil},
}
