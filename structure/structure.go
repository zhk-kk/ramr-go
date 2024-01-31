package structure

const Signature uint64 = 0xFCEA45B372616D72

type headerFlags uint32

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

func NewFormatVersionNumber() uint32 {
	return 0x1000_0001
}

// Exactly 512 bytes
type header struct {
	Signature       uint64
	FormatVersion   formatVersion
	ArchiveUID      uint64
	Flags           headerFlags
	AmountOfEntries uint64
	AmountOfHeaps   uint64
	_               [3]uint64
}

func NewHeader() header {
	return header{
		Signature: Signature,
	}
}

func (h *header) Checksum() {

}
