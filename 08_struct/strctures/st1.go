package strctures
import (
	"math"
	"fmt"
)

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

// interface example

type Human struct {
	Name string
	Age int
	Phone string
}

type Student struct {
	Human
	School string
	Loan float32
}

type Employee struct {
	Human
	Company string
	Money float32
}


type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderLyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.Name,h.Phone)
}

func (h Human) Sing(lyrics string){
	fmt.Println("La la,la la la, la la ....", lyrics)
}

func (h Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, is work at %s. Call me on %s\n", e.Name,
		e.Company, e.Phone)
}

func (s Student) BorrowMoney(amount float32){
	s.Loan += amount
}

func (e Employee) SpendSalary(amount float32){
	e.Money -= amount
}

//if a function uses an empty interface as its argument type, it can accept
// any type; if a function uses empty as it's return value type, it can
// return any type

// to define the slice with can have all type of values shoule
// be done with the help of empty interface.

}


