package main

import "fmt"

type Bus struct {
	Wheels  string
	Chassis string
	Seat    string
}

func (b *Bus) Show() {
	fmt.Println(b.Wheels + "---" + b.Chassis + "---" + b.Seat)
}

type BusBuilder struct {
	bus *Bus
}

func (bb *BusBuilder) NewProduct() {
	bb.bus = new(Bus)
}

func (bb *BusBuilder) BuildWheels() {
	bb.bus.Wheels = "Bus Wheels"
}

func (bb *BusBuilder) BuildChassis() {
	bb.bus.Chassis = "Bus Chassis"
}

func (bb *BusBuilder) BuildSeat() {
	bb.bus.Seat = "Bus Seat"
}

func (bb *BusBuilder) GetResult() interface{} {
	return bb.bus
}
