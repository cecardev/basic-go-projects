package main

import (
	"fmt"

	mypackage "example.com/example/mypackage"
)

func main() {
	var myCar mypackage.Car
	myCar.Brand = "Ford"
	myCar.Model = "Mustang"
	fmt.Println(myCar)
}
