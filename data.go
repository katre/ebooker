package data

import (
	"ebooker/proto"
)

type Book struct {
	Title    string
	Author   string
	Chapters []*Chapter
}

func NewBook(input proto.Book) *Book {
	return &Book{
		Title:    input.GetTitle(),
		Author:   input.GetAuthor(),
		Chapters: newChapters(input.GetChapters(), input.GetDefaultSelector()),
	}
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

func (b Book) AsProto() *proto.Book {
	p := &proto.Book{
		Title:  b.Title,
		Author: b.Author,
		// defaultSelector
		// Chapters
	}

	return p
}

type Chapter struct {
	Name     string
	content  []string
	urls     []string
	selector string
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
