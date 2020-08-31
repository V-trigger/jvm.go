package constants

import(
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

//无动作指令，什么都不做
type NOP struct{
	base.NoOperandsInstruction
} 

func (self *NOP) Execute(frame *rtda.Frame) {

}