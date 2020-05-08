package solutions

// IsCousins function is a solution for the
// Problem #993 - Cousins in Binary Tree
// from leetcode.com
func IsCousins(root *TreeNode, x int, y int) bool {
	hashMap := make(map[int]item)

	depthFirstSearch(root, 0, 0, hashMap)
	result := hashMap[x].level == hashMap[y].level && hashMap[x].parent != hashMap[y].parent
	return result
}

func depthFirstSearch(node *TreeNode, level int, parent int, hashMap map[int]item) {
	if node == nil {
		return
	}

	hashMap[node.Val] = item{level, parent}
	depthFirstSearch(node.Left, level+1, node.Val, hashMap)  // recursive call - left side
	depthFirstSearch(node.Right, level+1, node.Val, hashMap) // recursive call - right side
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type item struct {
	level  int
	parent int
}
