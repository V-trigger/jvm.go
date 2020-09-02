package comparisons

import(
	"jvmgo/ch07/instructions/base"
    "jvmgo/ch07/rtda"
)

//比较指令， 比较long型变量

//栈顶的两个long变量弹出，进行比较，然后把 比较结果（int型0、1或-1）推入栈顶
type LCMP struct{ base.NoOperandsInstruction }

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 { 
		stack.PushInt(0)
	} else {
		stack.PushInt(-1) 
	}
}