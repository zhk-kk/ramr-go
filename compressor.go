package ramrgo

import (
	"compress/gzip"
	"io"
)

type Compressor interface {
	Compress(r io.Reader, w io.Writer) error
	Decompress(r io.Reader, w io.Writer) error
}

type emptyCompressor struct{}

func NewEmptyCompressor() emptyCompressor                           { return emptyCompressor{} }
func (c emptyCompressor) Compress(r io.Reader, w io.Writer) error   { return nil }
func (c emptyCompressor) Decompress(r io.Reader, w io.Writer) error { return nil }

type gzipCompressor struct{}

func NewGzipCompressor() gzipCompressor {
	return gzipCompressor{}
}

func (c gzipCompressor) Compress(r io.Reader, w io.Writer) error {
	gw := gzip.NewWriter(w)
	var buf []byte
	if _, err := r.Read(buf); err != nil {
		return err
	}
	gw.Write(buf)
	gw.Close()
	return nil
}

func (c gzipCompressor) Decompress(r io.Reader, w io.Writer) error {
	gr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}

	if _, err := io.Copy(w, gr); err != nil {
		return err
	}

	gr.Close()

	return nil
}
