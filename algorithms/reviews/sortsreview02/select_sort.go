package sortsreview02

import "fmt"

func selectSort(nums []int) error {
	if nums == nil {
		return fmt.Errorf("the input nums cannot be null")
	}

	n := len(nums)
	if n < 2 {
		return nil
	}

	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
	return nil
}
