package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {

	fmt.Println("this is web ")

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response of type: %T\n", res)
	defer res.Body.Close() // caller's reponsibility to close the connections

	databyes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)

	}
	content := string(databyes)
	fmt.Println(content)
}
