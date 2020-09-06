package references

import(
	"jvmgo/ch07/instructions/base"
    "jvmgo/ch07/rtda"
    "jvmgo/ch07/rtda/heap"
)

//putstatic指令给类的某个静态变量赋值，它需要两个操作数。
//第一个操作数是uint16索引，来自字节码。
//通过这个索引可以从当前类的运行时常量池中找到一个字段符号引用，解析这个符号引用就可以知道要给类的哪个静态变量赋值。
//第二个操作数是要赋给静态变量的值，从操作数栈中弹出
type PUT_STATIC struct{ base.Index16Instruction }

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	//获取当前方法、类、常量池
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()

	//通过索引获取常量池数据
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)

	//解析符号引用
	field := fieldRef.ResolvedField()

	class := field.Class()

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//类初始化方法由编译器生成，名字是<clinit>
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
		case 'Z', 'B', 'C', 'S', 'I':
			slots.SetInt(slotId, stack.PopInt())
		case 'F':
			slots.SetFloat(slotId, stack.PopFloat())
		case 'J':
			slots.SetLong(slotId, stack.PopLong())
		case 'D':
			slots.SetDouble(slotId, stack.PopDouble())
		case 'L', '[':
			slots.SetRef(slotId, stack.PopRef())
	}
}