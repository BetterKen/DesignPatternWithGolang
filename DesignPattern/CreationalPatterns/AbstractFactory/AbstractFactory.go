package main

const (
	DELL = iota
	HP
)

type KeyBoInter interface {
	SayHi()
}

type MouseInter interface {
	SayHi()
}

type PcFactoryInter interface {
	CreateMouse() MouseInter
	CreateKeyBo() KeyBoInter
}
