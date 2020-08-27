package constants

import(
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//无动作指令，什么都不做
type NOP struct{
	base.NoOperandsInstruction
} 

func (self *NOP) Execute(frame *rtda.Frame) {

}