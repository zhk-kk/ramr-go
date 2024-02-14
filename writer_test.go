package ramrgo

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriter(t *testing.T) {
	b := bytes.NewBuffer(nil)
	w := NewWriter(WriterOptions{}, b)
	if err := w.AttachEntries(
		NewEntry("dir1/nested/file.txt").
			AttachPayload(strings.NewReader("Hey there")),
		NewEntry("root-file.txt").
			AttachPayload(strings.NewReader("Hey there")),
	); err != nil {
		t.Fatalf("%s\n", err)
	}
	if err := w.Close(); err != nil {
		t.Fatalf("%s\n", err)
	}
	t.Logf("bytes-out: %v\n", b)
}
