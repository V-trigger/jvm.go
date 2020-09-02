package classfile

//实现AttributeInfo接口
type UnparsedAttribute struct {
	name string
	length uint32
	info []byte
}

func (self *UnparsedAttribute) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.length)
}