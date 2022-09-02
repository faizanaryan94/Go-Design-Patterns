package main

import "fmt"

//INTERFACE SEGREGATION PRINCIPLE
// this principle states that try to break interface into parts that people will use

// Here we separate Print interface and Scan interface add them to Multifunction device and use
// that to build a multifunctional device

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (mfp *MultiFunctionPrinter) Print() {

}

func (mfp *MultiFunctionPrinter) Scan() {

}

func (mfp *MultiFunctionPrinter) Fax() {

}

type OldFashionedPrinter struct{}

func (ofp *OldFashionedPrinter) Print() {

}

// Deprecated: ...
func (ofp *OldFashionedPrinter) Scan() {
	panic("Operation not supported")
}

// Deprecated: ...
func (ofp *OldFashionedPrinter) Fax() {
	panic("Operation not supported")

}

// Solution
type Scanner interface {
	Scan(d Document)
}

type Printer interface {
	Print(d Document)
}

type Printer1 struct{}

type Scanner1 struct{}

func (pr1 *Printer1) Print(d Document) {
	panic("I Print for Printer1")
}

func (sc1 *Scanner1) Scan(d Document) {
	panic("I scan for Scanner1")
}

//Now lets assume we have a multifunctional device
// for that we create a multifunctional interface

type MultiFunctionalDevice interface {
	Printer
	Scanner
	//Fax
}

// Now multifunctional interface implements both printer and scanner interface
//Now lets say we have a multifunctional device

type MultiFunctionDevice1 struct {
	printer Printer
	scanner Scanner
}

func (mfd *MultiFunctionDevice1) Print(d Document) {
	mfd.printer.Print(d)
}

func (mfd *MultiFunctionDevice1) Scan(d Document) {
	mfd.scanner.Scan(d)
}

func main() {
	fmt.Printf("Integration Segregation Principle")
}
