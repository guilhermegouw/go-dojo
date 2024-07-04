package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(arrays ...[]int) []int {
	var sums []int

	for _, numbers := range arrays {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(arrays ...[]int) []int {
	var sums []int

	for _, numbers := range arrays {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return sums
}
