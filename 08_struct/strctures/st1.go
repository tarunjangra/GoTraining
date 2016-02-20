package strctures
import "math"

type Person struct {
	Name string
	Age int
}

func Older(p1, p2 Person) (Person, int) {
	if p1.Age>p2.Age {
		return p1, p1.Age-p2.Age
	}
	return p2, p2.Age-p2.Age
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width*r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
