package main

import "fmt"

func main() {
	fmt.Println("Struct is golang")

	// no inheritance in golang , no super or parent
	lovlesh := User{"Lovlesh", "lovlesh0107@gmail.com", true, 23}
	fmt.Printf("lovesh details are %+v\n:", lovlesh)

	// for one value you can do the following --
	fmt.Printf("lovesh email is %v:", lovlesh.Email)
	lovlesh.GetStatus()

	lovlesh.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// this is method
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@.dev"
	fmt.Println("email of this user is: ", u.Email)
}

// when you pass a function and method then you pass as copy of that
