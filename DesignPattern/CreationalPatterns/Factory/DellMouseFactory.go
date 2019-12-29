package main

import "fmt"

type DellMouse struct {
}

func (dell *DellMouse) SayHi() {
	fmt.Println("DellMouse")
}

type DellMouseFactory struct {
}

func (df *DellMouseFactory) CreateMouse() MouseInter {
	return &DellMouse{}
}
