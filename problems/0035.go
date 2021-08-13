package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 2, 3, 4, 5}, 4))
}

func searchInsert(nums []int, target int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		if target <= nums[i] {
			return i
		}
	}
	return l
}
