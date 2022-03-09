package main

import (
	"fmt"
	"strconv"
)

func main() {
	value := 10
	value2 := 20

	if value > 0 {
		fmt.Println("value is greater than 0")
	}
	if value < 0 {
		fmt.Println("value is less than 0")
	}

	//whith an else
	if value > 0 {
		fmt.Println("value is greater than 0")
	} else {
		fmt.Println("value is less than 0")
	}

	//with and else if
	if value > 0 {
		fmt.Println("value is greater than 0")
	} else if value < 0 {
		fmt.Println("value is less than 0")
	} else {
		fmt.Println("value is 0")
	}

	//whit and
	if value > 0 && value2 < 10 {
		fmt.Println("value is between 0 and 10")
	}

	//with or
	if value > 0 || value < 10 {
		fmt.Println("value is between 0 and 10")
	}

	//text to int
	number, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("value is not a number:", err)
	}
	fmt.Println("number:", number)

}
