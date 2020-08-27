package rtda

//表示一个局部变量
type Slot struct {
	num  int32
	ref *Object
}