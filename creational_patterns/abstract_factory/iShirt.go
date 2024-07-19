package main

import "fmt"

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) String() {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getSize() int {
	return s.size
}
