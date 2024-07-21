package main

import "log"

type Handler interface {
	execute(*Command)
	setNext(Handler)
}

type HandlerOne struct {
	next Handler
}

func (h *HandlerOne) execute(c *Command) {
	if c.handlerOneDone {
		log.Println("handler two done")
		h.next.execute(c)
	}
	c.handlerOneDone = true
	h.next.execute(c)
}

func (h *HandlerOne) setNext(next Handler) {
	h.next = next
}

type HandlerTwo struct {
	next Handler
}

func (h *HandlerTwo) execute(c *Command) {
	if c.handlerTwoDone {
		log.Println("handler two done")
		h.next.execute(c)
	}
	c.handlerTwoDone = true
	h.next.execute(c)
}

func (h *HandlerTwo) setNext(next Handler) {
	h.next = next
}

type Command struct {
	name           string
	handlerOneDone bool
	handlerTwoDone bool
}

func main() {
	one := HandlerOne{}

	two := HandlerTwo{}
	two.setNext(&one)

	command := Command{
		name: "c1",
	}

	two.execute(&command)
}
