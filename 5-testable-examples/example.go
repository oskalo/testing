package examples

import "fmt"

func Reverse(arg string) string {
	return arg
}

type User struct {
	Name string
}

func (u User) Hi() {
	fmt.Printf("Hi, my name is %s!", u.Name)
}
