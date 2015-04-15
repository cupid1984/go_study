package main

import (
	"fmt"
)

func main() {
	i := 19
	f := float64(i)
	u := uint(f)

	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("u is of type %T\n", u)
}