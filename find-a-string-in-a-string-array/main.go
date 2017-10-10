package main

import (
	"fmt"
	"sort"
)

func main() {

	var studentNamesPool = []string{
		"Jim",
		"Tom",
		"Lee",
		"Jack",
	}

	found := sort.SearchStrings(studentNamesPool, "Lee")

	fmt.Printf("found: %d\n", found)
}
