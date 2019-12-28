package q002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s = new(ListNode)
	ans := s
	a := 0
	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		s.Next = &ListNode{Val: (x + y + a) % 10}
		a = (x + y + a) / 10
		s = s.Next
	}
	if a > 0 {
		s.Next = &ListNode{Val: a}
	}
	return ans.Next
}
