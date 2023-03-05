package main

import "fmt"

func main() {
	fmt.Println("this is my pointer")
	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)
	mynumber := 23

	var ptr = &mynumber
	// reference we use & sign
	fmt.Println("Value of actual pointer is  ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)
	*ptr = *ptr * 2
	fmt.Println("new value of pointer is: ", *ptr)

}
