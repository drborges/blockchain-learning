package blockchain

import (
	"encoding/binary"
	"github.com/drborges/blockchain-learning/varint"
	"io"
)

type TxOut struct {
	Value    int64
	PKScript []byte
}

func (tx *TxOut) BtcEncode(w io.Writer) error {
	var value [8]byte
	binary.LittleEndian.PutUint64(value[:], uint64(tx.Value))
	if _, err := w.Write(value[:]); err != nil {
		return err
	}

	if _, err := w.Write(varint.Serialize(uint64(len(tx.PKScript)))); err != nil {
		return err
	}

	if _, err := w.Write(tx.PKScript); err != nil {
		return err
	}

	return nil
}

func (tx *TxOut) BtcDecode(r io.Reader) error {
	var value [8]byte
	if _, err := r.Read(value[:]); err != nil {
		return err
	}
	tx.Value = int64(binary.LittleEndian.Uint64(value[:]))

	pkscriptLen, err := varint.Deserialize(r)
	if err != nil {
		return err
	}

	pkscript := make([]byte, pkscriptLen)
	if _, err := r.Read(pkscript); err != nil {
		return err
	}
	tx.PKScript = pkscript

	return nil
}

func (tx *TxOut) SerializedSize() int {
	// Value(8) + Varint(len(PKScript))
	return 8 + varint.Size(uint64(len(tx.PKScript)))
}
