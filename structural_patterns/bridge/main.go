package main

import "fmt"

// Computer 抽象部分
type Computer interface {
	Print()
	SetPrinter(Printer)
}

// Mac 精确抽象，mac电脑
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

// Windows 精确抽象，windows电脑
type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

// Printer 抽象实施部分，打印机
type Printer interface {
	PrintFile()
}

// Epson 精确抽象，epson打印机
type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

// Hp 精确抽象，hp打印机
type Hp struct {
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

func main() {

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
