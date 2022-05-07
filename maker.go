package maker

import (
	"strings"

	epub "github.com/bmaupin/go-epub"

	"ebooker/data"
)

func Make(book *data.Book, filename string) error {
	e := epub.NewEpub(book.Title)
	e.SetAuthor(book.Author)

	for _, chapter := range book.Chapters {
		data := strings.Join(chapter.Content(), "\n<br />\n")
		e.AddSection(data, chapter.Name, "", "")
	}

	return e.Write(filename)
}
