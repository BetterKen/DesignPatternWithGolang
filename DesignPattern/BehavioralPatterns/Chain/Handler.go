package main

import "fmt"

type handlerInter interface {
	setNext(handler handlerInter)
	getNext() handlerInter
	handleRequest(request int)
}

type handlerBase struct {
	next handlerInter
}

func (h *handlerBase) setNext(base handlerInter) {
	h.next = base
}

func (h *handlerBase) getNext() handlerInter {
	return h.next
}

func (h *handlerBase) handleRequest(request int) {

}

type AHandler struct {
	handlerBase
}

func (h *AHandler) handleRequest(request int) {
	if request < 100 {
		fmt.Println("AHandler 处理了请求")
		return
	}
	if h.getNext() != nil {
		h.getNext().handleRequest(request)
		return
	}

	fmt.Println("没有人能处理该请求")
}

type BHandler struct {
	handlerBase
}

func (h *BHandler) handleRequest(request int) {
	if request < 200 {
		fmt.Println("BHandler 处理了请求")
		return
	}
	if h.getNext() != nil {
		h.getNext().handleRequest(request)
		return
	}

	fmt.Println("没有人能处理该请求")
}
