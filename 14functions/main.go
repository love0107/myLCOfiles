package main

import "fmt"

func main() {
	fmt.Println("Functuion to function in go lang ")
	greeter()
	greaterTwo()

	result := adder(3, 5)
	fmt.Println("Result is : ", result)

	proResult := proAdder(2, 3, 4, 5, 7)
	fmt.Println(proResult)
}
func greeter() {
	fmt.Println("Namste from golang")
}
func greaterTwo() {
	fmt.Println("Another method")
}

// you can not declear the fuction in another function

func adder(a int, b int) int {
	return a + b
}

// func function_name ( arg1 type_of_arg) return_type{

//}
func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

// if there is a two vaues need to be return then you shloud rap in the same bracjet

// func function_name ( arg1 type_of_arg) (return_type1,return_type2) {

//}
