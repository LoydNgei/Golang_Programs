// Is a func that takes a variable number of arguments of a specific type

package main

import "fmt"

func main() {
	variadicExample("red", "blue", "white", "Yellow", "Purple")
}

func variadicExample(s ...string){
	fmt.Println(s[0])
	fmt.Println(s[3])
}