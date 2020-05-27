package main

import (
	"fmt"
)

func BinarySort(needle int, haystack []int) int {
	left := 0
	right := len(haystack) - 1

	for left <= right {
		mid := (left+right) / 2
		if haystack[mid] < needle {
			left = mid + 1
		} else if haystack[mid] > needle {
			right = mid - 1
		} else { return mid }
	}
	return -1
}

func main() {
	var A = []int{-1, 1,2,3,4,5,6,7, 10}
	fmt.Printf("%v\n", BinarySort(12, A))
}