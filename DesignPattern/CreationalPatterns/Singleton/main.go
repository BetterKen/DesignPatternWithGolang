package main

import "fmt"

func main() {
	t11 := GetInstance()
	t12 := GetInstance()

	t21 := GetInstance2()
	t22 := GetInstance2()

	t31 := GetInstance3()
	t32 := GetInstance3()

	if t11 == t12 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等")
	}

	if t21 == t22 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等")
	}

	if t31 == t32 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等")
	}


}
