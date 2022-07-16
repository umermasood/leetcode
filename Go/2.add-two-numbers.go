package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current, carry := result, 0

	for l1 != nil || l2 != nil {
		val := carry

		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		carry, val = val/10, val%10
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	if carry == 1 {
		current.Next = &ListNode{Val: carry}
	}

	return result.Next
}
