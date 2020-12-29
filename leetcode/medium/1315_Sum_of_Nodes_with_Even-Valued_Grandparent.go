package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret int

func sumEvenGrandparent(root *TreeNode) int {
	ret = 0
	SerchTreePreOrder(root)

	return ret
}

func SerchTreePreOrder(node *TreeNode) {

	if node == nil {
		return
	}

	if node.Val%2 == 0 {
		childChild := GetChildChild(node)

		for _, n := range childChild {
			if n != nil {
				fmt.Printf("Node: %d\n", n.Val)
				ret += n.Val
			}
		}
	}

	SerchTreePreOrder(node.Left)
	SerchTreePreOrder(node.Right)

	return
}

func GetChildChild(node *TreeNode) []*TreeNode {

	leftChild := node.Left
	rightChild := node.Right

	childChild := make([]*TreeNode, 4)

	if leftChild != nil {
		childChild[0] = leftChild.Left
		childChild[1] = leftChild.Right
	}

	if rightChild != nil {
		childChild[2] = rightChild.Left
		childChild[3] = rightChild.Right
	}

	return childChild
}
