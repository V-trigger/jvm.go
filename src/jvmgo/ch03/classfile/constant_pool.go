package classfile

//保存常量信息
type ConstantInfo interface {
	// 读取常量信息，需要由具体的常量结构体实现。
	// readConstantInfo()函数先读出tag值，然后调用newConstantInfo（）函数创建具体的常量，
	// 最后调用常量的readInfo（）方法读取常量信息
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
		case CONSTANT_Integer:
			return &ConstantIntegerInfo{}
		case CONSTANT_Float:
			return &ConstantFloatInfo{}
		case CONSTANT_Long:
			return &ConstantLongInfo{}
		case CONSTANT_Double:
			return &ConstantDoubleInfo{}
		case CONSTANT_Utf8: 
			return&ConstantUtf8Info{}
		case CONSTANT_String:
			return &ConstantStringInfo{cp: cp}
		case CONSTANT_Class:
			return &ConstantClassInfo{cp: cp}
		case CONSTANT_Fieldref:
			return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
		case CONSTANT_Methodref: 
			return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
		case CONSTANT_InterfaceMethodref:
			return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
		case CONSTANT_NameAndType:
			return &ConstantNameAndTypeInfo{}
		case CONSTANT_MethodType: 
			return &ConstantMethodTypeInfo{}
		case CONSTANT_MethodHandle: 
			return &ConstantMethodHandleInfo{}
		case CONSTANT_InvokeDynamic:
			return &ConstantInvokeDynamicInfo{}
		default: panic("java.lang.ClassFormatError: constant pool tag!") }
}
//常量池
type ConstantPool []ConstantInfo

// 读取常量池
// 常量池实际上也是一个表，但是有三点需要特别注意。
// 第一， 表头给出的常量池大小比实际大1。假设表头给出的值是n，那么常 量池的实际大小是n–1。
// 第二，有效的常量池索引是1~n–1。0是无效索引，表示不指向任何常量。
// 第三，CONSTANT_Long_info和 CONSTANT_Double_info各占两个位置。也就是说，如果常量池中 存在这两种常量，实际的常量数量比n–1还要少，而且1~n–1的某些 数也会变成无效索引。
func readConstantPool(reader *ClassReader) ConstantPool {
    //读取表头
    cpCount := int(reader.readUint16())
    cp := make([]ConstantInfo, cpCount)
    //索引从1开始
    for i := 1; i < cpCount; i++ {
        cp[i] = readConstantInfo(reader, cp)
        switch cp[i].(type) {
            //CONSTANT_Long_info和 CONSTANT_Double_info各占两个位置
            case *ConstantLongInfo, *ConstantDoubleInfo:
                i++
        }
    }
    return cp
}

//按索引查找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
    if cpInfo := self[index]; cpInfo != nil {
        return cpInfo
    }
    panic("Invalid constant pool index!")
}
//按索引从常量池查找字段或方法的名字和描述符
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

//按索引从常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

//从常量池查找UTF-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}