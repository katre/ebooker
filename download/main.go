package main

import (
	"flag"
	"fmt"

	"ebooker/downloader"
)

// Define flags

var input = flag.String("input", "-", "The input textproto to process.")
var dir = flag.String("dir", ".", "The directory to store output in.")

func main() {
	flag.Parse()

	fmt.Println("downloader starting...")

	book, err := downloader.DownloadBook(*input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = book.Write(*dir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
