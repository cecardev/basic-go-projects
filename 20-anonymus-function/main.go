package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	y := func() int {
		return x
	}
	fmt.Println(y())

	c := make(chan int)
	go func() {
		fmt.Println("start")
		time.Sleep(3 * time.Second) // sleep for 1 second
		fmt.Println("done")
		c <- x
	}()
	<-c
}
