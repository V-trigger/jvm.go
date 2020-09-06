package math

import(
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

//位移指令,弹出两个变量进行位移操作，得到的结果再入操作数栈
// 次顶变量 << 栈顶变量 

type ISHL struct{ base.NoOperandsInstruction } // int左位移
type ISHR struct{ base.NoOperandsInstruction } // int算术右位移,正数高位补0，负数补1
type IUSHR struct{ base.NoOperandsInstruction } // int逻辑右位移, 高位都补0
type LSHL struct{ base.NoOperandsInstruction } // long左位移
type LSHR struct{ base.NoOperandsInstruction } // long算术右位移
type LUSHR struct{ base.NoOperandsInstruction } // long逻辑右位移

// int左位移
func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	//int变量只有32位，所以只取v2的前5个比特就 足够表示位移位数了
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

// int算术右位移,正数高位补0，负数补1
func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	//long变量有64位，所以取v2的前6个比特
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

// int逻辑右位移, 高位都补0
func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	//Go语言并没有Java语言中的>>>运算符，为了达到无符号位移 的目的，需要先把v1转成无符号整数，位移操作之后，再转回有符号整数。
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

//long左位移
func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	//int变量只有32位，所以只取v2的前5个比特就 足够表示位移位数了
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

// long算术右位移
func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	//long变量有64位，所以取v2的前6个比特
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// long逻辑右位移, 高位都补0
func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	//Go语言并没有Java语言中的>>>运算符，为了达到无符号位移 的目的，需要先把v1转成无符号整数，位移操作之后，再转回有符号整数。
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}