package package2
import (
	"github.com/tarunjangra/GoTraining/07_random/package3"
	"fmt"
)

const LastName string = "Jangra"

func Add1(a *int) int {
	*a = *a+1
	return *a
}
func init() {
	fmt.Println("package2 : %s", package3.Name)
}