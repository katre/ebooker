package downloader

import (
	"fmt"
	"io"
	"net/http"
)

func Download(urls []string) (map[string]string, error) {
	results := make(map[string]string)
	for _, url := range urls {
		content, err := downloadOne(url)
		if err != nil {
			return nil, err
		}
		results[url] = content
	}

	return results, nil
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
