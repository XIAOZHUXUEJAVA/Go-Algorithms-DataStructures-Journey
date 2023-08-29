package sortsreview02

import "fmt"

func bubbleSort(numbers []int) error {
	if numbers == nil {
		return fmt.Errorf("the input numbers can not be null")
	}
	if len(numbers) <= 1 {
		return nil
	}
	length := len(numbers)
	for i := length - 1; i > 0; i-- {
		hasSwapped := false
		for j := 0; j < i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				hasSwapped = true
			}
		}
		if !hasSwapped {
			break
		}
	}
	return nil
}
