package main

import (
	"flag"
	"fmt"
	"regexp"

	"ebooker/data"
	"ebooker/downloader"
	"ebooker/maker"
)

// Define flags

var input = flag.String("input", "-", "The input textproto to process.")
var dir = flag.String("dir", ".", "The directory to store output in.")

var download = flag.Bool("download", false, "Whether to download the content or assume it already exists.")
var generate = flag.Bool("generate", false, "Whether to generate the ebook.")

var textprotoRe = regexp.MustCompile(`\.textproto$`)

func main() {
	flag.Parse()

	fmt.Println("ebooker starting...")

	book, err := getBook()
	if err != nil {
		fmt.Printf("Error reading book data: %v\n", err)
		return
	}

	fmt.Printf("Processing book %s (by %s)\n", book.Title, book.Author)

	if *generate {
		// Generate output
		output := textprotoRe.ReplaceAllString(*input, ".epub")
		err := maker.Make(book, output)
		if err != nil {
			fmt.Printf("Error generating book data: %v\n", err)
			return
		}
	} else {
		err = book.Write(*dir)
		if err != nil {
			fmt.Printf("Error writng book data to %s: %v\n", *dir, err)
			return
		}
		fmt.Printf("Wrote data for %s to %s\n", book.Title, *dir)
	}
}

func getBook() (*data.Book, error) {
	if *download {
		book, err := downloader.DownloadBook(*input)
		if err != nil {
			return nil, err
		}

		return book, nil
	}

	// Assume the book data is in dir.
	return data.ReadBook(*dir)
}
