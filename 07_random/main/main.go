package main
import "fmt"


func main() {
	var i uint8
	_,i = 5,4
	fmt.Printf("%v\t%T\n",1<<i,i)
	var c complex64 = 5+5i
	fmt.Printf("%v\n",c)

	// Change a string by index
	s := "hello"
	d := []byte(s)
	d[0] = 't'
	s = string(d)
	fmt.Printf("%s\n",s)
}
