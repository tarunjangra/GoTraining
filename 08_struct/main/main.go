package main
import (
	. "github.com/tarunjangra/GoTraining/08_struct/strctures"
	"fmt"
)

func main() {
	var tom Person
	tom.Age, tom.Name = 33,"tarun"
	fmt.Printf("Name: %s, Age: %d", tom.Name, tom.Age)

	r1 := Rectangle{12, 2}
	//r2 := Rectangle{9, 4}
	c1 := Circle{10}
	//c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.Area())
	fmt.Println("Area of c1 is: ", c1.Area())
}
