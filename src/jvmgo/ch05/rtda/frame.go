package rtda

//用单链表来实现Java虚拟机栈
type Frame struct {

	//链表的next域
	lower    *Frame

	//局部变量表
	localVars LocalVars

	//操作数栈
	operandStack *OperandStack

	thread *Thread

	nextPC int
}

// func NewFrame(maxLocals, maxStack uint) *Frame {
// 	return &Frame{
// 		localVars: newLocalVars(maxLocals),
// 		operandStack: newOperandStack(maxStack),
// 	}
// }

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

//getter/setter 方法
func (self *Frame) LocalVars() LocalVars{
    return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
