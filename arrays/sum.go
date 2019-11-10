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

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, val := range numbersToSum {
		if len(val) == 0 {
			sums = append(sums, 0)
		} else {
			//syntax: slice[low:high], skipping one side captures that side
			tail := val[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return
}
