package main

import "fmt"

func main() {
	fmt.Println("this is my array")
	var myarray [4]string
	myarray[0] = "apple"
	myarray[1] = "mango"
	// myarray[2] = "peach"
	myarray[3] = "orage"
	fmt.Println("Fruit list is : ", myarray)
	fmt.Println("Fruit length is : ", len(myarray))

	var veglist = [5]string{"patato", "beans", "mushroom"}
	fmt.Println("veg list is : ", len(veglist))
}
