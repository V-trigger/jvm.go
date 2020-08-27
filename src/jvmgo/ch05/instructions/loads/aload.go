package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//加载引用类型

type ALOAD struct { base.Index8Instruction }
type ALOAD_0 struct { base.NoOperandsInstruction }
type ALOAD_1 struct { base.NoOperandsInstruction }
type ALOAD_2 struct { base.NoOperandsInstruction }
type ALOAD_3 struct { base.NoOperandsInstruction }

//通用的入栈方法
func _aload(frame *rtda.Frame, index uint) {
	//从局部变量表读取一个引用类型的值
	ref := frame.LocalVars().GetRef(index)
	//推入操作数栈
	frame.OperandStack().PushRef(ref)
}

func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, uint(self.Index))
}

func (self ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

func (self ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

func (self ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}
