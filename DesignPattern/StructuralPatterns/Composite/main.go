package main

//组合（Composite）模式的定义：有时又叫作部分-整体模式，
//它是一种将对象组合成树状的层次结构的模式，用来表示“部分-整体”的关系，
//使用户对单个对象和组合对象具有一致的访问性。
//组合模式包含以下主要角色。
//抽象构件（Component）角色：它的主要作用是为树叶构件和树枝构件声明公共接口，并实现它们的默认行为。在透明式的组合模式中抽象构件还声明访问和管理子类的接口；在安全式的组合模式中不声明访问和管理子类的接口，管理工作由树枝构件完成。
//树叶构件（Leaf）角色：是组合中的叶节点对象，它没有子节点，用于实现抽象构件角色中 声明的公共接口。
//树枝构件（Composite）角色：是组合中的分支节点对象，它有子节点。它实现了抽象构件角色中声明的接口，它的主要作用是存储和管理子部件，通常包含 Add()、Remove()、GetChild() 等方法。
func main()  {
	root := NewComponent(CompositeNode,"root")
	l1 := NewComponent(LeafNode,"l1")
	l2 := NewComponent(LeafNode,"l2")
	l3 := NewComponent(LeafNode,"l3")
	l4 := NewComponent(CompositeNode,"l4")

	ll1 := NewComponent(LeafNode,"ll1")
	ll2 := NewComponent(LeafNode,"ll2")
	ll3 := NewComponent(LeafNode,"ll3")

	root.AddChild(l1)
	root.AddChild(l2)
	root.AddChild(l3)
	root.AddChild(l4)

	l4.AddChild(ll1)
	l4.AddChild(ll2)
	l4.AddChild(ll3)

	root.Print("")

}