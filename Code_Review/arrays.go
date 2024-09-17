package main

import (
	"fmt"
	"reflect"
)

func main() {
	strArray := [5]string{"george", "Mark", "John", "Austine", "Junior"}
	fmt.Println(itemExists(strArray, "Mark"))
	fmt.Println(itemExists(strArray, "Ken"))
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

