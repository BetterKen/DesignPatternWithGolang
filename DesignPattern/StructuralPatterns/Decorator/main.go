package main

//为让自己的能力增强，使得增强后的自己能够使用更多的方法，拓展在自己基础之上的功能的，就是装饰器模式
//让别人帮助你做你并不关心的事情，这是代理模式

func main() {
	d := NewDecoCar()
	d.Run()
}