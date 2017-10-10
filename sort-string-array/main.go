package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	size := 20

	// prepare data:
	arr := make([]string, size)
	for i := 0; i < size; i++ {
		arr[i] = fmt.Sprintf("s%d", rand.Intn(size))
	}

	fmt.Printf("Before sort: %v\n", arr)

	// sort:
	// 注意：字符串排序默认是按ascii码排序，即 s1 s11 s12 ... s2 s21...
	sort.Strings(arr)
	fmt.Printf("After sort: %v\n", arr)

}
