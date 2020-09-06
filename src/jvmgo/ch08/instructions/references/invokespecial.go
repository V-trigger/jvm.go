package references

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	//拿到当前类、当前常量池、方法符号引用
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)

	//解析符号引用，拿到解析后的类和方法
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()

	//构造方法必须本类调用
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}

	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {

		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
			methodRef.Name(), methodRef.Descriptor())
	}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)

}