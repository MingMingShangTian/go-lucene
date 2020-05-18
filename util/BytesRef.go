package util

type BytesRef struct {
	Bytes []byte
	OffSet int
	Length int
}

func NewBytesRef(bytes []byte, offset int, length int) *BytesRef {
	return &BytesRef{
		Bytes:  bytes,
		OffSet: offset,
		Length: length,
	}
}
