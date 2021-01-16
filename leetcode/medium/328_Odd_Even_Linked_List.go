package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, even, cur := head, head.Next, head.Next.Next
	curOdd, curEven := odd, even
	isOdd := true
	for cur != nil {
		if isOdd {
			curOdd.Next = cur
			curOdd = curOdd.Next
		} else {
			curEven.Next = cur
			curEven = curEven.Next
		}
		isOdd = !isOdd
		cur = cur.Next
	}
	curEven.Next = nil
	curOdd.Next = even
	return odd
}
