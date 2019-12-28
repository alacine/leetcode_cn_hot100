package q002

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	s := &ListNode{Val: 2}
	s.Next = &ListNode{Val: 4}
	s.Next.Next = &ListNode{Val: 3}

	r := &ListNode{Val: 5}
	r.Next = &ListNode{Val: 6}
	r.Next.Next = &ListNode{Val: 4}

	exp := &ListNode{Val: 7}
	exp.Next = &ListNode{Val: 0}
	exp.Next.Next = &ListNode{Val: 8}

	var got = addTwoNumbers(s, r)
	for i := 0; got != nil && exp != nil; i++ {
		fmt.Printf("got[%d]=%d, exp[%d]=%d\n", i, got.Val, i, exp.Val)
		if got.Val != exp.Val {
			t.Errorf("got[%d]=%d, exp[%d]=%d", i, got.Val, i, exp.Val)
		}
		got = got.Next
		exp = exp.Next
	}
}
