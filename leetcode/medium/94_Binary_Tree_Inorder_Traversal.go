package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	InorderTraval = make([]int, 0)
	Inorder(root)

	return InorderTraval
}

var InorderTraval []int

func Inorder(node *TreeNode) {

	if node == nil {
		return
	}

	Inorder(node.Left)
	InorderTraval = append(InorderTraval, node.Val)
	Inorder(node.Right)
}