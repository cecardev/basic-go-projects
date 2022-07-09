package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was inorrect, expected %d, got %d", 10, total)
	}

}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}
	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorrect expected %d, got %d", item.n, fib)
		}
	}
}
