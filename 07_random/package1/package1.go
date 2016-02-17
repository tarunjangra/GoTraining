package package1
import (
	"github.com/tarunjangra/GoTraining/07_random/package2"
	"fmt"
)

const Variable string = "Variable valie package 1"

func Add1(a *int) int {
	*a = *a+1
	return *a
}

func init () {
	fmt.Println("Package1 : %s", package2.LastName)
}