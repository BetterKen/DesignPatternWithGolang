package main

//桥接（Bridge）模式的定义如下：将抽象与实现分离，使它们可以独立变化。
//它是用组合关系代替继承关系来实现，从而降低了抽象和实现这两个可变维度的耦合度
func main() {
	color := new(Red)
	size := new(Big)
	animal := new(Dog)

	species := NewSpecies()
	species.SetAnimal(animal)
	species.SetColor(color)
	species.SetSize(size)

	species.show()
}
