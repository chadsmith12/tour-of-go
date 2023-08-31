package main

import "fmt"

// streams the fibonacci sequence to the channel up to number.
// Example: Setting number to 10 will calculate up to the 10th fibonacci number.
func streamFibonacci(channel chan int) {
	x, y := 0, 1
	channelCapacity := cap(channel)
	for i := 0; i < channelCapacity; i++ {
		channel <- x
		x, y = y, x+y
	}

	close(channel)
}

func main() {
	fibonacciChannel := make(chan int, 10)
	fibSequence := make([]int, 0, 10) 
	go streamFibonacci(fibonacciChannel)
	
	// as results from the channel come in append them into the slice
	for i := range fibonacciChannel {
		fibSequence = append(fibSequence, i)
	}

	fmt.Println(fibSequence)
}
