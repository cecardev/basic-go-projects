package main

import "fmt"

func function(message string) {
	fmt.Println(message)
}

func returnValue(message string) string {
	return message
}

func returnDoubleValue(message1 string, message2 string) (string, string) {
	return message1, message2
}

func calculateArea(length float64, width float64) float64 {
	return length * width
}

func main() {
	function("Hello, World!")
	returnValue := returnValue("Return value")
	fmt.Println(returnValue)

	returnDoubleValue1, returnDoubleValue2 := returnDoubleValue("returnDoubleValue1", "returnDoubleValue2")
	fmt.Println(returnDoubleValue1, returnDoubleValue2)

	returnDoubleValue4, _ := returnDoubleValue("returnDoubleValue1", "returnDoubleValue2")
	fmt.Println(returnDoubleValue4)

	area := calculateArea(10, 20)
	fmt.Println(area)

}
