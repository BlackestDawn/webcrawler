package main

import (
	"fmt"
	"os"
)

func main() {
	argNum := len(os.Args)
	if argNum < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if argNum > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseUrl := os.Args[1]

	fmt.Println("starting crawl of:", baseUrl)

	traversedPages := make(map[string]int)
	crawlPage(baseUrl, baseUrl, traversedPages)

	fmt.Println("finished crawling, page hit results:")
	for key, val := range traversedPages {
		fmt.Printf("page %s was hit %d time.\n", key, val)
	}
}
