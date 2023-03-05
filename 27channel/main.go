package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels")
	myChal := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// for push the value
	// myChal <- 5
	// for printing the val
	// fmt.Println(<-myChal)

	wg.Add(2)
	// receive only
	//if channel is close then ypu do not allow to close channels
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, ischannelOpen := <-myChal

		fmt.Println(ischannelOpen)
		fmt.Println(val)
		fmt.Println(<-myChal)

		wg.Done()
	}(myChal, wg)
	//send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChal <- 5
		myChal <- 6
		close(myChal)
		wg.Done()

	}(myChal, wg)
	wg.Wait()

}
