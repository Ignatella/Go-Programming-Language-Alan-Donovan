package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan []byte)
	r, err := regexp.Compile("\\Ahttps?://")

	if err != nil {
		fmt.Printf("Error while regexp compiling.")
		return
	}

	for _, url := range os.Args[1:] {

		if r.FindString(url) == "" {
			url = "http://" + url
		}

		go fetch(url, ch)
	}
	for index, _ := range os.Args[1:]{
		err := ioutil.WriteFile(strconv.Itoa(index), <-ch, 0644)

		if err != nil {
			fmt.Printf("Error occured on file saving.")
		}
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- []byte) {

	resp, err := http.Get(url)

	if err != nil {
		ch <- []byte(fmt.Sprint("Error on request sending"))
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- []byte(fmt.Sprint("Error on request processing."))
		return
	}

	ch <- b
	err = resp.Body.Close()

	if err != nil {
		ch <- []byte(fmt.Sprint("Error on request processing."))
		return
	}
}
