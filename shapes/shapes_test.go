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
	rectangle := Rectangle{10.0, 10.0}
	expected := 100.0
	received := Area(rectangle)

	checkResults(t, expected, received)
}
