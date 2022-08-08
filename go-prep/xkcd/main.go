package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No search term given. Using 'walrus'.")
		searchTerm := "walrus"
	} else {
		searchTerm := args[0]
	}

	makeOrUpdate()
}
