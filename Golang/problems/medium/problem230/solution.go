package solutions

import "sort"

// kthSmallest function is a solution for the
// Problem #230 - Kth Smallest Element in a BST
// from leetcode.com
func kthSmallest(root *TreeNode, k int) int {
	var stack []int

	getElements(root, &stack)
	stack = append(stack, root.Val)
	sort.Ints(stack)
	return stack[k-1]
}

func getElements(node *TreeNode, stack *[]int) *TreeNode {
	if node == nil {
		return node
	}

	left := getElements(node.Left, stack)
	right := getElements(node.Right, stack)

	if left != nil && right != nil {
		(*stack) = append(*stack, left.Val)
		(*stack) = append(*stack, right.Val)
	} else if left == nil && right != nil {
		(*stack) = append(*stack, right.Val)
	} else if right == nil && left != nil {
		(*stack) = append(*stack, left.Val)
	}

	return node
}

// TreeNode for the Problem 230
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
