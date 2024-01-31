package structure

func NewSignature() [2]uint64 {
	return [2]uint64{0x4E46574853594C53, 0x4E465651706C6166}
}

type header struct {
	Signature [2]uint64
}

func NewHeader() header {
	return header{
		Signature: NewSignature(),
	}
}
