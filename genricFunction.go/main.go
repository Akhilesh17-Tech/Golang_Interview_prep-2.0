package main

import "fmt"

func main() {
	number1 := []int{1, 2, 3, 4, 5}
	number2 := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}

	number3 := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	number4 := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	number5 := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9}

	fmt.Println("int = ", genricSum(number1))
	fmt.Println("int32 = ", genricSum(number2))
	fmt.Println("int64 = ", genricSum(number3))
	fmt.Println("float32 = ", genricSum(number4))
	fmt.Println("float64 = ", genricSum(number5))

}

type Number interface {
	int | int32 | int64 | float32 | float64
}

func genricSum[T Number](numbers []T) T {
	var result T
	for _, number := range numbers {
		result += number
	}
	return result
}