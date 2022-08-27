package textfilter_test

import (
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nobishino/textfilter"
)

func TestReader(t *testing.T) {
	testcases := []struct {
		title  string
		remove string
		src    string
		want   string
	}{
		{"remove CR", "\r", "Hello\r\nWorld\r\n", "Hello\nWorld\n"},
		{"remove Hello", "Helo", "Hello\r\nWorld\r\n", "\r\nWrd\r\n"},
		{"handle japanese", "\rは", "こんにちは\r\nせかい\r\n", "こんにち\nせかい\n"},
	}
	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			src := strings.NewReader(tc.src)
			r := textfilter.NewReader(src, []rune(tc.remove)...)
			var got strings.Builder
			if _, err := io.Copy(&got, r); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.want, got.String()); diff != "" {
				t.Errorf("result does not match. (-want, +got):\n%s", diff)
			}
		})
	}
}
