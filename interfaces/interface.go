package interfaces

var (
	MyVal = interface{}("100")
)

type Area interface {
	GetArea() float64
}
type Perimeter interface {
	GetPerimeter() float64
}

type Shape interface {
	Area
	Perimeter
}
type InterfaceResult struct {
	AreaOfShape, PerimeterOfShape float64
}

func Calculate(s Shape) InterfaceResult {
	rs := InterfaceResult{
		AreaOfShape:      s.GetArea(),
		PerimeterOfShape: s.GetPerimeter(),
	}
	return rs
}
