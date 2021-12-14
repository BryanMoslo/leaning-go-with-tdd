package arrays

func Sum(array []int) int {
	var result int
	for i := range array {
		result += array[i]
	}

	return result
}

func SumAll(numbers ...[]int) []int {
	var result []int

	for i := range numbers {
		result = append(result, Sum(numbers[i]))
	}

	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	var result []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			tail := numbers[1:]
			result = append(result, Sum(tail))
		}
	}

	return result
}
