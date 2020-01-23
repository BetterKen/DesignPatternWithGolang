package main

type strategyInter interface {
	strategyMethod()
}

type contextInter interface {
	setStrategy(inter strategyInter)
	strategyMethod()
}

type Context struct {
	strategy strategyInter
}

func (c *Context) setStrategy(inter strategyInter) {
	c.strategy = inter
}

func (c *Context) strategyMethod() {
	c.strategy.strategyMethod()
}
