  
package conversions

import(
	"jvmgo/ch05/instructions/base"
    "jvmgo/ch05/rtda"
)

// double转float
type D2F struct{ base.NoOperandsInstruction }
// double转int
type D2I struct{ base.NoOperandsInstruction }
// double转long
type D2L struct{ base.NoOperandsInstruction }

func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}