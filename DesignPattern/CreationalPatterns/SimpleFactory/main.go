package main

func main() {
	f1 := newMouseFactory(Dell)
	f2 := newMouseFactory(HP)
	f1.SayHi()
	f2.SayHi()
}
