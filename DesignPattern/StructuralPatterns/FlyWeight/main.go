package main

import "fmt"
//运用共享技术来有効地支持大量细粒度对象的复用。它通过共享已经存在的又橡来大幅度减少需要创建的对象数量、
//避免大量相似类的开销，从而提高系统资源的利用率。
func main()  {
	imf := NewImageFlyWeightFactory()
	f1 := imf.Get("fA")
	fmt.Println(imf.GetMapCount())
	imf.Get("fB")
	fmt.Println(imf.GetMapCount())
	f3 := imf.Get("fA")
	fmt.Println(imf.GetMapCount())

	if f1 == f3{
		println("没有消耗多余内存")
	}

}
