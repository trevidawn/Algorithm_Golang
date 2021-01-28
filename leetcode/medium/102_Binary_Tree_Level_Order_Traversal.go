package medium

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