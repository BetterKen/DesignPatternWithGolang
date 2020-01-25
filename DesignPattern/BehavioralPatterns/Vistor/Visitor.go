package main

type visitorInter interface {
	Visit(element Element)
}

type Visitor struct {
	name string
}

func NewVisitor(name string) *Visitor {
	return &Visitor{name: name}
}

func (v *Visitor) Visit(element *Element) {
	element.Operation()
}
