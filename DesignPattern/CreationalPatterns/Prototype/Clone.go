package main

type Example struct {
	name string
}

func (e *Example) clone() *Example {
	cloneExample := *e
	return &cloneExample
}
