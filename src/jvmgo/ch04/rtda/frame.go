package rtda

//用单链表来实现Java虚拟机栈
type Frame struct {

	//链表的next域
	lower    *Frame

	//局部变量表
	localVars LocalVars

	//操作数栈
	operandStack *OperandStack
}

func newFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

