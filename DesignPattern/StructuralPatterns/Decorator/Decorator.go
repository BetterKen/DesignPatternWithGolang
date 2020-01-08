package main

import "fmt"

type DecoCar struct {
	car Car
}

func NewDecoCar() *DecoCar {
	return new(DecoCar)
}

func (d *DecoCar) Run() {
	fmt.Print("加速")
	d.car.Run()
}
