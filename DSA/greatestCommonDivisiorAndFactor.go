package main

import "fmt"

func main() {
	a, b := 3,5
	fmt.Println(greatedCommonDivisorAndFactor(a, b))
}

func greatedCommonDivisorAndFactor(a, b int) (int, int) {
	min := 0
	if a < b {
		min = a
	} else {
		min = b
	}

	greatestDivisior := 0

	for i := 1; i < min; i++ {
		if a%i == 0 && b%i == 0 {
			greatestDivisior = i
		}

	}

	//hcf =a*b/gcf(a,b)

	hcf := hcp(a, b, greatestDivisior)

	return greatestDivisior, hcf

}

func hcp(a, b, gcd int) int {
	return a * b / gcd
}
