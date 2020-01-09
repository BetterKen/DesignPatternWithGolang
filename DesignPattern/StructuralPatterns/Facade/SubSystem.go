package main

import "fmt"

type Chinese struct {
}

func (c *Chinese) SayHello() {
	fmt.Println("你好")
}

type Japanese struct {
}

func (j *Japanese) SayHello() {
	fmt.Println("库尼奇瓦")
}

type English struct {
}

func (e *English) SayHello() {
	fmt.Println("Hello")
}
