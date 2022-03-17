package main

import "fmt"

type figures2D interface {
	area() float64
}

type square struct {
	width float64
}

type rectangule struct {
	width, height float64
}

func (s square) area() float64 {
	return s.width * s.width
}

func (r rectangule) area() float64 {
	return r.width * r.height
}

func calculateArea(f figures2D) float64 {
	return f.area()
}

func main() {
	mySquare := square{width: 5}
	myRectangule := rectangule{width: 5, height: 10}
	fmt.Println(calculateArea(mySquare))
	fmt.Println(calculateArea(myRectangule))

	//interface list
	myInterfaceList := []figures2D{mySquare, myRectangule}
	fmt.Println(myInterfaceList)

}
