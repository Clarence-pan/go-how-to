package main

import "fmt"

func main() {
	// NOTE: [3]int{1,2,3} is type of Array [3]int
	//       BUT []int{1,2,3} is type of slice []int
	a := []int{1, 2, 3}
	var b []int

	/////////////////////////////////////////////////////////////////////////////
	b = a
	fmt.Printf(" ==== Simple assignment will NOT copy a Slice  === \n")
	fmt.Printf("[before]:\n  a: %v\n  b:%v\n", a, b)

	a[0] = 10

	fmt.Printf("[after a[0] = 10]:\n  a: %v\n  b: %v\n", a, b)

	/////////////////////////////////////////////////////////////////////////////
	b = make([]int, len(a))
	copy(b, a)
	fmt.Printf(" ==== make new Slice and copy(b, a) will work  === \n")
	fmt.Printf("[before]:\n  a: %v\n  b: %v\n", a, b)

	a[0] = 20

	fmt.Printf("[after a[0] = 20:\n  a: %v\n  b: %v\n", a, b)

	/////////////////////////////////////////////////////////////////////////////
	b = append([]int{}, a...)
	fmt.Printf(" ==== a more brief way is using append -- `b = append([]int, a...)`  === \n")
	fmt.Printf("[before]:\n  a: %v\n  b: %v\n", a, b)

	a[0] = 100

	fmt.Printf("[after a[0] = 100:\n  a: %v\n  b: %v\n", a, b)

}

/* output:
 ==== Simple assignment will NOT copy a Slice  ===
[before]:
  a: [1 2 3]
  b:[1 2 3]
[after a[0] = 10]:
  a: [10 2 3]
  b: [10 2 3]
 ==== make new Slice and copy(b, a) will work  ===
[before]:
  a: [10 2 3]
  b: [10 2 3]
[after a[0] = 20:
  a: [20 2 3]
  b: [10 2 3]
 ==== a more brief way is using append -- `b = append([]int, a...)`  ===
[before]:
  a: [20 2 3]
  b: [20 2 3]
[after a[0] = 100:
  a: [100 2 3]
  b: [20 2 3]
*/
