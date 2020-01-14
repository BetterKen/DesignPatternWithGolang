package main

import "fmt"

type Receiver struct {
}

func (r *Receiver) action() {
	fmt.Print("Receiver action")
}
