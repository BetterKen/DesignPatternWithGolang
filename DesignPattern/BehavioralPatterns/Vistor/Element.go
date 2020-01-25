package main

import "fmt"

type elementInter interface {
	Accept(visitor Visitor)
	Operation()
}

type Element struct {
	visitor *Visitor
	name    string
}

func NewElement(name string) *Element {
	return &Element{
		visitor: nil,
		name:    name,
	}
}

func (e *Element) Accept(visitor *Visitor) {
	visitor.Visit(e)
}

func (e *Element) Operation() {
	fmt.Println(e.name + "Operation")
}
