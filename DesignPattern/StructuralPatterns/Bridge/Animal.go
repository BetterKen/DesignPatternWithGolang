package main

import "fmt"

type Dog struct {
}

func (d *Dog) SayHi() {
	fmt.Print("我是一只狗")
}

type Cat struct {
}

func (c *Cat) SayHi() {
	fmt.Print("我是一只猫")
}
