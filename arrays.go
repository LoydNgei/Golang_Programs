package main

import ("fmt")

func main() {
	var theArray [3]string

	theArray[0] = "India"
	theArray[1] = "Canada"
	theArray[2] = "Japan"

	fmt.Println(theArray[0])
}


// Loop through indexed array

package main

import ("fmt")

func main() {
	intArray := [5]int{10, 20, 30, 40, 50}

	fmt.Println("\n----Example 1-----")
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	fmt.Println("\n----Example 2-----")
	for index, element := range intArray {
		fmt.Println(index, "=>", element)
	}

	fmt.Println("\n----Example 4-----")
	for _, value := range intArray {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("\n-----Example 4------")
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}






package main

import (
	"fmt"
	"reflect"
)

func main() {
	strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strArray, "Canada"))
	fmt.Println(itemExists(strArray, "Africa"))
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}