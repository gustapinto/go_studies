package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var result []int

	n := len(nums)

	if n == 2 {
		return []int{0, 1}
	}

out:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			if nums[i]+nums[j] == target {
				result = append(result, i, j)

				break out
			}
		}
	}

	return result
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5, 6}, 10))
}
