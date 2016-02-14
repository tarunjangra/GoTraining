package main

import (
	"fmt"
	"github.com/tarunjangra/GoTraining/03_package/Employee"
)

var e string = "cowboy"

func main() {
	fmt.Println(Employee.Name)
	fmt.Println(Employee.Address)


	// shorthand variables

	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \t  %T \n",a,a)
	fmt.Printf("%v \t  %T  \n",b, b)
	fmt.Printf("%v \t  %T  \n",c, c)
	fmt.Printf("%v \t  %T  \n",d, d)
	fmt.Printf("%v \t  %T  \n",e, e)

	// zero value (default value) of variables.

	var e string
	var f int
	var g float64
	var h bool

	fmt.Printf("%v\n",e)
	fmt.Printf("%v\n",f)
	fmt.Printf("%v\n",g)
	fmt.Printf("%v\n",h)
}
