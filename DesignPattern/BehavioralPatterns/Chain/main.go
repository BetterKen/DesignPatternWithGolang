package main

func main()  {
	ha := new(AHandler)
	hb := new(BHandler)
	ha.setNext(hb)

	ha.handleRequest(100)

}
