package main

import "fmt"

func main() {
	//defer
	//Is the last function to be executed
	//eg: Close db connection / close file
	defer fmt.Println("world")
	fmt.Println("hello")

	//Continue break
	//continue: is mostly for errrors even an error continue the loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			fmt.Println("continue")
			continue
		}
	}

	//break
	//break: is mostly for errrors even an error break the loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			fmt.Println("break")
			break
		}
	}
}
