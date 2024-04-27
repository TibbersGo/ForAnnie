package main

import "fmt"

type Tree struct {
	Left  *Tree
	Right *Tree
	Value int
}

func main() {
	// 3, 9, 20, null, null, 15, 7

	//root := &Tree{
	//	Value: 3,
	//	Left: &Tree{
	//		Value: 9,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &Tree{
	//		Value: 20,
	//		Left: &Tree{
	//			Value: 15,
	//		},
	//		Right: &Tree{
	//			Value: 7,
	//		},
	//	},
	//}
	//r := tree(root)
	//fmt.Println(r)

	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	root := buildTree(inorder, postorder)
	r := tree(root)
	fmt.Println(r)
}

func tree(root *Tree) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*Tree{root}
	for len(queue) != 0 {
		l := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < l; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			tmp = append(tmp, queue[0].Value)
			queue = queue[1:]
		}
		res = append(res, tmp)
	}
	ress := make([][]int, 0)
	for i := len(res) - 1; i >= 0; i-- {
		ress = append(ress, res[i])
	}

	return ress
}

func buildTree(inorder []int, postorder []int) *Tree {
	// 取后序数组最后一个元素
	// 分割中序数组
	// 分割后序数组
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := &Tree{
		Value: postorder[len(postorder)-1],
	}
	if len(postorder) == 1 {
		return root
	}
	leftSubIn := make([]int, 0)
	rightSubIn := make([]int, 0)
	leftSubPost := make([]int, 0)
	rightSubPost := make([]int, 0)
	for i, v := range inorder {
		if v == root.Value {
			leftSubIn = inorder[:i]
			rightSubIn = inorder[i+1:]
			leftSubPost = postorder[:i]
			rightSubPost = postorder[i : len(postorder)-1]
			break
		}
	}
	root.Left = buildTree(leftSubIn, leftSubPost)
	root.Right = buildTree(rightSubIn, rightSubPost)
	return root
}
