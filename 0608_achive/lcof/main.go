package main

import (
	"fmt"
	"math"
	"strings"
)

func replaceSpace(str string) string {
	ans := strings.Builder{}
	for _, s := range str {
		if s == ' ' {
			ans.WriteString("%20")
		} else {
			ans.WriteRune(s)
		}
	}
	return ans.String()
}

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

/*
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。

数值（按顺序）可以分成以下几个部分：

若干空格
一个 小数 或者 整数
（可选）一个 'e' 或 'E' ，后面跟着一个 整数
若干空格


小数（按顺序）可以分成以下几个部分：

（可选）一个符号字符（'+' 或 '-'）
下述格式之一：
至少一位数字，后面跟着一个点 '.'
至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
一个点 '.' ，后面跟着至少一位数字


整数（按顺序）可以分成以下几个部分：

（可选）一个符号字符（'+' 或 '-'）
至少一位数字
部分数值列举如下：

["+100", "5e2", "-123", "3.1416", "-1E-16", "0123"]
部分非数值列举如下：

["12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"]
*/

func isNumber(s string) bool {
	/*
		预备知识
		确定有限状态自动机（以下简称「自动机」）是一类计算模型。它包含一系列状态，这些状态中：
		有一种特殊的状态，被称作「初始状态」。
		还有一系列状态被称为「接受状态」，它们组成了一个特殊的集合。其中，一个状态可能既是「初始状态」，也是「接受状态」。
		起初，这个自动机处于「初始状态」。随后，它顺序地读取字符串中的每一个字符，并根据当前状态和读入的字符，按照某个事先约定好的「转移规则」，从当前状态转移到下一个状态；当状态转移完成后，它就读取下一个字符。当字符串全部读取完毕后，如果自动机处于某个「接受状态」，则判定该字符串「被接受」；否则，判定该字符串「被拒绝」。

		注意：如果输入的过程中某一步转移失败了，即不存在对应的「转移规则」，此时计算将提前中止。在这种情况下我们也判定该字符串「被拒绝」。

		一个自动机，总能够回答某种形式的「对于给定的输入字符串 S，判断其是否满足条件 P」的问题。在本题中，条件 P 即为「构成合法的表示数值的字符串」。

		自动机驱动的编程，可以被看做一种暴力枚举方法的延伸：它穷尽了在任何一种情况下，对应任何的输入，需要做的事情。

		自动机在计算机科学领域有着广泛的应用。在算法领域，它与大名鼎鼎的字符串查找算法「KMP」算法有着密切的关联；在工程领域，它是实现「正则表达式」的基础。
	*/

	/*
		解题思路：
		1. 首先去除字符串前后空格 strings.TrimLeft, strings.TrimRight
		2. 一个数值型字符串组成元素包括：数字，小数点，e或者E，正负号
		3. 首位确定符号位，首位包含了符号，后面所有的字符都不能包含符号
		4. 第二位可以为小数点，或者数字
		5. 小数点后可以结束，也可以为数字
		6. e后必须包含数字

	*/

	/*
		0.起始的空格
		1.符号位
		2.整数部分
		3.左侧有整数的小数点
		4.左侧无整数的小数点（根据前面的第二条额外规则，需要对左侧有无整数的两种小数点做区分）
		5.小数部分
		6.字符 e
		7.指数部分的符号位
		8.指数部分的整数部分
		9.末尾的空格
	*/

	// 状态转移表  status[i] 表示当前所处的状态, map[int32]int 表示输入key, 则会转移到的状态
	// 空格、sign(s), digital(d), dot(.), e

	status := []map[int32]int{
		{' ': 0, 's': 1, 'd': 2, '.': 4}, // 0 状态
		{'d': 2, '.': 4},                 // 1 状态
		{'d': 2, ' ': 9, 'e': 6, '.': 3}, // 2 状态
		{'e': 6, ' ': 9, 'd': 5},         // 3
		{'d': 5},                         // 4
		{'d': 5, 'e': 6, ' ': 9},         // 5
		{'s': 7, 'd': 8},                 // 6
		{'d': 8},                         // 7
		{'d': 8, ' ': 9},                 // 8
		{' ': 9},                         // 9
	}

	var (
		p = 0
		t int32
	)
	for _, c := range s {
		if c == ' ' {
			t = ' '
		} else if c == '+' || c == '-' {
			t = 's'
		} else if '0' <= c && c <= '9' {
			t = 'd'
		} else if c == '.' {
			t = '.'
		} else if c == 'e' || c == 'E' {
			t = 'e'
		} else {
			t = '?'
		}
		if v, ok := status[p][t]; ok {
			p = v
		} else {
			return false
		}
	}
	return p == 2 || p == 3 || p == 5 || p == 8 || p == 9
}

