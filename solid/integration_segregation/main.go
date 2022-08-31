package main

import "fmt"

//INTERFACE SEGREGATION PRINCIPLE

// this principle states that try to break interface into parts that people will use

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
func main() {
	fmt.Printf("Integration Segregation Principle")

}
