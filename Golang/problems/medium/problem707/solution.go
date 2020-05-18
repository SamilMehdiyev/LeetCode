package solutions

import "fmt"

// MyLinkedList - Custom Linked List Implementation
type MyLinkedList struct {
	Head   *Node
	Tail   *Node
	length int
}

// Node struct which used in Linked List
type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

// Constructor - Initialize your data structure here.
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get - Get the value of the index-th node in the linked list. If the index is invalid, return -1.
func (obj *MyLinkedList) Get(index int) int {
	node := obj.Head
	for i := 0; i <= index && i < obj.length; i++ {
		if node != nil {
			if i == index {
				return node.Val
			}
			node = node.Next
		}
	}
	return -1
}

// AddAtHead - Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
func (obj *MyLinkedList) AddAtHead(val int) {
	node := Node{val, obj.Head, nil}
	if obj.Head != nil {
		obj.Head.Prev = &node
	}
	obj.Head = &node

	if obj.Tail == nil {
		obj.Tail = &node
	}
	obj.length++
}

// AddAtTail - Append a node of value val to the last element of the linked list.
func (obj *MyLinkedList) AddAtTail(val int) {
	node := Node{val, nil, obj.Tail}
	if obj.Tail != nil {
		obj.Tail.Next = &node
	}
	obj.Tail = &node

	if obj.Head == nil {
		obj.Head = &node
	}
	obj.length++
}

// AddAtIndex - Add a node of value val before the index-th node in the linked list.
// If index equals to the length of linked list, the node will be appended to the end of linked list.
// If index is greater than the length, the node will not be inserted.
func (obj *MyLinkedList) AddAtIndex(index int, val int) {

	newNode := Node{val, nil, nil}
	if index == 0 {
		obj.AddAtHead(val)
		return
	} else if index == obj.length {
		obj.AddAtTail(val)
		return
	}

	node := obj.Head
	prev := node
	for i := 0; i <= index && i < obj.length; i++ {
		if node != nil {
			if i == index {
				newNode.Prev = prev
				newNode.Next = node

				prev.Next = &newNode
				node.Prev = &newNode
				obj.length++
				return
			}
			prev = node
			node = node.Next
		}
	}
}

// DeleteAtIndex - Delete the index-th node in the linked list, if the index is valid.
func (obj *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 && obj.Head != nil {
		obj.Head = obj.Head.Next
		obj.length--
		return
	} else if index == obj.length-1 && obj.Tail != nil {
		obj.Tail = obj.Tail.Prev
		obj.length--
		return
	}

	node := obj.Head
	prev := node
	for i := 0; i <= index && i < obj.length; i++ {
		if node != nil {
			if i == index {
				prev.Next = node.Next

				if node.Next != nil {
					node.Next.Prev = prev
				}

				node.Next = nil
				node.Prev = nil

				obj.length--
				return
			}
			prev = node
			node = node.Next
		}
	}
}

// Print - prints all the values
func (obj *MyLinkedList) Print() {
	node := obj.Head
	for node != nil {
		fmt.Print(node.Val, ", ")
		node = node.Next
	}
	fmt.Println()
}
