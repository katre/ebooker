package downloader

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"google.golang.org/protobuf/encoding/prototext"

	"ebooker/data"
	"ebooker/proto"
	"ebooker/selector"
)

func DownloadBook(filename string) (*data.Book, error) {
	book, err := readDataFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Unable to read book data from %q: %v", filename, err)
	}

	// Actually download the book files.
	fmt.Printf("Downloading %s, by %s\n", book.Title, book.Author)

	err = Download(book.Chapters)
	if err != nil {
		return nil, fmt.Errorf("Unable to download book: %v", err)
	}
	err = selector.Select(book.Chapters)
	if err != nil {
		return nil, fmt.Errorf("Unable to download book: %v", err)
	}

	return book, nil
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

func Download(downloadables []*data.Chapter) error {
	for _, d := range downloadables {
		var content []string
		for _, url := range d.Urls() {
			c, err := downloadOne(url)
			if err != nil {
				return err
			}
			content = append(content, c)
		}
		d.SetContent(content)
	}
	return nil
}

func downloadOne(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Invalid response code %s", resp.Status)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
