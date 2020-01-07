package main

import "fmt"

type Composite struct {
	component
	children []Component //叶子集合
}

//创建一个组合结构体
func NewComposite() *Composite {
	return &Composite{children: make([]Component, 0)}
}

func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.children = append(c.children, child) //加入孩子节点
}
func (c *Composite) Print(pre string) { //打印显示每一个节点
	fmt.Println(pre, c.name)
	pre += "  "
	for _, leaf := range c.children {
		leaf.Print(pre)
	}
}
