package search

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		}
	}
	return -1
}
