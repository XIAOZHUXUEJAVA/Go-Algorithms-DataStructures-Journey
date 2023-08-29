package sortsreview02

func merge(leftNumbers, rightNumbers []int) []int {
	if leftNumbers == nil || rightNumbers == nil {

	}
	leftIndex := 0
	rightIndex := 0
	mergedNumbers := make([]int, len(leftNumbers)+len(rightNumbers))
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

func mergrSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	middleIndex := len(numbers) >> 1
	leftNumbers := numbers[:middleIndex]
	rightNumbers := numbers[middleIndex:]
	return merge(mergrSort(leftNumbers), mergrSort(rightNumbers))
}
