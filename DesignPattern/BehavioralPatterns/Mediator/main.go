package main

func main()  {
	lianjia := NewConcreteMediator("链家")
	collA := NewCollage("中介从业者1")
	collA.SetMedium(lianjia)
	collB := NewCollage("中介从业者2")
	collB.SetMedium(lianjia)
	collC := NewCollage("买房人A")
	collC.SetMedium(lianjia)
	collC.Send()


	woaiwojia := NewConcreteMediator("我爱我家")
	collA1 := NewCollage("中介从业者11")
	collA1.SetMedium(woaiwojia)
	collB1 := NewCollage("中介从业者21")
	collB1.SetMedium(woaiwojia)
	collC1 := NewCollage("买房人A1")
	collC1.SetMedium(woaiwojia)
	collC1.Send()
}