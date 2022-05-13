package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"google.golang.org/protobuf/encoding/prototext"

	"ebooker/data"
	"ebooker/downloader"
	"ebooker/proto"
)

// Define flags

var input = flag.String("input", "-", "The input textproto to process.")
var dir = flag.String("dir", ".", "The directory to store output in.")

func main() {
	flag.Parse()

	fmt.Println("downloader starting...")

	book, err := readDataFile(*input)
	if err != nil {
		fmt.Printf("Unable to read book data from %q: %v\n", *input, err)
		return
	}

	// Actually download the book files.
	fmt.Printf("Downloading %s, by %s\n", book.Title, book.Author)

	err = downloader.Download(book.Chapters)
	if err != nil {
		fmt.Printf("Unable to download book: %v\n", err)
		return
	}

	// Write out the data so far.
	err = os.MkdirAll(*dir, 0755)
	if err != nil {
		fmt.Printf("Unable to create output directory %s: %v\n", *dir, err)

	}
	for i, chap := range book.Chapters {
		for j, cont := range chap.Content() {
			fileName := fmt.Sprintf("c%02d_s%02d.html", i, j)
			outfile := path.Join(*dir, fileName)

			err = ioutil.WriteFile(outfile, []byte(cont), 0644)
			if err != nil {
				fmt.Printf("Unable to write chapter %d, section %d to file %s: %v", i, j, outfile, err)
				return
			}
		}
	}

	// Write out the updated book data file.
	outfile := path.Join(*dir, "book.textproto")
	err = writeDataFile(book, outfile)
	if err != nil {
		fmt.Printf("Unable to write book data to %q: %v\n", outfile, err)
		return
	}
}

func openDataFile(name string) (io.Reader, error) {
	if name == "-" {
		return os.Stdin, nil
	}

	return os.Open(name)
}

func readDataFile(name string) (*data.Book, error) {
	r, err := openDataFile(name)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var input proto.Book
	if err := prototext.Unmarshal(content, &input); err != nil {
		return nil, err
	}

	book := data.NewBook(input)
	return book, nil
}

func writeDataFile(book *data.Book, name string) error {
	result, err := prototext.Marshal(book.AsProto())
	if err != nil {
		return err
	}

	if name == "-" {
		fmt.Printf("%s\n", string(result))
	} else {
		err := ioutil.WriteFile(name, result, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
