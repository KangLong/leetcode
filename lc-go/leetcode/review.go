package main

import "fmt"

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

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	size := removeDuplicates(nums)
	for i := 0; i < size; i++ {
		fmt.Print(nums[i], ",")
	}
}
