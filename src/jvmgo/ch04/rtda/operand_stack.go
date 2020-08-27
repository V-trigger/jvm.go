package rtda

import "math"

//与局部变量表一样，均以字长为单位的数组。不过局部变量表用的是索引，操作数栈是弹栈/压栈来访问。
//操作数栈可理解为java虚拟机栈中的一个用于计算的临时数据存储区。
//存储的数据与局部变量表一致含int、long、float、double、reference、returnType
//操作数栈中byte、short、char压栈前(bipush)会被转为int。
//数据运算的地方，大多数指令都在操作数栈弹栈运算，然后结果压栈。
//java虚拟机栈是方法调用和执行的空间，每个方法会封装成一个栈帧压入占中
//其中里面的操作数栈用于进行运算，当前线程只有当前执行的方法才会在操作数栈中调用指令（可见java虚拟机栈的指令主要取于操作数栈）。
//int类型在-1~5、-128~127、-32768~32767、-2147483648~2147483647范围分别对应的指令是iconst、bipush、sipush、ldc(这个就直接存在常量池了)
//操作数栈的大小是编译器已经确定的，所以可以用[]Slot实现。 size字段用于记录栈顶位置。
type OperandStack struct {
	//当前栈中的变量个数
	size    uint
	
	//栈内变量
	slots   []Slot 
}

func newOperandStack(maxStack uint) *OperandStack {
    if maxStack > 0 {
        return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

//int入栈
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

//int出栈
func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

//float入栈,先转换成int再处理
func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}

//float出栈
func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

//long入栈,需要拆成两个int
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val >> 32)
	self.size += 2
}

//long出栈
func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return int64(high)<<32 | int64(low)
}

//double入栈,先转成long再处理
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

//double入栈
func (self *OperandStack) PopDouble() float64 {
	val := uint64(self.PopLong())
	return math.Float64frombits(val)
}

//引用类型入栈
func (self *OperandStack) PushRef(ref *Object) {
	self.slots[self.size].ref = ref
	self.size++
}

//引用类型出栈
//把Slot结构体的ref字段设置成nil，这样 做是为了帮助Go的垃圾收集器回收Object结构体实例
func (self *OperandStack) PopRef() *Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil
	return ref
}