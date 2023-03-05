package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("this is a slices")

	var fruitList = []string{"mango", "banana", "peach"}
	fmt.Printf("Type of fruits list is %T\n", fruitList)
	fruitList = append(fruitList, "aalu", "muli", "began")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:])
	// it will start 1 to end
	// mysclies[start: end]
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 345
	highScore[2] = 464
	highScore[3] = 897
	// highScore[4] = 997
	fmt.Println(highScore)

	// when we append then it will create new memory for provided int
	highScore = append(highScore, 999, 666, 777)
	// for sorting int in increaing
	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)

	// for knowing weather slice is sorted or not
	fmt.Println(sort.IntsAreSorted(highScore))

	// remove the value from the slice based on the index

	var couses = []string{"reactjs", "javaScript", "swift", "ruby"}
	fmt.Println(couses)
	var index int = 2
	couses = append(couses[:index], couses[index+1:]...)
	fmt.Println(couses)
}
