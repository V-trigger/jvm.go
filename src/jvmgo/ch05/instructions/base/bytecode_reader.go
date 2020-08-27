package base

type BytecodeReader struct {
	code   []byte
	pc     int
}

func (self *BytecodeReader) Reset(code []byte, pc int ){
	self.code = code
	self.pc = pc
}

//读取uint8
func (self *BytecodeReader) ReadUint8() uint8{
	i := self.code[self.pc]
	self.pc++
	return i
}

//读取uint8再转换
func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

//uint16连续读取两字节,一个高位，一个低位
func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}

//int16类型可以读取出uint16类型再转换
func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

//int32类型，连续读取四个字节
func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}