func strToInt(str string) int {
	// 去除空格
	// 判断是否有符号位
	// 整数
	// 非整数字符返回
	if len(str) == 0 {
		return 0
	}
	i := 0
	for j, s := range str {
		if s != ' ' {
			i = j
			break
		}
	}
	sign := 1
	if str[i] == '-' {
		sign = -1
	}
	if str[i] == '-' || str[i] == '+' {
		i += 1
	}
	res := 0
	for _, c := range str[i:] {
		if c < '0' || c > '9' {
			break
		}
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && int(str[i]-'0') > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		res = res*10 + int(c-'0')

	}
	return sign * res

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	if head.Next != nil {
		res = reversePrint(head.Next)
	}

	res = append(res, head.Val)
	return res
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
本题要求我们对一个特殊的链表进行深拷贝。如果是普通的链表，我们可以直接按照遍历的顺序创建链表节点。
而本题中因为随机指针的存在，当我们拷贝节点时，「当前节点的随机指针指向的节点」可能还没创建，因此我们需要变换思路。
一个可行方案是，我们利用回溯的方式，让每个节点的拷贝操作相互独立。
对于当前节点，我们首先要进行拷贝，然后我们进行「当前节点的后继节点」和「当前节点的随机指针指向的节点」拷贝，拷贝完成后将创建的新节点的指针返回，即可完成当前节点的两指针的赋值。

具体地，我们用哈希表记录每一个节点对应新节点的创建情况。
遍历该链表的过程中，我们检查「当前节点的后继节点」和「当前节点的随机指针指向的节点」的创建情况。
如果这两个节点中的任何一个节点的新节点没有被创建，我们都立刻递归地进行创建。
当我们拷贝完成，回溯到当前层时，我们即可完成当前节点的指针赋值。
注意一个节点可能被多个其他节点指向，因此我们可能递归地多次尝试拷贝某个节点，为了防止重复拷贝，我们需要首先检查当前节点是否被拷贝过，如果已经拷贝过，我们可以直接从哈希表中取出拷贝后的节点的指针并返回即可。

在实际代码中，我们需要特别判断给定节点为空节点的情况。
*/

var cacheNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if cur, ok := cacheNode[node]; ok {
		return cur
	}
	newNode := &Node{Val: node.Val}
	cacheNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

/*
1. 遍历原始链表，复制原始节点
2. 遍历原始节点，构建random节点指向
3. 分离复制链表
*/

func copyList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 1. 遍历所有节点
	cur := head
	for cur != nil {
		newNode := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = newNode
		cur = newNode.Next
	}
	// 2. 构建random指向

	cur = head
	for cur != nil {
		cur.Next.Random = cur.Random.Next
		cur = cur.Next.Next
	}

	// 3. 拆解

	res := head.Next
	cur = head.Next
	pre := head
	for cur != nil {
		pre.Next = pre.Next.Next
		cur.Next = cur.Next.Next
		pre = pre.Next
		cur = cur.Next
	}
	return res

}

func copyRandomList(head *Node) *Node {
	return deepCopy(head)
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		if head.Next == nil {
			return nil
		} else {
			return head.Next
		}
	}
	pre := head
	cur := head.Next
	for cur != nil {
		if cur.Val == val {
			if cur.Next == nil {
				pre.Next = nil
			} else {
				pre.Next = cur.Next
			}
		}
		pre = pre.Next
		cur = cur.Next
	}
	return head

}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 双指针，先确定 k间距
	// 遍历前指针至列尾
	cur, pre := head, head
	for i := k; i == 1; i-- {
		cur = cur.Next
	}
	for cur != nil {
		cur, pre = cur.Next, pre.Next
	}
	return pre
}

// 合并两个排序的链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	res := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			res.Next, l1 = l1, l1.Next
		} else {
			res.Next, l2 = l2, l2.Next
		}
		res = res.Next
	}
	if l1 != nil {
		res.Next = l1
	} else {
		res.Next = l2
	}
	return head.Next
}

// 两个链表的第一个公共节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cacheMap := make(map[*ListNode]struct{})
	for headA != nil {
		cacheMap[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := cacheMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	tmpA := headA
	tmpB := headB
	for {
		if tmpA == nil && tmpB == nil || tmpA == tmpB {
			return tmpA
		}
		if tmpA == nil {
			tmpA = headB
		} else {
			tmpA = tmpA.Next
		}
		if tmpB == nil {
			tmpB = headA
		} else {
			tmpB = tmpB.Next
		}
	}
}

// 调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	//if len(nums) == 0 {
	//	return nil
	//}
	//oddList := make([]int, 0)
	//evenList := make([]int, 0)
	//for _, n := range nums {
	//	if n%2 == 0 {
	//		evenList = append(evenList, n)
	//	} else {
	//		oddList = append(oddList, n)
	//	}
	//}
	//res := append(oddList, evenList...)
	//return res
	if len(nums) == 0 {
		return nil
	}
	res := make([]int, 0, len(nums))
	for _, n := range nums {
		if n%2 == 0 {
			res = append(res, n)
		} else {
			res = append([]int{n}, res...)
		}
	}
	return res
}

func twoSum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for i := 0; i <= len(nums)/2; i++ {
		if target-nums[left] > nums[right] {
			left += 1
		} else if target-nums[left] < nums[right] {
			right -= 1
		} else {
			return []int{left, right}
		}
	}
	return nil
}

func main() {
	//res := replaceSpace("We are happy.")
	//fmt.Println(res)

	//res := reverseLeftWords("lrloseumgh", 6)
	//fmt.Println(res)

	//res := isNumber(" . ")
	//fmt.Println(res)
	//fmt.Println('3' - '0')

	//res := strToInt("  -42")
	//fmt.Println(res)

	//res := reversePrint()
	list := []int{1, 2}
	fmt.Println(list[1:])
}
