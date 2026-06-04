package shapes

import (
	"fmt"
	"math"
)

// Shape is an interface for geometric shapes that can calculate their area.
type Shape interface {
	Area() float64
}

// Rectangle represents a rectangle with width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represents a circle with a radius.
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle.
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// PrintArea prints the area of any given Shape.
func PrintArea(s Shape) string {
	area := s.Area()
	result := fmt.Sprintf("Area: %.2f", area)
	fmt.Println(result)
	return result
}
