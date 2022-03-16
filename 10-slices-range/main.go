package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	// your code goes here
	s = strings.ToLower(s)
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	for i, value := range slice {
		fmt.Println(i, value)
	}
	fmt.Println(isPalindrome("racecar"))
	fmt.Println(isPalindrome("Ama"))
	fmt.Println(isPalindrome("crazy"))
}
