package main

import (
	"fmt"
	"sort"

	"./palindrome"
)

func main() {
	a1 := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	a2 := []int{0, 1, 2, 3, 4}

	fmt.Printf("a1 IsPalindrome = %v\n", palindrome.IsPalindrome(sort.IntSlice(a1)))
	fmt.Printf("a2 IsPalindrome = %v\n", palindrome.IsPalindrome(sort.IntSlice(a2)))
}
