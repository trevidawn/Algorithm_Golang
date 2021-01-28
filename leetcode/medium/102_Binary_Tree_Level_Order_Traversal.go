package medium

/*
	Runtime: 4 ms, faster than 16.51% of Go online submissions for Binary Tree Level Order Traversal.
	Memory Usage: 2.9 MB, less than 31.86% of Go online submissions for Binary Tree Level Order Traversal.
	Bottom-Up 방식이나, 코드 최적화하면 시간 더 단축 될 듯?
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var treeLevel = make([][]int, 0)

	inorder(root, 0, &treeLevel)
	return treeLevel
}

func inorder(node *TreeNode, level int, levelTree *[][]int) {

	if node == nil {
		return
	}
	if len(*levelTree) <= level {
		(*levelTree) = append((*levelTree), []int{})
	}
	(*levelTree)[level] = append((*levelTree)[level], node.Val)

	inorder(node.Left, level+1, levelTree)
	inorder(node.Right, level+1, levelTree)
}