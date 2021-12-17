package fibonacci

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func Culc(firstNum, secondNum int) []int {
	numbers := make([]int, 0, secondNum-firstNum+1)
	for i := firstNum; i <= secondNum; i++ {
		numbers = append(numbers, fibonacci(i))
	}
	return numbers
}
