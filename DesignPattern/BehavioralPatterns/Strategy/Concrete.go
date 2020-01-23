package main

import "fmt"

type ConcreteA struct {
}

func (c *ConcreteA) strategyMethod() {
	fmt.Println("ConcreteA strategyMethod")
}

type ConcreteB struct {
}

func (c *ConcreteB) strategyMethod() {
	fmt.Println("ConcreteB strategyMethod")
}
