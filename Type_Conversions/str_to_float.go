package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	strVar := "200.89303657"

	floatVar, err := strconv.ParseFloat(strVar, 8)

	fmt.Println(floatVar, err, reflect.TypeOf(floatVar))
}