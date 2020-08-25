package classfile

// 由于常量池中存放的信息各不相同，所以每种常量的格式也不同。
// 常量数据的第一字节是tag，用来区分常量类型。
// 下面是Java虚拟机规范给出的常量结构
// cp_info { 
// 	u1 tag; 
// 	u1 info[]; 
// }

//tag标记值
const ( 
	CONSTANT_Class = 7
	CONSTANT_Fieldref = 9
	CONSTANT_Methodref = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String = 8
	CONSTANT_Integer = 3
	CONSTANT_Float = 4
	CONSTANT_Long = 5
	CONSTANT_Double = 6
	CONSTANT_NameAndType = 12
	CONSTANT_Utf8 = 1
	CONSTANT_MethodHandle = 15
	CONSTANT_MethodType = 16
	CONSTANT_InvokeDynamic = 18
)