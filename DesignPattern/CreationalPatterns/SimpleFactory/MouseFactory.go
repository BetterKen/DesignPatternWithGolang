package main

const (
	Dell = iota
	HP
)

type MouseInter interface {
	SayHi()
}

func newMouseFactory(mCase int) MouseInter {
	switch mCase {
	case Dell:
		return &DellMouse{}
	case HP:
		return &HpMouse{}
	default:
		return &HpMouse{}
	}
}
