package main

type HpKeyBo struct {
}

func (hp *HpKeyBo) SayHi() {
	println("HpKeyBo")
}

type HpMouse struct {
}

func (hp *HpMouse) SayHi() {
	println("HpMouse")
}

type HpFactory struct {
}

func (hf *HpFactory) CreateMouse() MouseInter {
	return new(HpMouse)
}

func (hf *HpFactory) CreateKeyBo() KeyBoInter {
	return new(HpKeyBo)
}
