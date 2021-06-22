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
	if result != response {
		t.Errorf("downloadOne(%s) = %s, wanted %s", url, result, response)
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
	if result != "" {
		t.Errorf("expected empty result, but was non-empty: %v", result)
	}
}
