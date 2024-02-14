package structure

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestHeaderChecksum(t *testing.T) {
	header := NewHeader()
	checksum, err := header.Checksum()
	if err != nil {
		t.Error(err)
	}
	t.Log(checksum)
}

func TestSerialize(t *testing.T) {
	h := NewHeader()
	h.FormatVersion.SetMajor(50)
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, h); err != nil {
		t.Error(err)
	}
	t.Log(buf)
}
