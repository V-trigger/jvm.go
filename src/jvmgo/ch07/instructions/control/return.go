package control

import(
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)
//返回指令
//areturn、ireturn、lreturn、freturn和dreturn 分别用于返回引用、int、long、float和double类型的值
//return 返回void
//6条返回指令都不需要操作数
type RETURN struct{ base.NoOperandsInstruction }
type ARETURN struct{ base.NoOperandsInstruction }
type DRETURN struct{ base.NoOperandsInstruction }
type FRETURN struct{ base.NoOperandsInstruction }
type IRETURN struct{ base.NoOperandsInstruction }
type LRETURN struct{ base.NoOperandsInstruction }
// Return void from method

//返回viod弹出栈帧即可
func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}