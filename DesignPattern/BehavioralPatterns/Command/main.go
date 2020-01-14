package main

func main()  {
	invoker := NewInvoker()
	command := NewCommand()
	invoker.setCommand(command)
	invoker.call()
}
