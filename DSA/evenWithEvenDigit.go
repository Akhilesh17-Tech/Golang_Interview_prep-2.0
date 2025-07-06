package main

import "fmt"

func main() {
	evenNumbers()
}

func evenNumbers() {
	for i := 1; i <= 1000; i++ {
		temp := i
		num := true
		for temp > 0 {
			l := temp % 10
			if l%2 != 0 {
				num = false
				break
			}
			temp = temp / 10

		}
		if num {
			fmt.Println(i)
		}

	}

}

// 1 3 5 7 9 11