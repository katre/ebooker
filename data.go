package data

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"google.golang.org/protobuf/encoding/prototext"

	"ebooker/proto"
)

type Book struct {
	Title           string
	Author          string
	Chapters        []*Chapter
	defaultSelector string
}

func NewBook(input proto.Book) *Book {
	return &Book{
		Title:           input.GetTitle(),
		Author:          input.GetAuthor(),
		Chapters:        newChapters(input.GetChapters(), input.GetDefaultSelector()),
		defaultSelector: input.GetDefaultSelector(),
	}
}

func (b Book) Write(dir string) error {
	// Write out the data so far.
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("Unable to create output directory %s: %v", dir, err)
	}

	for i, chap := range b.Chapters {
		for j, cont := range chap.Content() {
			fileName := fmt.Sprintf("c%02d_s%02d.html", i, j)
			outfile := path.Join(dir, fileName)

			err = ioutil.WriteFile(outfile, []byte(cont), 0644)
			if err != nil {
				return fmt.Errorf("Unable to write chapter %d, section %d to file %s: %v", i, j, outfile, err)
			}
		}
	}

	// Write out the updated book data file.
	outfile := path.Join(dir, "book.textproto")
	err = b.writeDataFile(outfile)
	if err != nil {
		return fmt.Errorf("Unable to write book data to %q: %v", outfile, err)
	}

	return nil
}

func (b Book) writeDataFile(name string) error {
	result, err := prototext.Marshal(b.AsProto())
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

func (b Book) AsProto() *proto.Book {
	return &proto.Book{
		Title:           b.Title,
		Author:          b.Author,
		DefaultSelector: b.defaultSelector,
		Chapters:        chaptersAsProtos(b.Chapters),
	}
}

type Chapter struct {
	Name     string
	content  []string
	urls     []string
	selector string
}

func newChapters(chapters []*proto.Chapter, defaultSelector string) []*Chapter {
	var results []*Chapter

	for _, chapter := range chapters {
		selector := chapter.GetSelector()
		if selector == "" {
			selector = defaultSelector
		}
		result := &Chapter{
			Name:     chapter.GetName(),
			urls:     chapter.GetUrl(),
			selector: selector,
		}

		results = append(results, result)
	}

	return results
}

func chaptersAsProtos(chapters []*Chapter) []*proto.Chapter {
	var results []*proto.Chapter

	for _, chapter := range chapters {
		results = append(results, chapter.AsProto())
	}

	return results
}

func (c *Chapter) Urls() []string {
	return c.urls
}

func (c *Chapter) Selector() string {
	return c.selector
}

func (c *Chapter) Content() []string {
	return c.content
}

func (c *Chapter) SetContent(newContent []string) {
	c.content = newContent
}

func (c *Chapter) AsProto() *proto.Chapter {
	return &proto.Chapter{
		Name:     c.Name,
		Selector: c.selector,
		Url:      c.urls,
	}
}
