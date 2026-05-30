package main

import (
	"fmt"
	"math"
)

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

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
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

type Shape interface {
	Area() float64
	Perimeter() float64
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	fmt.Println("=== Rectangle Methods ===")
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle area: %.2f\n", rect.Area())
	fmt.Printf("Rectangle perimeter: %.2f\n", rect.Perimeter())

	fmt.Println("\n=== Pointer Receiver ===")
	rect.Scale(2)
	fmt.Printf("After scaling: Width=%.1f, Height=%.1f\n", rect.Width, rect.Height)

	fmt.Println("\n=== Interface Polymorphism ===")
	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 7},
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		printShapeInfo(shape)
	}
}
