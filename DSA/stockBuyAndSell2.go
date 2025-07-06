// stock buy and sell to get the maximum profit as many time you can buy and sell but buy only can be done when you sell vice versa 
package main

import "fmt"

func main() {
	prices := []int{7, 8, 3, 9, 12, 23, 2, 10}
	maxProfit := maxProfit(prices)
	fmt.Println("Max Profit:", maxProfit)
}

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
