package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
    WalkRecursive(t, ch)
    close(ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
    if t != nil {
        WalkRecursive(t.Left, ch)
        ch <- t.Value
        WalkRecursive(t.Right, ch)
    }
}

func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go Walk(t1, ch1)
    go Walk(t2, ch2)
    
    for {
        value1, ok1 := <- ch1
        value2, ok2 := <- ch2
        
        if ok1 != ok2 {
            return false
        }
        if value1 != value2 {
            return false
        }

        if !ok1 {
            break
        }
    }

    return true
}

func main() {
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
