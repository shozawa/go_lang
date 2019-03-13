package main

import (
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"os"
	"strconv"
	"time"
)

func main() {
	// 1.10 キャッシュの確認
	fetchAll()
	fetchAll()
}

func fetchAll() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	dst, err := os.Create(neturl.PathEscape(url + strconv.FormatInt(time.Now().Unix(), 10)))
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	nbytes, err := io.Copy(dst, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d %s", secs, nbytes, url)
}
