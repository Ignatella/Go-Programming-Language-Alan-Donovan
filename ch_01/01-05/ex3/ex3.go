package main

import (
"fmt"
"io"
"net/http"
"os"
"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		url = "gopl.io"

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		written, err := io.Copy(os.Stdout, resp.Body)

		if written == 0 || err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}

		fmt.Printf("That is request status code: %d", resp.StatusCode)
	}
}
