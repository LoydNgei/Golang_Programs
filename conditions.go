package main

import (
	"fmt"
)

func main() {
	x := 100
	if x == 50 {
		fmt.Println("Germany")
	} else if x == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}

	if x := 100; x == 100 {
		fmt.Println("Germany")
	}
}

