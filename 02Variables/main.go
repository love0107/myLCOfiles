package main

import "fmt"

const LoginToken string = "ye mera elaka h" // public

func main() {
	var username string = "Lovlesh"
	fmt.Println(username)
	fmt.Printf("variavles is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variavles is of type : %T \n", isLoggedIn)

	var small_int uint8 = 255
	fmt.Println(small_int)
	fmt.Printf("variavles is of type : %T \n", small_int)

	var small_float float32 = 255.34252245234253554
	fmt.Println(small_float)
	fmt.Printf("variavles is of type : %T \n", small_float)

	//default  value and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variavles is of type : %T \n", anotherVariable)

	// implicit type
	var wesite = "hi there!"
	fmt.Println(wesite)

	// no var style

	numberOfUser := 30000
	fmt.Println(numberOfUser)
	fmt.Println(LoginToken)
}
