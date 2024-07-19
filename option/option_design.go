package main

import "fmt"

//主要还是golang语法糖太少了，option模式主要是为了将某个参数转换为函数可选项的模式
//这个设计模式在golang用的极为常见，诸如k8s等项目都使用了大量的option模式

type Message struct {
	id      int
	name    string
	address string
}

func (msg *Message) String() {
	fmt.Printf("ID:%d \n- Name:%s \n- Address:%s \n", msg.id, msg.name, msg.address)
}

var DefaultMessage = &Message{}

type Option func(msg *Message)

func WithID(id int) Option {
	return func(msg *Message) {
		msg.id = id
	}
}

func WithName(name string) Option {
	return func(msg *Message) {
		msg.name = name
	}
}

func WithAddress(address string) Option {
	return func(msg *Message) {
		msg.address = address
	}
}

func NewByOption(opts ...Option) *Message {
	msg := DefaultMessage
	for _, opt := range opts {
		opt(&Message{})
	}
	return msg
}

func main() {
	message2 := NewByOption(WithID(2), WithName("message2"), WithAddress("cache2"))
	message2.String()
}
