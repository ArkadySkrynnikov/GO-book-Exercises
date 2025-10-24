package main
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// go run '.\Chapter 1\1-10\e1.10.go' https://jsonplaceholder.typicode.com/posts https://jsonplaceholder.typicode.com/users https://jsonplaceholder.typicode.com/comments https://jsonplaceholder.typicode.com/albums

func main() {
	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer file.Close()

	fetchAll(file)
	fetchAll(file)
}

func fetchAll(file *os.File) {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for i, arg := range os.Args[1:] {
		strings := strings.Split(arg, "/")

		fmt.Fprintln(file, i+1, strings[len(strings)-2], "\n", <-ch)
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

	written, err := io.Copy(io.Discard, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, written, url)
	resp.Body.Close()
}
