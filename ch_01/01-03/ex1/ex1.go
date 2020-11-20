package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dup3()
}

func dup2() {
	counts := make(map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		//ToDo: implement
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}

func dup3() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "%v", err)
			if err != nil {
				fmt.Println("Where I am running, i can not even write lines in Stderr?!!")
			}
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
