package control

import(
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)
// goto指令 无条件跳转
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}