package main

type aggregateInter interface {
	Add(o interface{})
	Remove()
	GetIterator() *iterator
}

type Aggregate struct {
	list []interface{}
}

func NewAggregate() aggregateInter {
	return &Aggregate{list: make([]interface{}, 0)}
}

func (a *Aggregate) Add(o interface{}) {
	a.list = append(a.list, o)
}

func (a *Aggregate) Remove() {
	a.list = a.list[:len(a.list)-1]
}
func (a *Aggregate) GetIterator() *iterator {
	return NewIterator(a.list)
}
