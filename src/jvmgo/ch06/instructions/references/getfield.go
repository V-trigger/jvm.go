package references

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

//getfield指令获取对象的实例变量值，然后推入操作数栈，它需要两个操作数。
//第一个操作数是uint16索引。
//第二个操作数是对象引用，用法和putfield一样
type GET_FIELD struct{ base.Index16Instruction }

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()

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
		default:
			// todo
	}
}