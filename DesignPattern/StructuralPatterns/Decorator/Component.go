package main

import "fmt"

type CarInter interface {
	Run()
}

type Car struct {
}

func (c *Car) Run() {
	fmt.Print("行驶")
}


