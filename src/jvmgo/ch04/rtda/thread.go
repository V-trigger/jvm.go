package rtda

//当前线程的栈
type Thread struct {
	pc    int
    stack   *Stack
}

func newThread() *Thread{
    return &Thread{
		stack: newStack(1024),
	}
}

//栈帧入栈
func (self *Thread) PushFrame(frame *Frame) {
    self.stack.push(frame)
}

//出栈
func (self *Thread) PopFrame() *Frame {
    return self.stack.pop()
}

//获取当前栈帧
func (self *Thread) CurrentFrame() *Frame {
    return self.stack.top()
}


func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}