package classfile
import(
	"fmt"
	"unicode/utf16"
)

//CONSTANT_Utf8_info常量里放的是MUTF-8编码的字符串， 结构如下：
// CONSTANT_Utf8_info {
// 	u1 tag;
// 	u2 length;
// 	u1 bytes[length];
// }
//
//字符串在class文件中是以MUTF-8（Modified UTF-8）方式编码的
//MUTF-8和标准的UF8并不兼容
//差别有两点：
//一是null字符（代码点U+0000）会被编码成2字节： 0xC0、0x80
//二是补充字符（Supplementary Characters，代码点大于 U+FFFF的Unicode字符）是按UTF-16拆分为代理对（Surrogate Pair）分别编码的
type ConstantUtf8Info struct {
	str string
}

//读取出[]byte，然后调用decodeMUTF8()函数 把它解码成Go字符串
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	//读取长度
	length := uint32(reader.readUint16())
	//根据长度读取[]byte
	bytes := reader.readBytes(length)
	//将[]byte转换为utf8字符
	self.str = decodeMUTF8(bytes)
}

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}
