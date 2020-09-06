package references

import(
	"jvmgo/ch08/instructions/base"
    "jvmgo/ch08/rtda"
    "jvmgo/ch08/rtda/heap"
)
// anewarray创建引用类型数组
// anewarray指令也要两个操作数。
// 第一个操作数是uint16索引，来自字节码
// 通过这个索引可以从当前类的运行时常量池中找到一个类符号引用，解析这个符号引用就可以得到数组元素的类
// 第二个操作数是数组长度，从操作数栈中弹出
type ANEW_ARRAY struct{
	base.Index16Instruction
}

//根据数组元素的类型和数组长度创建引用类型数组
func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	//获取数组元素
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}