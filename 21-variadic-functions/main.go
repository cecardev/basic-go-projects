package main

import (
	"fmt"
	"math"
)

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (number int, square int, cube int) {
	number = 5
	square = int(math.Pow(float64(x), 2))
	cube = int(math.Pow(float64(x), 3))
	return
}

func main() {
	print(sum(1, 2, 3, 4, 5))
	printNames("John", "Paul", "George", "Ringo")
	println(getValues(2))
	println(getValues(5))
}
