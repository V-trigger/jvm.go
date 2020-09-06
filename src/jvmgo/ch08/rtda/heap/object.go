package heap

type Object struct {
	class *Class
	// fields  Slots
	data   interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data: newSlots(class.instanceSlotCount),
	}
}

//getter/setter 

func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

func (self *Object) Class() *Class {
	return self.class
}