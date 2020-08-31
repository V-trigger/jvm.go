package heap

//符号引用
type SymRef struct {
	//运行时常量池
	cp *ConstantPool
	//类的完全限定名
	className string
	//解析后的Class结构体
	class *Class
}

func (self *SymRef) ResolvedClass() *Class {
    if self.class == nil {
		self.resolvedClassRef()
	}
	return self.class
}

//通过当前类的加载器加载运行时常量池中的类
func (self *SymRef) resolvedClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegaAccessError")
	}
	self.class = c
}

