package main

import "fmt"

func main() {
	fmt.Println(mySqrt(3))
	fmt.Println(mySqrt(9))
	fmt.Println(mySqrt(17))
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	left := 1
	right := x
	for {
		mid := left + (right-left)/2
		if mid > x/mid {
			right = mid - 1
			continue
		}
		if mid+1 > x/(mid+1) {
			return mid
		}
		left = mid + 1
	}
}
