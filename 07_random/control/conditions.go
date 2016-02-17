package control
import "fmt"

type testInt func(int) bool // define a function type of variable

func Loopcontrol() {
	sum := 0
	for index:=0; index<10; index++ {
		sum += index
	}
	fmt.Printf("Sum is equal to %v\n", sum)

	person := map[string]string {"name":"tarun","address":"Hargobind Nagar"}

	for k,v := range person {
		fmt.Printf("My %s is '%s'\n",k,v)
	}
}

func Max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func Add1(a *int) int {
	*a = *a+1
	return *a
}

func Filter (slice []int, f testInt) ([]int,[]int) {
	var even,odd []int
	for _, value := range slice {
		if f(value) {
			even = append(even, value)
		} else {
			odd = append(odd, value)
		}
	}
	return even,odd
}