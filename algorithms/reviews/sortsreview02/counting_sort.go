package sortsreview02

import "math"

func countingSort(nums []int) {
	maxValue := float64(0)
	for _, num := range nums {
		maxValue = math.Max(float64(maxValue), float64(num))
	}
	counts := make([]int, int(maxValue)+1)

	for _, num := range nums {
		counts[num]++
	}

	sortedIndex := 0
	for num := 0; num < len(counts); num++ {
		for i := 0; i < counts[num]; i++ {
			nums[sortedIndex] = num
			sortedIndex++
		}
	}
}
