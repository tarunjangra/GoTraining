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

type Color byte //Use Color as the alias of the byte

type Box struct { // Define a struct Box which has fields height, width, lenght and color
	Width, Height, Depth float64
	Color Color
}

type BoxList []Box // Define a struct BoxList which has box as its field

func (b Box) Volume() float64 { // Volume() uses box as its receiver and returns volume of Box
	return b.Width * b.Height * b.Depth
}

func (b *Box) SetColor(c Color) { //SetColor(c Color) changes Box's color.
	b.Color = c
}

func (b1 BoxList) BiggestsColor() Color { //BiggestColor() returns the color which has the biggest volume.
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

func (b1 BoxList) PaintItBack() { //PaintBlack() sets color for all Box in BoxList to black.
	for i, _ := range b1 {
		b1[i].SetColor(BLACK)
	}
}

func (c Color) String() string { //Use Color as its receiver, returns the string format of color name.
	strings := []string {"WHITE","BLACK","BLUE","RED","YELLOW"}
	return strings[c]
}
