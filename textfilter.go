package textfilter

import (
	"bufio"
	"io"
	"unicode/utf8"
)

// Reader is an io.Reader which discards specified runes from the source reader.
type Reader struct {
	buf    *bufio.Reader
	remove []rune
}

// Read implements io.Reader.
func (rd *Reader) Read(p []byte) (int, error) {
	w := p[:0]
	var written int
	for written < len(p) {
		r, n, err := rd.buf.ReadRune()
		if err != nil {
			return written, err
		}
		if rd.shouldSkip(r) {
			continue
		}
		if written+n > len(p) {
			return written, nil
		}
		w = utf8.AppendRune(w, r)
		written += n
	}
	return written, nil
}

func (rd *Reader) shouldSkip(r rune) bool {
	for _, rm := range rd.remove {
		if rm == r {
			return true
		}
	}
	return false
}

// NewReader returns a new textfilter.Reader which discards runes in 'remove'.
func NewReader(src io.Reader, remove ...rune) io.Reader {
	buf := bufio.NewReader(src)
	return &Reader{
		buf:    buf,
		remove: remove,
	}
}
