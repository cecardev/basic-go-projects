package main

import "fmt"

func main() {
	//array  no mutables
	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	fmt.Println(array, len(array), cap(array))

	//slice  mutables
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))
	//slice methods
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:])
	fmt.Println(slice[2:4])

	//append
	slice = append(slice, 6)
	fmt.Println(slice)

	//append new list
	newSlice := []int{7, 8, 9}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
