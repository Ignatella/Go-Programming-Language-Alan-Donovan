package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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

	bytes, _ := ioutil.ReadFile("./ch_01/01-06/ex2/sites.txt")

	for _, url := range strings.Split(string(bytes), "\n") {
		url = strings.TrimRight(url, "\r")
		if r.FindString(url) == "" {
			url = "http://" + url
		}

		go fetch(url, ch)
	}

	for index, _ := range strings.Split(string(bytes), "\n") {

		err := ioutil.WriteFile("" + strconv.Itoa(index)+".html", <-ch, 0644)

		if err != nil {
			fmt.Printf("Error occured on file saving.")
		}
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- []byte) {
	resp, err := http.Get(url)

	if err != nil {
		ch <- []byte(err.Error())
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- []byte(err.Error())
		return
	}

	err = resp.Body.Close()

	if err != nil {
		ch <- []byte(err.Error())
		return
	}

	ch <- b
}
