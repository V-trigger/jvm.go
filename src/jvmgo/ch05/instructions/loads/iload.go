package loads

import(
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//加载指令从局部变量表获取变量，然后推入操作数栈顶。
//iload加载int型变量

type ILOAD struct { base.Index8Instruction }  //iload需要索引，继承Index8Instruction结构体
type ILOAD_0 struct { base.NoOperandsInstruction }  //固定索引0
type ILOAD_1 struct { base.NoOperandsInstruction }  //固定索引1
type ILOAD_2 struct { base.NoOperandsInstruction }  //固定索引2
type ILOAD_3 struct { base.NoOperandsInstruction }  //固定索引3

//一个通用的入操作数栈方法
func _iload(frame *rtda.Frame, index uint) {
	//从栈帧的局部变量表里读取一个int型变量
	val := frame.LocalVars().GetInt(index)
	//入操作数栈
	frame.OperandStack().PushInt(val)
}


func (self *ILOAD) Execute(frame *rtda.Frame) {
    _iload(frame, uint(self.Index))
}

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

