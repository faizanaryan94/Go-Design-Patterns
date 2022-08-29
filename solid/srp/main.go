package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (journal *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d : %s", entryCount, text)
	journal.entries = append(journal.entries, entry)
	return entryCount
}

func (journal *Journal) RemoveEntry(index int) {
	//////
}

//separation of concerns

//here we are breaking single responsibility as we are doing tasks for persistence also in journal class
//whereas these can be handled by another class for persistence

// Single Responsibility Principle state one class should be responsible for one type of methods only
//in this particular scenario we could define a separate struct for persistence and use that instead of add all methods into this file

func (journal *Journal) Save(filename string) {
	//_ = ioutil.WriteFile(filename, []byte(journal.String()), 0644)
}

func (journal *Journal) Load() {

}

func (journal *Journal) LoadWeb() {

}

//SOLUTION

type Persistence struct {
	LineSeparator string
}

func (p Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.LineSeparator)), 0644)
}

func main() {
	p := Persistence{}
	j := Journal{}
	j.AddEntry("Hello")
	j.AddEntry("World")
	p.SaveToFile(&j, "journal.txt")
}
