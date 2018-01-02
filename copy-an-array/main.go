package main

import "fmt"

func main() {
	// NOTE: [3]int{1,2,3} is type of Array [3]int
	//       BUT []int{1,2,3} is type of slice []int
	a := [3]int{1, 2, 3}
	var b [3]int

	b = a
	fmt.Printf(" ==== Simple assignment WILL copy an Array  === \n")
	fmt.Printf("[before]:\n  a: %v\n  b:%v\n", a, b)

	a[0] = 10

	fmt.Printf("[after a[0] = 10]:\n  a: %v\n  b:%v\n", a, b)
}

/*  output:
 ==== Simple assignment WILL copy an Array  ===
[before]:
  a: [1 2 3]
  b:[1 2 3]
[after a[0] = 10]:
  a: [10 2 3]
  b:[1 2 3]

*/
