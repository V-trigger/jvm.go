package classfile

//CONSTANT_Class_info常量表示类或者接口的符号引用，结构如下：
// CONSTANT_Class_info {
// 	u1 tag;
// 	u2 name_index;
// }
//和CONSTANT_String_info类似，name_index是常量池索引，指向CONSTANT_Utf8_info常量。
type ConstantClassInfo struct {
	cp ConstantPool
	//名称索引,指向CONSTANT_Utf8_info常量
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}