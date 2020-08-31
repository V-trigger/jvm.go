package heap

import(
	"fmt"
	"jvmgo/ch06/classfile"
	"jvmgo/ch06/classpath"
)

//类加载器
//ClassLoader依赖Classpath来搜索和读取class文件，cp字段保存Classpath指针
//classMap字段记录已经加载的类数据，key是类的完全限定名
type ClassLoader struct {
	cp *classpath.Classpath
    classMap    map[string]*Class   
}

//创建ClassLoader实例
func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
    return &ClassLoader{
		cp: cp,
		classMap: make(map[string]*Class),
	}
}

//把类数据加载到方法区
func (self *ClassLoader) LoadClass(name string) *Class {
    if class, ok := self.classMap[name]; ok {
		//类已加载
		return class
	}
	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	//找到class文件并把数据读取到内存
	data, entry := self.readClass(name)
	//解析class文件
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	//调用classpath的ReadClass方法,读取class文件数据
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

//解析Class文件数据
func (self *ClassLoader) defineClass(data []byte) *Class {
	//把class文件数解析成Class结构体
	class := parseClass(data)
	class.loader = self
	//继承链上的父类全部加载
	resolveSuperClass(class)
	//所有接口加载
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

//class文件数据解析成Class结构体
func parseClass(data []byte) *Class {
	//调用classfile的Parse()将[]byte数据解析成ClassFile结构体
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	//转成Class结构体再返回
	return newClass(cf)
}

//递归加载继承链上的所有类
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

//加载所有实现的接口
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	// todo
}

func prepare(class *Class) {
	//计算实例字段的个数，并编号
	calcInstanceFieldSlotIds(class)
	//计算静态字段的个数，并编号
	calcStaticFieldSlotIds(class)

	allocAndInitStaticVars(class)
}

//计算实例字段的个数，并编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}

	for _, field := range class.fields {
        if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

//计算静态字段的个数，并编号
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

//给变量分配空间并赋初值
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
            initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descriptor() {
			case "Z", "B", "C", "S", "I":
				val := cp.GetConstant(cpIndex).(int32)
				vars.SetInt(slotId, val)
			case "J":
				val := cp.GetConstant(cpIndex).(int64)
				vars.SetLong(slotId, val)
			case "F":
				val := cp.GetConstant(cpIndex).(float32)
				vars.SetFloat(slotId, val)
			case "D":
				val := cp.GetConstant(cpIndex).(float64)
				vars.SetDouble(slotId, val)
			case "Ljava/lang/String;":
				panic("string init todo")
		}
	}
}