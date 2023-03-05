package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name      string `json:"coursename"`
	Price     int
	Plateform string   `json:"website"`
	Password  string   `json:"-"`              // hide the password  use -
	Tags      []string `json:"tags,omitempty"` // if value is nil then it will not show
}

func main() {
	fmt.Println("welcome to json video")
	// EncodeJson()
	DecideJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"reactjs bootcamp", 299, "Learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"mern bootcamp", 199, "Learncodeonline.in", "sdf23", []string{"full-stak", "js"}},
		{"Angular bootcamp", 99, "Learncodeonline.in", "aasdf3", nil},
	}
	// package this data as json data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecideJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename":"ReactJs Bootcamp",
		"Price":299,
		"website":"LearnCodeOnline.in",
		"tags":["web-dev","js"]
	}
	`)
	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JOSN WAS NOT Valid")
	}
	// ther are some where you just want to add data to key vale
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
