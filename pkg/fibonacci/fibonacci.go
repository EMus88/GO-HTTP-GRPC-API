package fibonacci

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Culc(firstNum, secondNum int) []int {
	numbers := make([]int, 0, secondNum-firstNum+1)
	for i := firstNum; i <= secondNum; i++ {
		numbers = append(numbers, Fibonacci(i))
	}
	return numbers
}
