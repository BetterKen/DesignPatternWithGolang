package main

type Animal interface {
	SayHi()
}

type Color interface {
	SayColor()
}

type Size interface {
	SaySize()
}

type Species struct {
	animal Animal
	color  Color
	size   Size
}

func NewSpecies() *Species {
	return new(Species)
}

func (s *Species) SetAnimal(animal Animal) {
	s.animal = animal
}

func (s *Species) SetColor(color Color) {
	s.color = color
}

func (s *Species) SetSize(size Size) {
	s.size = size
}

func (s *Species) show() {
	s.animal.SayHi()
	s.color.SayColor()
	s.size.SaySize()
}
