// Using the FormatBool function

	// package main

	// import (
	// 	"fmt"
	// 	"reflect"
	// 	"strconv"
	// )

	// func main() {
	// 	b := true
	// 	fmt.Println(reflect.TypeOf(b))

	// 	s := strconv.FormatBool(true)
	// 	fmt.Println(reflect.TypeOf(s))
	// }



// Using the Sprintf() method. It formats according to a format specifier and returns the resulting string

package main

import (
	"fmt"
	"reflect"
)

func main() {
	b := true
	fmt.Println(reflect.TypeOf(b))
	s := fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}