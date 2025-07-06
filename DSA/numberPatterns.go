package main

import "fmt"

func main() {
	n := 123
	sum := 0

	for n != 0 {
		last := n % 10
		sum += last
		n /= 10
	}
	fmt.Println(sum)
	strongNum()
	perfectNumber()
	automorphicNumber()
	harshadNumber()
}

func strongNum() {
	n := 145
	temp := n
	sum := 0
	for temp != 0 {
		last := temp % 10
		sum += factorial(last)
		temp /= 10

	}
	if sum == n {
		fmt.Println(n, "is a strong number")
	}
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func perfectNumber() {
	n := 28
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}

	}
	if sum == n {
		fmt.Println(n, "is a perfect number")
	}
}

func automorphicNumber() {
	n := 376
	sq := n * n
	fmt.Println(sq, "=", n)
	if sq%10 == n%10 {
		fmt.Println(n, "is an automorphic number")
	}

}

func harshadNumber() {
	n := 153
	sum := 0
	temp := n

	for temp != 0 {
		last := temp % 10
		sum += last
		temp /= 10
	}
	if n%sum == 0 {
		fmt.Println(n, "is a harshad number")
	} else {
		fmt.Println(n, "is not a harshad number")
	}
}
