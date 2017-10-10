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

// in order to use sort.Sort, sort.Interface must be implemented
type studentList []student

// ensure studentList implements sort.Interface
var _ sort.Interface = make(studentList, 0)

func (list studentList) Len() int {
	return len(list)
}

func (list studentList) Less(i, j int) bool {
	if list[i].age < list[j].age {
		return true
	} else if list[i].age > list[j].age {
		return false
	}

	return list[i].name < list[j].name
}

func (list studentList) Swap(i, j int) {
	t := list[i]
	list[i] = list[j]
	list[j] = t
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
	arr := make(studentList, size)
	for i := 0; i < size; i++ {
		arr[i].age = 10 + rand.Intn(10)
		arr[i].name = studentNamesPool[rand.Intn(len(studentNamesPool))]
	}

	fmt.Printf("Before sort: %v\n", arr)

	// sort:
	// 注意：字符串排序默认是按ascii码排序，即 s1 s11 s12 ... s2 s21...
	sort.Sort(arr)
	fmt.Printf("After sort: %v\n", arr)

}
