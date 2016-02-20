package main
import (
	. "github.com/tarunjangra/GoTraining/08_struct/strctures"
	"fmt"
)

func main() {
	var tom Person
	tom.Age, tom.Name = 33,"tarun"
	fmt.Printf("Name: %s, Age: %d\n", tom.Name, tom.Age)

	r1 := Rectangle{12, 2}
	c1 := Circle{10}

	fmt.Println("Area of r1 is: ", r1.Area())
	fmt.Println("Area of c1 is: ", c1.Area())

	boxes := BoxList{
		Box{4,4,4,RED},
		Box{10,10,10,YELLOW},
		Box{14,41,24,BLACK},
		Box{49,46,44,BLUE},
		Box{84,74,6,WHITE},
		Box{24,34,44,YELLOW},
		Box{45,43,41,RED},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(),"cm3")
	fmt.Println("the color of hte last one is",boxes[len(boxes)-1].Color.String())
	fmt.Println("The biggest one is", boxes.BiggestsColor().String())
	boxes.PaintItBack()
	fmt.Println("The color of the second one is", boxes[1].Color.String())
	fmt.Println("Obviously the color of the biggest one is", boxes.BiggestsColor().String())
}
