package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func findTargetNode(root *TreeNode, cnt int) int {
	// 中序遍历后倒序的第cnt个节点
	//

	return 0
}

var cnt int
var res int

func dfs(node *TreeNode, cnt int) {
	if node == nil {
		return
	}
	dfs(node.Right, cnt)
	if cnt == 0 {
		return
	}
	cnt -= 1
	if cnt == 0 {
		res = node.Val
	}
	dfs(node.Left, cnt)
}
