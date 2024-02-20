package sum

func Sum(numbers []int) (sum int) {

	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersOfSum ...[]int) (sums []int) {
	for _, numbers := range numbersOfSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return
}
