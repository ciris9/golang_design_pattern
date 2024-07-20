package main

import "log"

type Facader struct {
	partOne   PartOne
	partTwo   PartTwo
	partThree PartThree
}

func NewFacoder() *Facader {
	return &Facader{
		partOne:   PartOne{},
		partTwo:   PartTwo{},
		partThree: PartThree{},
	}
}

func (d *Facader) FacaderDo() {
	d.partTwo.partTwoDo()
	d.partThree.partThreeDo()
	d.partOne.partOneDo()
}

type PartOne struct {
}

func (p *PartOne) partOneDo() {
	log.Println("part one")
}

type PartTwo struct {
}

func (p *PartTwo) partTwoDo() {
	log.Println("part two")
}

type PartThree struct {
}

func (p *PartThree) partThreeDo() {
	log.Println("part three")
}

func main() {
	newDecorator := NewFacoder()
	newDecorator.FacaderDo()
}
