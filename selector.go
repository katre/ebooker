package selector

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SelectContent assumes content is valid HTML markup, and applies selector to return a portion of the content.
// It also scrubs any links from the returned content.
func SelectContent(content string, selector string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return "", err
	}
	result := doc.Find(selector)

	// Remove any link tags, replace with the textual content.
	result.Find("a").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("inside Each: %d, %v\n", i, s.Text())
		s.ReplaceWithHtml(s.Text())
	})

	return result.Html()
}
