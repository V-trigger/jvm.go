package heap

import "jvmgo/ch06/classfile"

//字段和方法共有的信息
type MemberRef struct {
	SymRef
	name    string
	descriptor    string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

//getter/setter

func (self *MemberRef) Name() string {
	return self.name
}

func (self *MemberRef) Descriptor() string {
	return self.descriptor
}




