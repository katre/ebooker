package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/encoding/prototext"

	"ebooker/data"
	"ebooker/downloader"
	"ebooker/proto"
	"ebooker/selector"
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

	for _, chapter := range book.Chapters {
		fmt.Printf("Read chapter %s, content: %s\n", chapter.Name, shorten(chapter.Content(), 200))
	}

	// Generate output

	return nil
}

func getAllUrls(chapters []*proto.Chapter) []string {
	var urls []string
	for _, chapter := range chapters {
		urls = append(urls, chapter.GetUrl())
	}
	return urls
}

func shorten(input string, length int) string {
	if len(input) < length {
		return input
	}
	return input[0:length] + "..."
}
