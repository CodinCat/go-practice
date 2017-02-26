package practice

func BubbleSort(input []int) {
	for i := range input {
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				input[j], input[i] = input[i], input[j]
			}
		}
	}
}
