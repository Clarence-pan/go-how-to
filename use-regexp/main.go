package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "Today is Monday."

	// just test if match
	matched, err := regexp.MatchString(`day`, s)
	fmt.Printf("matched: %v, error: %v\n", matched, err)
	// => matched: true, error: <nil>

	// find all matched
	re, err := regexp.Compile(`(Mon)?day`)
	if err != nil {
		panic(err)
	}

	found := re.FindAllString(s, -1)
	fmt.Printf("found: %v\n", found)
	// => found: [day Monday]

	// find all matched and its groups
	found2 := re.FindAllStringSubmatch(s, -1)
	fmt.Printf("found: %v\n", found2)
	// => found: [[day ] [Monday Mon]]
}
