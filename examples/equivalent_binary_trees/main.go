package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// This will walk a tree recursively, inorder, and put the values in the channel
// When it is done walking then it will close the channel to signal that nothing else is coming.
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

// This will take in two trees and walk them
// If it finds a value in the two trees that don't belong then it will return false to signify they are not the Same
// if we haven't exited yet when a channel has closed, then we know that they are the same, so return true.
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
