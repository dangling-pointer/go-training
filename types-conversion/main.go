package main

import (
	"fmt"
)

var y int

type custom int

var z custom

func main() {
	y = 42
	z = 43

	fmt.Print(y)
	fmt.Printf("\t%T", y)
	fmt.Print("\n", z)
	fmt.Printf("\t%T", z)

	// type conversion
	y = int(z)
	fmt.Print("\nValue of y after conversion: ", y)
	fmt.Printf("\t%T", y)
}
