package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet_Size(t *testing.T) {
	tests := []struct {
		keys         []string
		expectedSize int
	}{
		{
			keys:         []string{},
			expectedSize: 0,
		},
		{
			keys: []string{
				"one", "two",
			},
			expectedSize: 2,
		},
		{
			keys: []string{
				"one", "two", "two", "one",
			},
			expectedSize: 2,
		},
	}
	for _, test := range tests {
		set := New[string](test.keys...)
		assert.Equal(t, test.expectedSize, set.Size())
	}
}

func TestSet_Put(t *testing.T) {
	tests := []struct {
		keys            []string
		expectedResults []bool
		expectedSize    int
	}{
		{
			keys:            []string{},
			expectedResults: []bool{},
			expectedSize:    0,
		},
		{
			keys: []string{
				"one", "two",
			},
			expectedResults: []bool{
				true, true,
			},
			expectedSize: 2,
		},
		{
			keys: []string{
				"one", "two", "two", "one",
			},
			expectedResults: []bool{
				true, true, false, false,
			},
			expectedSize: 2,
		},
	}

	for _, test := range tests {
		set := New[string]()

		for i, key := range test.keys {
			ok := set.Put(key)
			assert.Equal(t, test.expectedResults[i], ok)
		}

		assert.Equal(t, test.expectedSize, set.Size())
	}
}

func TestSet_Remove(t *testing.T) {
	tests := []struct {
		keys            []string
		keysToRemove    []string
		expectedResults []bool
	}{
		{
			keys: []string{},
			keysToRemove: []string{
				"one",
			},
			expectedResults: []bool{
				false,
			},
		},
		{
			keys: []string{
				"one", "two",
			},
			keysToRemove: []string{
				"one", "two", "three",
			},
			expectedResults: []bool{
				true, true, false,
			},
		},
		{
			keys: []string{
				"one", "two",
			},
			keysToRemove: []string{
				"one", "two", "two",
			},
			expectedResults: []bool{
				true, true, false,
			},
		},
	}
	for _, test := range tests {
		set := New[string](test.keys...)

		for i, key := range test.keysToRemove {
			ok := set.Remove(key)
			assert.Equal(t, test.expectedResults[i], ok)
		}
	}
}

func TestSet_Contains(t *testing.T) {
	tests := []struct {
		keys            []string
		keysToContain   []string
		expectedResults []bool
	}{
		{
			keys: []string{},
			keysToContain: []string{
				"one",
			},
			expectedResults: []bool{
				false,
			},
		},
		{
			keys: []string{
				"one", "two",
			},
			keysToContain: []string{
				"one", "two", "three",
			},
			expectedResults: []bool{
				true, true, false,
			},
		},
	}
	for _, test := range tests {
		set := New[string](test.keys...)

		for i, key := range test.keysToContain {
			ok := set.Contains(key)
			assert.Equal(t, test.expectedResults[i], ok)
		}
	}
}

func TestSet_Equals(t *testing.T) {
	tests := []struct {
		keysSrc []string
		keysDst []string

		expectedResult bool
	}{
		{
			keysSrc:        []string{},
			keysDst:        []string{},
			expectedResult: true,
		},
		{
			keysSrc: []string{
				"one", "two",
			},
			keysDst: []string{
				"one", "two",
			},
			expectedResult: true,
		},
		{
			keysSrc: []string{
				"one", "two", "three",
			},
			keysDst: []string{
				"three", "two", "one",
			},
			expectedResult: true,
		},
		{
			keysSrc: []string{
				"one", "two", "two", "one",
			},
			keysDst: []string{
				"two", "one", "two", "two",
			},
			expectedResult: true,
		},
		{
			keysSrc: []string{
				"one", "two",
			},
			keysDst: []string{
				"one", "two", "three",
			},
			expectedResult: false,
		},
	}
	for _, test := range tests {
		src := New[string](test.keysSrc...)
		dst := New[string](test.keysDst...)

		assert.Equal(t, test.expectedResult, src.Equals(dst))
		assert.Equal(t, test.expectedResult, dst.Equals(src))
	}
}
