package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[0]
	fmt.Println("starting crawl of:", baseURL)

	pages := make(map[string]int)

	crawlPage(baseURL, baseURL, pages)
	fmt.Println(pages)
}
