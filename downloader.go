package downloader

import (
	"fmt"
	"io"
	"net/http"

	"ebooker/data"
)

func Download(downloadables []*data.Chapter) error {
	for _, d := range downloadables {
		content, err := downloadOne(d.Url())
		if err != nil {
			return err
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
