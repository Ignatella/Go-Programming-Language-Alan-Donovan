package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	EchoFromZero()
}

func EchoFromZero() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
