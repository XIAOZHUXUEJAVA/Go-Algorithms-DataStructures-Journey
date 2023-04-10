package sorts

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	originalNums := []int{8, 1, 2, 4, 3, 7, 9, 5}
	expectedNums := []int{1, 2, 3, 4, 5, 7, 8, 9}
	t.Run("Test quickSort()", func(t *testing.T) {
		quickSort(originalNums, 0, len(originalNums)-1)
		sortedNums := originalNums
		if !reflect.DeepEqual(sortedNums, expectedNums) {
			t.Errorf("sortedNums: %d, expectedNums: %d", sortedNums, expectedNums)
		}
	})
	t.Run("Test mergeSort()", func(t *testing.T) {
		sortedNums := mergeSort(originalNums)

		if !reflect.DeepEqual(sortedNums, expectedNums) {
			t.Errorf("sortedNums: %d, expectedNums: %d", sortedNums, expectedNums)
		}
	})
	t.Run("Test bubbleSort()", func(t *testing.T) {
		bubbleSort(originalNums)
		sortedNums := originalNums
		if !reflect.DeepEqual(sortedNums, expectedNums) {
			t.Errorf("sortedNums: %d, expectedNums: %d", sortedNums, expectedNums)
		}
	})

	t.Run("Test insertSort()", func(t *testing.T) {
		insertSort(originalNums)
		sortedNums := originalNums
		if !reflect.DeepEqual(sortedNums, expectedNums) {
			t.Errorf("sortedNums: %d, expectedNums: %d", sortedNums, expectedNums)
		}
	})
}
