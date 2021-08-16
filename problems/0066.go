package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9}))
}

func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 10 {
			continue
		}
		digits[i] = 0
		if i > 0 {
			digits[i-1]++
		} else {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
