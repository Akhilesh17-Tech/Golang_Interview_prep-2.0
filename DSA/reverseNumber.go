package main

import "fmt"

func main() {
	num := 123789
	sum := 0
	for num != 0 {
		last := num % 10
		sum = sum*10 + last
		num = num / 10
	}

	fmt.Println(sum)
}
