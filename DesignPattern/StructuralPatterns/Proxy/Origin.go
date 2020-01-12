package main

import "fmt"

type Origin struct {
}

func NewOrigin() *Origin {
	return new(Origin)
}

func (o *Origin) Request() {
	fmt.Println("Origin Request Method")
}
