package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursiveBinarySearch(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isFounded := RecursiveBinarySearch(numbers, 11)
	assert.False(t, isFounded)

	isFounded = RecursiveBinarySearch(numbers, 0)
	assert.False(t, isFounded)

	isFounded = RecursiveBinarySearch(numbers, 1)
	assert.True(t, isFounded)

	isFounded = RecursiveBinarySearch(numbers, 10)
	assert.True(t, isFounded)

	isFounded = RecursiveBinarySearch(numbers, 7)
	assert.True(t, isFounded)
}
