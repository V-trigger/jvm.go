package references

import(
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
	"jvmgo/ch06/rtda/heap"
)

//new指令的操作数是一个uint16索引，来自字节码。
//通过这个索引，可以从当前类的运行时常量池中找到一个类符号引用。
//解析这个类符号引用，拿到类数据，然后创建对象，并把对象引用推入栈顶，
type NEW struct { base.Index16Instruction }

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	//接口和抽象类不能实例化
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}