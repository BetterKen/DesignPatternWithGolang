package main

type ReceiverInter interface {
	action()
}

type InvokerInter interface {
	setCommand(inter commandInter)
	call()
}

type commandInter interface {
	execute()
}

type command struct {
	receiver Receiver
}

func NewCommand() *command {
	receiver := Receiver{}
	c := &command{}
	c.receiver = receiver
	return c
}


func (c *command) execute() {
	c.receiver.action()
}
