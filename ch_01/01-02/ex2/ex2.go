package main

import (
	"fmt"
	"os"
)

func main() {
	EchoIndexAndValue()
}

func EchoIndexAndValue() {
	for i, val := range os.Args[1:] {
		fmt.Printf("%v : %v\n", i, val)
	}
}
