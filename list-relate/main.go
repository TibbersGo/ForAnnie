package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	// 判断头节点
	for head.Val == val {
		head = head.Next
	}
	tmp := head
	for tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		}
		tmp = tmp.Next
	}
	return head
}

type node struct {
	val  int
	next *node
}

type MyLinkedList struct {
	dummyHead *node
	size      int
}

func Constructor() MyLinkedList {
	head := &node{
		val:  0,
		next: nil,
	}
	return MyLinkedList{
		dummyHead: head,
		size:      0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil || this.size == 0 || index < 0 || index >= this.size {
		return -1
	}
	cur := this.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &node{
		val:  val,
		next: nil,
	}
	if this.size == 0 {
		this.dummyHead.next = newNode
		this.size++
	} else {
		curHead := this.dummyHead.next
		this.dummyHead.next = newNode
		newNode.next = curHead
		this.size++
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &node{
		val:  val,
		next: nil,
	}
	if this.size == 0 {
		this.dummyHead.next = newNode
		this.size++
		return
	}
	cur := this.dummyHead
	for i := 0; i < this.size; i++ {
		cur = cur.next
	}
	cur.next = newNode
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this == nil || this.size < index {
		return
	}
	newNode := &node{
		val:  val,
		next: nil,
	}
	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	if this.size == index {
		cur.next = newNode
	} else {
		tmp := cur.next
		cur.next = newNode
		newNode.next = tmp
	}
	this.size++

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this == nil || this.size == 0 || index < 0 || index >= this.size {
		return
	}
	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	tmp := cur.next
	cur.next = cur.next.next
	tmp.next = nil
	this.size--
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		if cur.Next == nil {
			cur.Next = pre
			head = cur
			break
		}
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return head
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := &ListNode{
		Val:  0,
		Next: head,
	} // 虚拟头节点
	for head != nil && head.Next != nil {
		tmp := head.Next.Next
		cur.Next = head.Next
		head.Next.Next = head.Next
		head.Next = tmp
	}
	return cur.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre, cur := dummy, dummy
	for i := 0; i <= n; i++ {
		cur = cur.Next
	}

	for cur.Next != nil {
		pre = pre.Next
		cur = cur.Next
	}
	tmp := pre.Next.Next
	pre.Next.Next = nil
	pre.Next = tmp
	return dummy.Next
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	aLen, bLen := 0, 0
	tmpA, tmpB := headA, headB
	for tmpA != nil {
		tmpA = tmpA.Next
		aLen++
	}
	for tmpB != nil {
		tmpB = tmpB.Next
		bLen++
	}
	if aLen >= bLen {
		for i := 0; i < aLen-bLen; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < bLen-aLen; i++ {
			headB = headB.Next
		}
	}
	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for slow != nil && fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}

type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (q *MyQueue) Front() int {
	return q.queue[0]
}

func (q *MyQueue) Back() int {
	return q.queue[len(q.queue)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.queue) == 0
}

func (q *MyQueue) Push(val int) {
	// 对比最后一个，如果比最后一个大，则替换最后一个，如果小，加到尾部
	for !q.Empty() && val > q.Back() {
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue, val)
}

func (q *MyQueue) Pop(val int) {
	// 如果与第一个值相同，则删除
	if !q.Empty() && val == q.Front() {
		q.queue = q.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	length := len(nums)
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	res = append(res, queue.Front())
	for i := k; i < length; i++ {
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
		res = append(res, queue.Front())
	}
	return res
}

func isValid(s string) bool {
	sList := make([]string, 0)
	for _, v := range s {
		target := string(v)
		switch target {
		case "(":
			sList = append(sList, ")")
		case "{":
			sList = append(sList, "}")
		case "[":
			sList = append(sList, "]")
		default:
			if target != sList[len(sList)-1] {
				return false
			} else {
				sList = sList[:len(sList)-1]
			}
		}

	}
	if len(sList) != 0 {
		return false
	}
	return true
}

func evalRPN(tokens []string) int {
	res := 0
	sym := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}
	sList := make([]int, 0)
	for _, s := range tokens {

		if _, ok := sym[s]; ok {
			a := sList[len(sList)-2]
			b := sList[len(sList)-1]
			sList = sList[:len(sList)-2]
			switch s {
			case "+":
				sList = append(sList, a+b)
			case "-":
				sList = append(sList, a-b)
			case "*":
				sList = append(sList, a*b)
			case "/":
				sList = append(sList, a/b)
			}
		} else {
			tmp, _ := strconv.Atoi(s)
			sList = append(sList, tmp)
		}
	}
	res = sList[0]
	return res
}

func main() {
	// 移除链表中的指定元素
	s := "(){}[]"
	for k, v := range s {
		fmt.Println(k, string(v))
	}
}
