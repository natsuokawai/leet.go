package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-1, -5, -3}))
	fmt.Println("")
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	maxSoFar, maxEndingHere := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if maxEndingHere+nums[i] > nums[i] {
			maxEndingHere = maxEndingHere + nums[i]
		} else {
			maxEndingHere = nums[i]
		}
		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}
