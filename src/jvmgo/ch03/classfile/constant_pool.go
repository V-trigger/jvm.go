package classfile

//定义ConstantInfo接口,各类型的常量都要事先这个接口
type ConstantInfo interface {
	// 读取常量信息，需要由具体的常量结构体实现。
	// readConstantInfo()函数先读出tag值，然后调用newConstantInfo（）函数创建具体的常量，
	// 最后调用常量的readInfo（）方法读取常量信息
	readInfo(reader *ClassReader)
}

// 常量池
// 常量池放得是一个一个常量,常量可以看做一个个结构体,比如Utf8常,放得是一个uft8编码的字符串,他的结构如下:
// CONSTANT_Utf8_info {
//     u1 tag;
// 	   u2 length;
// 	   u1 bytes[length];
// }
// tag指出他的数据类型，也就是CONSTANT_Utf8_info常量
// 每个常量都有一个u1类型的tag, 知道了类型就能转换为Golang对应的结构,并根据每个常量的结构组装数据
//
// 
// 有些类型是带有索引字段的,比如CONSTANT_Class_info,他的结构如下:
// CONSTANT_Class_info {
//     u1 tag;
// 	   u2 name_index;
// }
//
// 索引就是按顺序读取结构体的位置，这里是从1开始, 0是无效索引
// 比如CONSTANT_Class_info的name_index指向的就是一个CONSTANT_Utf8_info常量
// 假如name_index的值为1，那么常量池第一个读取到的一定是一个CONSTANT_Utf8_info常量
// 这里就可以把常量池转换成Golang的一个切片(slice), 索引对应的就是slice的下标
type ConstantPool []ConstantInfo


//读取单个常量池中的数据
//读取常量池的时候回循环调用，并把Java Class的结构转换为Golang对应的结构
//常量池中的数据结构为 tag:1byte, data: 2~n个字节
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	//读取tag, tag只在生成ConstantInfo接口的实现类的时候用
	//具体的实现类处理的时候从tag后一个字节开始读取处理
	tag := reader.readUint8()

	//根据tag生成常量池数据,具体为实现了ConstantInfo接口的对应的结构体
	c := newConstantInfo(tag, cp)

	//调用对应结构体的readInfo()的方法
	//对应的tag已经读取过了
    //所以结构体初始化的时候不用再去管tag,从tag后开始读取并处理就可以了
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
			return &ConstantUtf8Info{}
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
            //CONSTANT_Long_info和CONSTANT_Double_info各占两个位置
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
	//获取常量数据并转换为*ConstantNameAndTypeInfo类型
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

//按索引从常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {
	//获取常量数据并转换为*ConstantClassInfo类型
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

//从常量池查找UTF-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	//获取常量数据并转换为*ConstantUtf8类型
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}