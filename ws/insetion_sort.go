package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomArr(size int) []int {
	A := make([]int, size)
	for i := range A {
		A[i] = rand.Intn(100)
	}
	return A
}

func main() {
	// array creation removed from perf timing
	A := randomArr(1000000)
	// A := [...]int{5,2,4,6,1,3}
	// fmt.Println(A)
	start := time.Now()

	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i -= 1
		}
		A[i+1] = key
	}

	elapsed := time.Since(start)
	// fmt.Println(A)
	fmt.Printf("Elapsed time: %v for arrays length: %d \n", elapsed.Seconds(), len(A))
}
