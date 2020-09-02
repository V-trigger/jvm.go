package stack

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

//交换两个栈顶变量

type SWAP struct { base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}