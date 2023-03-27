package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	index := binarySearch(nums, 3)
	fmt.Println(index)
}
