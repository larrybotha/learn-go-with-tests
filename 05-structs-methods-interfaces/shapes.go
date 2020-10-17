package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * .2
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
	return 0
}

//func Perimeter(width, height float64) float64 {
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//func Area(width, height float64) float64 {
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
