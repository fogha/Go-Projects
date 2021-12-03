package main

import (
	"fmt"
)

func main() {
	a := 20
	var b *int = &a
	fmt.Println("the address is:", &a)
	fmt.Println("the value is:", a)

	fmt.Println(*b)
	*b = *b + 2
	fmt.Println(*b)
}
