package main

import "fmt"

// Subject 抽象主体
type Subject interface {
	registerObserver(observer Observer)
	deleteObserver(observer Observer)
	notifyAll() //通知所有观察者
}

// Item 具体主体
type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) updateAvailability() {
	i.inStock = true
	i.notifyAll()
}

func (i *Item) registerObserver(observer Observer) {
	i.observerList = append(i.observerList, observer)
}

func (i *Item) deleteObserver(removedObserver Observer) {
	observerList := i.observerList
	observerListLength := len(observerList)
	for index, observer := range observerList {
		if removedObserver.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[index] = observerList[index], observerList[observerListLength-1]
			i.observerList = observerList
			return
		}
	}
	i.observerList = observerList
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

// Observer 抽象观察者
type Observer interface {
	update(string)
	getID() string
}

// Customer 具体观察者
type Customer struct {
	id string
}

func (c *Customer) getID() string {
	return c.id
}

func (c *Customer) update(id string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, id)
}

func main() {

	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.registerObserver(observerFirst)
	shirtItem.registerObserver(observerSecond)

	shirtItem.updateAvailability()
}
