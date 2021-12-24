package main

import (
	"fmt"
)

/*
Goroutine is a go version of thread which is faster (due to its low latency of its build in channels).
*/

func main() {
	c := make(chan int) // Create a channel to pass ints
	for i := 0; i < 5; i++ {
		go cookingGopher(i, c) // Start a goroutine
	}

	for i := 0; i < 5; i++ {
		gopherID := <-c // Receive a value from a channel
		fmt.Println("gopher", gopherID, "finished the dish")
	} // All goroutines are finished at this point
}

/* Notice the channel as an argument */
func cookingGopher(id int, c chan int) {
	fmt.Println("gopher", id, "started cooking")
	c <- id // Send a value back to main
}

/*
We cannot predict how the code will run. Goroutine will wait for available
channel and execute the code by sending its id.
*/
