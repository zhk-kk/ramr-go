package structure

import (
	"bytes"
	"encoding/binary"

	"github.com/cespare/xxhash/v2"
)

const Signature uint64 = 0xFCEA45B372616D72

type HeapId uint8

// Exactly 512 bytes
type Header struct {
	signature       uint64
	FormatVersion   formatVersion
	ArchiveUID      uint64
	Flags           headerFlags
	AmountOfEntries uint64
	AmountOfHeaps   uint64
	_               [3]uint64
}

func NewHeader() Header {
	return Header{
		signature:       Signature,
		FormatVersion:   NewFormatVersion(1, 1),
		ArchiveUID:      NewArchiveUID(),
		Flags:           NewHeaderFlags(),
		AmountOfEntries: 0,
		AmountOfHeaps:   0,
	}
}

func (h *Header) Checksum() (uint64, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, h); err != nil {
		return 0, err
	}
	digest := xxhash.New()
	digest.Write(buf.Bytes())
	return digest.Sum64(), nil
}

func NewArchiveUID() uint64 {
	return 0
}

type headerFlags uint32

func NewHeaderFlags() headerFlags {
	return 0
}

type formatVersion uint32

func NewFormatVersion(major uint16, minor uint16) formatVersion {
	return formatVersion((uint32(major) << 16) | uint32(minor))
}

func (f *formatVersion) Minor() uint16 {
	return uint16(*f)
}

func (f *formatVersion) Major() uint16 {
	return uint16(*f >> 16)
}

func (f *formatVersion) SetMinor(value uint16) *formatVersion {
	*f = formatVersion(uint32(*f&0x11110000) | uint32(value))
	return f
}

func (f *formatVersion) SetMajor(value uint16) *formatVersion {
	*f = formatVersion(uint32(*f&0x00001111) | (uint32(value) << 16))
	return f
}

type centralIndexEntry struct {
	Flags              centralIndexEntryFlags
	EntryName          string
	HeapId             HeapId
	UncompressedOffset uint64
}

func NewCentralIndexEntry(entryName string) centralIndexEntry {
	return centralIndexEntry{
		Flags:     NewCentralIndexEntryFlags(),
		EntryName: entryName,
	}
}

type centralIndexEntryFlags uint16

func NewCentralIndexEntryFlags() centralIndexEntryFlags {
	return 0
}

type reservationIndexEntry struct {
	Flags             reservationIndexEntryFlags
	ReservationOffset uint64
	ReservationSize   uint64
}

func NewReservationIndexEntry() reservationIndexEntry {
	return reservationIndexEntry{
		Flags: NewReservationIndexEntryFlags(),
	}
}

type reservationIndexEntryFlags uint16

func NewReservationIndexEntryFlags() reservationIndexEntryFlags {
	return 0
}

type tableOfContents struct {
	Flags       tableOfContentsFlags
	EntryName   string
	EntryAmount uint16
	Index       []tocIndexEntry
}

func NewTableOfContents() tableOfContents {
	return tableOfContents{
		Flags: 0,
	}
}

type tableOfContentsFlags uint8

type tocIndexEntry struct {
	MetadataId         tocEntryId
	HeapId             HeapId
	UncompressedOffset uint64
	UncompressedSize   uint64
}

func NewTocIndexEntry() tocIndexEntry {
	return tocIndexEntry{}
}

type tocEntryId uint16
