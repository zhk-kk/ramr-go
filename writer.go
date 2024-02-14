package ramrgo

import (
	"errors"
	"io"
)

var (
	errPayloadAlreadyAttached = errors.New("ramr: payload is already attached")
)

type Writer struct {
	w            io.Writer
	closed       bool
	entryHandles map[string]*entryHandle
	compressors  map[uint16]Compressor
	options      WriterOptions
}

type WriterOptions struct {
}

// NewWriter returns a new Writer writing a ramr file to w.
func NewWriter(options WriterOptions, w io.Writer) *Writer {
	return &Writer{
		w:            w,
		closed:       false,
		entryHandles: make(map[string]*entryHandle),
		compressors:  make(map[uint16]Compressor),
		options:      options,
	}
}

// LoadDefaultCompressors loads the default ramr compressors.
func (w *Writer) LoadDefaultCompressors() {
	w.compressors[0] = NewEmptyCompressor()
	w.compressors[1] = NewGzipCompressor()
}

// Close finishes writing ramr file by writing the whole archive.
// It does not close the underlying writer.
func (w *Writer) Close() error {
	return errors.New("Close() is not yet implemented")
}

// AttachEntries adds entries to a writer.
func (w *Writer) AttachEntries(entries ...*entryHandle) []error {
	for _, e := range entries {
		if err := e.Complete(); err != nil {
			return err
		}

		e.w = w
		w.entryHandles[e.Name()] = e
	}

	return nil
}

// NewEntry returns an entry handle, which is used to attach data to.
func NewEntry(name string) *entryHandle {
	e := &entryHandle{isComplete: false, name: name}
	return e
}

type entryHandle struct {
	w          *Writer
	name       string
	payload    io.Reader
	metadata   []io.Reader
	isComplete bool
	errs       []error
}

func (e entryHandle) Name() string {
	return e.name
}

// AttachPayload adds a payload to an entry, turning it from default directory entry into a file one.
func (e *entryHandle) AttachPayload(r io.Reader) *entryHandle {
	if e.payload != nil {
		e.errs = append(e.errs, errPayloadAlreadyAttached)
	}
	e.payload = r
	return e
}

// AttachMetadata adds metadata to an entry.
func (e *entryHandle) AttachData(r io.Reader) *entryHandle {
	e.metadata = append(e.metadata, r)
	return e
}

// Complete finishes the creation process of an entry. Usually does **not** need to be called before using Writer.AttachEntries()
func (e *entryHandle) Complete() []error {
	if len(e.errs) != 0 {
		return e.errs
	}
	e.isComplete = true
	return nil
}
