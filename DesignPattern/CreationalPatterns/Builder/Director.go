package main

type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) Generate() interface{} {
	d.builder.NewProduct()
	d.builder.BuildWheels()
	d.builder.BuildChassis()
	d.builder.BuildSeat()
	return d.builder.GetResult()
}
