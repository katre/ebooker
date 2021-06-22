package downloader

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

const url = "http://example.com/chapter/1"
const response = "Example http content"

func TestDownloadOne(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, response))

	result, err := downloadOne(url)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Url != url {
		t.Errorf("downloadOne(%s).Url = %s, wanted %s", url, result.Url, url)
	}
	if result.Content != response {
		t.Errorf("downloadOne(%s).Content = %s, wanted %s", url, result.Content, response)
	}
}

func TestDownloadOne_missing(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(404, "Not found"))

	result, err := downloadOne(url)

	if err == nil {
		t.Fatal("Expected error, but did not receive one")
	}
	if result != nil {
		t.Errorf("expected nil result, but was non-nil: %v", result)
	}
}
