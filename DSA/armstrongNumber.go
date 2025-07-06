package main

import (
	"fmt"
	"math"
)

// armstrong number where the power of each digit with length of number and there sum is equal the actual number
func main() {
	n := 121 
	armstrong := isArmstrong(n)
	if armstrong {
		fmt.Println("Number is armstrong =", n)
	} else {
		fmt.Println("Number is not armstrong =", n)
	}
}

func isArmstrong(n int) bool {
	temp := n
	sum := 0
	totalDigit := countDigit(n)
	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(totalDigit)))
		temp /= 10
	}
	fmt.Println(n, "==", sum)
	return sum == n

}

func countDigit(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}
