package selector

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSelectContent(t *testing.T) {
	type test struct {
		name     string
		content  string
		selector string
		want     string
	}

	testCases := []test{
		{
			name: "basic",
			content: `
      <html><body>boilerplate<div class="post-content">This is the actual stuff</div></html>
      `,
			selector: "div.post-content",
			want:     "This is the actual stuff",
		},
		{
			name: "missing",
			content: `
      <html><body>boilerplate<div>This is the actual stuff</div></html>
      `,
			selector: "div.post-content",
			want:     "",
		},
		{
			name: "with-markup",
			content: `
      <html><body>boilerplate<div class="post-content">This is the <i>actual</i> stuff</div></html>
      `,
			selector: "div.post-content",
			want:     "This is the <i>actual</i> stuff",
		},
		{
			name: "with-link",
			content: `
      <html><body>boilerplate<div class="post-content">This is the <a href="example.com">actual</a> stuff</div></html>
      `,
			selector: "div.post-content",
			want:     "This is the actual stuff",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := selectContent(tc.content, tc.selector)
			if err != nil {
				t.Fatalf("SelectContent(%s) had unexpected error: %v", tc.name, err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("SelectContent(%s) had unexpected diff (-want +got): %s", tc.name, diff)
			}
		})
	}
}
