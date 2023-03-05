package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/love0107/mongoapi/router"
)

func main() {
	fmt.Println("Mongo db api")
	r := router.Router()
	fmt.Println("server is getting strted ....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listen at port 4000...")

}

/// import go tool

// this is where mongo db hel;pers files
