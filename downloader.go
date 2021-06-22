package downloader

import (
	"fmt"
	"io"
	"net/http"
)

type Result struct {
	Url     string
	Content string
}

func (r *Result) String() string {
	return fmt.Sprintf("Result(%s): %s", r.Url, shorten(r.Content, 10))
}

func shorten(input string, length int) string {
	if len(input) < length {
		return input
	}
	return input[0:length] + "..."
}

func Download(urls []string) ([]*Result, error) {
	var results []*Result
	for _, url := range urls {
		result, err := downloadOne(url)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func downloadOne(url string) (*Result, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid response code %s", resp.Status)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Result{
		Url:     url,
		Content: string(content),
	}, nil
}
