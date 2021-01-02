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

	return GetSumOfDeepestNode3(root)
}

/*
   DFS
   Runtime: 36 ms, faster than 9.41% of Go online submissions for Deepest Leaves Sum.
   Memory Usage: 7 MB, less than 84.71% of Go online submissions for Deepest Leaves Sum.
*/

func GetSumOfDeepestNode3(root *TreeNode) int {
	Sum := 0
	DeepestDepth := 0

	var DFS func(node *TreeNode, depth int)
	DFS = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if depth == DeepestDepth {
			Sum += node.Val
		} else if depth > DeepestDepth {
			Sum = node.Val
			DeepestDepth = depth
		}

		DFS(node.Left, depth+1)
		DFS(node.Right, depth+1)
	}

	DFS(root, 0)

	return Sum
}

/*
   BFD
   Runtime: 28 ms, faster than 77.65% of Go online submissions for Deepest Leaves Sum.
   Memory Usage: 7.4 MB, less than 12.94% of Go online submissions for Deepest Leaves Sum.
*/
func GetSumOfDeepestNode2(root *TreeNode) int {

	Sum := 0
	CurrentNodes := []*TreeNode{root}

	for len(CurrentNodes) != 0 {
		Sum = 0
		NextNodes := make([]*TreeNode, 0, len(CurrentNodes) * 2)

		for _, node := range CurrentNodes {
			Sum += node.Val

			if node.Left != nil {
				NextNodes = append(NextNodes, node.Left)
			}
			if node.Right != nil {
				NextNodes = append(NextNodes, node.Right)
			}
		}

		CurrentNodes = NextNodes
	}

	return Sum
}

/*
   Runtime: 36 ms, faster than 9.41% of Go online submissions for Deepest Leaves Sum.
   Memory Usage: 7 MB, less than 91.76% of Go online submissions for Deepest Leaves Sum.
*/
func GetSumOfDeepestNode1(root *TreeNode) int {

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