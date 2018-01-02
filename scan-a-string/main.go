package main

import "fmt"

func main() {
	fmt.Printf("%#v\n", []byte("Hello, 世界"))

	fmt.Println("=== scan string by rune/char ===")
	for i, c := range "Hello, 世界" {
		fmt.Println(i, c)
	}

	fmt.Println("=== scan string by rune/char 2 ===")
	for i, c := range []rune("Hello, 世界") {
		fmt.Println(i, c)
	}

	fmt.Println("=== scan string by byte ===")
	for i, c := range []byte("Hello, 世界") {
		fmt.Println(i, c)
	}
}

/*output:
[]byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c}
=== scan string by rune/char ===
0 72
1 101
2 108
3 108
4 111
5 44
6 32
7 19990
10 30028
=== scan string by rune/char 2 ===
0 72
1 101
2 108
3 108
4 111
5 44
6 32
7 19990
8 30028
=== scan string by byte ===
0 72
1 101
2 108
3 108
4 111
5 44
6 32
7 228
8 184
9 150
10 231
11 149
12 140
*/
