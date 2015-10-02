package blockchain

import (
	"encoding/binary"
	"github.com/drborges/blockchain-learning/iterator"
	"github.com/drborges/blockchain-learning/varint"
	"io"
)

type Transaction struct {
	Version  uint32
	In       []*TxIn
	Out      []*TxOut
	LockTime uint32
}

func NewTransaction(version uint32) *Transaction {
	return &Transaction{Version: version}
}

func (tx *Transaction) BtcEncode(w io.Writer) error {
	var version [4]byte
	binary.LittleEndian.PutUint32(version[:], tx.Version)
	if _, err := w.Write(version[:]); err != nil {
		return err
	}

	w.Write(varint.Serialize(uint64(len(tx.In))))

	for _, txin := range tx.In {
		txin.BtcEncode(w)
	}

	w.Write(varint.Serialize(uint64(len(tx.Out))))

	for _, txout := range tx.Out {
		txout.BtcEncode(w)
	}

	var lockTime [4]byte
	binary.LittleEndian.PutUint32(lockTime[:], tx.LockTime)
	if _, err := w.Write(lockTime[:]); err != nil {
		return err
	}

	return nil
}

func (tx *Transaction) BtcDecode(r io.Reader) error {
	var version [4]byte
	if _, err := r.Read(version[:]); err != nil {
		return err
	}
	tx.Version = binary.LittleEndian.Uint32(version[:])

	inCount, err := varint.Deserialize(r)
	if err != nil {
		return err
	}

	tx.In = make([]*TxIn, inCount)
	for i := 0; i < int(inCount); i++ {
		tx.In[i] = &TxIn{}
		tx.In[i].BtcDecode(r)
	}

	outCount, err := varint.Deserialize(r)
	if err != nil {
		return err
	}

	tx.Out = make([]*TxOut, outCount)
	for i := 0; i < int(outCount); i++ {
		tx.Out[i] = &TxOut{}
		tx.Out[i].BtcDecode(r)
	}

	var lockTime [4]byte
	if _, err := r.Read(lockTime[:]); err != nil {
		return err
	}
	tx.LockTime = binary.LittleEndian.Uint32(lockTime[:])

	return nil
}

func (tx *Transaction) SerializedSize() int {
	// Version(4) + LockTime(4) + Varint(len(In)) + Reduce(In, Size(TxIn)) + Reduce(Out, Size(TxOut))
	return 8 +
		varint.Size(uint64(len(tx.In))) +
		ReduceTotalSerializedSize(iterator.New(tx.In)) +
		varint.Size(uint64(len(tx.Out))) +
		ReduceTotalSerializedSize(iterator.New(tx.Out))
}

func ReduceTotalSerializedSize(iter iterator.Iterator) int {
	size := 0
	for iter.HasNext() {
		size += iter.Next().(Serializable).SerializedSize()
	}
	return size
}
