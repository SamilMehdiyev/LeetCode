package solutions

// DetectCycle function is a solution for the
// Problem #142 - Linked List Cycle II
// from leetcode.com
func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			for head != fast {
				fast = fast.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}

// ListNode - struct for the Problem
type ListNode struct {
	Val  int
	Next *ListNode
}
