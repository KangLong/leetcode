package main

import "fmt"

//1.双端队列可从头尾一起操作数据， 故需要记录头结点和尾结点
//2.队列有大小大小限制 故需要记录总大小和当前大小
//3.需要一个结构体做数据记录,同时记录前后节点的位置
//4.暂不考虑扩容问题
type MyCircularDeque struct {
	startNode *Node
	endNode   *Node
	totalSize int
	curSize   int
}

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{totalSize: k}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.curSize >= this.totalSize {
		return false
	}
	node := Node{data: value}
	if this.curSize == 0 {
		this.startNode = &node
		this.endNode = &node
	} else {
		this.startNode.prev = &node
		node.next = this.startNode
		this.startNode = &node
	}
	this.curSize++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.curSize >= this.totalSize {
		return false
	}
	node := Node{data: value}
	if this.curSize == 0 {
		this.startNode = &node
		this.endNode = &node
	} else {
		node.prev = this.endNode
		this.endNode.next = &node
		this.endNode = &node
	}
	this.curSize++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.curSize <= 0 {
		return false
	} else if this.curSize == 1 {
		this.startNode = nil
		this.endNode = nil
	} else {
		this.startNode.next.prev = nil
		this.startNode = this.startNode.next
	}
	this.curSize--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.curSize <= 0 {
		return false
	} else if this.curSize == 1 {
		this.startNode = nil
		this.endNode = nil
	} else {
		this.endNode.prev.next = nil
		this.endNode = this.endNode.prev
	}
	this.curSize--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.curSize <= 0 {
		return -1
	}
	re := this.startNode.data
	return re.(int)
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.curSize <= 0 {
		return -1
	}
	re := this.endNode.data
	return re.(int)
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.curSize == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {

	return this.curSize == this.totalSize
}

func main() {
	circularDeque := MyCircularDeque{totalSize: 8}
	fmt.Println(circularDeque.InsertLast(5))
	fmt.Println(circularDeque.GetFront())
	fmt.Println(circularDeque.DeleteLast())
	//fmt.Println(circularDeque.GetRear())
	fmt.Println(circularDeque.IsFull())
	fmt.Println(circularDeque.InsertFront(4))

}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
