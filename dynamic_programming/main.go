package main

import "fmt"

/*
动态规划适用于求最大/最小值，是否可行，可行个数
前提条件是不能排序或交换
*/

/*
[
[2],
[3, 4],
[6, 5, 7],
[4, 1, 8, 3]
]
*/

//自底向上

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	tempList := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if tempList[i] == nil {
				tempList[i] = make([]int, len(triangle[i]))
			}
			tempList[i][j] = triangle[i][j]
		}
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			tempList[i][j] = min(tempList[i+1][j], tempList[i+1][j+1]) + triangle[i][j]
		}
	}
	return tempList[0][0]
}

func minimumTotal2(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	tempList := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if tempList[i] == nil {
				tempList[i] = make([]int, len(triangle[i]))
			}
			tempList[i][j] = triangle[i][j]
		}
	}
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j-1 < 0 {
				tempList[i][j] = tempList[i-1][j] + triangle[i][j]
			} else if j >= len(triangle[i-1]) {
				tempList[i][j] = tempList[i-1][j-1] + triangle[i][j]
			} else {
				tempList[i][j] = min(tempList[i-1][j-1], tempList[i-1][j]) + triangle[i][j]
			}
		}
	}
	l := len(triangle)
	res := tempList[len(triangle)-1][0]
	for i := 1; i < len(tempList[len(triangle)-1]); i++ {
		res = min(res, tempList[l-1][i])
	}
	return res
}

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	res := minimumTotal1(triangle)
	fmt.Println(res)
	fmt.Println(minimumTotal2(triangle))
}
