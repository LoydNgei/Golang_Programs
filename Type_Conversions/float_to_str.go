package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	var f float64 = 3.1415926535
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(f)

	var s string = strconv.FormatFloat(f, 'E', -1, 32)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)
}

// Sprintf method. Formats according to a format specifier and returns the resulting string

package main

import (
	"fmt"
	"reflect"
)

func main() {
	b := 12.454
	fmt.Println(reflect.TypeOf(b))

	s := fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}