package math

import(
	"jvmgo/ch06/instructions/base"
    "jvmgo/ch06/rtda"
)

// 取反指令，从操作数栈弹出一个变量，取反后放回去
type DNEG struct{ base.NoOperandsInstruction }
type FNEG struct{ base.NoOperandsInstruction }
type INEG struct{ base.NoOperandsInstruction }
type LNEG struct{ base.NoOperandsInstruction }

func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}