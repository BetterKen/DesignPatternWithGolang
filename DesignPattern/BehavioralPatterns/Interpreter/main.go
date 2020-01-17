package main

import "fmt"

//解释器（Interpreter）模式的定义：给分析对象定义一个语言，并定义该语言的文法表示，
//再设计一个解析器来解释语言中的句子。也就是说，
//用编译语言的方式来分析应用中的实例。这种模式实现了文法表达式处理的接口，该接口解释一个特定的上下文。

func main() {
	p := &Parser{}
	fmt.Print("start\n")
	p.Parse("1 + 3 - 2 + 5 - 6 + 7 + 21")
	fmt.Println(p.Result().Interpret())

}
