package main

const (
	DELL = iota
	HP
)

type MouseInter interface {
	SayHi()
}

type FactoryInter interface {
	CreateMouse() MouseInter
}
