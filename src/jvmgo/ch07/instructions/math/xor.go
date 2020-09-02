
package math

import(
	"jvmgo/ch07/instructions/base"
    "jvmgo/ch07/rtda"
)

// 按位异或

type IXOR struct{ base.NoOperandsInstruction }  //int按位异或
type LXOR struct{ base.NoOperandsInstruction }  //long按位异或

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}