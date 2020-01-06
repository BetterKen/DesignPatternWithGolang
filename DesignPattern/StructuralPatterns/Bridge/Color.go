package main

import "fmt"

type Red struct {
}

func (r *Red) SayColor() {
	fmt.Print("我的颜色是红色")
}

type Black struct {
}

func (b *Black) SayColor() {
	fmt.Print("我的颜色是黑色")
}
