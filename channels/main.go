package main

import "fmt"

func main() {
	links := []string{
		"http://google.com.co",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		fmt.Println(link)
	}
}
