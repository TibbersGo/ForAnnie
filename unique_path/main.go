package main

import "fmt"

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。
机器人试图达到网格的右下角（在下图中标记为“Finish”）。
问总共有多少条不同的路径？
 */

func uniquePath(x, y int) int {
	f := make([][]int, x)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if f[i] == nil {
				f[i] = make([]int, y)
			}
			f[i][j] = 1
		}
	}

	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			f[i][j] = f[i][j-1] + f[i-1][j]
		}
	}
	return f[x-1][y-1]
}

func main() {
	maxPath := uniquePath(5,5)
	fmt.Println(maxPath)
}