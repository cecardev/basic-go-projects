package main

import "fmt"

type car struct {
	brand string
	model string
}

func main() {
	myCar := car{brand: "Ford", model: "Mustang"}
	fmt.Println(myCar)
}
