package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// 1.7, 1.8

func main() {
	for _, url := range os.Args[1:] {
		// FIXME: https対応
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
