package medium

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var SumOfNode map[int]int

func deepestLeavesSum(root *TreeNode) int {

	return GetSumOfDeepestNode(root)
}



func GetSumOfDeepestNode(root *TreeNode) int {

	SumOfNode = map[int]int{}
	keys := make([]int, 0, 10)

	Preorder(root, 0)

	for depth, _ := range SumOfNode {
		keys = append(keys, depth)
	}

	sort.Ints(keys)

	return SumOfNode[keys[len(keys)-1]]
}

func Preorder(node *TreeNode, depth int) {

	if node == nil {
		return
	}

	SumOfNode[depth] = SumOfNode[depth] + node.Val

	Preorder(node.Left, depth+1)
	Preorder(node.Right, depth+1)

	return
}