package main

import (
	"fmt"

	"github.com/elliotforbes/athena/port"
)

func main() {
	fmt.Println("A port scanner in go")
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
