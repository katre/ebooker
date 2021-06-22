package selector

import (
	"strings"

	"github.com/PuerkitoBio/goquery"

	"ebooker/data"
)

func Select(selectables []*data.Chapter) error {
	for _, s := range selectables {
		content, err := selectContent(s.Content(), s.Selector())
		if err != nil {
			return err
		}
		s.SetContent(content)
	}
	return nil
}

// selectContent assumes content is valid HTML markup, and applies selector to return a portion of the content.
// It also scrubs any links from the returned content.
func selectContent(content string, selector string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return "", err
	}
	result := doc.Find(selector)

	// Remove any link tags, replace with the textual content.
	result.Find("a").Each(func(i int, s *goquery.Selection) {
		s.ReplaceWithHtml(s.Text())
	})

	return result.Html()
}
