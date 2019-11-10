package main

func Sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, val := range numbersToSum {
		sums = append(sums, Sum(val))
	}
	return
}
