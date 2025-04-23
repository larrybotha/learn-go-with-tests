package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "rectangle",
			shape: Rectangle{Width: 10.0, Height: 10.0},
			want:  40.0,
		},
		{
			name:  "circle",
			shape: Circle{Radius: 10.0},
			want:  62.83185307179586,
		},
	}

	for _, x := range perimeterTests {
		t.Run(x.name, func(t *testing.T) {
			got := x.shape.Perimeter()

			if got != x.want {
				t.Errorf("%#v, got %g, want %g", x, got, x.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "rectangle",
			shape: Rectangle{Width: 10.0, Height: 10.0},
			want:  100.0,
		},
		{
			name:  "circle",
			shape: Circle{Radius: 10.0},
			want:  314.1592653589793,
		},
		{"triangle", Triangle{12, 6}, 36.0},
	}

	for _, x := range areaTests {
		t.Run(x.name, func(t *testing.T) {
			got := x.shape.Area()

			if got != x.want {
				t.Errorf("%#v, got %g, want %g", x, got, x.want)
			}
		})
	}
}
