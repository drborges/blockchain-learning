package blockchain

import "io"

type BtcEncoder interface {
	BtcEncode(w io.Writer) error
}

type BtcDecoder interface {
	BtcDecode(r io.Reader) error
}
