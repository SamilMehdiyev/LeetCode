package solutions

import (
	"fmt"
)

func swapNodes(head *ListNode, k int) *ListNode {
	left, right := swapKthNode(head, 1, &k, nil, nil)
	fmt.Println(left, right)

	return head
}

func swapKthNode(head *ListNode, length int, k *int, left *ListNode, right *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		*k -= 1
		return left, right
	}
	fmt.Println(head.Val, left, right, length, k)

	if length == *k && left == nil {
		left = head
	}

	swapKthNode(head.Next, length+1, k, left, right)

	*k -= 1

	if *k == 0 && right == nil {
		right = head
	}

	fmt.Println(head.Val, left, right, length, k)
	return left, right
}

type ListNode struct {
	Val  int
	Next *ListNode
}
