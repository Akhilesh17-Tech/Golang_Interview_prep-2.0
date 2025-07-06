package main

import (
	"fmt"
	"sort"
)

func permute(nums []int) [][]int {
	var result [][]int
	permuteHelper(nums, 0, len(nums)-1, &result)
	return result
}

func permuteHelper(nums []int, left, right int, result *[][]int) {
	if left == right {
		*result = append(*result, append([]int{}, nums...))
	} else {
		for i := left; i <= right; i++ {
			nums[left], nums[i] = nums[i], nums[left]
			permuteHelper(nums, left+1, right, result)
			nums[left], nums[i] = nums[i], nums[left]
		}
	}
}

func main() {
	num := 123
	nums := []int{}
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}

	sort.Ints(nums)

	fmt.Println(nums)

	permutations := permute(nums)

	// printing all the permutations
	for _, permutation := range permutations {
		for _, digit := range permutation {
			fmt.Print(digit)
		}
		fmt.Println()
	}
}
