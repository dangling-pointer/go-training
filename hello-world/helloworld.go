package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello everyone")
	callthis()

	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func callthis() {
	fmt.Println("This is from a function call")
}
