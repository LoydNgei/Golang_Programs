package main

import "fmt"

type rectangle struct {
	length float64
	breadth float64
	color string
}


func main() {
	fmt.Println(rectangle{10.34, 755.57, "Yellow"})
}