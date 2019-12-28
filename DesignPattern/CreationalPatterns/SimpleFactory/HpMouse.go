package main

import "fmt"

type HpMouse struct {
}

func (hp *HpMouse) SayHi() {
	fmt.Println("HpMouse")
}
