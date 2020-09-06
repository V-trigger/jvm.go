package control

import(
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low int32
	high int32
	jumpOffsets []int32
}

//tableswitch指令操作码的后面有0~3字节的padding，以保证defaultOffset在字节码中的地址是4的倍数
func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

//从操作数栈中弹出一个int变量，然后看它是 否在low和high给定的范围之内
//如果在，则从jumpOffsets表中查出偏移量进行跳转，否则按照defaultOffset跳转
func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}