package main

import (
	"flag"
	"fmt"

	"ebooker/data"
	"ebooker/downloader"
	//"ebooker/maker"
	//"ebooker/proto"
	//"ebooker/selector"
)

// Define flags

var input = flag.String("input", "-", "The input textproto to process.")
var dir = flag.String("dir", ".", "The directory to store output in.")

var download = flag.Bool("download", false, "Whether to download the content or assume it already exists.")
var generate = flag.Bool("generate", false, "Whether to generate the ebook.")

func main() {
	flag.Parse()

	fmt.Println("ebooker starting...")

	book, err := getBook()
	if err != nil {
		fmt.Printf("Error reading book data: %v\n", err)
		return
	}

	fmt.Printf("Processing book %s (by %s)\n", book.Title, book.Author)
}

func getBook() (*data.Book, error) {
	if *download {
		book, err := downloader.DownloadBook(*input)
		if err != nil {
			return nil, err
		}

		err = book.Write(*dir)
		if err != nil {
			return nil, err
		}

		return book, nil
	}

	// Assume the book data is in dir.
	// blarg
	return nil, fmt.Errorf("Unimplemented")
}

/*
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
*/
