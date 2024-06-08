package main

import (
	"fmt"
	"math"
)

func main() {
	a := 8
	res := sqrt(a)
	fmt.Println("result: ", res)
	fmt.Println("sqrt: ", int(math.Sqrt(float64(a))))
}
func sqrt(x int) int {
	left := 0
	right := x
	// 因此采用的是左闭右闭的区间

	// 左闭右闭时，left == right 有效

	mid := x
	for left < right {
		fmt.Printf("left: %d, right: %d\n", left, right)
		mid = (left + right) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid > x { // 目标值一定是在 left, mid
			right = mid
		} else {
			if (mid + 1) > x/(mid+1) {
				return mid
			}
			left = mid + 1
		}
		fmt.Printf("mid = %d\n", mid)
	}
	return mid
}
