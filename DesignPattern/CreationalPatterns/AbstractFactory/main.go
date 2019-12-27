package main

//抽象工厂模式
//提供一个创建一系列相关或相互以来的对象的接口
//当产品只有一个的时候抽象工厂模式变成工厂模式
//当工厂模式的产品变为多个的时,工厂模式变为抽象工厂模式
func main()  {
	var pcFactory PcFactoryInter
	pcFactory = new(HpFactory)
	pcFactory.CreateMouse().SayHi()
	pcFactory.CreateKeyBo().SayHi()

	pcFactory = new(DellFactory)
	pcFactory.CreateMouse().SayHi()
	pcFactory.CreateKeyBo().SayHi()


}
