package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Height*2 + rectangle.Width*2)
}

func Area(rectangle Rectangle) float64 {
	return (rectangle.Height * rectangle.Width)
}
