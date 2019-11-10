package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//1.两边只能从头节点依次遍历元素，比大小
	//2.有序链表，当 A 链表中的一个元素大于 B 链表中的元素时，则 B 链表应当继续找一下个元素与 A 链表中的元素比较，
	//直到 B > A ，则 A 找下一个元素 与 B 进行比较，直至所有元素比较完毕

	ln := new(ListNode)
	cur := ln
	if l1 != nil && l2 != nil {
		for l1 != nil {
			for l2 != nil {
				if l1 == nil || l2 == nil {
					break
				}
				if l1.Val > l2.Val {
					cur.Next = newListNode(l2.Val)
					cur = cur.Next
					l2 = l2.Next
				} else if l1.Val < l2.Val {
					cur.Next = newListNode(l1.Val)
					cur = cur.Next
					l1 = l1.Next
				} else {
					cur.Next = newListNode(l1.Val)
					cur.Next.Next = newListNode(l2.Val)
					cur = cur.Next.Next
					l1 = l1.Next
					l2 = l2.Next
				}
			}
			if l2 == nil && l1 != nil {
				cur.Next = l1
				break
			} else if l1 == nil && l2 != nil {
				cur.Next = l2
				break
			}
		}
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l2 != nil && l1 == nil {
		return l2
	}

	return ln.Next
}

func newListNode(val int) *ListNode {
	ln := new(ListNode)
	ln.Val = val
	return ln
}

func initData(nums []int) (l1 *ListNode) {
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

func main() {

	l1 := initData([]int{2})
	l2 := initData([]int{1})

	nl := mergeTwoLists(l1, l2)

	for nl != nil {
		fmt.Print(nl.Val, ",")
		nl = nl.Next
	}
}
