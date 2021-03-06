package classfile

//CONSTANT_String_info常量表示java.lang.String字面量，结构如下：
// CONSTANT_String_info {
// 	u1 tag;
// 	u2 string_index;
// }

type ConstantStringInfo struct {
	cp ConstantPool
	stringIndex uint16
}

//读取常量池索引
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
    self.stringIndex = reader.readUint16()
}

//按索引从常量池中查找字符串
func (self *ConstantStringInfo) String() string {
    return self.cp.getUtf8(self.stringIndex)
}