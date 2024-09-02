// String to integer in GO

	// Use the strconv package Atoi() function. The Atoi() function returns 2 results: Result of conversion & error if any

		package main

		import (
			"fmt"
			"strconv"
			"reflect"
		)

		func main() {
			strVar := "100"

			intVar, err := strconv.Atoi(strVar)
			fmt.Println(intVar, err, reflect.TypeOf(intVar))
		}


	// Using parseint() function. It interprets a string in a given base(0, 2 to 36)
							  //  and bit size(0 to 64) and returns corresponding value i

			
	package main

	import (
		"fmt"
		"reflect"
		"strconv"
	)


	func main() {
		strVar := "100"

		intVar, err := strconv.ParseInt(strVar, 0, 8)
		fmt.Println(intVar, err, reflect.TypeOf(intVar))

		intVar, err = strconv.ParseInt(strVar, 0, 16)
		fmt.Println(intVar, err, reflect.TypeOf(intVar))
	}

	// Using fmt.Sscan. Scans the string argument and store into variables
	
	package main

	import(
		"fmt"
		"reflect"
	)

	func main() {
		strVar := "100"

		intValue := 0
		_, err := fmt.Sscan(strVar, &intValue)
		fmt.Println(intValue, err, reflect.TypeOf(intValue))
	}
