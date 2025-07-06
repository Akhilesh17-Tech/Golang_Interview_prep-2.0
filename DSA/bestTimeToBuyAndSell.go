// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {

// 	prices := []int{7, 8, 3, 9, 12, 23, 2, 10}
// 	maxProfit := 0
// 	for i := 0; i < len(prices); i++ {
// 		for j := i + 1; j < len(prices); j++ {
// 			if prices[j] > prices[i] {
// 				maxProfit = Max(maxProfit, prices[j]-prices[i])
// 			}
// 		}
// 	}

// 	fmt.Println(maxProfit)

// }

package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{7, 8, 3, 9, 12, 23, 2, 10}
	maxProfit := 0
	buyIndex := 0
	sellIndex := 0

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > prices[i] {
				profit := prices[j] - prices[i]
				if profit > maxProfit {
					maxProfit = profit
					buyIndex = i
					sellIndex = j
				}
			}
		}
	}

	fmt.Println("Max Profit:", maxProfit)
	fmt.Println("Buy Index:", buyIndex)
	fmt.Println("Sell Index:", sellIndex)
	OptimalSolution(prices)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func OptimalSolution(prices []int) {
	maxProfit := 0
	minPrice := math.MaxInt64 // Set minPrice to the maximum possible int value
	start := -1
	end := -1
	st := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i] // Update minPrice if a lower price is found
			st = i
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice // Calculate the current profit and update maxProfit if it's higher
			start = st
			end = i
		}
	}

	fmt.Println("Max Profit:", maxProfit)
	fmt.Println(start, end)

}
