package traversal

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
