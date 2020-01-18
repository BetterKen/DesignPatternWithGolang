package main

import "fmt"

func main() {
	agg := NewAggregate()
	agg.Add("甲")
	agg.Add("已")
	agg.Add("丙")
	agg.Add("丁")
	agg.Remove()
	i := agg.GetIterator()
	for i.HasNext() {
		fmt.Println(i.Next())
	}
}
