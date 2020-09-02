package references

import(
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

//取出类的某个静态变量值，然后推入栈顶
type GET_STATIC struct{ base.Index16Instruction }

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	//根据字段类型,从静态变量中取出相应的值,然后推入操作数栈顶
	switch descriptor[0] {
		case 'Z', 'B', 'C', 'S', 'I':
			stack.PushInt(slots.GetInt(slotId))
		case 'F':
			stack.PushFloat(slots.GetFloat(slotId))
		case 'J':
			stack.PushLong(slots.GetLong(slotId))
		case 'D':
			stack.PushDouble(slots.GetDouble(slotId))
		case 'L', '[':
			stack.PushRef(slots.GetRef(slotId))
	}

}