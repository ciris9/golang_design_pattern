package main

type Product interface {
	GetName() string
}

type product1 struct {
	name string
}

func NewProduct1(name string) Product {
	return &product1{name: name}
}

func (p *product1) GetName() string {
	return p.name
}

type product2 struct {
	name string
}

func (p *product2) GetName() string {
	return p.name
}

func NewProduct2(name string) Product {
	return &product2{name: name}
}

// 工厂类，创建不同的工厂实例后可以生产不同的产品
type factory struct{}

func (f *factory) CreateProduct(productName string) Product {
	if productName == "product1" {
		return NewProduct1("product1")
	}
	if productName == "product2" {
		return NewProduct2("product2")
	}
	return nil
}

func main() {
	factory := &factory{}
	product1 := factory.CreateProduct("product1")
	println(product1.GetName())
	product2 := factory.CreateProduct("product2")
	println(product2.GetName())
}
