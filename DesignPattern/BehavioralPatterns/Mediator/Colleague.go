package main

import "fmt"

type CollageInter interface {
	SetMedium(m MediatorInter)
	Receive()
	Send()
}

type Collage struct {
	name     string
	mediator MediatorInter
}

func NewCollage(name string) *Collage {
	return &Collage{
		name: name,
	}
}

func (coll *Collage) SetMedium(m MediatorInter) {
	coll.mediator = m
	coll.mediator.Register(coll)
}

func (coll *Collage) Receive() {
	fmt.Println(coll.name + "收到消息")
}

func (coll *Collage) Send() {
	fmt.Println(coll.name + "发送消息")
	coll.mediator.Relay(coll)
}
