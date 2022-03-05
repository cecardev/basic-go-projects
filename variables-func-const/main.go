package main

import "fmt"

func main() {
	//Constants
	const pi float64 = 3.14
	const pi2 = 3.15
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Variables
	base := 10
	var height int = 5
	var area int
	area = base * int(height)
	fmt.Println("area:", area)

	//Zeros
	//When a variable is just declared, it is initialized to zero or equivalent.
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("a:", a)
	//0
	fmt.Println("b:", b)
	//0.0
	fmt.Println("c:", c)
	//""
	fmt.Println("d:", d)
	//false

	//Primitives

	//Integers
	//int, int8, int16, int32, int64
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//UIntegers > 0
	//uint, uint8, uint16, uint32, uint64
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//Floats
	//float32, float64
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//Text
	//string

	//Booleans
	//bool

	//Complex
	//complex64, complex128
	//Complex64 = each float32
	//Complex128 = each float64

}
