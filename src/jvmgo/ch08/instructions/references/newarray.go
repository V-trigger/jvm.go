package references

import(
	"jvmgo/ch08/instructions/base"
    "jvmgo/ch08/rtda"
	"jvmgo/ch08/rtda/heap"
)

const (
	//Array Type  atype
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

// newarray指令用来创建基本类型数组
// newarray指令需要两个操作数
// 第一个操作数是一个uint8整数,在字节码中紧跟在指令操作码后面
// 第二个操作数是count，从操作数栈中弹出，表示数组长度
type NEW_ARRAY struct {
	atype uint8
}

//第一个操作数跟在指令操作码后面
func (self *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.atype = reader.ReadUint8()
}

func (self *NEW_ARRAY) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	classLoader := frame.Method().Class().Loader()
	arrClass := getPrimitiveArrayClass(classLoader, self.atype)
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}