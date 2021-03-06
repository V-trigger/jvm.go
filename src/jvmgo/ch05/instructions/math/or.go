package math

import(
	"jvmgo/ch05/instructions/base"
    "jvmgo/ch05/rtda"
)

//按位或
type IOR struct{ base.NoOperandsInstruction }  //int按位或
type LOR struct{ base.NoOperandsInstruction }  //long按位或

func (self *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

func (self *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}