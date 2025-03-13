package Add_Two_Numbers

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	result := addTwoNumbers(l1, l2)
	if result.Val != 7 || result.Next.Val != 0 || result.Next.Next.Val != 8 {
		t.Error("addTwoNumbers error")
	}
}

func TestAddTwoNumbers1(t *testing.T) {
	fmt.Println("TestAddTwoNumbers1")
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 0}
	result := addTwoNumbers(l1, l2)
	if result.Val != 0 {
		t.Error("addTwoNumbers error")
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	l1 := &ListNode{Val: 9}
	l1.Next = &ListNode{Val: 9}
	l1.Next.Next = &ListNode{Val: 9}
	l1.Next.Next.Next = &ListNode{Val: 9}
	l1.Next.Next.Next.Next = &ListNode{Val: 9}
	l1.Next.Next.Next.Next.Next = &ListNode{Val: 9}
	l1.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}
	l2 := &ListNode{Val: 9}
	l2.Next = &ListNode{Val: 9}
	l2.Next.Next = &ListNode{Val: 9}
	l2.Next.Next.Next = &ListNode{Val: 9}
	result := addTwoNumbers(l1, l2)
	if result.Val != 8 || result.Next.Val != 9 || result.Next.Next.Val != 9 || result.Next.Next.Next.Val != 9 || result.Next.Next.Next.Next.Val != 0 || result.Next.Next.Next.Next.Next.Val != 0 || result.Next.Next.Next.Next.Next.Next.Val != 0 || result.Next.Next.Next.Next.Next.Next.Next.Val != 1 {
		t.Error("addTwoNumbers error")
	}
}
