package my_http

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintln(os.Stderr, "error request:", err)
			os.Exit(1)
		}

		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Fprintln(os.Stderr, "error read:", err)
			os.Exit(1)
		}

		bodyString := string(bodyBytes)
		resp.Body.Close()
		fmt.Println(bodyString)
	}
}
