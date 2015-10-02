package blockchain

import (
	"encoding/binary"
	"io"
)

// See contract at: https://bitcoin.org/en/developer-reference#outpoint
type Outpoint struct {
	Hash  [32]byte
	Index uint32
}

func (out *Outpoint) BtcEncode(w io.Writer) error {
	w.Write(out.Hash[:])
	var index [4]byte
	binary.LittleEndian.PutUint32(index[:], out.Index)
	w.Write(index[:])
	return nil
}

func (out *Outpoint) BtcDecode(r io.Reader) error {
	var hash [32]byte
	if _, err := r.Read(hash[:]); err != nil {
		return err
	}
	out.Hash = hash

	var index [4]byte
	if _, err := r.Read(index[:]); err != nil {
		return err
	}
	out.Index = binary.LittleEndian.Uint32(index[:])

	return nil
}

func (out *Outpoint) SerializedSize() int {
	return 36 // Hash(32) + Index(4) bytes
}
