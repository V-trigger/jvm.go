package math

import(
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

//按位与指令，从操作符栈弹出两个变量进行按位与操作,结果再推入操作符栈
type IAND struct{ base.NoOperandsInstruction }  //int按位与
type LAND struct{ base.NoOperandsInstruction }  //long按位与

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}