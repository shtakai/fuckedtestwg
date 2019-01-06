package math

func Sum(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		// _ is index
		sum += n
	}
	return sum
}

func Add(a, b int) int {
	return a + b
}