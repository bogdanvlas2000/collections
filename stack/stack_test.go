package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
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

func TestStack_Pop(t *testing.T) {
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

func TestStack_Peek(t *testing.T) {
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
