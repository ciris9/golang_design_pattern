package main

import "log"

type Client struct {
}

func (c *Client) Do(r IRequest) {
	log.Println("do request")
	r.DoRequest()
}

type IRequest interface {
	DoRequest()
}

type IRequestImpl struct {
}

func (i *IRequestImpl) DoRequest() {
	log.Println("[IRequestImpl] do request]")
}

type Adapter struct {
	request IRequest
}

func (a *Adapter) DoRequest() {
	a.request.DoRequest()
}

func main() {
	impl := &IRequestImpl{}
	a := &Adapter{
		request: impl,
	}
	client := &Client{}
	client.Do(a)
}
