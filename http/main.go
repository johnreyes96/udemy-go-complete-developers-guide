package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp := executeUrl("https://www.google.com")
	lw := logWriter{}

	if _, err := io.Copy(lw, resp.Body); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func executeUrl(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return resp
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
