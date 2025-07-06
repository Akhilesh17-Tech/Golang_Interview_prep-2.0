package main

import (
    "fmt"
    "time"
)


// Summary
// Select: Designed for handling multiple channel operations concurrently. Blocks until one channel operation can proceed.
// Switch: Used for conditional branching based on the value of an expression. Evaluates cases sequentially without blocking.

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "from ch1"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "from ch2"
    }()

    select {
    case msg := <-ch1:
        fmt.Println(msg)
    case msg := <-ch2:
        fmt.Println(msg)
    default:
        fmt.Println("No messages received")
    }
}



// package main

// import (
//     "fmt"
// )

// func main() {
//     x := 2

//     switch x {
//     case 1:
//         fmt.Println("x is 1")
//     case 2:
//         fmt.Println("x is 2")
//     case 3:
//         fmt.Println("x is 3")
//     default:
//         fmt.Println("x is not 1, 2, or 3")
//     }
// }
