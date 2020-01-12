package main

import "fmt"

type Proxy struct {
	origin *Origin
}

func NewProxy() *Proxy {
	origin := NewOrigin()
	return &Proxy{origin: origin}
}

func (p *Proxy) Request() {
	p.PreRequest()
	p.origin.Request()
	p.PostRequest()
}

func (p *Proxy) PreRequest() {
	fmt.Println("Proxy PreRequest")
}

func (p *Proxy) PostRequest() {
	fmt.Println("Proxy PostRequest")
}
