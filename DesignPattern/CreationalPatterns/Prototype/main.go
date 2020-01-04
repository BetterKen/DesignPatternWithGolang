package main

import "fmt"
//这是一个十分简单的设计模式,可以看做是其他语言中的克隆方法，
//例如 JAVA/PHP 中都有相关方法，从一个内存中已经存在的对象中，
//拷贝出一个一模一样的对象来，针对复杂对象或比较大的对象，要比使用各种设计模式new出来的对象要快的多
//深拷贝
func main() {
	t1 := new(Example)
	t1.name = "t1"
	t2 := t1.clone()

	fmt.Printf("type t1 = %T\n", t1)
	fmt.Printf("type t2 = %T\n", t2)
	fmt.Println(t1.name)
	fmt.Println(t2.name)
	if t1 == t2 {
		fmt.Println("浅拷贝")
	} else {
		fmt.Println("深拷贝")
	}
}
