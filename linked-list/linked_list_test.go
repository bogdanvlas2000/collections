package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Push(t *testing.T) {
	tests := []struct {
		nums         []int
		expectedSize int
	}{
		{
			nums:         []int{},
			expectedSize: 0,
		},
		{
			nums:         []int{1},
			expectedSize: 1,
		},
		{
			nums:         []int{1, 2, 3, 4, 5},
			expectedSize: 5,
		},
	}

	for _, test := range tests {
		s := New[int]()
		for _, n := range test.nums {
			s.Push(n)
		}

		assert.Equal(t, test.expectedSize, s.Size())
	}
}

func TestLinkedList_Pop(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{
			nums: []int{},
		},
		{
			nums: []int{1},
		},
		{
			nums: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		s := New[int]()

		for _, num := range test.nums {
			s.Push(num)
		}

		if len(test.nums) > 0 {
			for i := len(test.nums) - 1; i >= 0; i-- {
				n, ok := s.Pop()
				assert.True(t, ok)
				assert.Equal(t, test.nums[i], n)
			}
		}

		_, ok := s.Pop()
		assert.False(t, ok)
	}
}

func TestLinkedList_Peek(t *testing.T) {
	tests := []struct {
		nums          []int
		expectedValue int
		expectedOk    bool
	}{
		{
			nums:          []int{},
			expectedValue: 0,
			expectedOk:    false,
		},
		{
			nums:          []int{1},
			expectedValue: 1,
			expectedOk:    true,
		},
		{
			nums:          []int{1, 2, 3},
			expectedValue: 3,
			expectedOk:    true,
		},
	}

	for _, test := range tests {
		s := New[int]()

		for _, num := range test.nums {
			s.Push(num)
		}

		value, ok := s.Peek()
		assert.Equal(t, test.expectedValue, value)
		assert.Equal(t, test.expectedOk, ok)
		assert.Equal(t, len(test.nums), s.Size())
	}
}

func TestLinkedList_PullFirst(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{
			nums: []int{},
		},
		{
			nums: []int{1},
		},
		{
			nums: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		s := New[int]()

		for _, num := range test.nums {
			s.Push(num)
		}

		if len(test.nums) > 0 {
			for i := 0; i < len(test.nums); i++ {
				n, ok := s.PullFirst()
				assert.True(t, ok)
				assert.Equal(t, test.nums[i], n)
			}
		}

		_, ok := s.PullFirst()
		assert.False(t, ok)
	}
}

func TestLinkedList_First(t *testing.T) {
	tests := []struct {
		nums          []int
		expectedValue int
		expectedOk    bool
	}{
		{
			nums:          []int{},
			expectedValue: 0,
			expectedOk:    false,
		},
		{
			nums:          []int{1},
			expectedValue: 1,
			expectedOk:    true,
		},
		{
			nums:          []int{1, 2, 3},
			expectedValue: 1,
			expectedOk:    true,
		},
	}

	for _, test := range tests {
		s := New[int]()

		for _, num := range test.nums {
			s.Push(num)
		}

		value, ok := s.First()
		assert.Equal(t, test.expectedValue, value)
		assert.Equal(t, test.expectedOk, ok)
		assert.Equal(t, len(test.nums), s.Size())
	}
}
