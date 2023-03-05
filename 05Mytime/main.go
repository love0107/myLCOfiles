package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time stydy of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2019, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

// go env
// GOOS="darwin"-> example
// memoery management
// 1. new()
// allocatre memory but no init
// you will get a memory address
// zeroed storage
// 2. make()
// allocate memory and init
// you will get a memory address
// non-zeroed-storage
// GC happens atomatically
// Out oid scope or nil
