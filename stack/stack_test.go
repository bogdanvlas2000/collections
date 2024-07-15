package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push_Pop(t *testing.T) {
	nums := []int{1, 2, 3}
	s := New[int]()

	for _, num := range nums {
		s.Push(num)
	}

	var results []int
	for {
		if num, ok := s.Pop(); ok {
			results = append(results, num)
		} else {
			break
		}
	}

	expectedResults := []int{3, 2, 1}

	assert.Equal(t, expectedResults, results)
}

func TestStack_Size(t *testing.T) {
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
			nums:         []int{1, 2, 3},
			expectedSize: 3,
		},
	}

	for _, test := range tests {
		s := New[int]()

		for _, num := range test.nums {
			s.Push(num)
		}
		assert.Equal(t, test.expectedSize, s.Size())
	}
}
