package main


//访问者（Visitor）模式的定义：将作用于某种数据结构中的各元素的操作分离出来封装成独立的类，
//使其在不改变数据结构的前提下可以添加作用于这些元素的新的操作，
//为数据结构中的每个元素提供多种访问方式。它将对数据的操作与数据结构进行分离
//是行为类模式中最复杂的一种模式。
func main()  {
	os := NewObjectStruct()
	e1 := NewElement("A")
	e2 := NewElement("B")

	os.Add(e1)
	os.Add(e2)

	v := NewVisitor("v")
	os.Accept(v)

}
