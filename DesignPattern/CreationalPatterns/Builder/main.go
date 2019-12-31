package main

func main() {
	director := new(Director)
	builder := new(BusBuilder)
	director.SetBuilder(builder)
	bus := director.Generate().(*Bus)
	bus.Show()
}
