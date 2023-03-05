package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello  of in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	router := http.ListenAndServe(":4000", r)
	log.Fatal(router)

}
func greeter() {
	fmt.Println("hey there mod users")
}
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to golang</h1>"))
}

//go mod verify ----> verify all modules
// go mod why github.com/gorilla/mux  tell depens on
//go mod edit -go 1.17 --> chagen the version of mod
// go mod edit -module 1.6 change the versinon of the
// go mod vendor ----> give you a folder called vendor
/// go run -mod=vendor main.go ---> this will lookfirst in vendor folder in first and then it will look in web afterwards
