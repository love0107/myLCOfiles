package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}
	fmt.Println(days)
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }
	// second option
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// third option
	// for index, days := range days {
	// 	fmt.Printf("index is%v and value is %v\n", index, days)
	// }

	// forth option ---
	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			rougueValue++
			continue // skip test case 5
			// break after this statement this loop breaks
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
		// ++rougueValue this wil not work go do not support it
	}
lco:
	fmt.Println("Jumping at LearnCodeline.in")

}
