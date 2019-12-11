package test

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitData(nums []int) (l1 *ListNode) {
	l1 = new(ListNode)
	cur := l1
	for _, v := range nums {
		tmp := new(ListNode)
		tmp.Val = v
		cur.Next = tmp
		cur = tmp
	}
	return l1.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num := 0
	addNum := 0
	startLn := new(ListNode)
	curLn := startLn
	for {
		if l1 == nil || l2 == nil {
			break
		}
		addNum = l1.Val + l2.Val + num
		if addNum > 9 {
			num = 1
			addNum = addNum % 10
		} else {
			num = 0
		}
		curLn = addNode(curLn, addNum)
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		curLn = NodeNotNull(curLn, l1, &num)
	}
	if l2 != nil {
		curLn = NodeNotNull(curLn, l2, &num)
	}
	if num != 0 {
		ff := new(ListNode)
		ff.Val = num
		curLn.Next = ff
	}
	return startLn.Next
}

func NodeNotNull(curLn *ListNode, node *ListNode, num *int) *ListNode {
	if *num != 0 {
		for node != nil {
			nu := node.Val + *num
			if nu > 9 {
				*num = 1
				nu = nu % 10
			} else {
				*num = 0
			}
			curLn = addNode(curLn, nu)
			node = node.Next
		}
	} else {
		curLn.Next = node
	}
	return curLn
}

func addNode(curLn *ListNode, val int) *ListNode {
	newLn := new(ListNode)
	newLn.Val = val
	curLn.Next = newLn
	return newLn
}

func TestAdd(t *testing.T) {
	l1 := InitData([]int{1})
	l2 := InitData([]int{9, 9})
	ln := addTwoNumbers(l1, l2)
	for ln != nil {
		fmt.Println(ln.Val)
		ln = ln.Next
	}
}
