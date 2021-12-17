package main

import (
	"project/pkg/fibonacci"
	"testing"
)

func TestFibonacci(t *testing.T) {
	data := [20]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 234, 377, 610, 987, 1597, 2584, 4181}

	for i, d := range data {
		if fibonacci.Fibonacci(i) != d {
			t.Errorf("Invalid Fibonacci value for N: %d, want: %d", i, d)
		}
	}
}
