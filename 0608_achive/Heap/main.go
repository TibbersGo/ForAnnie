package main

import "container/heap"

type Heap [][2]int

// 对数组第一个元素排序

func NewHeap() Heap {
	res := make([][2]int, 0)
	return res
}

func (h Heap) swap(a, b int) {
	// 交换两个节点的值
	h[a], h[b] = h[b], h[a]
}

func (h Heap) less(a, b int) bool {
	// 比较两个值的大小
	if h[a][0] < h[b][0] {
		return true
	}
	return false
}

func (h Heap) Push(x [2]int) {
	// 第一步放到队尾
	// 向上比较
}

// 方法一：小顶堆
func topKFrequent(nums []int, k int) []int {
	map_num := map[int]int{}
	//记录每个元素出现的次数
	for _, item := range nums {
		map_num[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key, value := range map_num {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	//按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

// 构建小顶堆
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
