package main

import "fmt"

func main() {
	//Println
	hello := "hello"
	world := "world"
	fmt.Println(hello)
	fmt.Println(world)

	//Printf
	fmt.Printf("%s %s\n", hello, world)

	//Sprintf
	s := fmt.Sprintf("%s %s", hello, world)
	fmt.Println(s)

	//Data type
	fmt.Printf("hello : %T\n", hello)
	fmt.Printf("world : %T\n", world)
}
