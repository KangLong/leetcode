package main

import "fmt"

func main() {
	l1 := InitData([]int{1, 2, 4})
	l2 := InitData([]int{1, 3, 4})
	node := mergeTwoLists(l1, l2)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
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

/**
26.删除排序数组中的重复项
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成
*/
func removeDuplicates(nums []int) int {
	var index = 0
	//1.原地交换 那需要准备两个指针，一个指针 index 用于存储非重复元素在数组中的位置
	//另外一个用于遍历整个数组
	//2.有序数组 当出现一个不同的元素时，则将此元素 放入 index 的位置
	//因 i 必然大于>=index, 被覆盖的元素已经被对比过，可以被舍弃
	for i := 1; i < len(nums); i++ {
		if nums[index] != nums[i] {
			index++
			nums[index] = nums[i]
		}
	}
	return index + 1 //index代表角标， +1 则代表元素个数
}

/**
189. 旋转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
要求使用空间复杂度为 O(1) 的 原地 算法
*/
func rotate(nums []int, k int) {
	//旋转数组，即 将数组前面 k 个元素移动到数组末尾
	//暴力解法： 用两个循环 外循环为 k 次， 内循环将最开头的元素一次移动到末尾 时间复杂度为 O(k*n) 空间 O(1)
	//额外数组解法： 复制数组，用两个并列的循环，一次将 前面 k 个元素复制到元素末尾，再将 k 之后的元素复制到开头
	//时间复杂度为 O(n) 空间复杂度为O(n)
	//反转方法：由于有 k%n 个元素会被移动到数组头部， 故直接反转整个数组， 然后将 前面k个元素反转，再将后面 n-k 个元素反转即可

	k %= len(nums)
	//反转整个数组
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, len(nums)-k-1)
	reverse(nums, len(nums)-k, len(nums)-1)
}

//将数组指定区段的元素完全反转
func reverse(nums []int, start int, end int) {
	for start < end {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start++
		end--
	}
}

/**
21. 合并两个有序链表
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//暴力：依次用l1链表中的元素和l2中元素对比，如果 l2<l1 则将 l2的元素插入 l1 的当前位置 时间 O(n2)
	//递归：
	//迭代

	prevHead := &ListNode{Val: -1}
	prev := prevHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			prev = prev.Next
			l1 = l1.Next
		} else {
			prev.Next = l2
			prev = prev.Next
			l2 = l2.Next
		}
	}

	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}

	return prevHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
641. 设计循环双端队列
设计实现双端队列。
你的实现需要支持以下操作：

MyCircularDeque(k)：构造函数,双端队列的大小为k。
insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
isEmpty()：检查双端队列是否为空。
isFull()：检查双端队列是否满了。
*/
