package main

type MediatorInter interface {
	Register(c *Collage)
	Relay(c *Collage)
}

type ConcreteMediator struct {
	name     string
	collages map[string]*Collage
}

func NewConcreteMediator(name string) *ConcreteMediator {
	return &ConcreteMediator{
		name:     name,
		collages: make(map[string]*Collage),
	}
}

func (c *ConcreteMediator) Register(coll *Collage) {
	c.collages[coll.name] = coll
}

func (c *ConcreteMediator) Relay(coll *Collage) {
	for name, collage := range c.collages {
		if name != coll.name {
			collage.Receive()
		}
	}
}
