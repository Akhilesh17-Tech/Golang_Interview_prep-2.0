package main

import (
	"fmt"
	"sort"
)

func main() {
	denoms := []int{1, 5, 10, 20, 50, 100, 200, 500, 1000, 2000}
	target := 138
	// fmt.Println(minimumCoins(denoms, target))

	totalCoins, count, success := minimumCoins(denoms, target)

	if success {
		fmt.Printf("Total number of coins/notes required: %d\n", totalCoins)
		fmt.Println("Count of each denomination:")
		for i, cnt := range count {
			if cnt > 0 {
				fmt.Printf("%d x %d\n", denoms[i], cnt)
			}
		}
	} else {
		fmt.Printf("The target amount %d cannot be reached with the given denominations.\n", target)
	}
}

func minimumCoins(denoms []int, target int) (int, []int, bool) {

	if target == 0 || target < 0 {
		return 0, []int{}, false
	}
	// To store the count of each denomination
	count := make([]int, len(denoms))

	// decending order of denomination
	sort.Slice(denoms, func(i, j int) bool {
		return denoms[i] > denoms[j]
	})

	totalCoins := 0
	for i, denom := range denoms {
		if target == 0 {
			break
		}
		// Find the number of coins/notes of this denomination
		count[i] = target / denom
		// Update the total number of coins/notes
		totalCoins += count[i]
		// Update the remaining target amount
		target -= count[i] * denom
	}

	// If target is not zero, it means we couldn't reach the target with given denominations
	if target > 0 {
		return -1, count, false
	}

	return totalCoins, count, true
}
