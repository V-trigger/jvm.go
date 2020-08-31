package heap
import "math"

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot


func newSlots(slotCount uint) Slots{
    if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

//根据索引设置整型
func (self Slots) SetInt(index uint, val int32) {
    self[index].num = val
}

//根据索引从局部变量表查找整型
func (self Slots) GetInt(index uint) int32 {
	return self[index].num
}

//根据索引设置浮点型, float变量可以先转成int类型，然后按int变量来处理
func (self Slots) SetFloat(index uint, val float32) {
	//Float32bits()函数返回浮点数的IEEE 754格式二进制表示对应的4字节无符号整数。
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

//根据索引查找浮点型
func (self Slots) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	//Float32frombits函数返回无符号整数对应的IEEE 754格式二进制表示的4字节浮点数。
	return math.Float32frombits(bits)
}

//根据索引设置Long型变量,Long需要两个连续的int来表示, 一个存低位。一个存高位
func (self Slots) SetLong(index uint, val int64){
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}

//根据索引查找Long型变量
func (self Slots) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

//设置double类型变量,dobule可以先转换成long处理
func (self Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

//读取double类型的变量
func (self Slots) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

//设置引用类型变量
func (self Slots) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}

//读取引用类型变量
func (self Slots) GetRef(index uint) *Object{
	return self[index].ref
}