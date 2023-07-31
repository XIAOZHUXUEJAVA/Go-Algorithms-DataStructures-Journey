package sortsreview

// func partition(nums []int, left, right int) int {
// 	i := left
// 	j := right
// 	for i < j {
// 		for i < j && nums[j] >= nums[left] {
// 			j--
// 		}
// 		for i < j && nums[i] <= nums[left] {
// 			i++
// 		}
// 		nums[i], nums[j] = nums[j], nums[i]
// 	}
// 	nums[i], nums[left] = nums[left], nums[i]
// 	return i
// }
// func quickSort(nums []int, left, right int) {
// 	if left >= right {
// 		return
// 	}
// 	pivot := partition(nums, left, right)
// 	quickSort(nums, left, pivot-1)
// 	quickSort(nums, pivot+1, right)
// }

func partition(nums []int, left, right int) int {
	i := left
	j := right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}
