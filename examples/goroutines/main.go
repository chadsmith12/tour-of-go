package main

import "fmt"
import "time"

func say(str string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(str)
    }
}

func main()  {
    go say("world")
    say("hello")
}
