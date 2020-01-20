package main

type Caretaker struct {
	memento *Memento
}

func NewCaretaker() *Caretaker {
	return new(Caretaker)
}

func (c *Caretaker) SetMemento(memento *Memento) {
	c.memento = memento
}

func (c *Caretaker) GetMemento() *Memento {
	return c.memento
}
