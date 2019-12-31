package main

type Builder interface {
	NewProduct()
	BuildWheels()
	BuildChassis()
	BuildSeat()
	GetResult() interface{}
}

