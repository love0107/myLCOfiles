package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"string"}
var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

func main() {
	// fmt.Println("Helow there !")
	// go greeter("hello")
	// greeter("wordld")
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://fb.com",
		"https://google.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("strings")
}

// func greeter(s string) {

// 	for i := 0; i <= 5; i++ {
// 		time.Sleep(3 * time.Microsecond)
// 		fmt.Println(s)
// 	}

// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoints")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}
}
