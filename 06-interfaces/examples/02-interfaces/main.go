package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Stringer interface {
	String() string
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("%s by %s (%d pages)", b.Title, b.Author, b.Pages)
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	fmt.Println("=== Shape Interface ===")
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
	}

	for i, s := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		printShapeInfo(s)
	}

	fmt.Println("\n=== Stringer Interface ===")
	book := Book{"The Go Programming Language", "Donovan & Kernighan", 380}
	fmt.Println(book)
}
