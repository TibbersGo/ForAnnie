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

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	x := len(obstacleGrid)
	y := len(obstacleGrid[0])
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
		if obstacleGrid[i][0] == 1 || f[i-1][0] == 0 {
			f[i][0] = 0
		}
	}

	for j:= 1; j < y; j++ {
		if obstacleGrid[0][j] == 1 || f[0][j-1] == 0 {
			f[0][j] = 0
		}
	}

	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			if obstacleGrid[i][j] == 1 {
				f[i][j] = 0
			} else {
				f[i][j] = f[i][j-1] + f[i-1][j]
			}
		}
	}

	return f[x-1][y-1]
}


func main() {
	maxPath := uniquePath(3,3)
	fmt.Println(maxPath)
	obstacleGrid := [][]int{{0,0,0},{0,1,0},{0,0,0}}
	maxPath2 := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(maxPath2)
}