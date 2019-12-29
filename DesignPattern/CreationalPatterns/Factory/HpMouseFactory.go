package main

import "fmt"

type HpMouse struct {
}

func (hp *HpMouse) SayHi() {
	fmt.Println("HpMouse")
}

type HpMouseFactory struct {
}

func (hf *HpMouseFactory) CreateMouse() MouseInter {
	return &HpMouse{}
}
