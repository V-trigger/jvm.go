package classfile

//字段表和方法表，分别存储字段和方法信息。
//字段和方法的基本结构大致相同，差别仅在于属性表
//下面是Java虚拟机规范给出的字段结构定义。
// field_info { 
//     u2 access_flags; 
//     u2 name_index; 
//     u2 descriptor_index; 
//     u2 attributes_count;
//     attribute_info attributes[attributes_count]; 
// }
type MemberInfo struct {
    //常量池
    cp ConstantPool

    //访问标识符
    accessFlags uint16

    //字段或方法的名称索引.指向一个CONSTANT_Utf8_info常量
    nameIndex uint16
    
    //描述符索引,指向一个CONSTANT_Uft8_info
    // 描述符是一种对函数返回值和参数的编码。这种编码叫做JNI字段描述符（JavaNative Interface FieldDescriptors)。、
    // 比如一个数组int[]，就需要表示为这样"[I"。
    // 如果多个数组double[][][]就需要表示为这样 "[[[D"。也就是说每一个方括号开始，就表示一个数组维数。多个方框后面，就是数组的类型。
    // 如果以一个L开头的描述符，就是类描述符，它后紧跟着类的字符串，然后分号";"结束。
    // 具体描述符参照Java虚拟机规范
    descriptorIndex uint16

    //属性表
    attributes []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
    //读取字段表表头
    memberCount := reader.readUint16()
    members := make([]*MemberInfo, memberCount)
    for i := range members {
        members[i] = readMember(reader, cp)
    }
    return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
    return &MemberInfo{
        cp: cp,
        accessFlags: reader.readUint16(),
        nameIndex: reader.readUint16(),
        descriptorIndex: reader.readUint16(),
        attributes: readAttributes(reader, cp),
    }
}

//从常量池查找字段或方法名
func (self *MemberInfo) Name() string {
    return self.cp.getUtf8(self.nameIndex)
}

//从常量池查找字段或方法描述符
func (self *MemberInfo) Descriptor() string {
    return self.cp.getUtf8(self.descriptorIndex)
}

//getter方法,获取访问标识符
func (self *MemberInfo) AccessFlags() uint16 {
    return self.accessFlags
}



