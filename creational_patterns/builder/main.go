package main

import "fmt"

type product struct {
	part1 string
	part2 string
	part3 string
}

type Builder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetProduct() *product
}

type productBuidler struct {
	product *product
}

func (p *productBuidler) GetProduct() *product {
	return p.product
}

func (p *productBuidler) BuildPart1() {
	p.product.part1 = "1"
}
func (p *productBuidler) BuildPart2() {
	p.product.part2 = "2"
}
func (p *productBuidler) BuildPart3() {
	p.product.part3 = "3"
}

type Director struct {
	b Builder
}

func (d *Director) Construct() {
	d.b.BuildPart1()
	d.b.BuildPart2()
	d.b.BuildPart3()
}

func main() {
	buidler := productBuidler{product: &product{}}
	director := Director{b: &buidler}
	director.Construct()
	fmt.Printf("%#v", buidler.GetProduct())
}
