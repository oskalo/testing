package examples

import (
	"fmt"
)

func ExampleReverse() {
	fmt.Println(Reverse("hello"))
	// Output:olleh
}

func ExampleUser_Hi() {
	u := User{"Timmy"}
	u.Hi()
	// Output:
	// Hi, my name is Timmy!
}
