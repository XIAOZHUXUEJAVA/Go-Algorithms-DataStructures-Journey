package sortsreview

func mergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	// 将切片分为两半，递归的进行归并排序
	midIndex := len(numbers) >> 1
	leftNumbers := numbers[:midIndex]
	rightNumbers := numbers[midIndex:]
	return merge(mergeSort(leftNumbers), mergeSort(rightNumbers))
}

func merge(leftNumbers, rightNumbers []int) []int {
	leftIndex, rightIndex := 0, 0
	merged := make([]int, 0, len(leftNumbers)+len(rightNumbers))

	for leftIndex < len(leftNumbers) && rightIndex < len(rightNumbers) {
		if leftNumbers[leftIndex] < rightNumbers[rightIndex] {
			merged = append(merged, leftNumbers[leftIndex])
			leftIndex++
		} else {
			merged = append(merged, rightNumbers[rightIndex])
			rightIndex++
		}
	}
	// 将两个切片中的剩余元素追加到merged中，因为比较的过程中可能有剩余
	merged = append(merged, leftNumbers[leftIndex:]...)
	merged = append(merged, rightNumbers[rightIndex:]...)
	return merged
}
