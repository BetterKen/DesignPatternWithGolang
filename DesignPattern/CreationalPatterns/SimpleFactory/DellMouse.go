package main

import "fmt"

type DellMouse struct {
}

func (dell *DellMouse) SayHi() {
	fmt.Println("DellMouse")
}
