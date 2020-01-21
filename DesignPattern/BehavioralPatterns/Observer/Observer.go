package main

import "fmt"

type observerInter interface {
	Response()
}

type Observer struct {
	name string
}

func (o *Observer) Response() {

}

type AObserver struct {
	Observer
}

func NewAObserver(name string) *Observer {
	return &Observer{name: name}
}
func (o *AObserver) Response() {
	fmt.Println("AObserver Response")
}


type BObserver struct {
	Observer
}

func NewBObserver(name string) *Observer {
	return &Observer{name: name}
}
func (o *BObserver) Response() {
	fmt.Println("BObserver Response")
}

