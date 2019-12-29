package main

//工厂模式:定义一个用于创建对象的接口
//让子类决定实例化哪一个类
//使一个类的实例化延迟到其子类
func main() {
	var mf FactoryInter
	mf = new(DellMouseFactory)
	mouse := mf.CreateMouse()
	mouse.SayHi()


	mf = new(HpMouseFactory)
	mouse = mf.CreateMouse()
	mouse.SayHi()
}
