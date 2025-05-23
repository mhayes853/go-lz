package compression

type LZWCompression struct {
	Header LZWHeader
	Data   []byte
}

func LZWCompress(bytes []byte) LZWCompression {
	return LZWCompression{Header: LZWHeader{}, Data: make([]byte, 0)}
}
