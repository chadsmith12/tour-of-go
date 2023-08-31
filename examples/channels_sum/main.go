package main

import "fmt"

func sum(s []int, channel chan int) {
    sum := 0

    for _, value := range s {
        sum += value
    }

    channel <- sum // send the sum to the channel
}

func main() {
    slice := []int { 7, 2, 8, -9, 4, 0 }
    channel := make(chan int)
    
    go sum(slice[:len(slice) / 2], channel)
    go sum(slice[len(slice) / 2:], channel)

    firstPart, secondPart := <-channel, <-channel // receive from the channel

    fmt.Println(firstPart, secondPart, firstPart + secondPart)
}
