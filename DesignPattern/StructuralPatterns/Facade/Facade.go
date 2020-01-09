package main

type Facade struct {
	Japanese *Japanese
	Chinese  *Chinese
	English  *English
}

func NewFacade() *Facade {
	Japanese := new(Japanese)
	Chinese := new(Chinese)
	English := new(English)

	return &Facade{
		Japanese: Japanese,
		Chinese:  Chinese,
		English:  English,
	}
}

func (f *Facade) SayHello() {
	f.Chinese.SayHello()
	f.Japanese.SayHello()
	f.English.SayHello()
}
