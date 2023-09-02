package main

import (
    "fmt"
    "time"
)

func fibonacci(fibChannel, quitChannel chan int) {
    x, y := 0, 1
    
    for {
        select {
        case fibChannel <- x:
            x, y = y, x+y
        case <-quitChannel:
            fmt.Println("Quiting...")
            return
        }
    }
}

func streamFibonacci(fibChannel, quitChannel chan int){
    for i := 0; i < 10; i++ {
        fmt.Println(<-fibChannel)
    }

    quitChannel <- 0
}

func main() {
    //fibChannel := make(chan int)
    //quitChannel := make(chan int)

    //go streamFibonacci(fibChannel, quitChannel)

    //fibonacci(fibChannel, quitChannel)

    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)

    for {
        select {
        case <- tick:
            fmt.Println("tick")
        case <- boom:
            fmt.Println("BOOM!")
            return
        default:
            fmt.Print(".")
            time.Sleep(50 * time.Millisecond)
        }
    }
}

