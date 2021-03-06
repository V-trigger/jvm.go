package rtda

import "jvmgo/ch08/rtda/heap"

//用单链表来实现Java虚拟机栈
type Frame struct {

	//链表的next域
	lower    *Frame

	//局部变量表
	localVars LocalVars

	//操作数栈
	operandStack *OperandStack

	method      *heap.Method

	thread *Thread

	nextPC int
}

// func NewFrame(maxLocals, maxStack uint) *Frame {
// 	return &Frame{
// 		localVars: newLocalVars(maxLocals),
// 		operandStack: newOperandStack(maxStack),
// 	}
// }

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(method.MaxLocals()),
		method:       method,
		operandStack: newOperandStack(method.MaxStack()),
	}
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
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
func (self *Frame) Method() *heap.Method {
	return self.method
}


