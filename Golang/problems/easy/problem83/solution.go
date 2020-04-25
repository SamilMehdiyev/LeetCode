package solutions

// DeleteDuplicates function is a solution for the
// Problem #83 - Remove Duplicates from Sorted List
// from leetcode.com
func DeleteDuplicates(arr []int) *ListNode {

	head := generateListNode((arr))
	if head == nil || head.Next == nil {
		return head
	}

	current := head

	for current != nil {
		next := current.Next

		for next != nil && current.Val == next.Val {
			next = next.Next
		}
		current.Next = next
		current = current.Next
	}

	return head
}

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func generateListNode(arr []int) *ListNode {
	var listNode *ListNode = nil

	for i := len(arr) - 1; 0 <= i; i-- {
		var item = arr[i]

		if listNode == nil {
			listNode = &ListNode{item, nil}
			continue
		}

		node := ListNode{item, nil}
		node.Next = listNode

		listNode = &node
	}

	return listNode
}
