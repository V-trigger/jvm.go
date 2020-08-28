package base

import "jvmgo/ch05/rtda"

//各个指令都继承这个接口
type Instruction interface {
	//从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)

	//执行指令逻辑
	Execute(frame *rtda.Frame)
}

//为了避免重复代码，按照操作数类型定义一些结构体，并实现FetchOperands()方 法。
//这相当于Java中的抽象类，具体的指令继承这些结构体，然后专注实现Execute()方法即可

//nop指令，无动作
type NoOperandsInstruction struct {

}
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}


//跳转指令
type BranchInstruction struct {
	//跳转偏移量
	Offset int
}
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

//存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出。
//把这类指令抽象成Index8Instruction结构体，用 Index字段表示局部变量表索引
type Index8Instruction struct {
	Index uint 
}
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint 
}
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}