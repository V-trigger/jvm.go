package classfile

//CONSTANT_NameAndType_info给出字段或方法的名称和描述符
//CONSTANT_Class_info和CONSTANT_NameAndType_info加在一起可以唯一确定一个字段或者方法。
//其结构如下:
// CONSTANT_NameAndType_info {
// 	u1 tag;
// 	u2 name_index;
// 	u2 descriptor_index;
// }
//字段或方法名由name_index给出
//字段或方法的描述符由descriptor_index给出
//name_index和descriptor_index都是常量池索引，指向CONSTANT_Utf8_info常量。
//字段和方法名就是代码中出现的（或者编译器生成的）字段或方法的名字。
//Java虚拟机规范定义了一种简单的语法来描述字段和方法,可参照Java虚拟机规范4.3节
type ConstantNameAndTypeInfo struct{
	nameIndex uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}