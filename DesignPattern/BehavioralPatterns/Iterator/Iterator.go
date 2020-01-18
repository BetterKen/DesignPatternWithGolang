package main

type iteratorInter interface {
	First() interface{}
	Next() interface{}
	HasNext() bool
}

type iterator struct {
	index int
	list  []interface{}
}

func NewIterator(list []interface{}) *iterator {
	return &iterator{
		index: 0,
		list:  list,
	}
}

func (i *iterator) First() interface{} {
	return i.list[0]
}
func (i *iterator) Next() interface{} {
	if i.HasNext() {
		i.index = i.index + 1
		return i.list[i.index-1]
	}
	return nil
}

func (i *iterator) HasNext() bool {
	return i.index < len(i.list)
}
