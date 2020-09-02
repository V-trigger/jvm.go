package extended

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

// 根据引用是否是null进行跳转，ifnull和ifnonnull指令把栈顶的 引用弹出
type IFNULL struct{ base.BranchInstruction }
type IFNONNULL struct{ base.BranchInstruction }

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}