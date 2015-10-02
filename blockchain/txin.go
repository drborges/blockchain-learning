package blockchain

import (
	"encoding/binary"
	"github.com/drborges/blockchain-learning/varint"
	"io"
)

// See contract at: https://bitcoin.org/en/developer-reference#txin
type TxIn struct {
	PreviousTxOut   Outpoint
	ScriptSignature []byte
	Sequence        uint32
}

func (tx *TxIn) BtcEncode(w io.Writer) error {
	tx.PreviousTxOut.BtcEncode(w)

	if _, err := w.Write(varint.Serialize(uint64(len(tx.ScriptSignature)))); err != nil {
		return err
	}

	if _, err := w.Write(tx.ScriptSignature); err != nil {
		return err
	}

	var seq [4]byte
	binary.LittleEndian.PutUint32(seq[:], tx.Sequence)

	if _, err := w.Write(seq[:]); err != nil {
		return err
	}

	return nil
}

func (tx *TxIn) BtcDecode(r io.Reader) error {
	tx.PreviousTxOut.BtcDecode(r)
	scriptLen, err := varint.Deserialize(r)
	if err != nil {
		return err
	}

	sig := make([]byte, scriptLen)
	if _, err := r.Read(sig); err != nil {
		return err
	}
	tx.ScriptSignature = sig

	var seq [4]byte
	if _, err := r.Read(seq[:]); err != nil {
		return err
	}
	tx.Sequence = binary.LittleEndian.Uint32(seq[:])
	return nil
}

func (tx *TxIn) SerializedSize() int {
	return tx.PreviousTxOut.SerializedSize() +
		varint.Size(uint64(len(tx.ScriptSignature))) +
		len(tx.ScriptSignature) +
		varint.Size(uint64(tx.Sequence))
}
