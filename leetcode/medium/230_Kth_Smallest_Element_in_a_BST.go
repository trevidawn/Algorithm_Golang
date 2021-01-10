package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	sorted := &[]int{}
	inorder(root, sorted)

	return (*sorted)[k-1]
}

func inorder(node *TreeNode, sortedNode *[]int) {
	if node == nil {
		return
	}

	inorder(node.Left, sortedNode)
	*sortedNode = append(*sortedNode, node.Val)
	inorder(node.Right, sortedNode)

}