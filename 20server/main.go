package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to the video")
	// PerformGetRequest()
	// PerformPOSTRequest()
	PerformPOSTFromRequest()
}

// first lettor of function decide weather a function is public or not
func PerformGetRequest() {
	const myurl = "http://localhost:1111"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("content length:", response.ContentLength)
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	// fmt.Println(string(content))
	fmt.Println("Byte count is", byteCount)
	fmt.Println(responseString.String())
}
func PerformPOSTRequest() {
	const myurl = "http://localhost:1111/post"
	// fake json payload
	requestBody := strings.NewReader(`
	{
		"courseName":"lets go with golang",
		"price":0,
		"Plateform":"golearnonline.in"
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPOSTFromRequest() {
	const myurl = "http://localhost:1111/postfrom"

	//form data
	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "chadhuray")
	data.Add("email", "hitesh@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)

	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
