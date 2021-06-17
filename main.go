package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/encoding/prototext"

	"ebooker/data"
)

// Define flags

func main() {
	flag.Parse()

	fmt.Println("ebooker starting...")

	// Handle inputs
	for _, input := range flag.Args() {
		if err := createBook(input); err != nil {
			fmt.Printf("Error building book: %v\n", err)
			return
		}
	}
}

func createBook(inputfile string) error {
	fmt.Printf("Reading book data from %s\n", inputfile)

	// Read the file.
	contents, err := ioutil.ReadFile(inputfile)
	if err != nil {
		return err
	}

	var input data.Book
	if err := prototext.Unmarshal(contents, &input); err != nil {
		return err
	}

	fmt.Printf("Processing %s, by %s\n", input.GetTitle(), input.GetAuthor())

	// Download chapters

	// Generate output

	return nil
}
