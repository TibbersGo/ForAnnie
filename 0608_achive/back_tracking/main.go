package main

import "fmt"

// 回溯模板
// 回溯法一般是在集合中递归搜索
func backtracking() {
	/*
		if 终止条件 {
			存放结果
			return
		}
		for （选择：本层集合中元素（树中节点孩子的数量就是集合的大小）） {
			处理节点；
			backtracking(路径，选择列表)； 递归调用
			回溯，撤销处理结果
		}
		for循环理解为横向遍历，backtracking理解为纵向遍历
	*/
}

func backtrackingPermutation(n, k, index int, path []int, res [][]int) [][]int {
	if len(path) == k {
		dst := make([]int, len(path), len(path))
		copy(dst, path)
		res = append(res, dst)
		return res
	}
	for i := index; i <= n; i++ {
		if n-i+1 < k-len(path) {
			break
		}
		path = append(path, i)
		res = backtrackingPermutation(n, k, i+1, path, res)
		path = path[:len(path)-1]
	}
	return res
}

func combinationSum(k, n, index int, path []int, res [][]int) [][]int {
	if len(path) == k && sum(path) == n {
		tmp := make([]int, k)
		copy(tmp, path)
		res = append(res, tmp)
	}
	for i := index; i < 10; i++ {
		if len(path) < k {
			path = append(path, i)
			if sum(path) > n {
				break
			}
		} else {
			break
		}
		res = combinationSum(k, n, i+1, path, res)
		path = path[:len(path)-1]
	}
	return res
}

func sum(path []int) int {
	res := 0
	for _, v := range path {
		res += v
	}
	return res
}

func permutation(index int, nums, path []int, res [][]int, used []bool) [][]int {
	if len(nums) == index {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return res
	}
	fmt.Println("-------")
	for i := 0; i < len(nums); i++ {
		fmt.Println("---")
		if !used[i] {
			path = append(path, nums[i])
			fmt.Println(path)
			used[i] = true
			res = permutation(index+1, nums, path, res, used)
			fmt.Println(path)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	return res
}

func permute(nums []int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	res = permutation(0, nums, path, res, used)
	fmt.Println(res)
	return res
}

var (
	iList       = make([]int, 0)
	iSum    int = 0
	sum1Res     = make([][]int, 0)
)

func sum1(nums []int, target, index int) [][]int {
	fmt.Println(iList)
	if iSum == target {
		tmp := make([]int, len(iList))
		copy(tmp, iList)
		sum1Res = append(sum1Res, tmp)
		return sum1Res
	}
	for i := index; i < len(nums); i++ {
		if iSum+nums[i] <= target {
			iList = append(iList, nums[i])
			iSum += nums[i]
		} else {
			continue
		}
		sum1Res = sum1(nums, target, index)
		iSum -= iList[len(iList)-1]
		iList = iList[:len(iList)-1]
	}
	return sum1Res
}

func subsets() {
	nums := []int{1, 2, 3}
	sets := make([]int, 0)
	res := make([][]int, 0)
	res = subset(nums, sets, res, 0)
	fmt.Println(res)
}

func subset(nums []int, sets []int, res [][]int, index int) [][]int {
	if index == len(nums) {
		return res
	}
	for i := index; i < len(nums); i++ {
		sets = append(sets, nums[i])
		tmp := make([]int, len(sets))
		copy(tmp, sets)
		res = append(res, tmp)
		res = subset(nums, sets, res, i+1)
		sets = sets[:len(sets)-1]
	}
	return res
}

func main() {
	subsets()
}
