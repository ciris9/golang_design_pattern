package main

// ISomethingFactory 抽象工厂
type ISomethingFactory interface {
	makeOne() IOne
	makeTwo() ITwo
}

func GetSomethingFactory(tag string) ISomethingFactory {
	if tag == "one" {
		return &OneFactory{}
	}
	if tag == "two" {
		return &TwoFactory{}
	}
	return nil
}

// OneFactory 具体工厂1
type OneFactory struct {
}

func (o *OneFactory) makeOne() IOne {
	return &One{}
}

func (o *OneFactory) makeTwo() ITwo {
	return &Two{}
}

// TwoFactory 具体工厂2
type TwoFactory struct {
}

func (o *TwoFactory) makeOne() IOne {
	return &One{}
}

func (o *TwoFactory) makeTwo() ITwo {
	return &Two{}
}

// IOne 抽象产品1
type IOne interface {
	setOne(string)
	getOne() string
}

// One 具体产品1
type One struct {
	one string
}

func (o *One) setOne(one string) {
	o.one = one
}

func (o *One) getOne() string {
	return o.one
}

// ITwo 抽象产品2
type ITwo interface {
	setTwo(string)
	getTwo() string
}

type Two struct {
	Two string
}

func (t *Two) setTwo(two string) {
	t.Two = two
}

func (t *Two) getTwo() string {
	return t.Two
}

func main() {
	one := GetSomethingFactory("one")
	two := GetSomethingFactory("two")

	makeOne := one.makeOne()
	makeTwo := two.makeTwo()

	makeOne.setOne("one")
	makeTwo.setTwo("two")

	println(makeOne.getOne())
	println(makeTwo.getTwo())
}
