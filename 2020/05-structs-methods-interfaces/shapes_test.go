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

func assertShapeArea(t *testing.T, shape Shape, expected float64) {
	t.Helper()
	actual := shape.Area()

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
		//rectangle := Rectangle{10.0, 10.0}
		rectangle := Rectangle{10.0, 10.0}
		expected := 100.0

		assertShapeArea(t, rectangle, expected)
	})

	t.Run("calculates circle area", func(t *testing.T) {
		//actual := Area(10.0, 10.0)
		circle := Circle{10}
		expected := 314.1592653589793

		assertShapeArea(t, circle, expected)
	})
}

func TestArea_Table(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, want: 100},
		{name: "circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("expected %#v area %g to be %g", tt.shape, got, tt.want)
			}
		})
	}
}
