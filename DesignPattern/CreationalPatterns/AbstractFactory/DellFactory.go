package main

type DellKeyBo struct {
}

func (dell *DellKeyBo) SayHi() {
	println("DellKeyBo")
}

type DellMouse struct {
}

func (dell *DellMouse) SayHi() {
	println("DellMouse")
}

type DellFactory struct {
}

func (df *DellFactory) CreateMouse() MouseInter {
	return new(DellMouse)
}

func (df *DellFactory) CreateKeyBo() KeyBoInter {
	return new(DellKeyBo)
}
