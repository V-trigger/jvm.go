package rtda

import "jvmgo/ch07/rtda/heap"

//表示一个局部变量
type Slot struct {
	num  int32
	ref *heap.Object
}