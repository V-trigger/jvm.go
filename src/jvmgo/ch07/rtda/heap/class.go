package heap

import(
	"jvmgo/ch07/classfile"
	"strings"
)

//需要放进方法区的类信息

type Class struct {
	//访问标识符
	accessFlags         uint16

	//当前class完全限定名称
	//java/lang/Object的形式
	name                string

	//父类完全限定名称
	//java/lang/Object的形式
	superClassName      string

	//接口完全限定名称表
	//java/lang/Object的形式
	interfaceNames      []string

	//常量池
	constantPool        *ConstantPool
	//字段表
	fields              []*Field
	//方法表
	methods             []*Method

	//类加载器
	loader              *ClassLoader
	//父类指针
	superClass          *Class
	//接口指针
	interfaces          []*Class
	//接口变量数量 
	instanceSlotCount   uint
	//静态变量数量
	staticSlotCount     uint
	//静态变量
	staticVars          Slots
}


//ClassFile结构体转Class结构体
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

//是否设置public
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

//是否设置final
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
}

//是否为父类
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags & ACC_SUPER
}

//是否为接口
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags & ACC_INTERFACE
}

//是否是抽象类
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags & ACC_ABSTRACT
}

//是否为synthetic TODO
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags & ACC_SYNTHETIC
}

//是否为注解类
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags & ACC_ANNOTATION
}

//是否为枚举类
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags & ACC_ENUM
}


// getters
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

// java虚拟机规范5.4.4  是否可以访问
func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName() == other.getPackageName()
}

//获取包名
func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

//查找主方法
func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

//查找静态方法
func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

//实例化对象
func (self *Class) NewObject() *Object {
	return newObject(self)
}