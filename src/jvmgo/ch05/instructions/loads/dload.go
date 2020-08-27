package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//加载double类型

type DLOAD struct { base.Index8Instruction }
type DLOAD_0 struct { base.NoOperandsInstruction }
type DLOAD_1 struct { base.NoOperandsInstruction }
type DLOAD_2 struct { base.NoOperandsInstruction }
type DLOAD_3 struct { base.NoOperandsInstruction }

//通用加载方法
func _dload(frame *rtda.Frame, index uint) {
	//通过索引获取一个double类型的值
	val := frame.LocalVars().GetDouble(index)
	//加入操作数栈顶
	frame.OperandStack().PushDouble(val)
}

func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, uint(self.Index))
}

func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}