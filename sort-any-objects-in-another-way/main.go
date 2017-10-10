package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type student struct {
	age  int
	name string
}

var studentNamesPool = []string{
	"Jim",
	"Tom",
	"Lee",
	"Jack",
}

func main() {
	size := 5

	// prepare data:
	arr := make([]student, size)
	for i := 0; i < size; i++ {
		arr[i].age = 10 + rand.Intn(10)
		arr[i].name = studentNamesPool[rand.Intn(len(studentNamesPool))]
	}

	fmt.Printf("Before sort: %v\n", arr)

	// sort:
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].age < arr[j].age {
			return true
		} else if arr[i].age > arr[j].age {
			return false
		}

		return arr[i].name < arr[j].name
	})

	fmt.Printf("After sort: %v\n", arr)

}
