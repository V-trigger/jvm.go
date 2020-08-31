package heap

import "jvmgo/ch06/classfile"

//类的成员
type ClassMember struct {
	//访问标识符
	accessFlags  uint16
	//变量或方法名称
	name         string
	//变量或方法描述符
	descriptor   string
	//类指针
	class        *Class
}

//复制MemberInfo数据
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

//是否为static
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags & ACC_STATIC
}

//是否是final
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags & ACC_PROTECTED
}

func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags & ACC_PRIVATE
}

//是否有权限访问
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	//如果字段是public，则任何类都可以访问
	if self.IsPublic() {
		return true
	}
	c := self.class
	//如果字段是protected,则只有子类和同一个包下的类可以访问
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}
	//如果public、protected、private都不是，那就只有默认权限了
	//默认权限只有同包能访问
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}

	return d == c
}

func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}

func (self *ClassMember) Class() *Class {
	return self.class
}