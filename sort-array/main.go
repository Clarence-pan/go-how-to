package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	size := 10

	// prepare data:
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size)
	}

	fmt.Printf("Before sort: %v\n", arr)

	// sort:
	sort.Ints(arr)
	fmt.Printf("After sort: %v\n", arr)

}
