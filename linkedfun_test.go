package linkedfun

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const NUM_TESTS = 1000

func BuildList(size int) (head *Node, tail *Node) {
	var h *Node = nil
	var t *Node = nil

	if size > 0 {
		h = &Node{
			value: 0,
		}
		t = h
	}

	for i := 1; i < size; i = i + 1 {
		t.next = &Node{
			value: i,
		}
		t = t.next
	}

	return h, t
}

func Increasing(head *Node) bool {
	h := head
	n := head.next

	for n != nil {

		if h.value > n.value {
			return false
		}

		h = n
		n = h.next
	}

	return true
}

func Decreasing(head *Node) bool {
	h := head
	n := head.next

	for n != nil {

		if h.value < n.value {
			return false
		}

		h = n
		n = h.next
	}

	return true
}

func TestNoCycle(t *testing.T) {
	for i := 0; i < NUM_TESTS; i = i + 1 {
		head, _ := BuildList(i)
		assert.False(t, FindCycle(head))
	}
}

func TestFullCycle(t *testing.T) {
	for i := 1; i < NUM_TESTS; i = i + 1 {
		head, tail := BuildList(i)
		tail.next = head
		assert.True(t, FindCycle(head))
	}
}

func TestCycleWithPrefix(t *testing.T) {
	for i := 1; i < NUM_TESTS; i = i + 1 {
		head, tail := BuildList(i)
		cycleHead, cycleTail := BuildList(i)
		tail.next = cycleHead
		cycleTail.next = tail
		assert.True(t, FindCycle(head), fmt.Sprintf("There should be a cycle of %d", i))
	}
}

func TestReverse(t *testing.T) {
	for i := 1; i < NUM_TESTS; i = i + 1 {
		head, tail := BuildList(i)

		assert.True(t, Increasing(head))

		reversed := Reverse(head)
		assert.Nil(t, head.next)
		assert.Equal(t, reversed.value, tail.value)
		assert.True(t, Decreasing(reversed))

		if i != 1 {
			assert.NotNil(t, tail.next)
		}
	}
}
