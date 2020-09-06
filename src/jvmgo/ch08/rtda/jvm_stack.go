package rtda

type Stack struct {
	//最多可容纳多少栈帧
	maxSize    uint

	//当前栈帧数量
	size       uint

	//指向栈顶的栈帧
	_top       *Frame
}

func newStack(maxSize uint) *Stack{
	return &Stack{
		maxSize: maxSize,
	}
}

//栈帧入栈
func (self *Stack) push(frame *Frame) {
    if self.size > self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	//将新入栈的栈帧的下一个栈帧设为当前的栈顶栈帧
	if self._top != nil {
		frame.lower = self._top
	}
	//栈顶栈帧设为新入栈的栈帧
	self._top = frame
	self.size++
}

//栈帧出栈
func (self *Stack) pop() *Frame {
    if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top
}

//返回栈顶栈帧，单不弹出
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top 
}

func (self *Stack) isEmpty() bool {
	return self._top == nil
}