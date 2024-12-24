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

	htmlBody, err := getHTML(baseUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(htmlBody)
}
