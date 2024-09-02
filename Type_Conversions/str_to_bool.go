package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "true"

	f, _ := strconv.ParseBool(s)
	fmt.Printf("%T, %v\n", f, f)

	x := "F"

	y, _ := strconv.ParseBool(x)
	fmt.Printf("%T, %v\n", y, y)
}

// ParseBool returns the boolean value represented by the string.
//  It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.