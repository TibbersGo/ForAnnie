package main

func trap2(height []int) int {
	singleStack := make([]int, 0)
	res := 0
	for index, value := range height {
		for len(singleStack) > 0 && height[singleStack[len(singleStack)-1]] < value {
			// 弹出栈顶元素
			mid := singleStack[len(singleStack)-1]
			singleStack = singleStack[:len(singleStack)-1]
			// 计算雨水
			h := min(value, height[singleStack[len(singleStack)-1]]) - height[mid]
			w := index - singleStack[len(singleStack)-1] - 1
			res += w * h
		}
		if height[singleStack[len(singleStack)-1]] == value {
			singleStack[len(singleStack)-1] = index
		} else {
			singleStack = append(singleStack, index)
		}
	}
	return res
}

func trap(height []int) int {
	res := 0
	// 第一个柱子和最后一个不接雨水
	for i := 0; i < len(height); i++ {
		if i == 0 || i == len(height)-1 {
			continue
		}
		lHeight := height[i]
		rHeight := height[i]

		// 求左侧最高
		for l := i - 1; l >= 0; l-- {
			if height[l] > lHeight {
				lHeight = height[l]
			}
		}

		for r := i + 1; r < len(height); r++ {
			if height[r] > rHeight {
				rHeight = height[r]
			}
		}
		cur := min(lHeight, rHeight) - height[i]
		if cur > 0 {
			res += cur
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 单调栈
func dailyTemperatures(temperatures []int) []int {
	singleStack, res := make([]int, 0), make([]int, 0)
	for index, t := range temperatures {
		for len(singleStack) > 0 && temperatures[singleStack[len(singleStack)-1]] < t {
			res[singleStack[len(singleStack)-1]] = index - singleStack[len(singleStack)-1]
			singleStack = singleStack[:len(singleStack)-1]
		}
		singleStack = append(singleStack, index)
	}
	return res
}
