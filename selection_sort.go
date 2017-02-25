package practice

func SelectionSort(input []int) {
	for i := range input {
		minIndex := selectMin(input, i)
		input[i], input[minIndex] = input[minIndex], input[i]
	}
}

func selectMin(input []int, from int) int {
	minIndex := from
	for i := from; i < len(input); i++ {
		if input[i] < input[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}
