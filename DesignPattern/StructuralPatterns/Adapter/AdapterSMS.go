package main

type AdapterSMS struct {
	context  string
	receiver string
}

func NewAdapterSMS(context string, receiver string) *SMS {
	sender := "Default sender"
	return NewSMS(context, sender, receiver)
}
