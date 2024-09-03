// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var x, y = 35, 7

// 	fmt.Printf("x + y = %d\n", x + y)
// 	fmt.Printf("x - y = %d\n", x - y)
// 	fmt.Printf("x * y = %d\n", x * y)
// 	fmt.Printf("x / y = %d\n", x / y)
// 	fmt.Printf("x mod y = %d\n", x % y)
// }


package main

import ("fmt")

func main() {
	var x uint = 9
	var y uint = 65
	var z uint

	z = x & y
	fmt.Println("x & y =", z)

	z = x | y
	fmt.Println("x | y =", z)

	z = x ^ y
	fmt.Println("x ^ y =", z)

	z = x << 1
	fmt.Println("x << 1 =", z)

	z = x >> 1
	fmt.Println("x >> 1=", z)
}
