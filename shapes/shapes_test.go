package shapes

import (
	"testing"
)

func checkResults(t *testing.T, expected, received float64) {
	if expected != received {
		t.Errorf("expected %.2f received %.2f", expected, received)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	expected := 40.0
	received := Perimeter(rectangle)

	checkResults(t, expected, received)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		received := tt.shape.Area()
		checkResults(t, tt.want, received)
	}
}
