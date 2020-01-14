package main

type Invoker struct {
	command commandInter
}

func NewInvoker() *Invoker {
	return &Invoker{}
}

func (i *Invoker) setCommand(command commandInter) {
	i.command = command

}

func (i *Invoker) call() () {
	i.command.execute()

}
