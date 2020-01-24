package main

import "fmt"

type DressInter interface {
	Dress()
}

type PersonInter interface {
	SetName(name string)
	BeforeOut()
	Out()
}

type Person struct {
	Specific DressInter
	name     string
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) Out() {
	p.BeforeOut()
	fmt.Println(p.name + " go out...")
}

func (p *Person) BeforeOut() {
	if p.Specific == nil {
		return
	}

	p.Specific.Dress()
}
