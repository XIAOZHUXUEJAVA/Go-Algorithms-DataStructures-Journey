package sortsreview01

// // mergeSort 对数组进行归并排序
// func mergeSort(nums []int) []int {
// 	if len(nums) <= 1 {
// 		return nums
// 	}

// 	middle := len(nums) / 2
// 	left := mergeSort(nums[:middle])
// 	right := mergeSort(nums[middle:])

// 	return merge(left, right)
// }

// // merge 将两个有序数组合并为一个有序数组
// func merge(left, right []int) []int {
// 	result := make([]int, 0, len(left)+len(right))
// 	l, r := 0, 0

// 	for l < len(left) && r < len(right) {
// 		if left[l] < right[r] {
// 			result = append(result, left[l])
// 			l++
// 		} else {
// 			result = append(result, right[r])
// 			r++
// 		}
// 	}

// 	result = append(result, left[l:]...)
// 	result = append(result, right[r:]...)

// 	return result
// }

func merge(leftNumbers, rightNumbers []int) []int {
	leftIndex, rightIndex := 0, 0
	mergedNumbers := make([]int, 0, len(leftNumbers)+len(rightNumbers))
	for leftIndex < len(leftNumbers) && rightIndex < len(rightNumbers) {
		if leftNumbers[leftIndex] < rightNumbers[rightIndex] {
			mergedNumbers = append(mergedNumbers, leftNumbers[leftIndex])
			leftIndex++
		} else {
			mergedNumbers = append(mergedNumbers, rightNumbers[rightIndex])
			rightIndex++
		}
	}
	mergedNumbers = append(mergedNumbers, leftNumbers[leftIndex:]...)
	mergedNumbers = append(mergedNumbers, rightNumbers[rightIndex:]...)
	return mergedNumbers
}

func mergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	middleIndex := len(numbers) >> 1
	leftNumbers := numbers[:middleIndex]
	rightNumbers := numbers[middleIndex:]
	return merge(mergeSort(leftNumbers), mergeSort(rightNumbers))
}
