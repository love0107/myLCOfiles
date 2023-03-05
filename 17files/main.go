package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("this is file video")
	content := "This need to go in a file - learnCodeinline.in"
	file, err := os.Create("./mytext-file.txt")
	checkNillErr(err)
	// for creating a file
	length, err := io.WriteString(file, content)
	checkNillErr(err)
	fmt.Println("length is", length)
	defer file.Close()
	readFile("./mytext-file.txt")

}

// for creating the file
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNillErr(err)
	fmt.Println("Text data inside file is\n", databyte)

}

// reading the data from file is always happens in byte so we need to took care of that

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}

}
