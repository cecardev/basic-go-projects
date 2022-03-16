package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println(m)

	//iterate over map
	for key, value := range m {
		fmt.Println(key, value)
	}

	//find value
	value, ok := m["a"]
	fmt.Println(value, ok)

}
