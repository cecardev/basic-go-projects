package main

import "fmt"

func main() {

	//for
	fmt.Println("for:")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for inverse
	fmt.Println("for inverse:")
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	//for while
	fmt.Println("\n for while:")
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//for forever
	fmt.Println("\n for forever:")
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}

}
