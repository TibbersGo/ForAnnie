package main

import "fmt"

func main() {
	res := isPerfect(14)
	fmt.Println(res)
}

func isPerfect(num int) bool {
	left := 0
	right := num
	target := num
	for left <= right {
		fmt.Println("left: ", left)
		fmt.Println("right: ", right)
		mid := left + (right-left)/2
		if mid*mid == target {
			return true
		} else if mid*mid > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func searchRange(nums []int, target int) []int {
	leftBorder := getLeftBorder(nums, target)
	rightBorder := getRightBorder(nums, target)
	if leftBorder == -2 && rightBorder == -2 {
		return []int{-1, -1}
	} else if rightBorder-leftBorder >= 0 {
		return []int{leftBorder + 1, rightBorder - 1}
	} else {
		return []int{-1, -1}
	}
}

func getLeftBorder(nums []int, target int) int {
	left := 0
	right := len(nums)
	leftBorder := -2
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
			leftBorder = right
		} else {
			left = mid + 1
		}
	}
	fmt.Println("left: ", leftBorder)
	return leftBorder
}

func getRightBorder(nums []int, target int) int {
	left := 0
	right := len(nums)
	rightBorder := -2
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
			rightBorder = left
		} else {
			right = mid - 1
		}
	}
	fmt.Println("right: ", rightBorder)
	return rightBorder
}
