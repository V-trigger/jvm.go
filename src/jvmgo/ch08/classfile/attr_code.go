package classfile

//Code是变长属性，只存在于method_info结构中。
//Code属性中存放字节码等方法相关信息。
//相比前面介绍的几种属性，Code属性比较复杂，其结构定义如下：
// Code_attribute {
// 	u2 attribute_name_index;
// 	u4 attribute_length;
// 	u2 max_stack;
// 	u2 max_locals;
// 	u4 code_length;
// 	u1 code[code_length];
// 	u2 exception_table_length; 
// 	{ 
// 		u2 start_pc;
// 		u2 end_pc;
// 		u2 handler_pc;
// 		u2 catch_type;
// 	}exception_table[exception_table_length];
// 	u2 attributes_count;
// 	attribute_info attributes[attributes_count];
// }
//
//max_stack给出操作数栈的最大深度，max_locals给出局部变量表大小。
//接着是字节码，存在u1表中。最后是异常处理表和属性表。
//在第10章讨论异常处理时，会使用异常处理表

type CodeAttribute struct {
	cp    ConstantPool
	maxStack    uint16
	maxLocals   uint16
	code    []byte
	exceptionTable    []*ExceptionTableEntry
    attributes []AttributeInfo
}

func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}
func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}
func (self *CodeAttribute) Code() []byte {
	return self.code
}
func (self *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return self.exceptionTable
}

type ExceptionTableEntry struct {
	startPc uint16
	endPc uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLen := reader.readUint32()
	self.code = reader.readBytes(codeLen)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry{
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc: reader.readUint16(),
			endPc: reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}