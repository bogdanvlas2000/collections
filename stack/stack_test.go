package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
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
