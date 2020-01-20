package main

type Originator struct {
	state string
}

func NewOriginator() *Originator {
	return new(Originator)
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) CreateMemento() *Memento {
	return NewMemento(o.state)
}

func (o *Originator) RestoreMemento(m *Memento) {
	o.SetState(m.GetState())
}

