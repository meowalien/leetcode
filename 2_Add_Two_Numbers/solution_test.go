package __Add_Two_Numbers

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		x := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		y := 0
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	result := addTwoNumbers(l1, l2)
	if result.Val != 7 || result.Next.Val != 0 || result.Next.Next.Val != 8 {
		t.Error("addTwoNumbers error")
	}
}

func TestAddTwoNumbers1(t *testing.T) {
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
