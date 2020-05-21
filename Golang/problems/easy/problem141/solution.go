package solutions

// HasCycle function is a solution for the
// Problem #141 - Linked List Cycle
// from leetcode.com
func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	step, first, second := 1, head, head.Next
	length := 0

	for ; ; step++ {
		if first == nil || second == nil {
			return false
		}

		if first.Val == second.Val {
			if length == step {
				return true
			}
			length = step
			step = 1
		}

		first = first.Next
		if second.Next != nil {
			second = second.Next.Next
		}

	}
}

// ListNode - struct for the Problem
type ListNode struct {
	Val  int
	Next *ListNode
}
