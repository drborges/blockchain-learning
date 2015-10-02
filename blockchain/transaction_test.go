package blockchain_test

import (
	"bytes"
	"github.com/drborges/blockchain-learning/blockchain"
	"reflect"
	"testing"
)

func TestTransactionBtcEncodeDecode(t *testing.T) {
	expected := blockchain.Transaction{
		Version: 1,
		In: []*blockchain.TxIn{
			{
				PreviousTxOut: blockchain.Outpoint{
					Hash:  [32]byte{'1', '2', '3'},
					Index: 2,
				},
				ScriptSignature: []byte("abcd"),
				Sequence:        2,
			},
			{
				PreviousTxOut: blockchain.Outpoint{
					Hash:  [32]byte{'3', '2', '1'},
					Index: 1,
				},
				ScriptSignature: []byte("abcde"),
				Sequence:        3,
			},
		},
		Out: []*blockchain.TxOut{
			{Value: 123, PKScript: []byte("pkscript1")},
			{Value: 321, PKScript: []byte("pkscript2")},
		},
		LockTime: 123,
	}

	buf := bytes.NewBuffer(make([]byte, 0, expected.SerializedSize()))
	if err := expected.BtcEncode(buf); err != nil {
		t.Fatal(err)
	}

	var actual blockchain.Transaction
	if err := actual.BtcDecode(buf); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

//func TestTransactionDecoding(t *testing.T) {
//	data := []byte("660802c98f18fd34fd16d61c63cf447568370124ac5f3be626c2e1c3c9f0052d")
//
//	transaction := &blockchain.Transaction{}
//	if err := transaction.BtcDecode(bytes.NewReader(data)); err != nil {
//		log.Fatalln("Error: ", err)
//	}
//
//	fmt.Printf("%+v\n", transaction)
//}
