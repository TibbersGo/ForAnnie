package traversal

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func perOrder(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = append(res, node.Val)
	res = perOrder(node.Left, res)
	res = perOrder(node.Right, res)
	return res
}

func postOrder(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = postOrder(node.Left, res)
	res = postOrder(node.Right, res)
	res = append(res, node.Val)
	return res
}

func middleOrder(node *TreeNode, res []int) []int {
	if node == nil {
		return res
	}
	res = middleOrder(node.Left, res)
	res = append(res, node.Val)
	res = middleOrder(node.Right, res)
	return res
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	return traversal(root, "", res)
}

func traversal(node *TreeNode, s string, res []string) []string {
	if node.Left == nil && node.Right == nil {
		s = s + "->" + strconv.Itoa(node.Val)
		res = append(res, s)
		return res
	}
	s = s + "->" + strconv.Itoa(node.Val)
	if node.Left != nil {
		traversal(node.Left, s, res)
	}
	if node.Right != nil {
		traversal(node.Right, s, res)
	}
	return res
}
