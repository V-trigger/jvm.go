package stack

import(
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

//弹出操作数栈的栈顶变量
//pop指令只能用于弹出int、float等占用一个操作数栈位置的变量
//double和long变量在操作数栈中占据两个位置，需要使用pop2 指令弹出
type POP struct{ base.NoOperandsInstruction }
type POP2 struct{ base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}