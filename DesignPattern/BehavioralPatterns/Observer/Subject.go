package main

import "fmt"

type subjectInter interface {
	Add(o *Observer)
	Remove(o *Observer)
	Notify()
}

type Subject struct {
	name      string
	observers map[string]*Observer
}

func NewSubject(name string) *Subject {
	return &Subject{
		name:      name,
		observers: make(map[string]*Observer),
	}
}

func (s *Subject) Add(o *Observer) {
	s.observers[o.name] = o
}

func (s *Subject) Remove(o *Observer) {
	delete(s.observers, o.name)
}

func (s *Subject) Notify() {
	fmt.Println("主题:"+s.name)
	for name, object := range s.observers {
		fmt.Println("通知:" + name)
		object.Response()
	}
}

