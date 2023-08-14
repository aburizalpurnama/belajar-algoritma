package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	list := []int{5, 3, 4, 1, 2}
	sortedList := MergeSort(list)
	// fmt.Printf("sortedList: %v\n", sortedList)
	assert.False(t, VerifySortedAsc(list))
	assert.True(t, VerifySortedAsc(sortedList))
}

func TestVerifySortedAsc(t *testing.T) {
	list := []int{5, 3, 4, 1, 2}
	assert.False(t, VerifySortedAsc(list))
}
