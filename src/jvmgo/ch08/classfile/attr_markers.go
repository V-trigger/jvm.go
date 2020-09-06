package classfile

// Deprecated和Synthetic是最简单的两种属性，仅起标记作用，不包含任何数据。
// 这两种属性都是JDK1.1引入的，可以出现在ClassFile、field_info和method_info结构中，它们的结构定义如下：
// Deprecated_attribute {
// 	u2 attribute_name_index;
// 	u4 attribute_length; 
// }
// Synthetic_attribute {
// 	u2 attribute_name_index;
// 	u4 attribute_length;
// } 
// 由于不包含任何数据，所以attribute_length的值必须是0。
// Deprecated属性用于指出类、接口、字段或方法已经不建议使用
//Synthetic属性用来标记源文件中不存在、由编译器生成的类成 员，引入Synthetic属性主要是为了支持嵌套类和嵌套接口
type MarkerAttribute struct{}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}
//实现AttributeInfo接口, 由于这两个属性都没有数据，所以readInfo()方法是空的。
func (self *MarkerAttribute) readInfo(reader *ClassReader) { 
	// read nothing 
}