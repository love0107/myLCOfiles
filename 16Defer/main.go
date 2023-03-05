package main

import "fmt"

func main() {
	defer fmt.Println("hellow world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("Hello")
	mydefer()
}

// defer worok as last in first out
// if you the exuction line will put at the end of the curly bracket

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
