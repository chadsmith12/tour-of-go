package main

import (
    "fmt"
    "sync"
    "time"
)

type SafeCounter struct {
    mu sync.Mutex
    values map[string]int
}

func (c *SafeCounter) Inc(key string) {
    c.mu.Lock()
    c.values[key]++
    c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.values[key]
}

func main() {
    c := SafeCounter{values: make(map[string]int)}

    for i := 0; i < 1000; i++ {
        go c.Inc("somekey")
    }

    time.Sleep(time.Second)
    fmt.Println(c.Value("somekey"))
}
