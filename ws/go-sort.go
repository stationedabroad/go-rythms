package main

import (
	"fmt"
	"math/rand"
	"sort"
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
	A := randomArr(1000000000)
	// A := []int{5,2,4,6,1,3}
	// fmt.Println(A)
	start := time.Now()

	// sort.Slice(A, func(i, j int) bool {
	// 	return A[i] > A[j]
	// })

	sort.Ints(A)
	elapsed := time.Since(start)
	// fmt.Println(A)
	fmt.Printf("Elapsed time: %v for arrays length: %d \n", elapsed.Seconds(), len(A))
}
