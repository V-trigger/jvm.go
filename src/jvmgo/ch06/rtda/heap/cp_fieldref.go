package heap

import "jvmgo/ch06/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef{
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

//字段符号引用解析 java虚拟机规范5.4.3.2
func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	field := lookupField(c, self.name, self.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		// 名称和描述符都要相同
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	//到接口中找
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor);field != nil {
			return field
		}
	}
	//到父类中找
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}
	return nil
}

