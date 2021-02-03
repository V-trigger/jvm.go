package classfile

import "fmt"


// Java虚拟机规范使用一种类似C语言的结构体语法来描述class 文件格式。
// 整个class文件被描述为一个ClassFile结构，代码如下： 
//   ClassFile { 
//         u4 magic;
//         u2 minor_version;
//         u2 major_version;
//         //常量池计数器
//         u2 constant_pool_count;  
//         cp_info constant_pool[constant_pool_count-1];
//         u2 access_flags;
//         u2 this_class;
//         u2 super_class;
//         u2 interfaces_count;
//         u2 interfaces[interfaces_count];
//         //字段格式
//         u2 fields_count;
//         field_info fields[fields_count];
//         //方法计数器
//         u2 methods_count;
//         method_info methods[methods_count];
//         //附加属性计数器
//         u2 attributes_count;
//         attribute_info attributes[attributes_count]; 
//     }

type ClassFile struct {
    //u4    magic    魔数，识别Class文件格式    4个字节
    //magic uint32

    //u2    minor_version    副版本号    2个字节
    minorVersion  uint16
    //u2    major_version    主版本号    2个字节
    majorVersion  uint16
    //cp_info    constant_pool    常量池    n个字节
    constantPool  ConstantPool
    //u2    access_flags    访问标志    2个字节
    accessFlags   uint16
    //u2    this_class    类索引    2个字节
    thisClass     uint16
    //u2    super_class    父类索引    2个字节
    superClass    uint16
    //u2    interfaces    接口索引集合    2个字节
    interfaces    []uint16
    //field_info    fields    字段集合    n个字节
    fields        []*MemberInfo
    //method_info    methods    方法集合    n个字节
    methods       []*MemberInfo
    //attribute_info    attributes    附加属性集合    n个字节
    attributes    []AttributeInfo
}

//把[]byte解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile, err error) {
    defer func() { 
        if r := recover(); r != nil {
            var ok bool
            err, ok = r.(error)
            if !ok { 
                err = fmt.Errorf("%v", r) 
            } 
        } 
    }()
    cr := &ClassReader{classData}
    cf = &ClassFile{}
    cf.read(cr)
    return
}

func (self *ClassFile) read(reader *ClassReader) {

    //读取并验证验证魔数
    self.readAndCheckMagic(reader)
    //读取并验证验证版本号
    self.readAndCheckVersion(reader)

    //读取常量池 TODO
    self.constantPool = readConstantPool(reader)

    //读取访问标识
    self.accessFlags = reader.readUint16()

    //标志之后是两个u2类型的常量池索引，分别给出类名和父类名
    //class文件存储的类名类似完全限定名，但是把点换成了斜线，Java语言规范把这种名字叫作二进制名（binary names）
    //因为每个类都有名字，所以thisClass必须是有效的常量池索引。
    //除 java.lang.Object之外，其他类都有父类，所以superClass只在 Object.class中是0，在其他class文件中必须是有效的常量池索引。
    self.thisClass = reader.readUint16()
    self.superClass = reader.readUint16()

    //接口索引表
    //接口索引表，表中存放的也是常量池索引，给出该类实现的所有接口的名字
    self.interfaces = reader.readUint16s()
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
    self.fields = readMembers(reader, self.constantPool)
    self.methods = readMembers(reader, self.constantPool)
    self.attributes = readAttributes(reader, self.constantPool)
}

//验证魔数
//Java编译的class文件的魔数"CA FE BA BE", 16进制 0xCAFEBABE 四个字节
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
    magic := reader.readUint32()
    if magic != 0xCAFEBABE {
        //此处应该抛异常
        //暂时先用panic()终止程序
        panic("java.lang.ClassFormatError: magic!")
    }
}

// class文件的次版本号和主版本号，都是u2类型。
// 假设某class文件的主版本号是M，次版本号是m，那么完整的版本号可以表示成“M.m”的形式。
// 次版本号只在J2SE 1.2之前用过，从1.2开始基本上就没什么用了（都是0）。
// 主版本号在J2SE 1.2之前是45， 从1.2开始，每次有大的Java版本发布，都会加1
// 
// 特定的Java虚拟机实现只能支持版本号在某个范围内的class文件。
// Oracle的实现是完全向后兼容的，比如Java SE 8支持版本号为 45.0~52.0的class文件。
// 如果版本号不在支持的范围内，Java虚拟机 实现就抛出java.lang.UnsupportedClassVersionError异常。
// 这里参考 Java 8，支持版本号为45.0~52.0的class文件。如果遇到其他版本号，
// 暂时先调用panic（）方法终止程序执行
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
    //第五第六个字节是次版本号
    self.minorVersion = reader.readUint16()
    //第六第七个字节是主版本号
    self.majorVersion = reader.readUint16()
    switch self.majorVersion {
        case 45: return
        case 46, 47, 48, 49, 50, 51, 52:
            if self.minorVersion == 0 {
                return
            }
    }
    panic("java.lang.UnsupportedClassVersionError!")
}

//从常量池查找类名
func (self *ClassFile) ClassName() string {
    return self.constantPool.getClassName(self.thisClass)
}

//从常量池查找父类名
func (self *ClassFile) SuperClassName() string {
    if self.superClass > 0 {
        return self.constantPool.getClassName(self.superClass)
    }
    return ""
}

//从常量池查找接口名
func (self *ClassFile) InterfaceNames() []string {
    //接口可以有多个,用一个切片保存
    interfaceNames := make([]string, len(self.interfaces))
    for i, cpIndex := range self.interfaces {
        interfaceNames[i] = self.constantPool.getClassName(cpIndex)
    }
    return interfaceNames
}


//getter方法
func (self *ClassFile) MinorVersion() uint16 {
    return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
    return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
    return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
    return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo {
    return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
    return self.methods
}
