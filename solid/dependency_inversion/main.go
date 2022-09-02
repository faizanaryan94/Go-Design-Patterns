package main

// Dependency Inversion Principle

/* It states that High Level Modules should not depend on Low Level Modules and they both
should depend on abstractions*/

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	Name string
}

func main() {

}
