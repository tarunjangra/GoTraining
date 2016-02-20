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

const(
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	Width, Height, Depth float64
	Color Color
}

type BoxList []Box // a slice of boxes

func (b Box) Volume() float64 {
	return b.Width * b.Height * b.Depth
}

func (b *Box) SetColor(c Color) {
	b.Color = c
}

func (b1 BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)// default value of each color would be WHITE
	for _,b := range b1 {
		if b.Volume() > v {
			v = b.Volume()
			k = b.Color
		}
	}
	return k
}

func (b1 BoxList) PaintItBack() {
	for i, _ := range b1 {
		b1[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string {"WHITE","BLACK","BLUE","RED","YELLOW"}
	return strings[c]
}
