package sortsreview01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {

	unorderedArray := []int{4, 4, 2, 2, 1, 1, 3}
	bubbleSort(unorderedArray)
	expectedOrderedArray := []int{1, 1, 2, 2, 3, 4, 4}
	assert.Equal(t, expectedOrderedArray, unorderedArray)
}

func TestMergeSort(t *testing.T) {
	unorderedArray := []int{4, 4, 2, 2, 1, 1, 3}
	orderedArray := mergeSort(unorderedArray)
	expectedOrderedArray := []int{1, 1, 2, 2, 3, 4, 4}
	assert.Equal(t, expectedOrderedArray, orderedArray)
}

func TestQuickSort(t *testing.T) {

	unorderedArray := []int{4, 4, 2, 2, 1, 1, 3}
	quickSort(unorderedArray, 0, len(unorderedArray)-1)
	expectedOrderedArray := []int{1, 1, 2, 2, 3, 4, 4}
	assert.Equal(t, expectedOrderedArray, unorderedArray)
}
