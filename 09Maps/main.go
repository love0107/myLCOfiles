package main

import "fmt"

func main() {
	fmt.Println("Maps in go lang")
	language := make(map[string]string)
	language["js"] = "java script"
	language["rb"] = "ruby"
	language["cpp"] = "c++"
	fmt.Println(language)
	// if you want to grab only one value then this syntx woek fine
	fmt.Println("js shorts for:", language["js"])

	// for deleting a value from the map
	delete(language, "rb")
	fmt.Println("List after deleting some items from the map ", language)

	// loops are intersteing in golang
	for key, value := range language {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
