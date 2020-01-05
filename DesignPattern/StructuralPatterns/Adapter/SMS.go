package main

import "fmt"

type SMS struct {
	context  string
	sender   string
	receiver string
}

func NewSMS(context string, sender string, receiver string) *SMS {
	return &SMS{
		context:  context,
		sender:   sender,
		receiver: receiver,
	}
}

func (s *SMS) SendSms() {
	fmt.Println(s.sender + " send to " + s.receiver + " message:" + s.context)
}
