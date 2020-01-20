package main

import "fmt"

//备忘录（Memento）模式的定义：在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态，
//以便以后当需要时能将该对象恢复到原先保存的状态。该模式又叫快照模式。
//备忘录模式的主要角色如下。
//发起人（Originator）角色：记录当前时刻的内部状态信息，提供创建备忘录和恢复备忘录数据的功能，实现其他业务功能，它可以访问备忘录里的所有信息。
//备忘录（Memento）角色：负责存储发起人的内部状态，在需要的时候提供这些内部状态给发起人。
//管理者（Caretaker）角色：对备忘录进行管理，提供保存与获取备忘录的功能，但其不能对备忘录的内容进行访问与修改。
func main()  {
	o := NewOriginator()
	o.SetState("S1")
	fmt.Println(o.GetState())
	c := NewCaretaker()
	c.SetMemento(o.CreateMemento())
	o.SetState("S2")
	fmt.Println(o.GetState())
	o.RestoreMemento(c.GetMemento())
	fmt.Println(o.GetState())
}