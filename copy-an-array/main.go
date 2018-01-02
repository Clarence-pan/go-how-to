package main

import "fmt"

func main() {
	a := make([]int, 3)
	var b []int

	b = a
	fmt.Printf(" ==== Simple assignment will NOT copy an Array  === \n")
	fmt.Printf("[before]:\n  a: %v\n  b:%v\n", a, b)

	a[0] = 10

	fmt.Printf("[after a[0] = 10]:\n  a: %v\n  b:%v\n", a, b)

	b = make([]int, len(a))
	copy(b, a)
	fmt.Printf(" ==== make new Array and copy(b, a) will work  === \n")
	fmt.Printf("[before]:\n  a: %v\n  b:%v\n", a, b)

	a[1] = 100

	fmt.Printf("[after a[1] = 100:\n  a: %v\n  b:%v\n", a, b)

}
