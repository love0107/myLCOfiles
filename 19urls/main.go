package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gfbj456"

func main() {
	fmt.Println("welcome to the hadling the url in golagn")
	fmt.Println(myurl)
	// parsing url
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	qparams := result.Query()
	fmt.Printf("The type of query arms are: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("parms is ", val)
	}
	// pass url as reference
	partsofurl := &url.URL{
		Scheme:   "https",
		Host:     "loc.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}
	// you can also make string like this
	anotherURL := partsofurl.String()
	fmt.Println(anotherURL)
}
