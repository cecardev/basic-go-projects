package main

import "fmt"

func main() {

	//with coiditional
	switch module := 5 % 2; module {
	case 0:
		fmt.Println("module is pair")
	default:
		fmt.Println("module is not pair")
	}

	//withou conditional
	value := 10
	switch {
	case value > 0:
		fmt.Println("value is greater than 0")
	case value < 0:
		fmt.Println("value is less than 0")
	default:
		fmt.Println("value is 0")

	}

}
