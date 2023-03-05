package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	// fmt.Println("hellow")
	// mynum1 := 2
	// myNum2 := 4.08
	// fmt.Println("Then sum is:", mynum1+int(myNum2))

	// maths random
	// chage the seed at every time
	// rand.Seed(time.Now().UnixNano())

	// fmt.Println(rand.Intn(5) + 1)

	// random from crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
