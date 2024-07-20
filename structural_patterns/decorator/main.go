package main

import "log"

// Component 抽象零件
type Component interface {
	getNumber() int
}

// ComponentOne 具体零件1
type ComponentOne struct {
}

func (c *ComponentOne) getNumber() int {
	return 1
}

// ComponentTwo 具体零件2
type ComponentTwo struct{}

func (c *ComponentTwo) getNumber() int {
	return 2
}

type Decorator struct {
	component Component
}

func (d *Decorator) getNumber() int {
	return d.component.getNumber() + 10
}
func main() {
	one := &ComponentOne{}
	decoratorOne := Decorator{
		component: one,
	}
	two := &ComponentTwo{}
	DecoratorTwo := Decorator{
		component: two,
	}
	log.Println(decoratorOne.getNumber())
	log.Println(DecoratorTwo.getNumber())
}
