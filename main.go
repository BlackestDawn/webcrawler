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
}
