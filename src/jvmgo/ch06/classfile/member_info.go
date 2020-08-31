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
    cp ConstantPool
    accessFlags uint16
    nameIndex uint16
    descriptorIndex uint16
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

func (self *MemberInfo) AccessFlags() uint16 {
    return self.accessFlags
}

func (self *MemberInfo) CodeAttribute() *CodeAttribute {
    for _, attrInfo := range self.attributes {
        switch attrInfo.(type) {
            case *CodeAttribute:
                return attrInfo.(*CodeAttribute)
        }
    }
    return nil
}

func (self *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute{
    for _, attrInfo := range self.attributes {
        switch attrInfo.(type) {
        case *ConstantValueAttribute:
            return attrInfo.(*ConstantValueAttribute)
        }
    }
    return nil
}

