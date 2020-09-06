package classfile
import "math"

//CONSTANT_Integer_info使用4字节存储整数常量，其结构定义如下：
// CONSTANT_Integer_info {
// 	u1 tag; u4 bytes;
// }
//
//CONSTANT_Integer_info正好可以容纳一个Java的int型常量
//但实际上比int更小的boolean、byte、short和char类型的常量也放在CONSTANT_Integer_info中
type ConstantIntegerInfo struct {
	val int32
}

//先读取一个uint32数据，然后把它转型成int32类型
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}


//CONSTANT_Float_info使用4字节存储IEEE754单精度浮点数常量，结构如下： 
// CONSTANT_Float_info {
// 	u1 tag;
// 	u4 bytes;
// }
type ConstantFloatInfo struct {
	val float32
}

//先读取一个uint32数据，然后调用math包的Float32frombits()函数把它转换成float32类型
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}

//CONSTANT_Long_info使用8字节存储整数常量，结构如下： 
// CONSTANT_Long_info {
// 	u1 tag;
// 	u4 high_bytes;
// 	u4 low_bytes;
// }
type ConstantLongInfo struct {
	val int64
}

//先读取一个uint64数据，然后把它转型成int64类型
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

func (self *ConstantLongInfo) Value() int64 {
	return self.val
}

//CONSTANT_Double_info，使用8字节存储IEEE754双精度浮点数，结构如下：
// CONSTANT_Double_info {
// 	u1 tag;
// 	u4 high_bytes;
// 	u4 low_bytes;
// }
type ConstantDoubleInfo struct {
	val float64
}

//先读取一个uint64数据，然后调用math包的Float64frombits()函数把它转换成float64类型
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}

func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}




