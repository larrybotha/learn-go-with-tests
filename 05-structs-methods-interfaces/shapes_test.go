package shapes

import (
	"testing"
)

func assertWithFVerb(t *testing.T, expected, actual float64) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected %.2f to be %.2f", actual, expected)
	}
}

func assertWithGVerb(t *testing.T, expected, actual float64) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected %g to be %g", actual, expected)
	}
}

func TestPerimeter(t *testing.T) {
	t.Run("calculates rectangle perimeter", func(t *testing.T) {
		//actual := Perimeter(10.0, 10.0)
		rectangle := Rectangle{10.0, 10.0}
		actual := rectangle.Perimeter()
		expected := 40.0

		assertWithFVerb(t, expected, actual)
	})

	//t.Run("calculates circle perimeter", func(t *testing.T) {
	////actual := Perimeter(10.0, 10.0)
	//circle := Circle{10.0}
	//actual := circle.Circumference()
	//expected := 40.0
	//
	//assertWithFVerb(t, expected, actual)
	//})
}

func TestArea(t *testing.T) {
	t.Run("calculates rectangle area", func(t *testing.T) {
		//actual := Area(10.0, 10.0)
		rectangle := Rectangle{10.0, 10.0}
		actual := rectangle.Area()
		expected := 100.0

		assertWithFVerb(t, expected, actual)
	})

	t.Run("calculates circle area", func(t *testing.T) {
		//actual := Area(10.0, 10.0)
		circle := Circle{10}
		actual := circle.Area()
		expected := 314.1592653589793

		assertWithGVerb(t, expected, actual)
	})
}
