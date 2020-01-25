package main

type ObjectStruct struct {
	list []*Element
}

func NewObjectStruct() *ObjectStruct {
	return &ObjectStruct{list: make([]*Element, 0)}
}

func (ob *ObjectStruct) Add(element *Element) {
	ob.list = append(ob.list, element)
}
func (ob *ObjectStruct) Remove(element *Element) {
	for i, e := range ob.list {
		if element == e {
			ob.list = append(ob.list[:i], ob.list[i+1:]...)
		}
	}
}

func (ob *ObjectStruct) Accept(visitor *Visitor) {
	for _,e := range ob.list {
		e.Accept(visitor)
	}
}

