package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"

	"google.golang.org/protobuf/encoding/prototext"

	"ebooker/data"
	"ebooker/downloader"
	"ebooker/maker"
	"ebooker/proto"
	"ebooker/selector"
)

// Define flags

var textprotoRe = regexp.MustCompile(`\.textproto$`)

func main() {
	flag.Parse()

	fmt.Println("ebooker starting...")

	// Handle inputs
	for _, input := range flag.Args() {
		output := textprotoRe.ReplaceAllString(input, ".epub")
		if err := createBook(input, output); err != nil {
			fmt.Printf("Error building book: %v\n", err)
			return
		}
	}
}

func createBook(inputfile, filename string) error {
	fmt.Printf("Reading book data from %s\n", inputfile)

	// Read the file.
	contents, err := ioutil.ReadFile(inputfile)
	if err != nil {
		return err
	}

	var input proto.Book
	if err := prototext.Unmarshal(contents, &input); err != nil {
		return err
	}

	book := data.NewBook(input)

	fmt.Printf("Processing %s, by %s\n", input.GetTitle(), input.GetAuthor())

	// Download chapters.
	err = downloader.Download(book.Chapters)
	if err != nil {
		return err
	}

	// Get just the correct output.
	err = selector.Select(book.Chapters)

	// Generate output
	return maker.Make(book, filename)
}
