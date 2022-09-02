package main

import "fmt"

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

type Info struct {
	From         *Person
	Relationship Relationship
	To           *Person
}

// low-level module
type Relationships struct {
	Relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.Relations = append(r.Relations, Info{parent, Parent, child})
	r.Relations = append(r.Relations, Info{child, Child, parent})
}

// high-level module
//here we will be breaking the DIP by making high level module depend on low level module

type Research struct {
	//Break DIP
	//Deprecated : ...
	// relationships Relationships

	//solution property
	browser RelationshipBrowser
}

// func (r *Research) Investigate() {
// 	relations := r.relationships.Relations

// 	for _, rel := range relations {
// 		if rel.From.Name == "John" && rel.Relationship == Parent {
// 			fmt.Println("John has a child called : " + rel.To.Name)
// 		}
// 	}
// }

//solution
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, person := range r.Relations {
		if person.Relationship == Parent && person.From.Name == name {
			result = append(result, r.Relations[i].To)
		}
	}

	return result
}

func (r *Research) InvestigateSolution() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called : " + p.Name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := Research{&relationships}
	research.InvestigateSolution()
}
