package main

import "fmt"

func main() {
	k := 1
	for ; k <= 10; k++ {
		fmt.Println(k)
	}

	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}

	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}
}

package main

import "fmt"

func main() {
	i := 5
	for {
		fmt.Println("Hello")
		if i == 10 {
			break
		}
		i++
	}
}