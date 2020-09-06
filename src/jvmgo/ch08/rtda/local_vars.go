package rtda

import(
	"math"
	"jvmgo/ch08/rtda/heap"
)

//局部变量表
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars{
    if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

//根据索引设置整型
func (self LocalVars) SetInt(index uint, val int32) {
    self[index].num = val
}

//根据索引从局部变量表查找整型
func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

//根据索引设置浮点型, float变量可以先转成int类型，然后按int变量来处理
func (self LocalVars) SetFloat(index uint, val float32) {
	//Float32bits()函数返回浮点数的IEEE 754格式二进制表示对应的4字节无符号整数。
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

//根据索引查找浮点型
func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	//Float32frombits函数返回无符号整数对应的IEEE 754格式二进制表示的4字节浮点数。
	return math.Float32frombits(bits)
}

//根据索引设置Long型变量,Long需要两个连续的int来表示, 一个存低位。一个存高位
func (self LocalVars) SetLong(index uint, val int64){
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}

//根据索引查找Long型变量
func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

//设置double类型变量,dobule可以先转换成long处理
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

//读取double类型的变量
func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

//设置引用类型变脸
func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self[index].ref = ref
}

//读取引用类型变量
func (self LocalVars) GetRef(index uint) *heap.Object{
	return self[index].ref
}

//设置变量
func (self LocalVars) SetSlot(index uint, slot Slot) {
	self[index] = slot
}