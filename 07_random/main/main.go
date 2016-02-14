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

	//you can not change string value by index. But you
	// can get string character by index

	s = "c" + s[:1]
	fmt.Printf("%s\n",s)

	//string multiple linkes
	s = `Hello!
	Terry`
	fmt.Printf("%s\n",s)
}
