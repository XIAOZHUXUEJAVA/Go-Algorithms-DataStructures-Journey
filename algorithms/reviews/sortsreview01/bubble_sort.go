package sortsreview01

func bubbleSort(array []int) {
	if array == nil {
		panic("the input array is null")
	}
	if len(array) <= 1 {
		return
	}
	length := len(array)
	for i := length - 1; i > 0; i-- {
		hasSwapped := false
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				hasSwapped = true
			}
		}
		if !hasSwapped {
			break
		}
	}

}
