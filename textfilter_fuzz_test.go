package textfilter_test

import (
	"io"
	"strings"
	"testing"

	"github.com/nobishino/textfilter"
)

func FuzzReader(f *testing.F) {
	testcases := []struct {
		in     string
		remove string
	}{
		{"Hello\nWorld\n", "ol\n"},
	}
	for _, tc := range testcases {
		f.Add(tc.in, tc.remove)
	}
	f.Fuzz(func(t *testing.T, in, remove string) {
		src := strings.NewReader(in)
		r := textfilter.NewReader(src, []rune(remove)...)
		var got strings.Builder
		_, err := io.Copy(&got, r)
		if err != nil {
			t.Fatal(t)
		}
	})
}
