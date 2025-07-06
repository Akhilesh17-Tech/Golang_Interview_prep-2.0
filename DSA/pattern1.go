package main

import "fmt"

func main() {
    count := 0
    for i := 5; i > 0; i-- {
        for j := i; j > 0; j-- {
            fmt.Print(j, " ")
        }
        count++
        fmt.Println()
    }
    count = 0
    for i := 5; i > 0; i-- {
        for k := 0; k < count; k++ {
            fmt.Print("  ")
        }
        for j := i; j > 0; j-- {
            fmt.Print(j, " ")
        }
        count++
        fmt.Println()
    }
}