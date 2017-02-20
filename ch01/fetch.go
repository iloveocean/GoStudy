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
		if !strings.HasPrefix(url, "http://") {
			fmt.Printf("url: %s has no http:// was found, added one to it\n", url)
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Http Status is: %s\n", resp.Status)
		bytes, err := io.Copy(os.Stdout, resp.Body) // io.Copy has much less memory consumption than ioutil.ReadAll
		//bytes, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("total bytes read %d\n", bytes)
	}
}
