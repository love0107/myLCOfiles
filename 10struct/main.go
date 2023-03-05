package main

import "fmt"

func main() {
	fmt.Println("Struct is golang")

	// no inheritance in golang , no super or parent
	lovlesh := User{"Lovlesh", "lovlesh0107@gmail.com", true, 23}
	fmt.Printf("lovesh details are %+v\n:", lovlesh)

	// for one value you can do the following --
	fmt.Printf("lovesh email is %v:", lovlesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
