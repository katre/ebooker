package maker

import (
	epub "github.com/bmaupin/go-epub"

	"ebooker/data"
)

func Make(book *data.Book, filename string) error {
	e := epub.NewEpub(book.Title)
	e.SetAuthor(book.Author)

	for _, chapter := range book.Chapters {
		e.AddSection(chapter.Content(), chapter.Name, "", "")
	}

	return e.Write(filename)
}
