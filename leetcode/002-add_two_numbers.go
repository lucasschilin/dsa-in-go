package leetcode

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example 1:
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.

// Example 2:
// Input: l1 = [0], l2 = [0]
// Output: [0]

// Example 3:
// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

// Constraints:
// - The number of nodes in each linked list is in the range [1, 100].
// - 0 <= Node.val <= 9
// - It is guaranteed that the list represents a number that does not have leading zeros.

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

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) map[string]func() any {
	fns := map[string]func() any{
		"Resolution": func() any { return addTwoNumbers1(l1, l2) },
	}

	return fns
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var currentNode *ListNode

	carry := 0

	l1CurrentNode := l1
	l2CurrentNode := l2

	for l1CurrentNode != nil || l2CurrentNode != nil || carry > 0 {

		l1Val := 0
		if l1CurrentNode != nil {
			l1Val = l1CurrentNode.Val
			l1CurrentNode = l1CurrentNode.Next
		}

		l2Val := 0
		if l2CurrentNode != nil {
			l2Val = l2CurrentNode.Val
			l2CurrentNode = l2CurrentNode.Next
		}

		sum := l1Val + l2Val + carry
		digit := sum % 10
		carry = sum / 10

		newNode := &ListNode{Val: digit}

		if head == nil {
			head = newNode
		} else {
			currentNode.Next = newNode
		}
		currentNode = newNode
	}

	return head
}
