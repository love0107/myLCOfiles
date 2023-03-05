package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wecome := "welcome to user inpiut"
	fmt.Println(wecome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our Pizza:")

	// comma ok \\ error ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of rating %T", input)
}
