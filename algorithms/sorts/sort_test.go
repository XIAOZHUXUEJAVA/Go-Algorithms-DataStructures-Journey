package sorts

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	originalNums := []int{8, 1, 2, 4, 3, 7, 9, 5}

	t.Run("Test quickSort()", func(t *testing.T) {
		quickSort(originalNums, 0, len(originalNums)-1)
		sortedNums := originalNums
		expectedNums := []int{1, 2, 3, 4, 5, 7, 8, 9}
		if !reflect.DeepEqual(sortedNums, expectedNums) {
			t.Errorf("sortedNums: %d, expectedNums: %d", sortedNums, expectedNums)
		}
	})
	t.Run("Test mergeSort()", func(t *testing.T) {
		sortedNums := mergeSort(originalNums)
		expectedNums := []int{1, 2, 3, 4, 5, 7, 8, 9}

		if !reflect.DeepEqual(sortedNums, expectedNums) {
			t.Errorf("sortedNums: %d, expectedNums: %d", sortedNums, expectedNums)
		}
	})
	t.Run("Test bubbleSort()", func(t *testing.T) {
		expectedNums := []int{1, 2, 3, 4, 5, 7, 8, 9}
		bubbleSort(originalNums)
		sortedNums := originalNums
		if !reflect.DeepEqual(sortedNums, expectedNums) {
			t.Errorf("sortedNums: %d, expectedNums: %d", sortedNums, expectedNums)
		}
	})
}
