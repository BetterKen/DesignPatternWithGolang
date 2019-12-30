package main

import "sync"

type Single struct {
	data int
}

var singleton *Single
var once sync.Once
var mux sync.Mutex

//最好
func GetInstance() *Single {
	once.Do(func() {
		singleton = &Single{data: 100}
	})
	return singleton
}

//次好 双重检查机制
func GetInstance2() *Single {
	if singleton == nil {
		mux.Lock()
		defer mux.Unlock()
		if singleton == nil {
			singleton = new(Single)
			singleton.data = 100
		}
	}
	return singleton
}

//最差 并发下有问题
func GetInstance3() *Single {
	if singleton == nil {
		singleton = new(Single)
		singleton.data = 100
	}
	return singleton
}
