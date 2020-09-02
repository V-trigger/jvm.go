package classfile


// Java虚拟机规范定义了u1、u2和u4三种数据类型来表示1、 2和4字节无符号整数
// 分别对应Go语言的uint8、uint16和uint32类型。
// 相同类型的多条数据一般按表（table）的形式存储在class文件中。
// 表由表头和表项（item）构成，表头是u2或u4整数。假设表头是n，后面就紧跟着n个表项数据。

import(
    "encoding/binary"
)

type ClassReader struct {
    data []byte
}

//读取u1表头
func (self *ClassReader) readUint8() uint8 {
    val := self.data[0]
    self.data = self.data[1:]
    return val
}

//读取u2表头
func (self *ClassReader) readUint16() uint16 {
    val := binary.BigEndian.Uint16(self.data)
    self.data = self.data[2:]
    return val
}

//读取u4表头
func (self *ClassReader) readUint32() uint32 {
    val := binary.BigEndian.Uint32(self.data)
    self.data = self.data[4:]
    return val
}

//readUint64（）读取uint64（Java虚拟机规范并没有定义u8）类型
func (self *ClassReader) readUint64() uint64 {
    val := binary.BigEndian.Uint64(self.data)
    self.data = self.data[8:]
    return val
}

//读取uint16表，表的大小由开头的uint16数据指出
func (self *ClassReader) readUint16s() []uint16 {
    //u2表头
    n := self.readUint16()
    //创建uint16切片，保存u1数据项
    s := make([]uint16, n)
    //读取u1数据项
    for i := range s {
        s[i] = self.readUint16()
    }
    return s
}

//读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte {
    bytes := self.data[:n]
    self.data = self.data[n:]
    return bytes
}