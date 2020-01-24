package main

import "fmt"

type Boy struct {
	Person
}

func (_ *Boy) Dress() {
	fmt.Println("wash face")
}
