package test

/**
26.删除排序数组中的重复项
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成
*/
func removeDuplicates(nums []int) int {
	//1.已排序数组，当碰到不一样的元素时，其后面的必然不一样
	//2.用index表示当前非重复数组最后的位置，用 i 表示处理到元素的第几个元素
	curIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[curIndex] != nums[i] {
			curIndex++
			nums[curIndex] = nums[i]
		}
	}
	return curIndex + 1
}

/**
189. 旋转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
*/
func rotate(nums []int, k int) {

}

/**
21. 合并两个有序链表
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
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
