package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	WhatEchoIsFaster()
}

func EchoEfficientJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func EchoInefficientLoop() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)
}

func WhatEchoIsFaster() {

	//region EchoFromZeroTime
	start := time.Now()
	EchoEfficientJoin()
	fmt.Printf("Echo from efficient join (microseconds: %v)", time.Since(start).Microseconds())
	//endregion

	//region EchoInefficientTime
	start = time.Now()
	EchoInefficientLoop()
	fmt.Printf("Echo Inefficient loop (microseconds: %v)", time.Since(start).Microseconds())
	//endregion

	//result EchoFromZero is minimum 2 times faster than EchoInefficient
}
