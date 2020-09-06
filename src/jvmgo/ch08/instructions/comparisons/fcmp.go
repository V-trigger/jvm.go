package comparisons

import(
	"jvmgo/ch08/instructions/base"
    "jvmgo/ch08/rtda"
)
//fcmpg和fcmpl指令用于比较float变量
//这两条指令和lcmp指令很像，但是除了比较的变量类型不同以外，还有一个重要的区别。
//由于浮点数计算有可能产生NaN（Not a Number）值，所以比较两个浮点数时，除了大于、等于、小于之外，还有第4种结果：无法比较。
//fcmpg和fcmpl指令的区别就在于对第4种结果的定义
type FCMPG struct{ base.NoOperandsInstruction }
type FCMPL struct{ base.NoOperandsInstruction }

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (self *FCMPG) Execute(frame *rtda.Frame) {
    _fcmp(frame, false)
}

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}