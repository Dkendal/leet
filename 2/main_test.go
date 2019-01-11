package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"reflect"
	"testing"
	"testing/quick"
)

var _ = fmt.Printf

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func (t *ListNode) String() string {
	result := fmt.Sprintf("%d", t.Val)
	t = t.Next
	for t != nil {
		result += fmt.Sprintf(" -> %d", t.Val)
		t = t.Next
	}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * You are given two non-empty linked lists representing two non-negative
 * integers. The digits are stored in reverse order and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 *
 * Example:
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{}
	n := &result
	carry := 0

	test := func() bool {
		return l1 != nil || l2 != nil || carry != 0
	}

	next := func(l *ListNode) *ListNode {
		if l == nil {
			return nil
		}
		return l.Next
	}

	val := func(l *ListNode) int {
		if l == nil {
			return 0
		}
		return l.Val
	}

	for {
		v := val(l1) + val(l2) + carry
		r := (v % 10)
		carry = v / 10
		n.Val = r

		l1 = next(l1)
		l2 = next(l2)
		if test() {
			n.Next = &ListNode{}
			n = n.Next
		} else {
			break
		}
	}

	return &result
}

func NewInt(l *ListNode) int {
	result := 0
	for e := 1; l != nil; e *= 10 {
		result += l.Val * e
		l = l.Next
	}
	return result
}

func NewListNode(x int) *ListNode {
	if x == 0 {
		return &ListNode{}
	}

	result := &ListNode{}
	n := result
	size := int(math.Log10(float64(x))) + 1

	for i := 0; i < size; i++ {
		r := x % 10
		x = x / 10

		n.Val = r

		if i < (size - 1) {
			n.Next = &ListNode{}
			n = n.Next
		}
	}

	return result
}

func TestNewInt(t *testing.T) {
	assert.Equal(t, 0, NewInt(&ListNode{0, nil}))

	assert.Equal(t,
		123,
		NewInt(&ListNode{3, &ListNode{2, &ListNode{1, nil}}}),
	)
}

func TestNewListNode(t *testing.T) {
	assert.Equal(
		t,
		&ListNode{},
		NewListNode(0),
	)

	assert.Equal(
		t,
		&ListNode{3, &ListNode{2, &ListNode{1, nil}}},
		NewListNode(123),
	)

	assert.Equal(t, NewInt(NewListNode(183958634923505057)), 183958634923505057)

	assert.Equal(t, NewInt(NewListNode(0)), 0)
}

func TestAddTwoNumbers(t *testing.T) {
	genTest := func(i, j int) {
		assert.Equal(t, NewListNode(i+j),
			addTwoNumbers(NewListNode(i), NewListNode(j)))

		assert.Equal(t, NewListNode(i+j),
			addTwoNumbers(NewListNode(j), NewListNode(i)))
	}

	genTest(342, 465)
	genTest(1234570912398740, 1234)
	genTest(0, 12345)
	genTest(0, 0)
	genTest(1, 1)
	genTest(5, 5)
	genTest(2822802455101979716, 202761764396069819)
	genTest(3789433100770299615, 8367414889222326951)

	err := quick.Check(func(i, j int) bool {
		if i >= 0 && j >= 0 {
			return reflect.DeepEqual(NewListNode(i+j),
				addTwoNumbers(NewListNode(i), NewListNode(j)))
		}

		return true

	}, nil)
	assert.NoError(t, err)
}
