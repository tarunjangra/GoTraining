package main
import "fmt"
import "errors"


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

	//Error type
	err := errors.New("This is error trying to get it work")
	if err != nil {
		fmt.Print(err);
	}

	// Multiple varialbe definition in Groups

	const (
		ii = 100
		pi = 3.1415
	)

	// Array declarations

	var arr [10]int
	arr[0] = 42
	fmt.Printf("The first element is %d\n",arr[0])

	//shortend array definition

	arr1 := [3]int{1,2,3}
	fmt.Printf("The second element in arr1 is %d\n",arr1[2])

	// autoassignement of lenght of an array

	arr2 := [...][2]int{{23,4},{54,13},{32,88}}

	fmt.Printf("Ths array of array %d\n",arr2[1][1])

	// map are like dictionary in python and in php these are
	// associative arrays

	/* - map is disorderly. Every time you print map you
	     will get different result.
	   - map doestn't have a fixed length. It's a reference type
	     just like slice.
	   - len works for map also. it returns how many key S that
	     map has.
	   - it's quite easy to change the value thorugh map. Simply use
	     numbers["one"] = 11 to change the vlaue of key one to 11

	*/
	mp := make(map[string]string)
	mp["one"] = "Yogesh"
	fmt.Printf("This value is from map %v\n",mp["one"])

	rating := map[string]float32 {"C": 5, "Go":4.5, "Python": 4.5, "C++": 2}
	// map has two return values. For the second return value, it
	// the key doestn't exist, 'ok' returns false. It returns true
	// otherwise
	rating["C#"] = 3.4
	csharpRating, ok := rating["C#"]

	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C#")

	// maps are pointer. And if two map variable have been
	// assigned to same map structure. they will not be copy
	// of each other, they will be pointing to same data structure
	// in the memory
	mp1 := mp
	fmt.Printf("This is the value of older map %s\n",mp1["one"])

	//Make Vs New

	var age *[]int = new([]int)
	fmt.Println(age)
	fmt.Println(*age)

	var bday []int = make([]int, 1, 2)
	fmt.Println(bday)

	var rate *int = new(int)
	fmt.Println(rate)
	fmt.Println(*rate)

	var class *string = new(string)
	fmt.Println(class)
	fmt.Println(*class)

	var member *bool = new(bool)
	fmt.Println(member)
	fmt.Println(*member)

}
