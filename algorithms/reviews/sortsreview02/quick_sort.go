package sortsreview02

func partition(numbers []int, leftIndex, rightIndex int) int {
	i, j := leftIndex, rightIndex
	for i < j {
		for i < j && numbers[j] >= numbers[leftIndex] {
			j--
		}
		for i < j && numbers[i] <= numbers[leftIndex] {
			i++
		}
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	numbers[i], numbers[leftIndex] = numbers[leftIndex], numbers[i]
	return i
}

func quickSort(numbers []int, leftIndex, rightIndex int) {
	if leftIndex >= rightIndex {
		return
	}
	pivotIndex := partition(numbers, leftIndex, rightIndex)
	quickSort(numbers, leftIndex, pivotIndex-1)
	quickSort(numbers, pivotIndex+1, rightIndex)
}
