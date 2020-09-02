package heap

import "jvmgo/ch07/classfile"

//方法信息
type Method struct {
	ClassMember
	//操作数栈最大深度
	maxStack    uint
	//局部变量表大小
	maxLocals   uint
	//放方法字节码
	code        []byte
}

//方法表
func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute();codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}

//getter/setter

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) Code() []byte {
	return self.code
}