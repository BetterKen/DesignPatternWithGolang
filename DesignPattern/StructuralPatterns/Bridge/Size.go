package main

import "fmt"

type Big struct {
}

func (b *Big) SaySize() {
	fmt.Print("我是个大号的")
}

type Small struct {
}

func (s *Small) SaySize() {
	fmt.Print("我是个小号的")
}
