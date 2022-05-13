package main

import (
	"bufio"
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
var output = flag.String("output", "-", "The file to write the result file to.")
var outDir = flag.String("out_dir", ".", "The directory to store output in.")

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
	for i, chap := range book.Chapters {
		for j, cont := range chap.Content() {
			fileName := fmt.Sprintf("c%02d_s%02d.html", i, j)
			outfile := path.Join(*outDir, fileName)

			err = ioutil.WriteFile(outfile, []byte(cont), 0644)
			if err != nil {
				fmt.Printf("Unable to write chapter %d, section %d to file %s: %v", i, j, outfile, err)
				return
			}
		}
	}

	// Write out the updated book data file.
	err = writeDataFile(book, *output)
	if err != nil {
		fmt.Printf("Unable to write book data to %q: %v\n", *output, err)
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

func openOutputFile(name string) (io.Writer, error) {
	if name == "-" {
		return os.Stdout, nil
	}

	return os.Open(name)
}

func writeDataFile(book *data.Book, name string) error {
	w, err := openOutputFile(name)
	if err != nil {
		return err
	}

	result, err := prototext.Marshal(book.AsProto())
	if err != nil {
		return err
	}

	bw := bufio.NewWriter(w)
	n, err := bw.Write(result)
	if err != nil {
		return err
	}
	if n != len(result) {
		return fmt.Errorf("Didn't write entire output!")
	}
	bw.Flush()
	return nil
}
