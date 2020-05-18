package solutions

// OddEvenList function is a solution for the
// Problem #328 - Odd Even Linked List
// from leetcode.com
func OddEvenList(head *ListNode) *ListNode {
	index := 1

	var oddHead *ListNode
	var evenHead *ListNode
	var oddNode *ListNode
	var evenNode *ListNode

	for head != nil {
		if index == 1 {
			oddHead = head
			oddNode = head
		} else if index == 2 {
			evenHead = head
			evenNode = head
		} else if index%2 == 1 {
			oddNode.Next = head
			oddNode = oddNode.Next
		} else if index%2 == 0 {
			evenNode.Next = head
			evenNode = evenNode.Next
		}

		head = head.Next
		index++
	}

	if evenNode != nil {
		evenNode.Next = nil
	}

	if oddNode != nil {
		oddNode.Next = evenHead
		return oddHead
	}

	return head
}

// ListNode struct for solving the problem
type ListNode struct {
	Val  int
	Next *ListNode
}
