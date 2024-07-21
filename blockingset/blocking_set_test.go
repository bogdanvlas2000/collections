package blockingset

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Test_BlockingSet_Put(t *testing.T) {
	blockingSet := New[string]()

	wg := sync.WaitGroup{}

	goroutinesCount := 100

	value := "test"

	results := make(chan bool)

	wg.Add(goroutinesCount)
	for i := 0; i < goroutinesCount; i++ {
		go func() {
			defer wg.Done()
			result := blockingSet.Put(value)

			results <- result
		}()
	}

	succeeded := 0
	for i := 0; i < goroutinesCount; i++ {
		result := <-results
		if result {
			succeeded++
		}
	}

	expectedSucceeded := 1
	assert.Equal(t, expectedSucceeded, succeeded)
}
