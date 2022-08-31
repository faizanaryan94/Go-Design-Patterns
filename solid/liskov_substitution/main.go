package main

import "fmt"

// Liskov Substitution Principle
// It states that if there is an api that works perfectly with
// base class then it should also work perfectly with derived class

//Simply put, the Liskov Substitution Principle (LSP) states that
//objects of a superclass should be replaceable with objects of its subclasses without breaking the application.

// In Our scenario Rectangle is parent and Square is child and UseIt method for Rectangle should also work
// for Square but because we use setheight and setwidth method and we set both height and widht in the same method
// to enforce it so we break it

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (rect *Rectangle) GetWidth() int {
	return rect.width
}

func (rect *Rectangle) SetWidth(width int) {
	rect.width = width
}

func (rect *Rectangle) GetHeight() int {
	return rect.height
}

func (rect *Rectangle) SetHeight(height int) {
	rect.height = height
}

// Again Explanation
// In Our scenario Rectangle is parent and Square is child and UseIt method for Rectangle should also work
// for Square but because we use setheight and setwidth method and we set both height and widht in the same method
// to enforce it so we break it
// {BELOW EXAMPLE WILL DEMONSTRATE HOW}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (sq *Square) SetWidth(width int) {
	sq.width = width
	sq.height = width
}

func (sq *Square) SetHeight(height int) {
	sq.height = height
	sq.width = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetHeight() * sized.GetWidth()
	fmt.Print("Expected area of ", expectedArea, ", but got ", actualArea, ".\n")
}

func main() {
	rect := Rectangle{2, 3}
	UseIt(&rect)

	fmt.Print("Now we are doing it for Square \n")
	sq := NewSquare(5)
	UseIt(sq)
}
