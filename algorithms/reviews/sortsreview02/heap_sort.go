package sortsreview02

func heapSort(nums []int) {
	n := len(nums)
	for i := 0; i < n/2-1; i++ {
		heapify(nums, n, i)
	}
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
}

func heapify(nums []int, n, i int) {
	max := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && nums[left] > nums[max] {
		max = left
	}
	if right < n && nums[right] > nums[max] {
		max = right
	}

	if max != i {
		nums[i], nums[max] = nums[max], nums[i]
		heapify(nums, n, max)
	}
}
